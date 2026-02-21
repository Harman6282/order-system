CREATE TYPE order_status AS ENUM (
    'created',
    'paid',
    'processing',
    'completed',
    'failed',
    'cancelled'
);