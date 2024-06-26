FROM docker.io/golang:1.22.3

# set the working directory
WORKDIR /usr/src/app

# copy bun from bun base image onto current golang base image
COPY --from=docker.io/oven/bun:1.1.7-alpine /usr/local/bin/bun /usr/local/bin/bun

# set the ENV for adding bun to PATH
ENV PATH="/usr/local/bin:${PATH}"

# copy all the project files onto the working directory
COPY . .

# basic container setup (according to my liking)
RUN apt-get update && \
	apt-get install -y curl && \
	apt-get install -y tmux psmisc && \
	bash -c "echo 'PATH="/usr/local/cargo/bin:$PATH"' >> ~/.bashrc" && \
	bash -c "source ~/.bashrc" && \
	git config --global --add safe.directory /usr/src/app

# install needed dependencies
RUN go mod download && \
	go install github.com/a-h/templ/cmd/templ@latest && \
 	go install github.com/air-verse/air@latest && \
 	go install github.com/go-task/task/v3/cmd/task@latest
RUN bun install --yarn

# run the tail command to keep the container running
CMD ["tail", "-f", "/dev/null"]
