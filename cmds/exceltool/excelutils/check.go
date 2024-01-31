package excelutils

import (
	"fmt"
	"strings"
)

// map[表名_字段名]map[字段的值]bool
var allValidSheetData = make(map[string]map[any]bool)

// []string{"表名_字段名"}
var supportValidList = []string{"Item_Id"}

// 检测数值表
func CheckExcel(excelDataMap map[string]map[string]SheetData) {
	//获取所有等待验证的数据
	for _, sheetDataList := range excelDataMap {
		for _, sheetData := range sheetDataList {
			for _, supportValid := range supportValidList {
				splitN := strings.SplitN(supportValid, "_", 2)
				tableName := splitN[0]
				fieldName := splitN[1]
				if tableName != sheetData.StructName {
					continue
				}
				var fieldExist bool
				for _, field := range sheetData.Fields {
					if field.Name != fieldName {
						continue
					}
					fieldExist = true
					var data = make(map[any]bool)
					for _, value := range field.ValueList {
						data[value] = true
					}
					allValidSheetData[supportValid] = data
					break
				}
				if !fieldExist {
					panic(fmt.Sprintf("supportValid 未找到 field", supportValidList, tableName, fieldName))
				}
			}
		}
	}

	for _, sheetDataList := range excelDataMap {
		for _, sheetData := range sheetDataList {
			for _, field := range sheetData.Fields {
				for _, valid := range field.ValidList {
					switch valid {
					case "none":
						continue
					case "Item_Id":
						for _, value := range field.ValueList {
							_, ok := allValidSheetData[valid][value]
							if !ok {
								panic(fmt.Sprintf("验证失败 %v %v %v %v",
									valid, sheetData.FileName, sheetData.SheetName, field.Name))
							}
						}
					case "unique":
						var uniqueMap = make(map[any]bool)
						for _, value := range field.ValueList {
							_, ok := uniqueMap[value]
							if !ok {
								uniqueMap[value] = true
								continue
							}
							panic(fmt.Sprintf("unique失败：fileName=%v sheetName=%v fieldName=%v value=%v ",
								sheetData.FileName, sheetData.SheetName, field.Name, value))
						}
					}
				}
			}
		}
	}
}
