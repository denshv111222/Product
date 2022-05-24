
CREATE TABLE Products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR,
    brand_id BIGINT,
    sku VARCHAR(50),
    short_description TEXT,
    full_description TEXT,
    sort INTEGER,
    CONSTRAINT fk_brand_id FOREIGN KEY (brand_id) REFERENCES brands(id) on delete cascade
);