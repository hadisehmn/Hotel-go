package models

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	ID       int      `db:"id"`
	Name     string   `db:"name"`
	Password string   `db:"password"`
	Age      int      `db:"age"`
	Phone    string   `db:"phone"`
	Role     UserRole `db:"role"`
}

// admin
