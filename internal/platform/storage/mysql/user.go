package mysql

type sqlUser struct {
	id        string `db:"id"`
	name      string `db:"name"`
	firstname string `db:"firstname"`
}
