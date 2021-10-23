package main

import (
	"database/sql"
	"log"
	"strings"
)

var database *sql.DB

func DB_command() []string{
	db, err := sql.Open("mysql", "root:1234@/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT command,command_description FROM productdb.command_bot")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	command := []Command{}

	for rows.Next(){
		p := Command{}
		err := rows.Scan(&p.command_name,&p.command_description)
		if err != nil{
			log.Println(err)
			continue
		}
		command = append(command, p)
	}
	array := []string{}
	for _, p := range command {
		array=append(array,p.command_name)
		array=append(array,p.command_description)
	}
	return array
}

func Create_article(massage string) int{
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		return 1
	}
	name, description :=  splitting_string_and_delete_article(massage,"/create_article")
	if name == ""{
		return 3
	}
	if checking_title_article(name) == 3{
		return 3
	}
	_, err = db.Exec("insert into productdb.article_bot (name, description) values (?, ?)",name,description)
	if err != nil {
		panic(err)
		return 2
	}
	defer db.Close()
	return 0
}

func Create_link_article(message string)  int{
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		return 1
	}
	name_article, link, description :=splitting_string_and_delete_link(message,"/save_link")
	if name_article == "" {
		return 3
	}
	if  checking_link_article(link)==3 || checking_title_article(name_article) != 3{
		return 3
	}
	_, err = db.Exec("insert into productdb.article_bot_links (name_article_bot, link, description) values (?, ?, ?)",name_article, link,description)
	if err != nil {
		panic(err)
		return 2
	}
	defer db.Close()
	return 0
}

func splitting_string_and_delete_article(st string,del string) (string,string){
	words := strings.Fields(st)
	if len(words)==1{
		return "",""
	}
	name := words[1]
	i:=0
	for index:=0; index<2;index++{
		copy(words[i:],words[i+1:])
		words[len(words)-1]=""
		words=words[:len(words)-1]
	}
	resolt_string := strings.Join(words," ")
	return name,resolt_string
}

func splitting_string_and_delete_link(st string,del string) (string,string,string){
	words := strings.Fields(st)
	if len(words)==1{
		return "","",""
	}
	name := words[1]
	link :=words[2]
	i:=0
	for index:=0; index<3;index++{
		copy(words[i:],words[i+1:])
		words[len(words)-1]=""
		words=words[:len(words)-1]
	}
	resolt_string := strings.Join(words," ")
	return name,link,resolt_string
}
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

