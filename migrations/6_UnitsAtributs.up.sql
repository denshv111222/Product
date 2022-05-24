CREATE TABLE IF NOT EXISTS Units(
    id_unit BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR
);

CREATE TABLE IF NOT EXISTS atributes(
    id_atribute BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR,
    unit_id BIGINT,
    CONSTRAINT fk_unit_id FOREIGN KEY (unit_id) REFERENCES Units(id_unit) on delete cascade
);

CREATE TABLE IF NOT EXISTS atributes_values(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    atribute_id BIGINT,
    CONSTRAINT fk_atribute_id FOREIGN KEY (atribute_id) REFERENCES atributes(id_atribute) on delete cascade
);

CREATE TABLE IF NOT EXISTS atributes_values_products (
products_id BIGINT,
atributes_values_id BIGINT,
sort INTEGER ,
CONSTRAINT fk_products_id FOREIGN KEY (products_id) REFERENCES Products(id),
CONSTRAINT fk_atributes_values_id FOREIGN KEY (atributes_values_id) REFERENCES atributes_values(id)


);
CREATE UNIQUE INDEX atribut_value_product
    ON atributes_values_products (products_id, atributes_values_id)