-- Cleanup
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS score;
DROP TABLE IF EXISTS friendship;

CREATE TABLE user (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    email text UNIQUE,
    nickname text UNIQUE,
    password text NOT NULL,
    fullname text
);

CREATE TABLE score (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id integer NOT NULL,
    score integer NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE friendship (
    id1 integer NOT NULL,
    id2 integer NOT NULL,
    FOREIGN KEY (id1) REFERENCES user(id)
    FOREIGN KEY (id2) REFERENCES user(id)
);