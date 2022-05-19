CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT  NOT NULL
);
CREATE TABLE lists (
	id  SERIAL PRIMARY KEY,
	name text  NOT NULL,
	user_id INT REFERENCES users(id)  NOT NULL
);
CREATE TABLE tasks (
	id SERIAL PRIMARY KEY,
	text text  NOT NULL,
	completed boolean  NOT NULL,
	list_id INT REFERENCES lists(id)  NOT NULL
);
	
-- name: GetAllTasks :many
SELECT * FROM tasks;
-- name: GetAllLists :many
SELECT * FROM lists;
-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? ;
-- name: GetList :one
SELECT * FROM lists
WHERE id = ? ;
-- name: GetTasksInsideOfList :many
SELECT * FROM tasks
WHERE list_id = ?;
-- name: GetTextOfTasksInsideOfList :many
SELECT text FROM tasks
WHERE list_id = ? ;
-- name: GetUserLists :many
SELECT * FROM lists
WHERE user_id = ? ;
-- name: GetUser :one
SELECT * FROM users
WHERE username = ? ;
-- name: GetUserPassword :one
SELECT password FROM users
WHERE username = ? ;
-- name: DeleteAllTasks :exec
DELETE FROM tasks;
-- name: DeleteAllLists :exec
DELETE FROM lists;
-- name: CreateTask :execresult
INSERT INTO tasks(id,list_id, text, completed) VALUES (?,?, ?, ?);
-- name: CreateUser :execresult
INSERT INTO users(id,username, password) VALUES (?,?,?);
-- name: CreateList :execresult
INSERT INTO lists (id,name, user_id) VALUES (?,?,?);
-- name: DeleteList :exec
DELETE FROM lists WHERE id=?;
-- name: DeleteTasksInsideList :exec
DELETE FROM tasks WHERE list_id=?;
-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=?;
-- name: ToggleTask :exec
UPDATE tasks
SET completed=?
WHERE id=?;

-- name: MaxIdlist :one
SELECT MAX(id) FROM lists;

-- name: MaxIdtask :one
SELECT MAX(id) FROM tasks;
