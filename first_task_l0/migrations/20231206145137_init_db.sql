-- +goose Up
CREATE TABLE IF NOT EXISTS providers
(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS banks
(
    id serial PRIMARY KEY NOT NULL UNIQUE ,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions
(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    request_id VARCHAR,
    currency VARCHAR,
    provider_id INTEGER,
    FOREIGN KEY (provider_id) REFERENCES providers (id) ON DELETE CASCADE,
    amount INTEGER,
    payment_dt INT,
    bank_id INTEGER,
    FOREIGN KEY (bank_id) REFERENCES banks (id) ON DELETE CASCADE,
    delivery_cost BIGINT,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE IF NOT EXISTS brands
(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS items
(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    track_number VARCHAR NOT NULL UNIQUE,
    price BIGINT NOT NULL,
    SIZE VARCHAR default 0,
    total_price BIGINT NOT NULL,
    nm_id SERIAL NOT NULL,
    brand_id INTEGER,
    FOREIGN KEY (brand_id) REFERENCES brands (id) ON DELETE CASCADE,
    status INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS customers
(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    date_created TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS addresses
(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    zip VARCHAR,
    city TEXT,
    address TEXT,
    region TEXT,
    owner_id UUID,
    FOREIGN KEY (owner_id) REFERENCES customers (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS delivery_services
(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS orders
(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    track_number VARCHAR NOT NULL UNIQUE,
    entry VARCHAR NOT NULL,
    delivery JSONB NOT NULL,
    payment_id UUID,
    FOREIGN KEY (payment_id) REFERENCES transactions (id) ON DELETE CASCADE,
    items INTEGER,
    FOREIGN KEY (items) REFERENCES items (id) ON DELETE CASCADE,
    locale VARCHAR NOT NULL,
    internal_signature VARCHAR,
    delivery_service_id INTEGER,
    FOREIGN KEY (delivery_service_id) REFERENCES delivery_services (id) ON DELETE CASCADE,
    date_created TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down


