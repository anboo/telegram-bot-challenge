CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,
    telegram_id VARCHAR(255) NOT NULL,
    chat_id VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,

    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX users_chat_id_telegram_id ON users (telegram_id, chat_id);