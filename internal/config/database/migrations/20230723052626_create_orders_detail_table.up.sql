CREATE TABLE orders_detail (
   id SERIAL PRIMARY KEY,
   order_id UUID NOT NULL,
   product_id UUID NOT NULL,
   product_name varchar(100) NOT NULL,
   qty int NOT NULL DEFAULT 0,
   unit_price numeric NOT NULL DEFAULT 0,
   total_price numeric NOT NULL DEFAULT 0,
   created_at timestamptz NULL,
   updated_at timestamptz NULL,
   CONSTRAINT fk_orders_customer FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);