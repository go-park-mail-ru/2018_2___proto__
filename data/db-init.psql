-- Cleanup
DROP TABLE IF EXISTS player;
DROP TABLE IF EXISTS score;
DROP TABLE IF EXISTS friendship;

CREATE TABLE player (
    id SERIAL NOT NULL PRIMARY KEY,
    email text UNIQUE,
    nickname text UNIQUE,
    password text NOT NULL,
    fullname text
);

-- CREATE TYPE player AS (
--     Id
-- 	Nickname 
-- 	Password 
-- 	Email 
-- 	Fullname 
-- )

-- CREATE TYPE session AS (
--     id text,
--     ttl INTEGER,
--     user player,
-- )

-- CREATE TABLE user_session (
--     id serial not null PRIMARY key,
--     token text UNIQUE,
    
-- )

CREATE TABLE score (
    id serial NOT NULL PRIMARY KEY,
    player_id integer NOT NULL,
    score integer NOT NULL,
    FOREIGN KEY (player_id) REFERENCES player(id)
);

CREATE TABLE friendship (
    id1 integer NOT NULL,
    id2 integer NOT NULL,
    FOREIGN KEY (id1) REFERENCES player(id),
    FOREIGN KEY (id2) REFERENCES player(id)
);