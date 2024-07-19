package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type (
	KeyboardBindings struct {
		XMLName    xml.Name   `xml:"KeyboardBindings"`
		Categories Categories `xml:"Categories"`
	}
	Categories struct {
		XMLName    xml.Name   `xml:"Categories"`
		Categories []Category `xml:"Category"`
	}
	Category struct {
		XMLName     xml.Name    `xml:"Category"`
		Identifier  string      `xml:"Identifier"`
		KeyBindings KeyBindings `xml:"KeyBindings"`
	}
	KeyBindings struct {
		XMLName     xml.Name     `xml:"KeyBindings"`
		KeyBindings []KeyBinding `xml:"KeyBinding"`
	}
	KeyBinding struct {
		XMLName xml.Name `xml:"KeyBinding"`
		Topic   string   `xml:"Topic"`
		Binding string   `xml:"Binding"`
		Key     string   `xml:"Key"`
	}
)

func parseXML(path string) []string {
	out := []string{}
	xmlFile, err := os.Open(path)
	if err != nil {
		log.Println("didn't open kb file")
	}
	defer xmlFile.Close()
	b, _ := io.ReadAll(xmlFile)
	var kbs KeyboardBindings
	xml.Unmarshal(b, &kbs)

	cat := ""
	for _, kb := range kbs.Categories.Categories {
		if kb.Identifier != cat {
			cat = kb.Identifier
		}
		curr := ""
		for _, k := range kb.KeyBindings.KeyBindings {
			if len(k.Key) > 0 {
				if k.Topic != curr {
					curr = k.Topic
				}
				out = append(out, cat+","+curr+","+k.Binding+","+k.Key)
			}
		}
	}
	return out
}

func main() {
	http.NewRequest("GET", "http://example.com", nil)
	var parsed_xml = parseXML(os.Args[1])
	for _, line := range parsed_xml {
		fmt.Println(line)
	}
}
