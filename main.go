package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"text/template"
)

//go:embed  template/*
var tmplFS embed.FS

const (
	TypeArray  = "array"
	TypeObject = "object"
)

type Tables struct {
	Object []Object
}
type Object struct {
	Name  string
	Items []Item
}

type Item struct {
	Name        string
	DataType    string
	IsNullable  string
	SampleValue string
	Description string
	OriginValue interface{}
}

func main() {
	fileFlag := flag.String("file", "", "Specify the file path")
	flag.Parse()

	if *fileFlag == "" {
		fmt.Println("Usage: jsonmd --file='c/d/e/f.json'")
		os.Exit(0)
	}

	jsonMap := parseJsonFile(*fileFlag)
	genResponseTable(jsonMap, "result")
}

func parseJsonFile(filePath string) (result map[string]interface{}) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		panic(err)
	}

	return
}

func genResponseTable(jsonMap map[string]interface{}, name string) {
	file, err := os.OpenFile(name+".md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmpl, err := template.ParseFS(tmplFS, "template/response.tmpl")
	if err != nil {
		panic(err)
	}

	table := Tables{}

	obj := Object{}
	obj.Name = name
	items := scanObjectTable("", jsonMap)
	obj.Items = items
	table.Object = append(table.Object, obj)

	nestedCount := countNestedObjects(jsonMap, "", &table, 5)
	fmt.Printf("Total nested objects: %d\n", nestedCount)

	err = tmpl.Execute(file, table)
	if err != nil {
		panic(err)
	}

}

func countNestedObjects(m map[string]interface{}, parentKey string, table *Tables, depth int) int {
	count := 0
	for s, value := range m {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			obj := Object{}
			obj.Name = parentKey + "." + s
			items := scanObjectTable(parentKey+"."+s, value.(map[string]interface{}))
			obj.Items = items
			table.Object = append(table.Object, obj)

			count += 1 + countNestedObjects(nestedMap, parentKey+"."+s, table, depth+1)
		}
	}
	return count
}

func scanObjectTable(parentKey string, jsonMap map[string]interface{}) (items []Item) {
	for s, i := range jsonMap {
		t := reflect.TypeOf(i)
		fmt.Println(s, t)

		ts := getTypeAsString(i)
		ref := fmt.Sprintf("Table : %s.%s", parentKey, s)
		if ts == TypeObject {
			item := Item{
				Name:        fmt.Sprintf("| %s ", s),
				DataType:    fmt.Sprintf("| %s ", getTypeAsString(i)),
				IsNullable:  "| ",
				Description: "| ",
				SampleValue: fmt.Sprintf("| %s |", ref),
				OriginValue: i,
			}
			items = append(items, item)
		} else if ts == TypeArray {
			item := Item{
				Name:        fmt.Sprintf("| %s ", s),
				DataType:    fmt.Sprintf("| %s ", getTypeAsString(i)),
				IsNullable:  "| ",
				Description: "| ",
				SampleValue: fmt.Sprintf("| %s |", ref),
				OriginValue: i,
			}
			items = append(items, item)
		} else {
			item := Item{
				Name:        fmt.Sprintf("| %s ", s),
				DataType:    fmt.Sprintf("| %s ", getTypeAsString(i)),
				IsNullable:  "| ",
				Description: "| ",
				SampleValue: fmt.Sprintf("| %v |", i),
				OriginValue: i,
			}
			items = append(items, item)
		}
	}

	return
}

func getTypeAsString(i interface{}) string {
	if i == nil {
		return ""
	}

	if _, ok := i.([]interface{}); ok {
		return TypeArray
	} else if _, ok := i.(map[string]interface{}); ok {
		return TypeObject
	} else {
		t := reflect.TypeOf(i)
		return t.String()
	}
}
