-- name: InsertUser :one
INSERT INTO Users(
    Id,
    FirstName,
    LastName,
    EmailAddresss,
    Password
) VALUES ($1, $2, $3, $4, $5)

RETURNING *;

-- name: GetUserById :one
SELECT * FROM Users
WHERE Id = $1;