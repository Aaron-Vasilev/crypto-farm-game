CREATE SCHEMA farm AUTHORIZATION postgres;

CREATE TABLE farm.user
(
  id BIGINT PRIMARY KEY,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64),
  username VARCHAR(32) UNIQUE,
  usd FLOAT DEFAULT 0 
);

CREATE TABLE farm.pot
(
  id SERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES farm.user(id) NOT NULL,
  position INT NOT NULL
);

CREATE TABLE farm.pot_count
(
  user_id BIGINT REFERENCES farm.user(id) PRIMARY KEY,
  count INT DEFAULT 1
);


CREATE TYPE tickers AS ENUM ('BTC', 'TON', 'ETH', 'DOGE', 'SOL', 'NEAR');

CREATE TABLE farm.plant
(
  id SERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES farm.user(id) NOT NULL,
  pot_id INT REFERENCES farm.pot(id) NOT NULL,
  coin tickers NOT NULL,
  enter_date TIMESTAMP DEFAULT NOW(),
  exit_date TIMESTAMP,
  enter_price FLOAT NOT NULL,
  exit_price FLOAT,
  profit FLOAT
);
