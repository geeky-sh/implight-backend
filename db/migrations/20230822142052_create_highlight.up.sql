CREATE TABLE IF NOT EXISTS highlights (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    text TEXT NOT NULL,
    url VARCHAR(255) NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
)
