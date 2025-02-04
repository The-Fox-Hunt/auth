CREATE TABLE IF NOT EXISTS participants (
    id SERIAL PRIMARY KEY,
    user_uuid UUID NOT NULL,
    short_name VARCHAR(255),
    full_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
    );

ALTER TABLE participants RENAME COLUMN short_name TO login;
ALTER TABLE participants RENAME COLUMN full_name TO password;