CREATE TABLE IF NOT EXISTS Orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,  -- Assuming a foreign key relationship with a users table
    total DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
