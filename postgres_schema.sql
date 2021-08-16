CREATE TABLE users
(
    user_id                    SERIAL,
    user_username              TEXT,
    user_email                 TEXT,
    user_password_hash_encoded TEXT,
    PRIMARY KEY (user_id)
);

CREATE TABLE sessions
(
    session_id            TEXT,
    session_ip            TEXT,
    session_creation_time TIMESTAMP,
    session_expire_time   BIGINT,
    user_username         TEXT
);

ALTER TABLE sessions
    ADD FOREIGN KEY (user_username) REFERENCES users (user_username);
