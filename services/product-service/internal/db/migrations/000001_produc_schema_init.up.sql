CREATE TYPE product_status AS ENUM ('ACTIVE', 'INACTIVE', 'OUT_OF_STOCK', 'DISCONTINUED');
CREATE TYPE availability_status AS ENUM ('AVAILABLE', 'BACKORDER', 'PREORDER');
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    price DOUBLE PRECISION NOT NULL,
    discount_price DOUBLE PRECISION,
    currency TEXT NOT NULL,

    status product_status NOT NULL,
    availability availability_status NOT NULL,

    stock_quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)