package postgresql

import (
	"fmt"
	"regulus/app/utils/config"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/jmoiron/sqlx"
	//sqldriver import
	_ "github.com/lib/pq"
)

// Postgresql DB接続のための構造体
type Postgresql struct {
	url      string
	port     string
	user     string
	password string
	dbname   string
}

// WithDbContext DB処理のためのラッパー関数
func (postgresql *Postgresql) WithDbContext(fn func(db *sqlx.DB) error) error {
	d, err := postgresql.Open()
	defer d.Close()
	if err != nil {
		return err
	}

	return fn(d)
}

// Open 接続情報は設定ファイルから読み込み
func (postgresql *Postgresql) Open() (*sqlx.DB, error) {
	dataSourceName := "host=" + postgresql.url +
		" port=" + postgresql.port +
		" user=" + postgresql.user +
		" password=" + postgresql.password +
		" dbname=" + postgresql.dbname +
		" sslmode=disable"
	return sqlx.Open("postgres", dataSourceName)

}

// Migrate DBスキーマ設定
func Migrate() {

	migrations := &migrate.FileMigrationSource{
		Dir: "./resources/db/migrations/postgres",
	}

	p := NewPostgresql()

	db, err := p.Open()
	if err != nil {
		fmt.Println("db connect error")
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

}

// NewPostgresql コンストラクタ
func NewPostgresql() *Postgresql {
	p := &Postgresql{
		url:      config.DbUrl(),
		port:     config.DbPort(),
		user:     config.DbUser(),
		password: config.DbPassword(),
		dbname:   config.DbName(),
	}

	return p
}
