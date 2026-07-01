CREATE TABLE users(

    id UUID PRIMARY KEY,
    role_id UUID REFERENCES roles(id),
    full_name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);