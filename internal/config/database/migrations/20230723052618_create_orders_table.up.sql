CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL,
    total_items int NOT NULL DEFAULT 0,
    total_price numeric NOT NULL DEFAULT 0,
    status smallint NOT NULL DEFAULT 0,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    CONSTRAINT fk_orders_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);