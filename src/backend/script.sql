DROP TABLE IF EXISTS raffles;
DROP TABLE IF EXISTS gambles;

-- Create a table with a primary key using the sequence
CREATE TABLE IF NOT EXISTS raffles (
    raffle_id SERIAL PRIMARY KEY,
    active BOOLEAN DEFAULT true,
    prize INTEGER DEFAULT 500000,
    numbers INTEGER[]
);

INSERT INTO raffles DEFAULT VALUES;

CREATE TABLE IF NOT EXISTS gambles (
    gamble_id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    cpf CHAR(11) NOT NULL,
    numbers INTEGER[] NOT NULL,
    raffle_id INTEGER REFERENCES raffles(raffle_id)
);

CREATE SEQUENCE gamble_id_sequence START WITH 1000;

ALTER TABLE gambles ALTER COLUMN gamble_id SET DEFAULT NEXTVAL('gamble_id_sequence');

CREATE TABLE IF NOT EXISTS winners (
    winner_id SERIAL PRIMARY KEY,
    prize INTEGER NOT NULL,
    gamble_id INTEGER REFERENCES gambles(gamble_id)
);

CREATE TABLE IF NOT EXISTS accumulated_prize (
    prize_id SERIAL PRIMARY KEY,
    prize_amount INTEGER DEFAULT 0
);

INSERT INTO accumulated_prize DEFAULT VALUES;