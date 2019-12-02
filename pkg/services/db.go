package services

import (
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/graphlog/heimdall/pkg/config"
	"github.com/jmoiron/sqlx"
)

func NewDBConnection(c *config.Config) *sqlx.DB {
	connect, err := sqlx.Open("clickhouse", c.DataStores.ClickHouseURI)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return connect
}
