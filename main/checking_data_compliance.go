package main

import (
	"database/sql"
	"log"
)

func checking_title_article(article string) int{
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		return 1
	}
	rows, err := db.Query("SELECT name FROM productdb.article_bot")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	name_from_db := []string{}
	tmp := ""
	for rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			log.Println(err)
		}
		name_from_db=append(name_from_db,tmp)
	}
	for _,name_db := range name_from_db{
		if name_db==article{
			return 3
		}
	}
	return 0
}

func checking_link_article(link string)  int{
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		return 1
	}
	rows, err := db.Query("SELECT link FROM productdb.article_bot_links")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	name_from_db := []string{}
	tmp := ""
	for rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			log.Println(err)
		}
		name_from_db=append(name_from_db,tmp)
	}
	for _,name_db := range name_from_db{
		if name_db==link{
			return 3
		}
	}
	return 0
}
