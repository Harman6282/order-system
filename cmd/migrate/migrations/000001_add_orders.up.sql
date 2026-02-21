CREATE TYPE order_status AS ENUM (
    'created',
    'paid',
    'processing',
    'completed',
    'failed',
    'cancelled'
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    product_name TEXT NOT NULL,
    price BIGINT NOT NULL CHECK (price >= 0),

    status order_status NOT NULL DEFAULT 'created',

    processing_by TEXT,
    processing_started_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_status ON orders(status);


CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER orders_set_updated_at
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();