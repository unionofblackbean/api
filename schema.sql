CREATE TABLE users
(
    id                    SERIAL PRIMARY KEY,
    username              TEXT,
    email                 TEXT,
    password_hash_encoded TEXT
);

CREATE TABLE sessions
(
    id       TEXT,
    username TEXT,
    ip       TEXT
);
