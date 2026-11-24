CREATE DATABASE IF NOT EXISTS aws_devops_lab;

CREATE TABLE IF NOT EXISTS aws_devops_lab.items (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

INSERT IGNORE INTO aws_devops_lab.items (id, name) VALUES
    (1, 'Docker setup'),
    (2, 'Go backend'),
    (3, 'MySQL connection');

CREATE USER IF NOT EXISTS 'mysql'@'%' IDENTIFIED BY 'mysql';
GRANT SELECT ON aws_devops_lab.items TO 'mysql'@'%';
