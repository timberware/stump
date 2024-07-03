CREATE TABLE user (
                  user_id TEXT PRIMARY KEY,
                  username TEXT NOT NULL,
                  twitch_token TEXT NOT NULL,
                  refresh_token TEXT NOT NULL
);

CREATE TABLE follows (
                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                     user_id TEXT NOT NULL,
                     username TEXT NOT NULL,
                     FOREIGN KEY (user_id) REFERENCES User (user_id)
);
