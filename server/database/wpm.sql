CREATE TABLE IF NOT EXISTS `users` (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    mail varchar(255) NOT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
);

INSERT INTO users (name, password, mail, created_at)
VALUES ('John', '123456', 'John@gmail.com', "2023-01-12 06:31:06");

CREATE TABLE IF NOT EXISTS `posts` (
    id int PRIMARY KEY AUTO_INCREMENT,
    title nvarchar(255) NOT NULL,
    descriptions nvarchar(255),
    content nvarchar(10000) NOT NULL,
    user_id int NOT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO posts (title, descriptions, content, user_id, created_at)
VALUES ('Hello World!', 'hello', 'Hello World! Hello World! Hello World! Hello World! Hello World!', 1, "2023-01-12 06:31:06");

CREATE TABLE IF NOT EXISTS `points` (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int NOT NULL,
    post_id int NOT NULL,
    wpm int DEFAULT 0,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);

INSERT INTO points (user_id, post_id, wpm, created_at)
VALUES (1, 1, 40, "2023-01-12 06:31:06");