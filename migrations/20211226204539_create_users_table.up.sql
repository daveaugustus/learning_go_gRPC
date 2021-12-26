CREATE TABLE users (
    id int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, 
    email VARCHAR(100) NOT NULL UNIQUE, 
    password VARCHAR(100) NOT NULL,
    visible BOOLEAN,
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(60) NOT NULL
)