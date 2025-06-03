CREATE TYPE order_status AS ENUM (
    'PENDING',
    'CONFIRMED',
    'PROCESSING',
    'SHIPPED',
    'DELIVERED',
    'CANCELLED',
    'RETURNED'
);

CREATE TYPE payment_status AS ENUM (
    'UNPAID',
    'PAID',
    'REFUNDED',
    'FAILED'
);

CREATE TYPE shipping_method AS ENUM (
    'STANDARD',
    'EXPRESS',
    'SAME_DAY',
    'PICKUP'
);

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL,
    order_number VARCHAR(50) NOT NULL,
    order_status order_status NOT NULL DEFAULT 'PENDING',
    payment_status payment_status NOT NULL DEFAULT 'UNPAID',
    shipping_method shipping_method NOT NULL DEFAULT 'STANDARD',
    total_amount NUMERIC(15, 4) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'INR'
);
