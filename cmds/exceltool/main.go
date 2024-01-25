package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"zinx-zero/cmds/exceltool/auto"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

var dirPath = "./excel/"

// 输出自定义文件夹
var dirGo = []string{"./output/go/cfg/"}
var dirJson = []string{"./output/json/", "./output/go/conf/excel/"}
var dirTs = []string{"./output/ts/"}

func main() {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var structNameList []string
	for _, info := range dir {
		if info.IsDir() {
			continue
		}
		if strings.Contains(info.Name(), "~$") {
			continue
		}
		if strings.Contains(info.Name(), ".~") {
			continue
		}
		if info.Name() == ".DS_Store" {
			continue
		}
		data, err := excelize.OpenFile(dirPath + info.Name())
		if err != nil {
			panic(err)
		}
		if info.Name() == "Template.xlsm" {
			continue
		}
		fmt.Println(info.Name())
		//获取`go`和`json`文件数据
		autocode, jsonData := GetExcelData(data)

		//首字母大写
		structName := strings.Title(autocode.StructName)
		structNameList = append(structNameList, structName)
		//生成go文件
		for _, dir := range dirGo {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			create, err := os.Create(dir + structName + ".go")
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}
			files, err := template.ParseFiles("./config.go.tpl")
			if err != nil {
				panic("create template failed, err:" + err.Error())
			}
			err = files.Execute(create, autocode)
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}
		}
		//生成json文件
		for _, dir := range dirJson {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			jsonFile, err := os.Create(dir + structName + ".json")
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}

			_, err = jsonFile.Write([]byte(jsonData))
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}
			jsonFile.Close()
		}
		//生成ts文件
		for _, dir := range dirTs {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			jsonFile, err := os.Create(dir + structName + ".ts")
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}

			jsonData = `import { ConfigList } from '../common/Init'; ` +
				`ConfigList.` + autocode.StructName + " = " + jsonData

			_, err = jsonFile.Write([]byte(jsonData))
			if err != nil {
				panic("create file failed, err:" + err.Error())
			}
			jsonFile.Close()
		}
	}

	//生成 ainit.go 文件
	// 初始版本自动化代码工具
	// type Ainit struct {
	// 	StructNameList []string `json:"structNameList"`
	// }
	// ainit := Ainit{StructNameList: structNameList}
	for _, dir := range dirGo {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		create, err := os.Create(dir + "ainit.go")
		if err != nil {
			panic("create file failed, err:" + err.Error())
		}
		files, err := template.ParseFiles("./config_ainit.go.tpl")
		if err != nil {
			panic("create template failed, err:" + err.Error())
		}
		err = files.Execute(create, structNameList)
		if err != nil {
			panic("create file failed, err:" + err.Error())
		}
	}
	fmt.Println("解析配置完成")
}

// 获取excel数据
func GetExcelData(data *excelize.File) (autocode auto.AutoCodeStruct, jsonData string) {
	if len(data.GetSheetList()) < 1 {
		panic("no sheet data")
	}
	getRows, err := data.GetRows(data.GetSheetList()[0])
	if err != nil {
		panic(err.Error())
	}

	autocode.StructName = strings.Title(getRows[0][0])
	//存放object内数据
	objectData := make(map[string]map[string]string)
	//结构体
	//第一行是描述
	//第二行是json字段名
	//第三行是字段类型
	for i, desc := range getRows[0] {
		field := auto.Field{}
		field.FieldJson = getRows[1][i]
		field.FieldName = strings.Title(field.FieldJson)
		field.FieldType = autocode.GetFieldType(getRows[2][i])
		field.FieldDesc = desc

		var fieldType = getRows[2][i]
		// 结构体 sheet
		if fieldType == "t" {
			field.FieldType = "map[int32]" + autocode.StructName + field.FieldName
			//	查找sheet 解析
			objectAutoCode, subObjectData := GetObjectData(autocode.StructName+field.FieldName, field.FieldJson, data)
			objectData[field.FieldJson] = subObjectData
			autocode.Structs = append(autocode.Structs, objectAutoCode)
		}
		autocode.Fields = append(autocode.Fields, field)
	}

	//json 数据
	jsonData = "{"
	jsonRows := getRows[4:]
	//数据 索引4开始
	for rowIndex, row := range jsonRows {
		if len(row) < 1 || strings.TrimSpace(row[0]) == "" {
			break
		}
		if rowIndex == 0 {
			toInt32 := cast.ToInt32(row[0])
			if toInt32 > 0 {
				autocode.IDType = "int32"
			} else {
				autocode.IDType = "string"
			}
		}
		//id的类型

		objectKey := ""
		jsonRow := "{"
		for i2, value := range row {
			if i2 == 0 {
				jsonRow = fmt.Sprintf(`"%v":{`, value)
				objectKey = value
			}
			if i2 >= len(getRows[1]) {
				break
			}
			//获取此列字段类型
			fieldType := getRows[2][i2]
			fieldJson := getRows[1][i2]
			//处理字典object特殊数据
			if fieldType == "t" {
				//检测是否有数据
				subObjectData, ok := objectData[fieldJson]
				if !ok {
					panic("字典没有对应的数据")
				}
				value2, ok := subObjectData[objectKey]
				if ok {
					value = value2
					cell := fmt.Sprintf(`"%v":%v`, getRows[1][i2], value)
					jsonRow = jsonRow + cell + ","
					continue
				} else {
					value = value2
					cell := fmt.Sprintf(`"%v":{}`, getRows[1][i2])
					jsonRow = jsonRow + cell + ","
					continue
				}
			} else {
				cell := GetCell(getRows[1][i2], value, fieldType)
				jsonRow = jsonRow + cell + ","
			}
		}
		//去除逗号
		jsonRow = jsonRow[0 : len(jsonRow)-1]
		jsonRow += "}"
		jsonData = jsonData + jsonRow + ","
	}
	jsonData = jsonData[0 : len(jsonData)-1]
	jsonData += "}"
	return autocode, jsonData
}

