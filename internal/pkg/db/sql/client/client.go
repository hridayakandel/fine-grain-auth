package client

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
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
	Conn   *sqlx.DB
	Config Config // <-- Added this field
}

func NewSqlClient(cfg Config) *SqlClient {
	return &SqlClient{Config: cfg} // <-- Set the config here
}

func (sc *SqlClient) Init() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		sc.Config.HostName,
		sc.Config.PortNos,
		sc.Config.UserName,
		sc.Config.Password,
		sc.Config.Database,
		sc.Config.SSLMode,
	)

	conn, err := sqlx.Open("pgx", connStr)
	// Note the use of "postgres" instead of "pgx"
	if err != nil {
		return err
	}

	sc.Conn = conn
	return sc.Conn.PingContext(context.Background())
}
func (sc *SqlClient) NamedQueryExecution(props *PrepareNamedWithContextProps) error {
	// Convert generic interface{} Args to named.Arg

	// Using named query to bind parameters to SQL query
	rows, err := sc.Conn.NamedQueryContext(props.Ctx, props.Query, props.Args)
	if err != nil {
		return err
	}
	defer rows.Close()

	// If a destination is provided, scan the result into it
	if props.Dest != nil {
		if rows.Next() {
			err = rows.StructScan(props.Dest)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
