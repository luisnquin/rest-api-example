DROP SCHEMA public CASCADE;

CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    topics TEXT []
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    birth_date DATE,
    enter_date DATE,
    creation_date DATE NOT NULL DEFAULT NOW(),
    gender CHAR(1),
    picture_hash TEXT
);

CREATE TABLE countries (
    code VARCHAR(2) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE cities (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    country_code VARCHAR(2) NOT NULL,
    name VARCHAR(50) NOT NULL,
    state_name VARCHAR(30) NOT NULL,
    CONSTRAINT id_tbl PRIMARY KEY (id)
    -- CONSTRAINT cc_ FOREIGN KEY (country_code) REFERENCES countries(code)
);

CREATE TABLE team_employees (
    employee_id INT NOT NULL,
    team_id INT NOT NULL,
    CONSTRAINT te_employee FOREIGN KEY (employee_id) REFERENCES employees(id),
    CONSTRAINT te_team FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE employees_location (
    employee_id INT NOT NULL,
    city_id UUID NOT NULL,
    postal_code TEXT,
    CONSTRAINT el_employee FOREIGN KEY (employee_id) REFERENCES employees(id),
    CONSTRAINT el_city FOREIGN KEY (city_id) REFERENCES cities(id)
);