package excel

import (
	"github.com/xuri/excelize/v2"
	"maps"
	"strconv"
)

func MakeSheet(filename string, data []map[string]any) {
	file := excelize.NewFile()
	headerKeys := maps.Keys(data[0])
	var headers []string
	for header := range headerKeys {
		headers = append(headers, header)
	}
	for index, header := range headers {
		file.SetCellValue("Sheet1", string(rune(65+index))+strconv.Itoa(1), header)
	}

	for index, row := range data {
		for index2, header := range headers {
			file.SetCellValue("Sheet1", string(rune(65+index2))+strconv.Itoa(index+2), row[header])
		}
	}

	file.SaveAs(filename + ".xlsx")
}
