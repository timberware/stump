-- name: GetUserID :one
SELECT user_id FROM user WHERE user_id = ?;

-- name: GetUsername :one
SELECT username FROM user WHERE user_id = ?;

-- name: GetTwitchToken :one
SELECT twitch_token FROM user WHERE user_id = ?;

-- name: GetRefreshToken :one
SELECT refresh_token FROM user WHERE user_id = ?;

-- name: GetFollowers :many
SELECT username FROM follows WHERE user_id = ?;

-- name: InsertUser :exec
INSERT INTO user (user_id, username, twitch_token) VALUES (?, ?, ?);

-- name: InsertUserId :exec
INSERT INTO user (user_id) VALUES (?);

-- name: InsertUsername :exec
INSERT INTO user (username) VALUES (?);

-- name: InsertTwitchToken :exec
INSERT INTO user (twitch_token) VALUES (?);

-- name: InsertRefreshToken :exec
INSERT INTO user (refresh_token) VALUES (?);

-- name: UpdateUserTwitchToken :exec
UPDATE user SET twitch_token = (?) WHERE user_id = (?);

-- name: UpdateUserRefreshToken :exec
UPDATE user SET refresh_token = (?) WHERE user_id = (?);

-- name: DeleteFollowsByUserId :exec
DELETE FROM follows WHERE user_id = (?);

-- name: DeleteUserByUserId :exec
DELETE FROM user WHERE user_id = (?);
