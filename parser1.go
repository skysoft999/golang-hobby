package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type post struct {
	XMLNAME xml.Name `xml:"post"`
	Id		string 	 `xml:"id,attr"`
	Content	string	  `xml:"content"`
	Authour Authour   `xml:"author"`
	Xml     string 	  `xml: ",innerxml"`
}

type Authour struct {
	Id string `xml:"id,attr`
	Name string `xml:",chardata"`
}


func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Errro Reading xml data:",err)
		return
	}
	defer xmlFile.Close()
	xmlData,err := ioutil.ReadAll(xmlFile)
	if err != nil {

		fmt.Println("error reading XML data:",err)
		return
	}

	var post post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

}