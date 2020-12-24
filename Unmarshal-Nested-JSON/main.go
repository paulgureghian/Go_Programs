package main

import (
	"encoding/json"
	"log"
)

const jsonText = `
{
 "libraries": [
	{
	 "name": "Paul Gureghian",
	 "city": "los Angeles",
	 "country": "USA",	
	 "books": [
		{
		 "title": "Harry Potter",
		 "author": "J. K. Rowling",
		 "available": true			  
		}
	  ]
	},
	{
	 "name": "Libreria Lowcost",
	 "city": "Pamploma",
	 "country": "Spain",
	 "books": [
		{
		 "title": "The Subtle art of Caring",
		 "author": "Mark Manson",
		 "available": false		
		}
	  ]
	}
  ]
}`

// Libraries is exported
type Libraries struct {
	LibrariesList []struct {
		Name    string `json:"name"`
		City    string `json:"city"`
		Country string `json:"country"`
		Books   []struct {
			Title     string `json:"title"`
			Author    string `json:"author"`
			Available bool   `json:"available"`
		} `json:"books"`
	} `json:"libraries"`
}

func main() {
	var librariesInformation Libraries
	err := json.Unmarshal([]byte(jsonText), &librariesInformation)
	if err != nil {
		log.Fatal("Error Unmarshaling json: ", err)
	}
	println(" ")
	log.Printf("librariesInformation: %+v", librariesInformation)
	println(" ")
}
