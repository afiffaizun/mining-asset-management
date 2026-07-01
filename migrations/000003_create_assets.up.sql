CREATE TABLE assets(
    id UUID PRIMARY KEY,
    asset_code VARCHAR(30) UNIQUE,
    asset_name VARCHAR(100),
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW()
);