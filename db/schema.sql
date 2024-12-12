CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cards (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    picture VARCHAR(64) NOT NULL,
    card_name VARCHAR(64) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_selected (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    card_id BIGINT NOT NULL,
    attribute ENUM(
        'red',
        'blue',
        'green',
        'kamekame',
        'nankuru',
        'random'
    ) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (card_id) REFERENCES cards(id)
);

CREATE TABLE IF NOT EXISTS rooms (
    room_id BIGINT NOT NULL AUTO_INCREMENT,
    user1_id BIGINT NOT NULL,
    user2_id BIGINT NOT NULL,
    PRIMARY KEY (room_id),
    FOREIGN KEY (user1_id) REFERENCES users(id),
    FOREIGN KEY (user2_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS battles (
    battle_id BIGINT NOT NULL AUTO_INCREMENT,
    room_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    red_card_id BIGINT,
    blue_card_id BIGINT,
    green_card_id BIGINT,
    kamekame_card_id BIGINT,
    nankuru_card_id BIGINT,
    random_card_id BIGINT,
    random_attribute ENUM('red', 'blue', 'green') NOT NULL,
    shogun_id BIGINT NOT NULL DEFAULT 1,
    hp INT NOT NULL,
    result ENUM('win', 'lose', 'draw', 'pending') NOT NULL DEFAULT 'pending',
    PRIMARY KEY (battle_id),
    FOREIGN KEY (room_id) REFERENCES rooms(room_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (red_card_id) REFERENCES cards(id),
    FOREIGN KEY (blue_card_id) REFERENCES cards(id),
    FOREIGN KEY (green_card_id) REFERENCES cards(id),
    FOREIGN KEY (kamekame_card_id) REFERENCES cards(id),
    FOREIGN KEY (nankuru_card_id) REFERENCES cards(id),
    FOREIGN KEY (random_card_id) REFERENCES cards(id)
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
    user_selected (user_id, card_id, attribute)
VALUES
    (1, 2, 'red'),
    (1, 3, 'blue'),
    (1, 5, 'green'),
    (1, 7, 'kamekame'),
    (1, 9, 'nankuru'),
    (1, 10, 'random'),
    (2, 11, 'red'),
    (2, 12, 'blue'),
    (2, 14, 'green'),
    (2, 16, 'kamekame'),
    (2, 18, 'nankuru'),
    (2, 19, 'random'),
    (3, 20, 'red'),
    (3, 21, 'blue'),
    (3, 23, 'green'),
    (3, 25, 'kamekame'),
    (3, 27, 'nankuru'),
    (3, 29, 'random');

INSERT INTO
    rooms(user1_id, user2_id)
VALUES
    (1, 2),
    (1, 2);

INSERT INTO
    battles (
        room_id,
        user_id,
        red_card_id,
        blue_card_id,
        green_card_id,
        kamekame_card_id,
        nankuru_card_id,
        random_card_id,
        random_attribute,
        hp,
        result
    )
VALUES
    (1, 1, 2, 3, 4, 5, 6, 7, 'red', 3, 'win'),
    (1, 2, 13, 14, 15, 16, 17, 18, 'green', 0, 'lose'),
    (2, 1, 4, 5, 6, 7, 8, 9, 'blue', 1, 'draw'),
    (2, 2, 14, 15, 16, 17, 18, 19, 'red', 1, 'draw');
