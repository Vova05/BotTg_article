package main

import "strings"

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

func splitting_string_name_article(start_string string)string{
	name_article:=strings.Fields(start_string)
	return name_article[1]
}
