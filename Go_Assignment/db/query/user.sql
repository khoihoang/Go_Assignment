-- name: GetUserInfoByEmail :many
SELECT * from users
WHERE
    email=$1;

-- name: CreateUser :one
INSERT INTO users(
    id,
    username,
	password,
	first_name,
	last_name,
	address,
	email,
	phone_number,
	active,
	create_at,
	update_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
	username=$2,
	first_name=$3,
	last_name=$4,
	address=$5,
	email = $6,
	phone_number = $7
WHERE
	id=$1
RETURNING *;