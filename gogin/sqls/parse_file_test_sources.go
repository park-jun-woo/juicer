package sqls

var parseFileValidRepoSrc = `package repo

import "context"
import "database/sql"

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) GetAll(ctx context.Context) (*sql.Rows, error) {
	return r.db.QueryContext(ctx, "SELECT id, name FROM users WHERE id > $1")
}

func (r *UserRepo) Create(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES ($1)")
	return err
}
`

var parseFileEmptyRepoSrc = `package repo

type EmptyRepo struct{}

func (r *EmptyRepo) NotSQL() string {
	return "hello"
}
`

var parseFileMiscRepoSrc = `package repo

import "database/sql"

var db *sql.DB

type Repo struct{}

func standalone() {}
`

var parseFileOrderRepoSrc = `package repo

import "context"
import "database/sql"

type OrderRepo struct {
	db *sql.DB
}

func (r *OrderRepo) Create(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO orders (id) VALUES ($1)")
	return err
}
`

var parseFileItemRepoSrc = "package repo\n\nimport \"context\"\nimport \"database/sql\"\n\ntype ItemRepo struct {\n\tdb *sql.DB\n}\n\nfunc (r *ItemRepo) CreateAndReturn(ctx context.Context) *sql.Row {\n\treturn r.db.QueryRowContext(ctx, `INSERT INTO items (name) VALUES ($1) RETURNING id`)\n}\n"
