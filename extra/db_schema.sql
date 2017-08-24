CREATE TABLE users (
    id serial PRIMARY KEY,
    login text NOT NULL UNIQUE,
    password text NOT NULL,
    salt text
);

CREATE TABLE rooms (
    id serial PRIMARY KEY,
    title text NOT NULL,
    url text NOT NULL
);

CREATE TABLE access_keys (
    user_id integer REFERENCES users ON DELETE CASCADE,
    key text NOT NULL
    -- date_end date,
);

CREATE INDEX index_access_keys ON access_keys (key);
