-- eater schema
CREATE SCHEMA IF NOT EXISTS support;

-- ratings
CREATE TABLE IF NOT EXISTS support.ratings (
	id varchar(36) PRIMARY KEY,
    order_id VARCHAR(16),
    rating VARCHAR(255),
    comment VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);

-- idx_order_rating
CREATE INDEX IF NOT EXISTS idx_order_rating ON support.ratings(order_id);




