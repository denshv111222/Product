CREATE TABLE IF NOT EXISTS Categories (
    id_categories BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR,
    parent_id BIGINT
);

CREATE TABLE IF NOT EXISTS Categories_Products (
products_id BIGINT,
categories_id BIGINT,
sort INTEGER ,
CONSTRAINT fk_products_id FOREIGN KEY (products_id) REFERENCES Products(id) on delete cascade  ,
CONSTRAINT fk_categories_id FOREIGN KEY (categories_id) REFERENCES Categories(id_categories) on delete cascade
);
CREATE UNIQUE INDEX Category_Product
    ON Categories_Products (products_id, categories_id)