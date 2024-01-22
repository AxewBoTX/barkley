<script>
  import toast, { Toaster } from "svelte-french-toast";
  import Icon from "@iconify/svelte";

  // Fetch Todo List
  const fetchTodos = async () => {
    const res = await fetch("/api/todo");
    const data = await res.json();
    return data;
  };
</script>

<div class="flex flex-col items-center mt-[50px] px-4">
  <div class="card w-full max-w-[600px] bg-base-200 shadow-2xl">
    <div class="card-body items-center">
      <h1 class="text-5xl">Dondu</h1>
      <h3 class="text-xl italic mb-[45px]">Your personal manager</h3>
      <!-- Create Todo Form -->
      <form
        class="join w-full max-w-[800px] flex items-center justify-end"
        on:submit|preventDefault={async (e) => {
          const form_data = new FormData(e.currentTarget);
          const formData = {};
          for (let i of form_data) {
            const [key, value] = i;
            formData[key] = value;
          }
          if (
            formData.title.trim().length == 0 ||
            formData.title.trim() == ""
          ) {
            toast.error("Input is needed!", {
              position: "top-right",
            });
          } else {
            const res = await fetch("/api/todo", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify(formData),
            });
            if (res.ok) {
              location.reload();
            } else {
              toast.error("Something went wrong!", {
                position: "top-right",
              });
            }
          }
        }}
      >
        <input
          name="title"
          class="input input-bordered join-item w-full max-w-[200px]"
          placeholder="Title"
        />
        <button
          class="btn hover:btn-primary btn-outline join-item"
          type="submit"
          ><Icon icon="material-symbols:add" class="text-2xl" /></button
        >
      </form>
      <!-- Todo List -->
      <div class="w-full mt-[25px]">
        {#await fetchTodos()}
          <div class="text-2xl">Loading...</div>
        {:then todos}
          <div class="flex flex-col items-center gap-[20px] w-full">
            {#each todos as todo}
              <div
                class="bg-base-300 rounded flex items-center py-2 px-3 w-full justify-between max-w-[350px] hover:outline outline-secondary transition-all"
              >
                {todo.title}
                <form
                  on:submit|preventDefault={async () => {
                    const res = await fetch(`/api/todo/${todo.id}`, {
                      method: "DELETE",
                    });
                    if (res.ok) {
                      location.reload();
                    } else {
                      toast.error("Something went wrong!", {
                        position: "top-right",
                      });
                    }
                  }}
                >
                  <button class="btn btn-outline hover:btn-error"
                    ><Icon icon="ic:outline-delete" class="text-2xl" /></button
                  >
                </form>
              </div>
            {/each}
          </div>
        {/await}
      </div>
    </div>
  </div>
</div>
<Toaster />
