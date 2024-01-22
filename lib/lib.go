package lib

const (
	DB_PATH string = "database/base.db"
	PORT    string = ":8080"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
