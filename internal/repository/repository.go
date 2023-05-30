package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db               *sql.DB
	statementBuilder sq.StatementBuilderType
	chats            []int64
}

func New(builderType sq.PlaceholderFormat, driver, connectString string) (*Repository, error) {
	db, err := sql.Open(driver, connectString)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db:               db,
		statementBuilder: sq.StatementBuilder.PlaceholderFormat(builderType),
	}, nil
}

func (r *Repository) GetAllSubscribedChat() ([]int64, error) {
	return r.chats, nil
}

func (r *Repository) SubscribeChat(chatId int64) error {
	r.chats = append(r.chats, chatId)
	return nil
}

func (r *Repository) UnsubscribeChat(chatId int64) error {
	return nil
}
