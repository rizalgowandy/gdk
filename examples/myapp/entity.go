package myapp

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
}
