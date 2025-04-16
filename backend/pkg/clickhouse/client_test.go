package clickhouse

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "Valid config",
			config: Config{
				Host:     "localhost",
				Port:     9000,
				Database: "default",
				Username: "default",
				Password: "password",
				Secure:   false,
			},
			wantErr: false,
		},
		{
			name: "Invalid host",
			config: Config{
				Host:     "",
				Port:     9000,
				Database: "default",
				Username: "default",
				Password: "password",
				Secure:   false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetTables(t *testing.T) {
	// This is an integration test that requires a running ClickHouse instance
	// Skip it if we're not running integration tests
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	client, err := NewClient(Config{
		Host:     "localhost",
		Port:     9000,
		Database: "default",
		Username: "default",
		Password: "password",
		Secure:   false,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	tables, err := client.GetTables()
	if err != nil {
		t.Errorf("GetTables() error = %v", err)
	}
	if len(tables) == 0 {
		t.Log("No tables found in database")
	}
}

func TestClient_GetColumns(t *testing.T) {
	// This is an integration test that requires a running ClickHouse instance
	// Skip it if we're not running integration tests
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	client, err := NewClient(Config{
		Host:     "localhost",
		Port:     9000,
		Database: "default",
		Username: "default",
		Password: "password",
		Secure:   false,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// First, get the tables
	tables, err := client.GetTables()
	if err != nil {
		t.Fatalf("Failed to get tables: %v", err)
	}
	if len(tables) == 0 {
		t.Skip("No tables available for testing")
	}

	// Test getting columns for the first table
	columns, err := client.GetColumns(tables[0])
	if err != nil {
		t.Errorf("GetColumns() error = %v", err)
	}
	if len(columns) == 0 {
		t.Errorf("GetColumns() returned no columns")
	}
} 