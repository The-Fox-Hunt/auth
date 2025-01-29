CREATE TABLE IF NOT EXISTS `users` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_vvid UUID NOT NULL,
    short_name VARCHAR(255),
    full_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NULL,)
    