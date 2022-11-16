package repository

type Account struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type UpdateRequest struct {
	AccountID int64
	Email     string
	Name      string
	Password  string
	Role      string
}
