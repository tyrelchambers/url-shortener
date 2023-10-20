package services

import (
	"fmt"
	"url/db"
	"url/types"
)

type Url struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	PublicId string `json:"public_id"`
}

func GetAllUrls() []types.Url {
	var urls []types.Url

	rows, err := db.Db.Query("SELECT * FROM urls")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var url types.Url

		if err = rows.Scan(&url.Id, &url.PublicId, &url.Url); err != nil {
			fmt.Println(err)
		}

		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	return urls
}

func GetUrlById(id string) types.Url {
	var url types.Url

	row := db.Db.QueryRow("SELECT * FROM urls WHERE public_id = $1", id)

	if err := row.Scan(&url.Id, &url.PublicId, &url.Url); err != nil {
		fmt.Println(err)
	}

	return url
}
