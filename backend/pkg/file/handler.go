package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	FilePath  string
	Delimiter string
}

type Handler struct {
	config Config
}

func NewHandler(cfg Config) *Handler {
	return &Handler{config: cfg}
}

func (h *Handler) GetColumns() ([]string, error) {
	file, err := os.Open(h.config.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = []rune(h.config.Delimiter)[0]

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read headers: %v", err)
	}

	return headers, nil
}

func (h *Handler) ReadData(columns []string) ([][]string, error) {
	file, err := os.Open(h.config.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = []rune(h.config.Delimiter)[0]

	// Read headers
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read headers: %v", err)
	}

	// Create column index map
	columnIndexes := make(map[string]int)
	for i, header := range headers {
		columnIndexes[header] = i
	}

	// If no columns specified, use all columns
	if len(columns) == 0 {
		columns = headers
	}

	// Validate requested columns exist
	for _, col := range columns {
		if _, exists := columnIndexes[col]; !exists {
			return nil, fmt.Errorf("column %s not found in file", col)
		}
	}

	var data [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %v", err)
		}

		// Extract only requested columns
		row := make([]string, len(columns))
		for i, col := range columns {
			row[i] = record[columnIndexes[col]]
		}
		data = append(data, row)
	}

	return data, nil
}

func (h *Handler) WriteData(data [][]string, headers []string) error {
	file, err := os.Create(h.config.FilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = []rune(h.config.Delimiter)[0]
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %v", err)
	}

	// Write data
	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %v", err)
		}
	}

	return nil
} 