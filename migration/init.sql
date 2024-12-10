CREATE SCHEMA farm AUTHORIZATION postgres;

CREATE TABLE farm.user
(
  id BIGINT PRIMARY KEY,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64),
  username VARCHAR(32) UNIQUE,
  balance FLOAT DEFAULT 0 
);

CREATE TYPE farm.tickers AS ENUM ('BTC', 'TON', 'ETH', 'DOGE', 'SOL', 'NEAR');

CREATE TABLE farm.pot
(
  id SERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES farm.user(id) NOT NULL,
  coin tickers,
  plant_time TIMESTAMP,
  harvest_time TIMESTAMP,
  plant_price FLOAT,
  harvest_price FLOAT
);

