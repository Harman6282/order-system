DROP TRIGGER IF EXISTS orders_set_updated_at ON orders;

DROP FUNCTION IF EXISTS set_updated_at;

DROP INDEX IF EXISTS idx_orders_status;

DROP TABLE IF EXISTS orders;

DROP TYPE IF EXISTS order_status;