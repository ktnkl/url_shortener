package database

import (
	"context"
	"log"
	"time"
)

func CreateShortenedUrlQuery(url string) (int, error) {
	var id int

	if db == nil {
		log.Fatal("Connection pool doesn't exist")
	}
	err := db.QueryRow(context.TODO(), `insert into shortened_urls(url, created_at, updated_at) values ($1, $2, $3) returning id`, url, time.Now(), time.Now()).Scan(&id)

	return id, err
}

func DeleteShorenedUrlQuery(shortened int) (int, error) {
	var id int

	if db == nil {
		log.Fatal("Connection pool doesn't exist")
	}
	err := db.QueryRow(context.TODO(), `delete from shortened_urls where id = $1 returning id`, shortened).Scan(&id)

	return id, err
}

func GetOriginalUrlQuery(shortened string) (string, error) {
	var url string

	if db == nil {
		log.Fatal("Connection pool doesn't exist")
	}
	err := db.QueryRow(context.TODO(), `select url from shortened_urls where id = $1`, shortened).Scan(&url)

	return url, err
}
