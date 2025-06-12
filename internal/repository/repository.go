package repository

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db               *sql.DB
	statementBuilder sq.StatementBuilderType
}

func New(builderType sq.PlaceholderFormat, driver, connectString string) (*Repository, error) {
	db, err := sql.Open(driver, connectString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}

	return &Repository{
		db:               db,
		statementBuilder: sq.StatementBuilder.PlaceholderFormat(builderType),
	}, nil
}

func (r *Repository) GetAllSubscribedChat() ([]int64, error) {
	sql, args, err := r.statementBuilder.Select(CHAT_ID_COLUMS).
		From(SUBSCRIBED_CHATS_TABLE).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		sql,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var chats []int64
	for rows.Next() {
		var chatid int64
		err := rows.Scan(&chatid)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		chats = append(chats, chatid)
	}
	return chats, nil
}

func (r *Repository) SubscribeChat(chatId int64) error {
	sql, args, err := r.statementBuilder.Insert(SUBSCRIBED_CHATS_TABLE).
		Columns(CHAT_ID_COLUMS).
		Values(chatId).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := r.db.Query(
		sql,
		args...,
	)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	return nil
}

func (r *Repository) UnsubscribeChat(chatId int64) error {
	sql, args, err := r.statementBuilder.Delete(SUBSCRIBED_CHATS_TABLE).
		Where(sq.Eq{CHAT_ID_COLUMS: chatId}).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := r.db.Query(
		sql,
		args...,
	)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	return nil
}
