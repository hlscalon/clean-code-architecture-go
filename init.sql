CREATE UNLOGGED TABLE "order" (
	id SERIAL PRIMARY KEY,
	customer_id INTEGER NOT NULL DEFAULT 0,
	total_price DECIMAL(11, 2) NOT NULL DEFAULT 0,
	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_order_customer_id ON "order" (customer_id);

-- valores default

INSERT INTO "order" (customer_id, total_price)
VALUES
    (1, 559.90),
    (2, 1250.90),
    (3, 1750.90),
    (4, 720.90);
