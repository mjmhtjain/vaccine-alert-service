CREATE USER 'root'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

CREATE DATABASE db;
USE db;

CREATE TABLE IF NOT EXISTS vaccine(
    id BINARY(16),
    name VARCHAR(16),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS center_info(
    center_id INT,
    name VARCHAR(16),
    address VARCHAR(32),
    state_name VARCHAR(16),
    district_name VARCHAR(16),
    pincode INT,
    PRIMARY KEY(center_id)
);

CREATE TABLE IF NOT EXISTS appointment_session(
    session_id BINARY(36),
    center_idfk INT, 
    date VARCHAR(16),
    available_capacity INT,
    min_age_limit INT,
    vaccine_idfk  BINARY(16),
    available_capacity_dose1 INT,
    available_capacity_dose2 INT,
    PRIMARY KEY(session_id),
    FOREIGN KEY (vaccine_idfk)
        REFERENCES vaccine(id)
        ON DELETE CASCADE,
    FOREIGN KEY (center_idfk)
        REFERENCES center_info(center_id)
        ON DELETE CASCADE
);