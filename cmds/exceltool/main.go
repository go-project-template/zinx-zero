package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	"zinx-zero/cmds/exceltool/excelutils"

	"github.com/xuri/excelize/v2"
)

var dirPath = "./excel/"

// map[文件名]map[sheet名]解析后的数据
var excelDataMap = make(map[string]map[string]excelutils.SheetData)

// 输出自定义文件夹
var dirGo = []string{"./output/go/cfg/"}
var dirJson = []string{"./output/json/", "./output/go/conf/excel/"}
var dirTs = []string{"./output/ts/"}

func main() {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	excelPattern := regexp.MustCompile(`^.*\.xlsx$`)
	// 解析所有的excel文件到内存中
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			continue
		}
		fileName := fileInfo.Name()
		if !excelPattern.MatchString(fileName) {
			continue
		}
		if strings.Contains(fileName, "~$") {
			continue
		}
		if strings.Contains(fileName, ".~") {
			continue
		}
		//打开 excel 文件
		data, err := excelize.OpenFile(dirPath + fileName)
		if err != nil {
			panic(err)
		}
		fmt.Println(fileName)
		//解析 excel 中的数据到内存中
		excelDataMap[fileName] = excelutils.GetExcelAllSheetData(fileName, data)
	}
	
	//检测数值表
	excelutils.CheckExcel(excelDataMap)

	//转换成文件
	type FileRelation struct {
		FileName string
		Relation map[string]string
	}
	//获取所有的 structName 初始化时使用
	var structNameList []string
	//将内存中的数据 转换为 具体文件
	for fileName, sheetDataList := range excelDataMap {
		fileRelation := FileRelation{FileName: fileName, Relation: map[string]string{}}
		for sheetName, sheetData := range sheetDataList {
			structNameList = append(structNameList, sheetData.StructName)
			fileRelation.Relation[sheetName] = sheetData.StructName
			//生成go文件
			for _, dir := range dirGo {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					panic(err)
				}
				create, err := os.Create(dir + sheetData.StructName + ".go")
				if err != nil {
					panic("create file failed, err:" + err.Error())
				}
				files, err := template.ParseFiles("./config.go.tpl")
				if err != nil {
					panic("create template failed, err:" + err.Error())
				}
				err = files.Execute(create, sheetData)
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
				jsonFile, err := os.Create(dir + sheetData.StructName + ".json")
				if err != nil {
					panic("create file failed, err:" + err.Error())
				}
				_, err = jsonFile.Write([]byte(sheetData.JsonData))
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
				jsonFile, err := os.Create(dir + sheetData.StructName + ".ts")
				if err != nil {
					panic("create file failed, err:" + err.Error())
				}

				jsonData := `import { ConfigList } from '../common/Init'; ` +
					`ConfigList.` + sheetData.StructName + " = " + sheetData.JsonData

				_, err = jsonFile.Write([]byte(jsonData))
				if err != nil {
					panic("create file failed, err:" + err.Error())
				}
				jsonFile.Close()
			}
		}
	}
	//生成 ainit.go 文件
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
