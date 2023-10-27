// client.go in the path: internal/pkg/db/sql/client

package client

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Config struct {
	Database string
	HostName string
	PortNos  int
	UserName string
	Password string
	SSLMode  string
}

type SqlClient struct {
	sync.Mutex
	Config Config
	Conn   *sqlx.DB
}

func NewSqlClient(cfg Config) *SqlClient {
	return &SqlClient{Config: cfg}
}

func (sc *SqlClient) Init(ctx context.Context) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		sc.Config.HostName,
		sc.Config.PortNos,
		sc.Config.UserName,
		sc.Config.Password,
		sc.Config.Database,
		sc.Config.SSLMode,
	)

	conn, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return err
	}

	sc.Conn = conn
	return sc.Conn.PingContext(ctx)
}

// ... You can expand this file with more methods, error handling, etc.
