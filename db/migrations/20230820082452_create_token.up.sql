CREATE TABLE IF NOT EXISTS tokens (
    id_token VARCHAR(64) NOT NULL,
    issued_at timestamp NOT NULL,
    expires_at timestamp NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY(id_token),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
)
