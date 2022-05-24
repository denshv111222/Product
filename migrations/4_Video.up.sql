CREATE TABLE IF NOT EXISTS Videos (
id_video BIGSERIAL PRIMARY KEY NOT NULL,
name VARCHAR(100),
storage VARCHAR(100),
path VARCHAR(100)
);
CREATE TABLE IF NOT EXISTS Videos_Products (
products_id BIGINT,
videos_id BIGINT,
sort INTEGER ,
CONSTRAINT fk_products_id FOREIGN KEY (products_id) REFERENCES Products(id) on delete cascade  ,
CONSTRAINT fk_videos_id FOREIGN KEY (videos_id) REFERENCES Videos(id_video) on delete cascade

);
CREATE UNIQUE INDEX Vide_Product
    ON Videos_Products (products_id, videos_id)