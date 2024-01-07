-- name: GetEmployeeByID :one
SELECT
  *
FROM
  employees
WHERE
  id = $1
LIMIT
  1;

-- name: CreateEmployee :one
INSERT INTO
  employees (
    name,
    last_name,
    birth_date,
    enter_date,
    gender
  )
VALUES
  ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateEmployeeLocation :exec
INSERT INTO
  employees_location (employee_id, city_id, postal_code)
VALUES
  ($1, $2, $3);

-- name: AddEmployeeToTeam :exec
INSERT INTO
  team_employees (employee_id, team_id)
VALUES
  ($1, $2);

-- name: CreateTeam :one
INSERT INTO
  teams (name, topics)
VALUES
  ($1, $2) RETURNING *;

-- name: GetCountries :many
SELECT
  *
FROM
  countries;

-- name: GetCitiesByCountry :many
SELECT
  country_code,
  name,
  state_name
FROM
  cities
WHERE
  (country_code = $1)
LIMIT $2;


-- name: GetCities :many
SELECT
  country_code,
  name,
  state_name
FROM
  cities
LIMIT $1;