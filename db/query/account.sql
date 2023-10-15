-- return 1 single account object, this is create operation
-- name: CreateAccount :one
INSERT INTO accounts (
  owner, 
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;
-- RETURNING * to tell postgresql to return valuer of all columns
--this is read operation 

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR  NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;
-- limit to set number of rows we want to get and offset to skip this many rows before starting to return the result

--exec does not return any data, it updates. we will use one to get the value instead of exec 
-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
  set balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *; 

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;