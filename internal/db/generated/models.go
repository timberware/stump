// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type Follow struct {
	ID       int64
	UserID   string
	Username string
}

type User struct {
	UserID       string
	Username     string
	TwitchToken  string
	RefreshToken string
}
