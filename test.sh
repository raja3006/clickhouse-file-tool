#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting ClickHouse File Tool Test Script${NC}"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}Docker is not running. Please start Docker and try again.${NC}"
    exit 1
fi

# Create necessary directories
echo "Creating necessary directories..."
mkdir -p uploads test-data

# Start the application
echo "Starting the application..."
docker-compose down -v # Clean up any existing containers and volumes
docker-compose up -d

# Wait for services to be ready
echo "Waiting for services to be ready..."
sleep 10

# Check if services are running
if ! docker-compose ps | grep -q "Up"; then
    echo -e "${RED}Services failed to start. Please check docker-compose logs for more information.${NC}"
    exit 1
fi

echo -e "${GREEN}Services are running!${NC}"
echo "Frontend: http://localhost:3000"
echo "Backend: http://localhost:8080"
echo "ClickHouse HTTP: http://localhost:8123"

# Test ClickHouse connection
echo -e "\nTesting ClickHouse connection..."
if curl -s "http://localhost:8123/ping" | grep -q "Ok."; then
    echo -e "${GREEN}ClickHouse is responding${NC}"
else
    echo -e "${RED}ClickHouse is not responding${NC}"
fi

# Test backend health
echo -e "\nTesting backend health..."
if curl -s "http://localhost:8080/health" | grep -q "ok"; then
    echo -e "${GREEN}Backend is healthy${NC}"
else
    echo -e "${RED}Backend is not responding${NC}"
fi

echo -e "\n${GREEN}Test data is available:${NC}"
echo "1. ClickHouse tables:"
echo "   - users (id, name, email, age, created_at)"
echo "   - orders (id, user_id, product_name, amount, order_date)"
echo "2. Example CSV file:"
echo "   - test-data/example.csv (employees data)"

echo -e "\n${GREEN}You can now:${NC}"
echo "1. Open http://localhost:3000 in your browser"
echo "2. Try ingesting data from ClickHouse to CSV:"
echo "   - Source: ClickHouse (users or orders table)"
echo "   - Target: Flat File (specify a path in the uploads directory)"
echo "3. Try ingesting data from CSV to ClickHouse:"
echo "   - Source: Flat File (use test-data/example.csv)"
echo "   - Target: ClickHouse (new table will be created)"

echo -e "\n${GREEN}To stop the application:${NC}"
echo "docker-compose down"

echo -e "\n${GREEN}To view logs:${NC}"
echo "docker-compose logs -f" 