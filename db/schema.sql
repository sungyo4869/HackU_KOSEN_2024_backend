CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (username) VALUES ('user1'), ('user2'), ('user3');
