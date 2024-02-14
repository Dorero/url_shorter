package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"url_shorter/db"
	"url_shorter/model"
)

type IUrlRepository interface {
	Find(id string) (model.Url, error)
	Create(path string) (model.Url, error)
}

type UrlRepository struct {
	Db *db.Db
}

func (u *UrlRepository) Create(path string) (model.Url, error) {
	var url model.Url

	err := u.Db.Pool.QueryRow(context.Background(), "insert into urls (path) values ($1) RETURNING id, path", path).Scan(&url.Id, &url.Path)

	if err != nil {
		return url, fmt.Errorf("Internal Server Error")
	}

	return url, nil
}

func (u *UrlRepository) Find(id string) (model.Url, error) {
	var url model.Url

	err := u.Db.Pool.QueryRow(context.Background(), "SELECT path from urls where id = $1", id).Scan(&url.Id, &url.Path)

	if err != nil {
		log.Println("Error querying database:", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return url, fmt.Errorf("URL not found for ID %s", id)
		} else {
			return url, fmt.Errorf("Internal Server Error: %w", err)
		}
	}

	return url, nil
}
