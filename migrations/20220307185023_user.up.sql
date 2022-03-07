CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    login VARCHAR(50),
    password VARCHAR(255),
    role INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (login)
);