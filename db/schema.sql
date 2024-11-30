CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cards (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    picture VARCHAR(64) NOT NULL,
    card_name VARCHAR(64) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO
    users (username)
VALUES
    ('user1'),
    ('user2'),
    ('user3');

INSERT INTO
    cards (user_id, picture, card_name)
VALUES
    (1, 'card1.jpg', 'カード1'),
    (2, 'card2.jpg', 'カード2'),
    (3, 'card3.jpg', 'カード3'),
    (1, 'card4.jpg', 'カード4'),
    (2, 'card5.jpg', 'カード5'),
    (3, 'card6.jpg', 'カード6');
