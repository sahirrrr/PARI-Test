package infra

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type (
	Databases struct {
		dig.Out
		Pg *sqlx.DB
	}

	DatabaseCfgs struct {
		dig.In
		Pg *DatabaseCfg
	}

	DatabaseCfg struct {
		DBName string `envconfig:"NAME" required:"true" default:"name"`
		DBUser string `envconfig:"USER" required:"true" default:"user"`
		DBPass string `envconfig:"PASS" required:"true" default:"pass"`
		Host   string `envconfig:"HOST" required:"true" default:"localhost"`
		Port   string `envconfig:"PORT" required:"true" default:"5432"`

		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"20" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"5" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"15m" required:"true"`
	}
)

func NewDatabases(cfgs DatabaseCfgs) Databases {
	return Databases{
		Pg: openPostgres(cfgs.Pg),
	}
}

func openPostgres(p *DatabaseCfg) *sqlx.DB {
	//nolint:nosprintfhostport
	sqlConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.DBUser, p.DBPass, p.Host, p.Port, p.DBName,
	)

	db, err := sqlx.Connect("postgres", sqlConn)
	if err != nil {
		logrus.Fatalf("postgres: %s", err.Error())
	}

	db.SetConnMaxLifetime(p.ConnMaxLifetime)
	db.SetMaxIdleConns(p.MaxIdleConns)
	db.SetMaxOpenConns(p.MaxOpenConns)

	if err = db.Ping(); err != nil {
		logrus.Fatalf("postgres: %s", err.Error())
	}

	return db
}
