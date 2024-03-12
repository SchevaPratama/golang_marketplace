CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    image_url VARCHAR(2048), -- Allow longer URLs for images
    stock INT NOT NULL,
    condition VARCHAR(50),
    is_purchasable BOOLEAN DEFAULT TRUE,
    user_id VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS product_tag (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    product_id uuid NOT NULL
);