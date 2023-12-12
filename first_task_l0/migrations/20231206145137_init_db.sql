-- +goose Up
CREATE TABLE IF NOT EXISTS customers
(
    id VARCHAR PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    zip VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    email VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    id VARCHAR PRIMARY KEY NOT NULL UNIQUE,
    track_number VARCHAR NOT NULL UNIQUE,
    entry VARCHAR NOT NULL,
    customer_id VARCHAR,
    FOREIGN KEY (customer_id) REFERENCES customers,
    payment JSONB NOT NULL,
    locale VARCHAR NOT NULL,
    internal_signature VARCHAR,
    delivery_service VARCHAR NOT NULL,
    shard_key VARCHAR,
    sm_id BIGINT,
    date_created TIMESTAMP WITH TIME ZONE NOT NULL,
    oof_shard VARCHAR
);

CREATE TABLE IF NOT EXISTS items
(
    id BIGINT PRIMARY KEY NOT NULL UNIQUE,
    track_number VARCHAR NOT NULL,
    price BIGINT NOT NULL,
    rid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    sale BIGINT NOT NULL,
    size VARCHAR NOT NULL default 0,
    total_price BIGINT NOT NULL,
    nm_id BIGINT NOT NULL,
    brand VARCHAR,
    status BIGINT,
    order_id VARCHAR NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders
);
-- +goose Down


