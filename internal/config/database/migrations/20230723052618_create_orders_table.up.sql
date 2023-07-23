CREATE TYPE status AS ENUM ('unpaid','paid');

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL,
    total_items int NOT NULL DEFAULT 0,
    total_price numeric NOT NULL DEFAULT 0,
    status status NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by varchar(100) NOT NULL,
    updated_by varchar(100) NOT NULL,
    deleted_at timestamptz NULL, 
    CONSTRAINT fk_orders_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE INDEX idx_orders_deleted_at ON orders USING btree (deleted_at);