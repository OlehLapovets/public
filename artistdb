DROP DATABASE IF EXISTS Artist;

CREATE DATABASE Artist;

USE Artist;

CREATE TABLE user(
    id INT AUTO_INCREMENT,
    id_telegram INT UNIQUE,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(80) UNIQUE,
    country VARCHAR(30) NOT NULL DEFAULT "Earth",
    language_code VARCHAR(2) NOT NULL DEFAULT 'en',
    registration_date DATE NOT NULL,
    about VARCHAR(200),
    PRIMARY KEY(id)
);

INSERT INTO user(id_telegram, username, email, country, language_code, registration_date, about) VALUES
    (947401843, "Taras", "taras821@gmail.com", "Ukraine", "ua", "2021-08-02", "musician from Lviv, Ukraine"),
    (985388135, "Oleh", "oleh314@gmail.com", "Ukraine", "ua", "2021-08-14", "musician from Kyiv, Ukraine"),
    (987424248, "Olia", "olia49@gmail.com", "Ukraine", "ua", "2021-08-15", "musician from Lviv, Ukraine"),
    (984994724, "Joe", "joejoe@gmail.com", "USA", "en", "2021-08-30", "musician from Washington, USA"),
    (934889194, "Peter", "peterkwachowski@gmail.com", "Poland", "en", "2021-08-30", "musician from Krakow, Poland"),
    (936587326, "Sarah", "sarahbond@gmail.com", "England", "en", "2021-09-01", "musician from London, England"),
    (935253453, "Sam", "samsmith@gmail.com", "England", "en", "2021-09-01", "musician from London, England"),
    (985892376, "Clara", "funlife23@gmail.com", "England", "en", "2021-09-03", "musician from London, England"),
    (932756753, "Jill", "jillvalentine@gmail.com", "England", "en", "2021-09-05", "musician from Birmingham, England"),
    (923789551, "Alek", "alekboldwin@gmail.com", "England", "en", "2021-09-10", "musician from London, England"),
    (957328523, "David", "davidadams@gmail.com", "USA", "en", "2021-09-10", "musician from New Your, USA"),
    (985327512, "Helen", "helenanders@gmail.com", "USA", "en", "2021-09-10", "musician from Utah, USA"),
    (948731355, "Aurora", "chalisi121@gmail.com", "USA", "en", "2021-09-11", "musician from Iowa, USA"),
    (937285751, "Andriy", "andriytarko12@gmail.com", "Ukraine", "ua", "2021-09-14", "musician from Lviv, Ukraine"),
    (993257875, "Ivan", "ivandrago@gmail.com", "Ukraine", "ua", "2021-09-15", "musician from Kharkiv, Ukraine"),
    (953275825, "Ania", "aniakarelina@gmail.com", "Ukraine", "ua", "2021-09-15", "musician from Kharkiv, Ukraine"),
    (958742376, "Pavlo", "pavlotychyna@gmail.com", "Ukraine", "ua", "2021-09-24", "musician from Chernigiv, Ukraine"),
    (953975231, "Grzegorz", "grzegorzbrzeczyszczhykiewicz@gmail.com", "Poland", "en", "2021-09-30", "artist manager from Krakow, Poland"),
    (978252358, "Pavel", "pavel934@gmail.com", "Poland", "en", "2021-09-30", "musician from Wroclaw, Poland"),
    (985728593, "Oleksandr", "oleksa9943@gmail.com", "Ukraine", "ua", "2021-10-01", "artist photographer from Kyiv, Ukraine"),
    (994728531, "Olena", "olena24@gmail.com", "Ukraine", "ua", "2021-10-03", "artist manager from Ternopil, Ukraine"),
    (954925815, "John", "smiler11@gmail.com", "USA", "en", "2021-10-05", "artist manager from Texas, USA"),
    (953728509, "William", "willsmith51@gmail.com", "England", "en", "2021-10-09", "artist manager from Manchester, England"),
    (952389358, "Helga", "helgapataki@gmail.com", "England", "en", "2021-10-11", "artist photographer from London, England"),
    (995738257, "Joe", "joesimons@gmail.com", "England", "en", "2021-10-15", "musician from London, England"),
    (972758923, "Alek", "alekos@gmail.com", "England", "en", "2021-10-19", "musician from Birmingham, England"),
    (989572359, "Aurora", "aurorasim@gmail.com", "Poland", "en", "2021-10-23", "artist photographer from Krakow, Poland"),
    (953287528, "Olena", "olenateliga@gmail.com", "Ukraine", "ua", "2021-10-29", "musician from Poltava, Ukraine"),
    (909953025, "Ivan", "ivansmereka2@gmail.com", "Ukraine", "ua", "2021-10-30", "musician from Poltava, Ukraine"),
    (984936823, "Andriy", "andriysmolko24@gmail.com", "Ukraine", "ua", "2021-11-01", "musician from Dnipro, Ukraine"),
    (958328681, "Mike", "mikethompson32@gmail.com", "USA", "en", "2021-11-05", "musician from Iowa, USA"),
    (942385813, "Sam", "samsmithers@gmail.com", "USA", "en", "2021-11-07", "musician from Illinois, USA"),
    (996236287, "Sergiy", "sergiymanager@gmail.com", "Ukraine", "ua", "2021-11-10", "artist manager from Kyiv, Ukraine"),
    (993523758, "Anjey", "musicianist@gmail.com", "Poland", "en", "2021-11-16", "musician from Krakow, Poland"),
    (988728681, "Sarah", "nameit14@gmail.com", "Poland", "en", "2021-11-16", "musician from Poznan, Poland"),
    (975962516, "Clair", "clairdonald@gmail.com", "England", "en", "2021-11-20", "musician from Chelsea, England"),
    (957818538, "Stan", "stanbush@gmail.com", "USA", "en", "2021-11-21", "musician from Dallas, USA"),
    (985328751, "Oleh", "olehrulez22@gmail.com", "Ukraine", "ua", "2021-11-24", "musician from Kyiv, Ukraine"),
    (985789785, "Oleksandr", "oleksandr95@gmail.com", "Ukraine", "ua", "2021-11-27", "musician from Kyiv, Ukraine"),
    (989895713, "David", "davidbackhamfake@gmail.com", "England", "en", "2021-11-28", "musician from London, England");

