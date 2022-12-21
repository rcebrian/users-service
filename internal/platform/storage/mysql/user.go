package mysql

const (
	sqlUserTable = "user"
)

type sqlUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Firstname string `db:"firstname"`
}
