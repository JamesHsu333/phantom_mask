DROP TABLE IF EXISTS pharmacies CASCADE;
DROP TABLE IF EXISTS masks CASCADE;
DROP TABLE IF EXISTS mask_prices CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS purchase_histories CASCADE;

CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE pharmacies
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(512) UNIQUE        NOT NULL CHECK ( name <> '' ),
    opening_hours int4multirange,
    opening_hours_description VARCHAR(512)   NOT NULL CHECK ( opening_hours_description <> '' ),
    cash_balance  NUMERIC(5,2)
);

CREATE TABLE masks
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(512) UNIQUE         NOT NULL CHECK ( name <> '' )
);

CREATE TABLE mask_prices
(
    mask_id      INT NOT NULL REFERENCES     masks (id) ON DELETE CASCADE,
    pharmacy_id  INT NOT NULL REFERENCES     pharmacies (id) ON DELETE CASCADE,
    price        NUMERIC(5,2)                NOT NULL,
    PRIMARY KEY (mask_id, pharmacy_id)
);

CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(512)               NOT NULL CHECK ( name <> '' ),
    cash_balance  NUMERIC(5,2) 
);

CREATE TABLE purchase_histories
(
    id                 SERIAL PRIMARY KEY,
    user_id            INT NOT NULL REFERENCES  users (id) ON DELETE CASCADE,
    pharmacy_id        INT NOT NULL REFERENCES  pharmacies (id) ON DELETE CASCADE,
    mask_id            INT NOT NULL REFERENCES  masks (id) ON DELETE CASCADE,
    transaction_amount NUMERIC(5,2)             NOT NULL,
    transaction_date   TIMESTAMP WITH TIME ZONE NOT NULL
);
