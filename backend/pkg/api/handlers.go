package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajaverma/clickhouse-file-tool/pkg/clickhouse"
	"github.com/rajaverma/clickhouse-file-tool/pkg/file"
)

type Handlers struct {
	clickhouseClient *clickhouse.Client
	fileHandler     *file.Handler
}

func NewHandlers(chClient *clickhouse.Client, fHandler *file.Handler) *Handlers {
	return &Handlers{
		clickhouseClient: chClient,
		fileHandler:     fHandler,
	}
}

type ClickHouseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWTToken string `json:"jwtToken"`
	Secure   bool   `json:"secure"`
}

type FileConfig struct {
	FilePath  string `json:"filePath"`
	Delimiter string `json:"delimiter"`
}

type ColumnRequest struct {
	Columns []string `json:"columns"`
}

func (h *Handlers) ConnectClickHouse(c *gin.Context) {
	var config ClickHouseConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := clickhouse.NewClient(clickhouse.Config{
		Host:     config.Host,
		Port:     config.Port,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
		JWTToken: config.JWTToken,
		Secure:   config.Secure,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.clickhouseClient = client
	c.JSON(http.StatusOK, gin.H{"message": "Connected to ClickHouse successfully"})
}

func (h *Handlers) GetTables(c *gin.Context) {
	if h.clickhouseClient == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not connected to ClickHouse"})
		return
	}

	tables, err := h.clickhouseClient.GetTables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tables": tables})
}

func (h *Handlers) GetColumns(c *gin.Context) {
	table := c.Param("table")
	if table == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table name is required"})
		return
	}

	if h.clickhouseClient == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not connected to ClickHouse"})
		return
	}

	columns, err := h.clickhouseClient.GetColumns(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"columns": columns})
}

func (h *Handlers) GetFileColumns(c *gin.Context) {
	var config FileConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler := file.NewHandler(file.Config{
		FilePath:  config.FilePath,
		Delimiter: config.Delimiter,
	})

	columns, err := handler.GetColumns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"columns": columns})
}

func (h *Handlers) IngestData(c *gin.Context) {
	source := c.Param("source")
	target := c.Param("target")

	if source == "clickhouse" && target == "file" {
		h.ingestClickHouseToFile(c)
	} else if source == "file" && target == "clickhouse" {
		h.ingestFileToClickHouse(c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source or target"})
	}
}

func (h *Handlers) ingestClickHouseToFile(c *gin.Context) {
	var req struct {
		Table    string   `json:"table"`
		Columns  []string `json:"columns"`
		FileConfig
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.clickhouseClient.QueryData(req.Table, req.Columns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Convert rows to [][]string
	var data [][]string
	for rows.Next() {
		values := make([]interface{}, len(req.Columns))
		pointers := make([]interface{}, len(req.Columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		row := make([]string, len(values))
		for i, v := range values {
			row[i] = fmt.Sprintf("%v", v)
		}
		data = append(data, row)
	}

	handler := file.NewHandler(file.Config{
		FilePath:  req.FilePath,
		Delimiter: req.Delimiter,
	})

	if err := handler.WriteData(data, req.Columns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data ingested successfully",
		"count":   len(data),
	})
}

func (h *Handlers) ingestFileToClickHouse(c *gin.Context) {
	// Implementation for file to ClickHouse ingestion
	// This would involve reading the file and inserting into ClickHouse
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
} 