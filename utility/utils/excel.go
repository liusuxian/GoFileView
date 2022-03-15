package utils

import (
	"github.com/tealeg/xlsx"
	"log"
)

// ExcelParse excel解析
func ExcelParse(filePath string) []map[string]interface{} {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Println("ExcelParseError: <", err.Error(), ">")
	}
	var resData []map[string]interface{}

	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		tmp := map[string]interface{}{}
		//遍历每一行
		var title []string
		var resourceArr [][]string
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			if rowIndex == 0 {
				for _, cell := range row.Cells {
					text := cell.String()
					title = append(title, text)
				}
				continue
			}
			//遍历每一个单元
			var result []string
			for _, cell := range row.Cells {
				text := cell.String()
				result = append(result, text)
			}
			resourceArr = append(resourceArr, result)
		}

		tmp["title"] = title
		tmp["resourceArr"] = resourceArr

		resData = append(resData, tmp)
	}
	return resData
}
