package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/tealeg/xlsx"
)

// Data is a struct that conforms to the JSON required by the display frontend.
type Data struct {
	LastModify string   `json:"lastmodify"`
	License    string   `json:"license"`
	Source     string   `json:"source"`
	People     []Person `json:"data"`
}

// Person is a struct that conforms to the  latest schema:
// 縣市別	醫院別	編號	檢傷編號	姓名	性別	年齡	國籍	醫療檢傷	救護檢傷	即時動向
type Person struct {
	Param0  string `json:"縣市別"`
	Param1  string `json:"收治單位"`
	Param2  string `json:"編號"`
	Param3  string `json:"檢傷編號"`
	Param4  string `json:"姓名"`
	Param5  string `json:"性別"`
	Param6  string `json:"國籍"`
	Param7  string `json:"年齡"`
	Param8  string `json:"醫療檢傷"`
	Param9  string `json:"救護檢傷"`
	Param10 string `json:"即時動向"`
}

func main() {
	excelFileName := "./people.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}
	var people []Person
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 || i == 1 {
				continue
			}
			var line []string
			for _, cell := range row.Cells {
				line = append(line, cell.String())
			}
			fmt.Println(line)
			people = append(people, Person{
				Param0:  line[0],
				Param1:  line[1],
				Param2:  line[2],
				Param3:  line[3],
				Param4:  line[4],
				Param5:  line[5],
				Param6:  line[6],
				Param7:  line[7],
				Param8:  line[8],
				Param9:  line[9],
				Param10: line[10],
			})
		}
	}
	data := Data{
		LastModify: time.Now().Format("2006/01/02 15:04"),
		Source:     "衛福部",
		License:    "cc0",
		People:     people,
	}
	dataJSON, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(dataJSON))
	ioutil.WriteFile("people.json", dataJSON, 0644)
}
