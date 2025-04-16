-- Create example tables
CREATE TABLE IF NOT EXISTS users
(
    id UInt32,
    name String,
    email String,
    age UInt8,
    created_at DateTime
) ENGINE = MergeTree()
ORDER BY id;

CREATE TABLE IF NOT EXISTS orders
(
    id UInt32,
    user_id UInt32,
    product_name String,
    amount Decimal(10,2),
    order_date DateTime
) ENGINE = MergeTree()
ORDER BY id;

-- Insert example data into users table
INSERT INTO users (id, name, email, age, created_at) VALUES
(1, 'John Doe', 'john@example.com', 30, '2024-01-01 10:00:00'),
(2, 'Jane Smith', 'jane@example.com', 25, '2024-01-02 11:00:00'),
(3, 'Bob Johnson', 'bob@example.com', 35, '2024-01-03 12:00:00'),
(4, 'Alice Brown', 'alice@example.com', 28, '2024-01-04 13:00:00'),
(5, 'Charlie Wilson', 'charlie@example.com', 32, '2024-01-05 14:00:00');

-- Insert example data into orders table
INSERT INTO orders (id, user_id, product_name, amount, order_date) VALUES
(1, 1, 'Laptop', 1299.99, '2024-01-01 10:30:00'),
(2, 2, 'Smartphone', 799.99, '2024-01-02 11:30:00'),
(3, 1, 'Monitor', 399.99, '2024-01-03 12:30:00'),
(4, 3, 'Keyboard', 99.99, '2024-01-04 13:30:00'),
(5, 4, 'Mouse', 49.99, '2024-01-05 14:30:00'),
(6, 5, 'Headphones', 199.99, '2024-01-06 15:30:00'),
(7, 2, 'Tablet', 599.99, '2024-01-07 16:30:00'),
(8, 3, 'Printer', 299.99, '2024-01-08 17:30:00'),
(9, 4, 'Scanner', 199.99, '2024-01-09 18:30:00'),
(10, 5, 'Speaker', 149.99, '2024-01-10 19:30:00'); 