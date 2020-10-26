package myapp

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
}
