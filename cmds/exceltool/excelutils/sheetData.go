package excelutils

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

// 初始版本自动化代码工具
type SheetData struct {
	FileName   string //文件名
	SheetName  string //sheet名称
	StructName string //结构体名称
	IdType     string //id类型 string 或者 int32
	Fields     []*Field
	JsonData   string
}

type Field struct {
	Name      string //字段名
	Desc      string //字段描述
	RawType   string //excel 中配置的类型值
	Type      string //解析后的类型如 int32 string...
	InServer  bool   //字段是否在服务器使用
	InClient  bool   //字段是否在客户端使用
	ValidList []string
	ValueList []any
}

// 获取excel 中所有 sheet 数据
func GetExcelAllSheetData(excelFileName string, data *excelize.File) map[string]SheetData {
	if len(data.GetSheetList()) < 1 {
		panic("no sheet data")
	}
	var sheetDataMap = make(map[string]SheetData)
	for _, sheet := range data.GetSheetList() {
		sheetRows, err := data.GetRows(sheet)
		if err != nil {
			panic(err.Error())
		}
		sheetData := SheetData{}
		sheetData.FileName = excelFileName
		sheetData.SheetName = sheet
		sheetData.StructName = sheetRows[0][0]
		//结构体
		//第一行是描述
		var descRow = sheetRows[0]
		//第二行是字段名
		var fieldNameRow = sheetRows[1]
		//第三行是字段类型
		var fieldTypeRow = sheetRows[2]
		//第四行是字段检测规则
		var fieldValidRow = sheetRows[3]
		for i, desc := range descRow {
			field := &Field{}
			field.Desc = desc
			field.Name = fieldNameRow[i]
			field.RawType = fieldTypeRow[i]
			field.ValidList = strings.Split(fieldValidRow[i], ",")
			field.Init()
			sheetData.Fields = append(sheetData.Fields, field)
		}

		//json 数据
		jsonData := "["
		valueRows := sheetRows[4:]
		//遍历行数据 索引4开始
		for rowIndex, row := range valueRows {
			// 行数据为空
			if len(row) < 1 || strings.TrimSpace(row[0]) == "" {
				break
			}
			//先算出行数据唯一 id 的类型
			if rowIndex == 0 {
				toInt32 := cast.ToInt32(row[0])
				if toInt32 > 0 {
					sheetData.IdType = "int32"
				} else {
					sheetData.IdType = "string"
				}
			}
			// 一行的数据转为 json 字符串
			jsonRow := "{"
			// 遍历一行中所有列数据
			for colIndex, colCell := range row {
				// 列的大小大于字段列的大小(说明已经读取完所有字段的数据了)
				if colIndex >= len(fieldNameRow) {
					break
				}
				field := sheetData.Fields[colIndex]
				field.ValueList = append(field.ValueList, colCell)
				//根据此列数据转为json
				cell := GetCellToJson(fieldNameRow[colIndex], colCell, fieldTypeRow[colIndex])
				jsonRow = jsonRow + cell + ","
			}
			//去除逗号
			jsonRow = jsonRow[0 : len(jsonRow)-1]
			jsonRow += "}"
			jsonData = jsonData + jsonRow + ","
		}
		jsonData = jsonData[0 : len(jsonData)-1]
		jsonData += "]"
		// json 数据
		sheetData.JsonData = jsonData
		sheetDataMap[sheet] = sheetData
	}
	return sheetDataMap
}

func GetCellToJson(key, val, fieldType string) string {
	cell := fmt.Sprintf(`"%v":"%v"`, key, val)
	switch fieldType {
	case "i32", "i64":
		if val == "" {
			val = "0"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	}
	return cell
}

func (a *Field) Init() {
	rawType := strings.Split(a.RawType, "_")
	switch rawType[0] {
	case "i64":
		a.Type = "int64"
	case "i32":
		a.Type = "int32"
	case "str":
		a.Type = "string"
	default:
		panic(fmt.Sprintf("未知的类型 %v %+v", rawType, a))
	}
	switch len(rawType) {
	case 1:
		a.InServer = true
		a.InClient = true
	case 2:
		if strings.Contains(rawType[1], "s") {
			a.InServer = true
		}
		if strings.Contains(rawType[1], "c") {
			a.InClient = true
		}
	default:
		panic(fmt.Sprintf("未知的类型 %v %+v", rawType, a))
	}
}
