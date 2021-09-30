CREATE TABLE products
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    model   VARCHAR(30) NOT NULL,
    company VARCHAR(30) NOT NULL,
    price   INT         NOT NULL
);

INSERT INTO db.products (model, company, price)
VALUES ('iPhone X', 'Apple', 74000),
       ('Pixel 2', 'Google', 62000),
       ('Galaxy S9', 'Samsung', 65000);
