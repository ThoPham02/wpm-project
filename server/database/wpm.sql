CREATE TABLE IF NOT EXISTS `users` (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    mail varchar(255) NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);

INSERT INTO users (name, password, mail)
VALUES ('John', '123456', 'John@gmail.com');

CREATE TABLE IF NOT EXISTS `languages` (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `posts` (
    id int PRIMARY KEY AUTO_INCREMENT,
    title nvarchar(255) NOT NULL,
    discription nvarchar(255),
    content nvarchar(10000) NOT NULL,
    user_id int NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS `points` (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int NOT NULL,
    post_id int NOT NULL,
    wpm int DEFAULT 0,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);