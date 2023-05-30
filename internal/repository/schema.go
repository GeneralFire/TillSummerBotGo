package repository

const _ = `
USE DB_NAME;

CREATE TABLE sq (
	id BIGINT,
	UNIQUE (id)
)
`

const (
	SUBSCRIBED_CHATS_TABLE = "sq"
	CHAT_ID_COLUMS         = "id"
)
