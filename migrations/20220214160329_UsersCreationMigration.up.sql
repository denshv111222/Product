

CREATE TABLE IF NOT EXISTS users(
id_user BIGSERIAL NOT NULL PRIMARY KEY,
login_user VARCHAR NOT NULL UNIQUE,
password_user VARCHAR NOT NULL
);