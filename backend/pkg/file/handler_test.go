package file

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestHandler_GetColumns(t *testing.T) {
	// Create a temporary CSV file for testing
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")
	content := "name,age,city\nJohn,30,New York\nJane,25,London\n"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name       string
		config     Config
		want       []string
		wantErr    bool
		setupFile  bool
		fileExists bool
	}{
		{
			name: "Valid CSV file",
			config: Config{
				FilePath:  tmpFile,
				Delimiter: ",",
			},
			want:       []string{"name", "age", "city"},
			wantErr:    false,
			setupFile:  true,
			fileExists: true,
		},
		{
			name: "Non-existent file",
			config: Config{
				FilePath:  "nonexistent.csv",
				Delimiter: ",",
			},
			want:       nil,
			wantErr:    true,
			setupFile:  false,
			fileExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.config)
			got, err := h.GetColumns()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_ReadData(t *testing.T) {
	// Create a temporary CSV file for testing
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")
	content := "name,age,city\nJohn,30,New York\nJane,25,London\n"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name       string
		config     Config
		columns    []string
		want       [][]string
		wantErr    bool
		setupFile  bool
		fileExists bool
	}{
		{
			name: "Read all columns",
			config: Config{
				FilePath:  tmpFile,
				Delimiter: ",",
			},
			columns: []string{"name", "age", "city"},
			want: [][]string{
				{"John", "30", "New York"},
				{"Jane", "25", "London"},
			},
			wantErr:    false,
			setupFile:  true,
			fileExists: true,
		},
		{
			name: "Read specific columns",
			config: Config{
				FilePath:  tmpFile,
				Delimiter: ",",
			},
			columns: []string{"name", "city"},
			want: [][]string{
				{"John", "New York"},
				{"Jane", "London"},
			},
			wantErr:    false,
			setupFile:  true,
			fileExists: true,
		},
		{
			name: "Invalid column",
			config: Config{
				FilePath:  tmpFile,
				Delimiter: ",",
			},
			columns:    []string{"invalid"},
			want:       nil,
			wantErr:    true,
			setupFile:  true,
			fileExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.config)
			got, err := h.ReadData(tt.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_WriteData(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "output.csv")

	tests := []struct {
		name    string
		config  Config
		data    [][]string
		headers []string
		wantErr bool
	}{
		{
			name: "Write valid data",
			config: Config{
				FilePath:  tmpFile,
				Delimiter: ",",
			},
			data: [][]string{
				{"John", "30", "New York"},
				{"Jane", "25", "London"},
			},
			headers: []string{"name", "age", "city"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.config)
			if err := h.WriteData(tt.data, tt.headers); (err != nil) != tt.wantErr {
				t.Errorf("Handler.WriteData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the written file
			if !tt.wantErr {
				content, err := os.ReadFile(tmpFile)
				if err != nil {
					t.Errorf("Failed to read output file: %v", err)
					return
				}

				expectedContent := "name,age,city\nJohn,30,New York\nJane,25,London\n"
				if string(content) != expectedContent {
					t.Errorf("WriteData() wrote %v, want %v", string(content), expectedContent)
				}
			}
		})
	}
} 