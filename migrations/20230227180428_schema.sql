Add migration script here
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message VARCHAR NOT NULL
);

INSERT INTO messages (message) VALUES ('Hello world!');