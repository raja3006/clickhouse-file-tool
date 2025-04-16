# ClickHouse-FlatFile Data Ingestion Tool

A web-based application that facilitates bidirectional data ingestion between ClickHouse database and Flat Files.

## Features

- Bidirectional data flow (ClickHouse ↔ Flat File)
- JWT token-based authentication for ClickHouse
- Column selection for data ingestion
- Record count reporting
- Multi-table join support (bonus feature)
- Data preview functionality
- Progress tracking

## Tech Stack

- Backend: Go
- Frontend: React
- Database: ClickHouse
- Authentication: JWT

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn
- Docker (for local ClickHouse instance)

## Project Structure

```
.
├── backend/           # Go backend server
├── frontend/         # React frontend application
├── docker/           # Docker configuration files
└── docs/            # Documentation
```

## Setup Instructions

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm start
   ```

### ClickHouse Setup

1. Start ClickHouse using Docker:
   ```bash
   docker-compose up -d
   ```

2. Access ClickHouse at http://localhost:8123

## Usage

1. Open the application in your browser (default: http://localhost:3000)
2. Select the data source (ClickHouse or Flat File)
3. Configure connection parameters
4. Select tables and columns
5. Start the ingestion process
6. Monitor progress and view results

## Testing

Run the test suite:
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm test
```

## License

MIT 