
CREATE TABLE IF NOT EXISTS Images (
id_Image BIGSERIAL PRIMARY KEY NOT NULL,
name VARCHAR(100),
storage VARCHAR(100),
path VARCHAR(100)
);



CREATE TABLE IF NOT EXISTS Images_Products (
products_id BIGINT,
images_id BIGINT,
sort INTEGER ,
CONSTRAINT fk_products_id FOREIGN KEY (products_id) REFERENCES Products(id) on delete cascade ,
CONSTRAINT fk_images_id FOREIGN KEY (images_id) REFERENCES Images(id_Image) on delete cascade

);
CREATE UNIQUE INDEX Image_Product
ON Images_Products (products_id, images_id)