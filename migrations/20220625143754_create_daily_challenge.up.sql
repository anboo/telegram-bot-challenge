CREATE TABLE IF NOT EXISTS challenge_result (
    id UUID NOT NULL UNIQUE,
    chat_id VARCHAR(255) NOT NULL,
    telegram_id VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,

    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX challenge_result_chat_id_telegram_id ON users (telegram_id, chat_id);