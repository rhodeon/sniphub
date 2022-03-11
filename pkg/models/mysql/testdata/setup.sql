-- snips
DROP TABLE IF EXISTS snips;
CREATE TABLE snips
(
    id      INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user    VARCHAR(255) NOT NULL DEFAULT '',
    title   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL
);

CREATE INDEX idx_snip_created on snips (created);

-- users
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id              INT                  NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username        VARCHAR(255)         NOT NULL,
    email           VARCHAR(255)         NOT NULL,
    hashed_password CHAR(60)             NOT NULL,
    created         DATETIME             NOT NULL,
    active          TINYINT(1) DEFAULT 1 NOT NULL,

    CONSTRAINT users_uc_username UNIQUE (username),
    CONSTRAINT users_uc_email UNIQUE (email)
);

INSERT INTO users(username, email, hashed_password, created)
VALUES ('rhodeon',
        'rhodeon@mail.com',
        '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
        '2022-03-08 23:00:00')