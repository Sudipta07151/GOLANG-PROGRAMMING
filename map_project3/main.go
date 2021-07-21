package main

import "fmt"

func main() {
	 
	//[keys]values //keys typem string & values type string
	
	//syntax1

	
	players := map[string]string{
		"Ramos":   "CB",
		"Laporte": "CB",
		"Morata":  "CF",
		"Olmo":    "CMF",
		"Alba":"LB",
		"Sarabia":"RWF",
	}

	
	fmt.Println(players)
	//syntax2
	// var players map[string]string

	//syntax3
	// players:=make(map[string]string)
	// players["Pedri"]="CMF"
	
	//delete 
	// delete(players,"Ramos")
	// fmt.Println(players)
	printMap(players)
	 
}

func printMap(players map[string]string){
	for name,pos:=range players{
		fmt.Println(name," plays as ",pos)
	} 
}