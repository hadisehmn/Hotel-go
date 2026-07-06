package models

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
	Role     string `db:"role"`
}
