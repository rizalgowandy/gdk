package myapp

// User represents user data model.
type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
}
