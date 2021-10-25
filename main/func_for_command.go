package main

import (
	"database/sql"
	"log"
)

var database *sql.DB

func Get_command() []string{
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

func Get_article(name_article string) (int,[]string){
	name_article=splitting_string_name_article(name_article)
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		//return 1, _ //поправить
	}
	if name_article == "" {
		//return 3 //поправить
	}
	//от сюда правки
	defer db.Close()
	rows, err := db.Query("SELECT link,description FROM productdb.article_bot_links WHERE name_article_bot=?",name_article)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	article_link := []Article_link{}

	for rows.Next(){
		tmp := Article_link{}
		err := rows.Scan(&tmp.link,&tmp.description)
		if err != nil{
			log.Println(err)
			continue
		}
		article_link = append(article_link, tmp)
	}
	array := []string{}
	for _, link := range article_link {
		array=append(array,"Ссылка = "+link.link+"\n")
		array=append(array,"Описание: "+link.description+"\n\n")
	}
	return 0,array
}

func Get_all_article()(int,[]string){
	error := []string{"error"}
	db, err := sql.Open("mysql", "root:1234@/productdb")
	if err != nil {
		panic(err)
		return 1, error
	}
	//от сюда правки
	defer db.Close()
	rows, err := db.Query("SELECT name FROM productdb.article_bot ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	name_article:= []string{}

	for rows.Next(){
		tmp := ""
		err := rows.Scan(&tmp)
		if err != nil{
			log.Println(err)
			continue
		}
		name_article = append(name_article, tmp)
	}
	resolt_data := []string{}
	for idx,name := range name_article{
		resolt_data = append(resolt_data,"Материалы по "+name_article[idx]+"\n ------------- \n")
		_, tmp :=Get_article("/command "+name)
		resolt_data = append(resolt_data,tmp...)
	}
	return 0,resolt_data
}



