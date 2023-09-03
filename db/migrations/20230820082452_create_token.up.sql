CREATE TABLE IF NOT EXISTS tokens (
    uuid VARCHAR(36) NOT NULL,
    id_token TEXT NOT NULL,
    issued_at timestamp NOT NULL,
    expires_at timestamp NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY(uuid),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
)
