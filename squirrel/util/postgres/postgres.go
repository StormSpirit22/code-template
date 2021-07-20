package postgres

import (
	"net/url"
	"strings"
	"time"

	"git.aimap.io/go/config"
	"git.aimap.io/go/logs"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
)

type PgSqlConfig struct {
	Url     string `json:"url"`
	MaxIdle int    `json:"maxIdle" default:"10"`
	MaxOpen int    `json:"maxOpen" default:"100"`
}

func NewPgSqlConfig(path string) (*PgSqlConfig, error) {
	cfg := new(PgSqlConfig)
	err := config.Get(path).Scan(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewPostgres(conf *PgSqlConfig) (*sqlx.DB, error) {
	logs.Debug("connecting postgres...")
	uri, err := url.Parse(conf.Url)
	if err != nil {
		logs.Fatalf("parse %s err:%s", conf.Url, err.Error())
		return nil, err
	}
	logs.Infof("uri %s ", uri)

	pg, err := sqlx.Open(uri.Scheme, conf.Url)
	if err != nil {
		logs.Fatalf("open %s err:%s", conf.Url, err.Error())
		return nil, err
	}

	err = pg.Ping()
	if err != nil {
		logs.Fatalf("ping %s err:%s", conf.Url, err.Error())
		return nil, err
	}
	pg.SetMaxIdleConns(conf.MaxIdle)
	pg.SetMaxOpenConns(conf.MaxOpen)
	pg.SetConnMaxLifetime(2 * time.Minute)
	pg.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	logs.Debugf("%s is connected", conf.Url)
	return pg, nil
}
