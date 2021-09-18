CREATE USER 'root'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

CREATE DATABASE db;
USE db;

CREATE TABLE IF NOT EXISTS center_info(
    id INT,
    name VARCHAR(40),
    address TEXT,
    state_name VARCHAR(16),
    district_name VARCHAR(16),
    pincode INT,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS appointment_session(
    id BINARY(36),
    center_idfk INT, 
    date VARCHAR(16),
    available_capacity INT,
    min_age_limit INT,
    vaccine  VARCHAR(16),
    available_capacity_dose1 INT,
    available_capacity_dose2 INT,
    PRIMARY KEY(id),
    FOREIGN KEY (center_idfk)
        REFERENCES center_info(id)
        ON DELETE CASCADE
);