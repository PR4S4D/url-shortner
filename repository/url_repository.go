package repository

import (
	"url-shortner/models"

	"github.com/jmoiron/sqlx"
)

type UrlRepository struct {
	db *sqlx.DB
}

const urlInsertQuery = `INSERT INTO url (url, short_url) VALUES (?, ?)`
const urlSelectAllQuery = `select * from url order by id desc limit 500`
const findUrlByShortUrlQuery = `select * from url where short_url = ?`

var UrlRepo *UrlRepository

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{db}
}

func (r *UrlRepository) QueryAll() ([]models.UrlData, error) {
	results := []models.UrlData{}
	err := r.db.Select(&results, urlSelectAllQuery)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *UrlRepository) Query(shortUrl string) (models.UrlData, error) {
	result := models.UrlData{}
	err := r.db.Get(&result, findUrlByShortUrlQuery, shortUrl)
	if err != nil {
		result.Id = -1
		return result, err
	}

	return result, nil
}

func (r *UrlRepository) Save(url string, shortUrl string) (int64, error) {
	res := r.db.MustExec(urlInsertQuery, url, shortUrl)
	return res.RowsAffected()
}
