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

type Response struct {
	Name        string
	DataType    string
	IsNullable  string
	SampleValue string
	Description string
}

type ResponseObjChild struct {
	ObjName     string
	Name        string
	DataType    string
	IsNullable  string
	SampleValue string
	Description string
}

func main() {
	fileFlag := flag.String("file", "", "Specify the file path")
	flag.Parse()

	if *fileFlag == "" {
		fmt.Println("Usage: jsonmd --file='c/d/e/f.json'")
		os.Exit(0)
	}

	jsonMap := parseJsonFile(*fileFlag)
	genResponseTable(jsonMap, "root")
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

func genResponseTable1(jsonMap map[string]interface{}, name string) {
	var objs []map[string]interface{}
	var items []Response
	for s, i := range jsonMap {
		t := reflect.TypeOf(i)
		fmt.Println(s, t)

		item := Response{
			Name:        s,
			DataType:    getTypeAsString(i),
			IsNullable:  "",
			Description: "",
		}

		ts := getTypeAsString(i)
		if ts == TypeObject {
			item.SampleValue = fmt.Sprintf("Table %s %v ", s, ts)
			objs = append(objs, jsonMap[s].(map[string]interface{}))
		} else if ts == TypeArray {
			item.SampleValue = fmt.Sprintf("Table %s %v ", s, ts)
		} else {
			item.SampleValue = fmt.Sprintf("%v", i)
		}

		items = append(items, item)
	}

	for _, o := range objs {
		for s, i := range o {
			t := reflect.TypeOf(i)
			fmt.Println(s, t)

			item := Response{
				Name:        s,
				DataType:    getTypeAsString(i),
				IsNullable:  "",
				Description: "",
			}

			ts := getTypeAsString(i)
			if ts == TypeObject {
				item.SampleValue = fmt.Sprintf("Table %s %v ", s, ts)
				objs = append(objs, map[string]interface{}{s: i})
			} else if ts == TypeArray {
				item.SampleValue = fmt.Sprintf("Table %s %v ", s, ts)
			} else {
				item.SampleValue = fmt.Sprintf("%v", i)
			}

			items = append(items, item)
		}
	}

	tmpl, err := template.ParseFS(tmplFS, "template/response.tmpl")
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(name+".md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, items)
	if err != nil {
		panic(err)
	}
}

func genResponseTable(jsonMap map[string]interface{}, name string) {
	var objs []map[string]map[string]interface{}
	var items []Response
	var itemsChild []Response

	for s, i := range jsonMap {
		t := reflect.TypeOf(i)
		fmt.Println(s, t)

		item := Response{
			Name:        fmt.Sprintf("| %s ", s),
			DataType:    fmt.Sprintf("| %s ", getTypeAsString(i)),
			IsNullable:  "| ",
			Description: "| ",
		}

		ts := getTypeAsString(i)
		if ts == TypeObject {
			objs = append(objs, map[string]map[string]interface{}{s: jsonMap[s].(map[string]interface{})})
			item.SampleValue = fmt.Sprintf("| Table %s %v |", s, ts)
		} else if ts == TypeArray {
			item.SampleValue = fmt.Sprintf("| Table %s %v |", s, ts)
		} else {
			item.SampleValue = fmt.Sprintf("| %v |", i)
		}

		items = append(items, item)
	}

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

	err = tmpl.Execute(file, items)
	if err != nil {
		panic(err)
	}
	items = nil

	for len(objs) != 0 {
		for i, o := range objs {
			itemsChild = nil
			var key string
			for s, _ := range o {
				key = s
			}
			items = append(items, Response{Name: fmt.Sprintf("### Table %s object", key)})
			items = append(items, Response{})
			err = tmpl.Execute(file, items)
			if err != nil {
				panic(err)
			}
			items = nil
			objs = removeByIndex(objs, i)

			for s, i := range o[key] {
				t := reflect.TypeOf(i)
				fmt.Println(s, t)

				itemChild := Response{
					Name:        fmt.Sprintf("| %s", s),
					DataType:    fmt.Sprintf("| %s", getTypeAsString(i)),
					IsNullable:  fmt.Sprintf("| "),
					Description: fmt.Sprintf("|  |"),
				}

				ts := getTypeAsString(i)
				if ts == TypeObject {
					itemChild.SampleValue = fmt.Sprintf("| Table %s %v |", s, ts)
					objs = append(objs, map[string]map[string]interface{}{s: i.(map[string]interface{})})
				} else if ts == TypeArray {
					itemChild.SampleValue = fmt.Sprintf("| Table %s %v |", s, ts)
				} else {
					itemChild.SampleValue = fmt.Sprintf("| %v |", i)
				}

				itemsChild = append(itemsChild, itemChild)
			}

			err = tmpl.Execute(file, itemsChild)
			if err != nil {
				panic(err)
			}
		}

	}

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

func removeByIndex(slice []map[string]map[string]interface{}, index int) []map[string]map[string]interface{} {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func removeByValue(slice []Response, valueToRemove Response) []Response {
	var result []Response
	for _, p := range slice {
		if p != valueToRemove {
			result = append(result, p)
		}
	}
	return result
}
