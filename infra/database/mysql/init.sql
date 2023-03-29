DROP DATABASE IF EXISTS task_cards;    

CREATE DATABASE task_cards;

USE task_cards;

CREATE TABLE roles(
    id INT PRIMARY KEY AUTO_INCREMENT,
    role_ ENUM('technician', 'manager') NOT NULL DEFAULT 'technician'
);

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  userpass VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  role_id INT NOT NULL,
  FOREIGN KEY (role_id) REFERENCES roles(id) 
);

CREATE TABLE cards (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  summary VARCHAR(2500),
  due_date VARCHAR(10),
  card_status ENUM('to do', 'in progress', 'done') NOT NULL DEFAULT 'to do',
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO roles (role_) VALUES ('technician');
INSERT INTO roles (role_) VALUES ('manager');
