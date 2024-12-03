CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(255) NOT NULL,
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

CREATE TABLE IF NOT EXISTS user_selected (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    card_id INT NOT NULL,
    attribute ENUM('red', 'blue', 'green', 'kamekame', 'nankuru') NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (card_id) REFERENCES cards(id)
);

CREATE TABLE IF NOT EXISTS rooms (
    room_id INT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    PRIMARY KEY (room_id),
    FOREIGN KEY (user1_id) REFERENCES users(id),
    FOREIGN KEY (user2_id) REFERENCES users(id)
);

INSERT INTO
    users (username, password)
VALUES
    ('user1', 1),
    ('user2', 2),
    ('user3', 3);

INSERT INTO
    cards (user_id, picture, card_name)
VALUES
(1, 'card1.jpg', 'カード1'),
(1, 'card2.jpg', 'カード2'),
(1, 'card3.jpg', 'カード3'),
(1, 'card4.jpg', 'カード4'),
(1, 'card5.jpg', 'カード5'),
(1, 'card6.jpg', 'カード6'),
(1, 'card7.jpg', 'カード7'),
(1, 'card8.jpg', 'カード8'),
(1, 'card9.jpg', 'カード9'),
(1, 'card10.jpg', 'カード10'),

(2, 'card11.jpg', 'カード11'),
(2, 'card12.jpg', 'カード12'),
(2, 'card13.jpg', 'カード13'),
(2, 'card14.jpg', 'カード14'),
(2, 'card15.jpg', 'カード15'),
(2, 'card16.jpg', 'カード16'),
(2, 'card17.jpg', 'カード17'),
(2, 'card18.jpg', 'カード18'),
(2, 'card19.jpg', 'カード19'),
(2, 'card20.jpg', 'カード20'),

(3, 'card21.jpg', 'カード21'),
(3, 'card22.jpg', 'カード22'),
(3, 'card23.jpg', 'カード23'),
(3, 'card24.jpg', 'カード24'),
(3, 'card25.jpg', 'カード25'),
(3, 'card26.jpg', 'カード26'),
(3, 'card27.jpg', 'カード27'),
(3, 'card28.jpg', 'カード28'),
(3, 'card29.jpg', 'カード29'),
(3, 'card30.jpg', 'カード30');

INSERT INTO
    user_selected (user_id, card_id)
VALUES
(1,2),
(1,3),
(1,5),
(1,7),
(1,9),

(2, 11),
(2, 12),
(2, 14),
(2, 16),
(2, 18),

(3, 20),
(3, 21),
(3, 23),
(3, 25),
(3, 27);
