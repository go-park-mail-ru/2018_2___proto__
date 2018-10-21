-- Cleanup
DROP TABLE IF EXISTS session;
DROP TABLE IF EXISTS answer;
DROP TABLE IF EXISTS question;
DROP TABLE IF EXISTS score;
DROP TABLE IF EXISTS friendship;
DROP TABLE IF EXISTS player;

-- Creation
CREATE TABLE player (
    id       serial NOT NULL PRIMARY KEY,
    avatar   text DEFAULT 'default.png',
    fullname text DEFAULT '',
    password text NOT NULL,
    email    text UNIQUE,
    nickname text UNIQUE
);

CREATE TABLE user_session (
    id           serial NOT NULL PRIMARY KEY,
    token        text NOT NULL UNIQUE,
    player_id    integer NOT NULL,
    expired_date integer NOT NULL,
    FOREIGN KEY (player_id) REFERENCES player(id)
);

CREATE TABLE score (
    id        serial NOT NULL PRIMARY KEY,
    score     integer NOT NULL,
    player_id integer NOT NULL,
    FOREIGN KEY (player_id) REFERENCES player(id)
);

CREATE TABLE friendship (
    id  serial NOT NULL PRIMARY KEY,
    id1 integer NOT NULL,
    id2 integer NOT NULL,
    FOREIGN KEY (id1) REFERENCES player(id),
    FOREIGN KEY (id2) REFERENCES player(id)
);

CREATE TABLE question (
    id            serial NOT NULL PRIMARY KEY,
    question_text text NOT NULL UNIQUE,
    question_type integer NOT NULL
);

CREATE TABLE answer (
    id          serial NOT NULL PRIMARY KEY,
    question_id integer NOT NULL,
    answer_text text NOT NULL,
    is_correct  boolean,
    FOREIGN KEY (question_id) REFERENCES question(id)
);