CREATE TABLE user_social(
    id INT AUTO_INCREMENT,
    user_id INT,
    social_type VARCHAR(30) NOT NULL,
    social_username VARCHAR(50) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

INSERT INTO user_social(user_id, social_type, social_username)
SELECT id,
       "instagram",
       REPLACE(email, "@gmail.com", "")
FROM user 
WHERE id BETWEEN 1 AND 40;

INSERT INTO user_social(user_id, social_type, social_username)
SELECT id,
       "facebook",
       REPLACE(email, "@gmail.com", "")
FROM user 
WHERE id BETWEEN 1 AND 40;

CREATE TABLE manager (
    id INT AUTO_INCREMENT,
    user_id INT,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

INSERT INTO manager(user_id)
SELECT id
FROM user
WHERE about LIKE "%manager%";

CREATE TABLE musician (
    id INT AUTO_INCREMENT,
    user_id INT,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

INSERT INTO musician(user_id)
SELECT id
FROM user
WHERE about LIKE "musician%";

CREATE TABLE photographer (
    id INT AUTO_INCREMENT,
    user_id INT,
    portfolio_url VARCHAR(50) DEFAULT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

INSERT INTO photographer(user_id)
SELECT id
FROM user
WHERE about LIKE "%photographer%";

CREATE TABLE band (
    id INT AUTO_INCREMENT,
    manager_id INT,
    band_name VARCHAR(30),
    style VARCHAR(30),
    PRIMARY KEY(id),
    FOREIGN KEY(manager_id) REFERENCES manager(id)
);

INSERT INTO band(manager_id, band_name, style) VALUES
    (1, "Scarlett dreams", "dream-pop"), #Poland
    (2, "Atomic kittens", "dream-pop"), #Ukraine
    (3, "Yankee doodles", "punk-rock"), #USA
    (4, "Tea party", "brit-pop"), #England
    (5, "Fogsters", "reggie"), #Ukraine
    (NULL, "Ghosthound", "gothic rock"), #England
    (NULL, "Skittles", "brit-pop"), #England
    (NULL, "Avengers", "hard-core punk"), #USA
    (NULL, "Bubble bee", "pop"); #Ukraine

CREATE TABLE band_selector (
    id INT AUTO_INCREMENT,
    musician_id INT,
    band_id INT,
    PRIMARY KEY(id),
    FOREIGN KEY(band_id) REFERENCES band(id)
);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       1 #Scarlett dreams / Poland
FROM user
WHERE id IN (5, 19, 34, 35);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       2 #Atomic kittens / Ukraine
FROM user
WHERE id IN (1, 2, 3, 14);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       3 #Yankee doodles / USA
FROM user
WHERE id IN (4, 11, 12, 13);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       4 #Tea party / England
FROM user
WHERE id IN (6, 7, 8, 9);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       5 #Fogsters / Ukraine
FROM user
WHERE id IN (15, 16, 17, 28);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       6 #Ghosthound / England
FROM user
WHERE id IN (10, 25, 26);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       7 #Skittles / England
FROM user
WHERE id IN (36, 40);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       8 #Avengers / USA
FROM user
WHERE id IN (31, 32, 37);

INSERT INTO band_selector(musician_id, band_id)
SELECT id,
       9 #Bubble bee / Ukraine
FROM user
WHERE id IN (29, 30, 38, 39);

CREATE TABLE band_social (
    id INT AUTO_INCREMENT,
    band_id INT,
    social_type VARCHAR(30) NOT NULL,
    social_bandname VARCHAR(50) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(band_id) REFERENCES band(id)
);

INSERT INTO band_social(band_id, social_type, social_bandname)
SELECT id,
       "instagram",
       LOWER(REPLACE(band_name, ' ', ''))
FROM band
WHERE id BETWEEN 1 AND 9;

INSERT INTO band_social(band_id, social_type, social_bandname)
SELECT id,
       "facebook",
       LOWER(REPLACE(band_name, ' ', ''))
FROM band
WHERE id BETWEEN 1 AND 9;

# Показує імена користувачів, назви гуртів у яких вони виступають, їх країни та опис,
# у проміжку реєстрації між 15 серпня та 16 листопада 2021 року.
CREATE VIEW musician_user_details AS
SELECT u.id,
       u.username,
       b.band_name,
       u.country,
       u.about,
       u.registration_date
FROM user AS u 
    INNER JOIN musician AS m
    ON u.id = m.user_id
    INNER JOIN band_selector AS bs
    ON m.id = bs.musician_id
    INNER JOIN band AS b
    ON b.id = bs.band_id
WHERE registration_date BETWEEN "2021-08-15" AND "2021-11-16"
ORDER BY id;

SELECT *
FROM musician_user_details;

# Показує країни та сумарну кількість виконавців з цих країн.
CREATE VIEW most_active_cities AS
SELECT country AS "Країна",
       COUNT(country) AS "Кількість виконавців"
FROM user
GROUP BY country;

SELECT *
FROM most_active_cities;