func GetObjectData(structName, objectName string, data *excelize.File) (auto.AutoCodeStruct, map[string]string) {
	var autocode auto.AutoCodeStruct
	autocode.StructName = structName
	getRows, err := data.GetRows(objectName)
	if err != nil {
		panic(err.Error())
	}
	//结构体
	//第一行是描述
	//第二行是json字段名
	//第三行是字段类型
	for i, desc := range getRows[0] {
		field := auto.Field{}
		field.FieldJson = getRows[1][i]
		field.FieldName = strings.Title(field.FieldJson)
		field.FieldType = autocode.GetFieldType(getRows[2][i])
		field.FieldDesc = desc
		autocode.Fields = append(autocode.Fields, field)
	}

	resData := map[string]string{}

	//json 数据
	jsonData := "{"
	jsonRows := getRows[4:]
	//数据 索引4开始
	for _, row := range jsonRows {
		if len(row) < 1 || strings.TrimSpace(row[0]) == "" {
			break
		}
		objectKey := ""
		jsonRow := "{"
		for i2, value := range row {
			if i2 == 0 {
				objectKey = value
			}
			if i2 == 1 {
				jsonRow = fmt.Sprintf(`"%v":{`, value)
			}
			if i2 >= len(getRows[1]) {
				break
			}
			//获取此列字段类型
			fieldType := getRows[2][i2]
			//fieldJson := getRows[1][i2]
			//处理字典object特殊数据
			cell := GetCell(getRows[1][i2], value, fieldType)
			jsonRow = jsonRow + cell + ","
		}
		//去除逗号
		jsonRow = jsonRow[0 : len(jsonRow)-1]
		jsonRow += "}"
		if objectKey != "" {
			s, ok := resData[objectKey]
			if ok {
				resData[objectKey] = s + "," + jsonRow
			} else {
				resData[objectKey] = jsonRow
			}
		}
		jsonData = jsonData + jsonRow + ","
	}

	for s, s2 := range resData {
		resData[s] = "{" + s2 + "}"
	}
	jsonData = jsonData[0 : len(jsonData)-1]
	jsonData += "}"

	return autocode, resData
}

func GetObjectData2(structName, objectName string, data *excelize.File) (auto.AutoCodeStruct, map[string]string) {
	var autocode auto.AutoCodeStruct
	autocode.StructName = structName
	getRows, err := data.GetRows(objectName)
	if err != nil {
		panic(err.Error())
	}
	//结构体
	//第一行是描述
	//第二行是json字段名
	//第三行是字段类型
	for i, desc := range getRows[0] {
		field := auto.Field{}
		field.FieldJson = getRows[1][i]
		field.FieldName = strings.Title(field.FieldJson)
		field.FieldType = autocode.GetFieldType(getRows[2][i])
		field.FieldDesc = desc
		autocode.Fields = append(autocode.Fields, field)
	}

	// [objectID]value
	jsonMap := map[string]string{}
	jsonRows := getRows[4:]
	//数据 索引4开始
	for _, row := range jsonRows {
		objectKey := ""
		jsonRow := "{"
		for i2, value := range row {
			//object不存第一个数据
			//第一个为id索引
			if i2 == 0 {
				objectKey = value
				continue
			}
			if i2 >= len(getRows[1]) {
				break
			}
			fieldType := getRows[2][i2]
			cell := GetCell(getRows[1][i2], value, fieldType)
			jsonRow = jsonRow + cell + ","
		}
		//去除逗号
		jsonRow = jsonRow[0 : len(jsonRow)-1]
		jsonRow += "}"
		if objectKey != "" {
			jsonMap[objectKey] = jsonRow
		}
	}
	return autocode, jsonMap
}

func GetCell(key, val, fieldType string) string {
	cell := fmt.Sprintf(`"%v":"%v"`, key, val)
	switch fieldType {
	case "ai":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "ai64":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "ac":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "i":
		if val == "" {
			val = "0"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "aai":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "aai64":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "aac":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "aaf":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "af":
		if val == "" {
			val = "[]"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	case "f":
		if val == "" {
			val = "0"
		}
		cell = fmt.Sprintf(`"%v":%v`, key, val)
		break
	}
	return cell
}
