CREATE TABLE IF NOT EXISTS Orders_Items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES Orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES Products(id) ON DELETE CASCADE,
    quantity INT NOT NULL DEFAULT 1,
    price DECIMAL(10, 2) NOT NULL
);
