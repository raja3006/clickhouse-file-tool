package clickhouse

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	JWTToken string
	Secure   bool
}

type Client struct {
	conn driver.Conn
}

func NewClient(cfg Config) (*Client, error) {
	opts := &clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		Settings: clickhouse.Settings{
			"jwt": cfg.JWTToken,
		},
		Secure: cfg.Secure,
	}

	conn, err := clickhouse.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ClickHouse: %v", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping ClickHouse: %v", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetTables() ([]string, error) {
	rows, err := c.conn.Query(context.Background(), "SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("failed to query tables: %v", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %v", err)
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func (c *Client) GetColumns(table string) ([]string, error) {
	query := fmt.Sprintf("DESCRIBE TABLE %s", table)
	rows, err := c.conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query columns: %v", err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var name, type_, defaultType, defaultExpression, comment, codecExpression, ttlExpression string
		if err := rows.Scan(&name, &type_, &defaultType, &defaultExpression, &comment, &codecExpression, &ttlExpression); err != nil {
			return nil, fmt.Errorf("failed to scan column: %v", err)
		}
		columns = append(columns, name)
	}

	return columns, nil
}

func (c *Client) QueryData(table string, columns []string) (driver.Rows, error) {
	columnList := "*"
	if len(columns) > 0 {
		columnList = ""
		for i, col := range columns {
			if i > 0 {
				columnList += ", "
			}
			columnList += col
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s", columnList, table)
	return c.conn.Query(context.Background(), query)
} 