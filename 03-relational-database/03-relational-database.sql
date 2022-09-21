/*
    Nama: Rahman Insani
    Tugas Relational Database - PostgreSQL

    Rabu, 21 September 2022
*/

-- Create Database
CREATE DATABASE efishery_db_task;

-- Create Tables
CREATE TABLE customers
(
    id            INT      NOT NULL PRIMARY KEY,
    customer_name CHAR(50) NOT NULL
);

CREATE TABLE products
(
    id   INT      NOT NULL PRIMARY KEY,
    name CHAR(50) NOT NULL
);

CREATE TABLE orders
(
    order_id    INT   NOT NULL PRIMARY KEY,
    customer_id INT   NOT NULL,
    product_id  INT   NOT NULL,
    order_date  DATE  NOT NULL,
    total       FLOAT NOT NULL,
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers (id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id)
);

-- Insert Data
INSERT INTO customers (id, customer_name)
VALUES (1, 'Alice'),
       (2, 'Bob'),
       (3, 'Charlie')
;

INSERT INTO products (id, name)
VALUES (1, 'Gourami'),
       (2, 'Catfish'),
       (3, 'Goldfish')
;

INSERT INTO orders (order_id, customer_id, product_id, order_date, total)
VALUES (1, 1, 1, '2022-09-01', 100000),
       (2, 2, 2, '2022-09-02', 200000),
       (3, 3, 3, '2022-09-03', 300000)
;

-- Update Data
UPDATE customers
SET customer_name = 'David'
WHERE id = 3;

UPDATE products
SET name = 'Shrimp'
WHERE id = 1;

UPDATE orders
SET total = 650000
WHERE order_id = 2;

-- Delete Data
DELETE
FROM orders
WHERE order_id = 3;

DELETE
FROM customers
WHERE id = 3;

DELETE
FROM products
WHERE id = 3;


-- Show Data
SELECT id, customer_name
FROM customers;

SELECT id, name
FROM products;

SELECT order_id, customer_id, product_id, order_date, total
FROM orders;

-- Show Data with Where Clause
SELECT customer_name
FROM customers
WHERE id = 2;

SELECT name
FROM products
WHERE id = 1;

SELECT order_date, total
FROM orders
WHERE order_id = 2
  AND customer_id = 2;

SELECT order_date, total
FROM orders
WHERE order_id = 2
   OR order_date = '2022-09-01';

-- Inner Join
SELECT c.customer_name AS "Customer Name",
       p.name          AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
;

SELECT c.customer_name AS "Customer Name",
       p.name          AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
WHERE order_date = '2022-09-02'
;

SELECT c.customer_name AS "Customer Name",
       p.name          AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
ORDER BY o.total DESC
LIMIT 1;

-- Aggregate Function
SELECT COUNT(total) AS "Order Count"
FROM orders;

SELECT SUM(total) AS "Total Sales"
FROM orders;

SELECT AVG(total) AS "Average Total Sales"
FROM orders;

SELECT MIN(total) AS "Min Total Sales"
FROM orders;

SELECT MAX(total) AS "Max Total Sales"
FROM orders;