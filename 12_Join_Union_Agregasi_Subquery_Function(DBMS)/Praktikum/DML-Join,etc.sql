DROP DATABASE IF EXISTS alta_online_shop;

CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE users (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	fullname VARCHAR(255) NOT NULL,
	address VARCHAR(255) NOT NULL,
	birthDate DATE,
	userStatus ENUM('Active', 'Deactive'),
	gender ENUM('Female', 'Male'),
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	typeID INT NOT NULL,
	operatorID INT NOT NULL,
	code INT(50) NOT NULL,
	name VARCHAR(255) NOT NULL,
	status SMALLINT,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE productTypes (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE productDesc (
	id INT PRIMARY KEY,
	description TEXT NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE operators (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	fullname VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	userId INT NOT NULL,
	paymentMethodId INT NOT NULL,
	status ENUM('Success', 'Pending', 'Cancelled'),
	totalQuantity INT NOT NULL,
	totalPrice DECIMAL(25,2),
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactionDetails (
	transactionID INT NOT NULL,
	productID INT NOT NULL,
	status ENUM('Success', 'Pending', 'Cancelled'),
	quantity INT NOT NULL,
	price INT NOT NULL,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE paymentMethods (
	id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	status SMALLINT,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updateAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE transactions ADD FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE transactions ADD FOREIGN KEY (paymentMethodId) REFERENCES paymentMethods(id) ON DELETE CASCADE;
ALTER TABLE transactionDetails ADD FOREIGN KEY (transactionID) REFERENCES transactions(id) ON DELETE CASCADE;
ALTER TABLE transactionDetails ADD FOREIGN KEY (productID) REFERENCES products(id) ON DELETE CASCADE;
ALTER TABLE productDesc ADD FOREIGN KEY (id) REFERENCES products(id) ON DELETE CASCADE;
ALTER TABLE products ADD FOREIGN KEY (typeID) REFERENCES productTypes(id) ON DELETE CASCADE;
ALTER TABLE products ADD FOREIGN KEY (operatorID) REFERENCES operators(id) ON DELETE CASCADE;

--DML--
-- 1.A --
INSERT INTO operators (fullname) VALUES ("Operator1"),("Operator2"),("Operator3"),("Operator4"),("Operator5");

-- 1.B --
INSERT INTO productTypes (name) VALUES ("Sepatu"),("Celana"),("Baju");

-- 1.C --
INSERT INTO products (typeID, operatorID, code, name, status) VALUES
	(1, 3, 70001, "Nike Air Force 1", 10),
	(1, 3, 70002, "Adidas Superstar", 12);

-- 1.D --
INSERT INTO products (typeID, operatorID, code, name, status) VALUES
	(2, 1, 80001, "Celana Cargo", 12),
	(2, 1, 80002, "Celana Chinos", 24),
	(2, 1, 80003, "Celana Jeans", 12);

-- 1.E --
INSERT INTO products (typeID, operatorID, code, name, status) VALUES
	(3, 1, 90001, "Baju Lengan Pendek", 24),
	(3, 1, 90002, "Baju Lengan Panjang", 12),
	(3, 1, 90003, "Baju Oversize", 12);

-- 1.F --
INSERT INTO productdesc (id, description) VALUES 
	(1, "Description Nike Air Force 1"),
	(2, "Description Adidas Superstar"),
	(3, "Description Celana Cargo"),
	(4, "Description Celana Chinos"),
	(5, "Description Celana Jeans"),
	(6, "Description Baju Lengan Pendek"),
	(7, "Description Baju Lengan Panjang"),
	(8, "Description Baju Oversize");

-- 1.G ---
INSERT INTO paymentMethods (name, status) VALUES
	("Bank Transfer", 1),
	("GOPAY", 2),
	("OVO", 3);

-- 1.H --
INSERT INTO users (fullname, address, birthDate, userStatus, gender) VALUES
	("Sidik", "Jl. Pahlawan 1", "2002-7-01", "Active", "Male"),
	("Anggi", "Jl. Pahlawan 2", "2002-6-02", "Active", "Female"),
	("Mike", "Jl. Pahlawan 3", "2002-8-07", "Active", "Male"),
	("Jessica", "Jl. Pahlawan 6", "2002-3-02", "Active", "Female"),
	("Dylee", "Jl. Pahlawan 18", "2002-9-03", "Deactive", "Male");

-- 1.I --
INSERT INTO transactions (userId, paymentMethodId, status, totalQuantity, totalPrice) VALUES
	(1, 1, "Success", 3, 500000),
	(1, 2, "Cancelled", 4, 600000),
	(1, 2, "Success", 3, 450000),
	(2, 3, "Pending", 3, 400000),
	(2, 2, "Success", 4, 500000),
	(2, 1, "Success", 3, 300000),
	(3, 1, "Success", 5, 500000),
	(3, 1, "Pending", 3, 300000),
	(3, 3, "Cancelled", 4, 400000),
	(4, 2, "Pending", 3, 450000),
	(4, 2, "Cancelled", 3, 450000),
	(4, 1, "Pending", 3, 450000),
	(5, 1, "Success", 3, 500000),
	(5, 2, "Cancelled", 4, 600000),
	(5, 2, "Pending", 3, 450000);

-- 1.J --
INSERT INTO transactionDetails (transactionID, productID, status, quantity, price) VALUES
	(1, 3, "Success", 1, 150000),
	(1, 5, "Success", 1, 150000),
	(1, 1, "Success", 1, 200000),
	(2, 3, "Cancelled", 1, 150000),
	(2, 5, "Cancelled", 1, 150000),
	(2, 4, "Cancelled", 2, 300000),
	(3, 3, "Success", 1, 150000),
	(3, 4, "Success", 1, 150000),
	(3, 5, "Success", 1, 150000),
	(4, 3, "Pending", 1, 150000),
	(4, 4, "Pending", 1, 150000),
	(4, 6, "Pending", 1, 100000),
	(5, 6, "Success", 2, 200000),
	(5, 3, "Success", 1, 150000),
	(5, 4, "Success", 1, 150000),
	(6, 6, "Success", 1, 100000),
	(6, 7, "Success", 1, 100000),
	(6, 8, "Success", 1, 100000),
	(7, 7, "Success", 3, 100000),
	(7, 6, "Success", 1, 100000),
	(7, 8, "Success", 1, 100000),
	(8, 6, "Pending", 1, 100000),
	(8, 7, "Pending", 1, 100000),
	(8, 8, "Pending", 1, 100000),
	(9, 8, "Cancelled", 1, 100000),
	(9, 6, "Cancelled", 2, 200000),
	(9, 7, "Cancelled", 1, 100000),
	(10, 7, "Pending", 1, 100000),
	(10, 1, "Pending", 1, 200000),
	(10, 3, "Pending", 1, 150000),
	(11, 7, "Cancelled", 1, 100000),
	(11, 1, "Cancelled", 1, 200000),
	(11, 3, "Cancelled", 1, 150000),
	(12, 8, "Pending", 1, 100000),
	(12, 1, "Pending", 1, 200000),
	(12, 5, "Pending", 1, 150000),
	(13, 1, "Success", 1, 200000),
	(13, 3, "Success", 1, 150000),
	(13, 5, "Success", 1, 150000),
	(14, 1, "Cancelled", 2, 400000),
	(14, 8, "Cancelled", 1, 100000),
	(14, 6, "Cancelled", 1, 100000),
	(15, 3, "Pending", 1, 150000),
	(15, 7, "Pending", 1, 100000),
	(15, 1, "Pending", 1, 200000);

-- 2.A --
SELECT * FROM users WHERE gender = "Male";

-- 2.B --
SELECT * FROM products WHERE id = 3;

-- 2.C --
SELECT * FROM users WHERE createdAt > DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY) AND fullname LIKE '%a%';

-- 2.D --
SELECT gender, COUNT(gender) AS count FROM users WHERE gender = "Female";

-- 2.E --
SELECT * FROM users ORDER BY fullname ASC;

-- 2.F --
SELECT * FROM products WHERE id LIMIT 5;

-- 3.A --
SELECT * FROM products;
UPDATE products SET name = "product dummy" WHERE id = 1;
SELECT * FROM products;

-- 3.B --
SELECT * FROM transactionDetails WHERE productID = 1;
UPDATE transactionDetails SET quantity = 3 WHERE productID = 1;
SELECT * FROM transactionDetails WHERE productID = 1;

-- 4.A --
SELECT * FROM products;
DELETE FROM products WHERE id = 1;
SELECT * FROM products;


-- 4.B --
SELECT * FROM products;
DELETE FROM products WHERE typeID = 1;
SELECT * FROM products;

-- Join, Union, Sub query, Function --
-- 1 --
SELECT u.fullname,t.* FROM users u INNER JOIN transactions t ON u.id = t.userId WHERE u.id = 1 OR u.id = 2;

-- 2 --
SELECT u.id, u.fullname, SUM(t.totalPrice) as jumlah_harga_transactions FROM users u INNER JOIN transactions t ON u.id = t.userId WHERE u.id = 1;

-- 3 --
SELECT p.typeID, SUM(td.price) as total_transaksi FROM transactionDetails td INNER JOIN products p ON p.id=td.productID WHERE p.typeID = 2;

-- 4 --
SELECT pt.name, p.* FROM productTypes pt RIGHT JOIN products p ON pt.id=p.typeID; 

-- 5 --
SELECT t.*, pt.name, u.fullname FROM users u RIGHT JOIN transactions t ON u.id=t.userId RIGHT JOIN transactionDetails td ON t.id=td.transactionID RIGHT JOIN products p ON td.productID=p.id RIGHT JOIN productTypes pt ON p.typeID=pt.id;

-- 6 --
DELIMITER //
CREATE TRIGGER afterDeleteTransaction
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transactionDetails WHERE transactionID = OLD.id;
END;
//
DELIMITER ;
-- Check --
SELECT * FROM transactions;
SELECT * FROM transactionDetails;
DELETE FROM transactions where id=1;
SELECT * FROM transactions;
SELECT * FROM transactionDetails;

-- 7 --
DELIMITER //

CREATE TRIGGER update_total_qty
AFTER DELETE ON transactionDetails
FOR EACH ROW
BEGIN
  
    DECLARE total_quantity INT(11);

    SELECT totalQuantity
    INTO total_quantity
    FROM transactions
    WHERE id = OLD.transactionID;

    UPDATE transactions
    SET totalQuantity = total_quantity - OLD.quantity
    WHERE id = OLD.transactionID;

END;
//
DELIMITER ;
-- Check --
SELECT * FROM transactions;
SELECT * FROM transactionDetails;
DELETE FROM transactionDetails WHERE transactionID = 2 and productID = 3;
SELECT * FROM transactions;
SELECT * FROM transactionDetails;

-- 8 --
SELECT * FROM products WHERE id NOT IN (select productID from transactionDetails)
