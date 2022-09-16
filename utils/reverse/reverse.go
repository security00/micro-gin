package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"gorm.io/gorm"
	"io"
	. "micro-gin/bootstrap"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Field struct {
	Field      string `gorm:"column:Field"`
	Type       string `gorm:"column:Type"`
	Null       string `gorm:"column:Null"`
	Key        string `gorm:"column:Key"`
	Default    string `gorm:"column:Default"`
	Extra      string `gorm:"column:Extra"`
	Privileges string `gorm:"column:Privileges"`
	Comment    string `gorm:"column:Comment"`
}

type Table struct {
	Name    string `gorm:"column:Name"`
	Comment string `gorm:"column:Comment"`
}

var (
	db        *gorm.DB
	dbNames   = "test"
	modelPath = "./utils/reverse/result/models/"
	jsonPath  = "./utils/reverse/result/jsons/"
	wg        = sync.WaitGroup{}
)

func main() {
	//var reTmp = "([a-z]+)(\\(?)(\\d+)?(\\)?)(\\s?)([a-z]+)?"
	//var contens = "int(10) unsigned"
	//re, _ := regexp.Compile(reTmp)
	//res := re.FindAllStringSubmatch(contens, 10)
	//for i, v := range res[0] {
	//	fmt.Println(i, "       ", v)
	//}
	//fmt.Println(strings.Contains(contens, "unsigned"))
	//return
	//for _, t := range getTables("test") {
	//	fmt.Println(t.Name, "->", t.Comment)
	//	for _, f := range getFields(t.Name) {
	//		var typeInfo = reflect.TypeOf(f)
	//		var valueInfo = reflect.ValueOf(f)
	//		num := typeInfo.NumField()
	//		for i := 0; i < num; i++ {
	//			key := typeInfo.Field(i).Name
	//			val := valueInfo.Field(i).Interface()
	//			fmt.Printf("   %v --- %v \n", key, val)
	//		}
	//	}
	//}
	Generate(dbNames)
}

func Generate(dbNames string) {
	tables := getTables(dbNames) //生成所有表信息

	for _, table := range tables {
		wg.Add(2)
		fields := getFields(table.Name)
		generateModel(table, fields)
		generateJSON(table, fields)
	}
	wg.Wait()
}

func getTables(dbNames string) []Table {
	var tables []Table

	if len(dbNames) > 0 {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema = '" + dbNames + "';").Find(&tables)
	}
	return tables
}

func init() {
	InitializeConfig()
	db = InitializeDB()
}

func getFields(tableName string) []Field {
	var fields []Field
	if len(tableName) > 0 {
		db.Raw("SHOW FULL COLUMNS FROM " + tableName + ";").Find(&fields)
	}
	return fields
}

// 生成Model
func generateModel(table Table, fields []Field) {
	defer wg.Done()
	content := "package models\n\n"
	//表注释
	if len(table.Comment) > 0 {
		content += "// " + table.Comment + "\n"
	}
	content += "type " + generator.CamelCase(table.Name) + " struct {\n"
	//生成字段
	for _, field := range fields {
		fieldName := generator.CamelCase(field.Field)
		//fieldJson := getFieldJson(field)
		fieldGorm := getFieldGorm(field)
		fieldType := getFiledType(field)
		fieldComment := getFieldComment(field)
		content += "	" + fieldName + " " + fieldType + " `" + fieldGorm + "` " + fieldComment + "\n"
	}
	content += "}\n"

	content += "func (entity *" + generator.CamelCase(table.Name) + ") TableName() string {\n"
	content += "	" + `return "` + table.Name + `"`
	content += "\n}"

	filename := modelPath + table.Name + ".go"
	var f *os.File
	var err error
	if checkFileIsExist(filename) {
		fmt.Println(table.Name + " 已存在，需删除才能重新生成...")
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666) //打开文件
		if err != nil {
			panic(err)
		}
	}
	f, err = os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(generator.CamelCase(table.Name) + " 已生成...")
	}

}

// 获取字段类型
func getFiledType(field Field) string {
	typeArr := strings.Split(field.Type, "(")
	typeArr1 := strings.Split(field.Type, ")")

	switch typeArr[0] {
	case "int":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint32"
		} else {
			return "*int32"
		}
	case "integer":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint32"
		} else {
			return "*int32"
		}
	case "mediumint":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint32"
		} else {
			return "*int32"
		}
	case "bit":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint32"
		} else {
			return "*int32"
		}
	case "year":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint32"
		} else {
			return "*int32"
		}
	case "smallint":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint16"
		} else {
			return "*int16"
		}
	case "tinyint":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint8"
		} else {
			return "*int8"
		}
	case "bigint":
		if len(typeArr1) > 1 && typeArr1[1] == " unsigned" {
			return "*uint64"
		} else {
			return "*int64"
		}
	case "decimal":
		return "*float64"
	case "double":
		return "*float32"
	case "float":
		return "*float32"
	case "real":
		return "*float32"
	case "numeric":
		return "*float32"
	case "timestamp":
		return "*time.Time"
	case "datetime":
		return "*jsontime.JsonTime"
	case "time":
		return "*time.Time"
	case "date":
		return "*time.Time"
	default:
		return "*string"
	}
}

// 获取字段json描述
func getFieldJson(field Field) string {
	return `json:"` + Lcfirst(generator.CamelCase(field.Field)) + `"`
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 获取字段gorm描述
func getFieldGorm(field Field) string {
	fieldContext := `gorm:"column:` + field.Field

	if field.Key == "PRI" {
		fieldContext = fieldContext + `;primaryKey`
	}
	if field.Key == "UNI" {
		fieldContext = fieldContext + `;unique`
	}
	if field.Extra == "auto_increment" {
		fieldContext = fieldContext + `;autoIncrement`
	}
	if field.Null == "NO" {
		fieldContext = fieldContext + `;not null`
	}
	return fieldContext + `"`
}

// 获取字段说明
func getFieldComment(field Field) string {
	if len(field.Comment) > 0 {
		//return "// " + field.Comment
		return "//" + strings.Replace(strings.Replace(field.Comment, "\r", "\\r", -1), "\n", "\\n", -1)
	}
	return ""
}

// 检查文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 生成JSON
func generateJSON(table Table, fields []Field) {
	defer wg.Done()
	content := "package reply\n\n"

	content += "type " + generator.CamelCase(table.Name) + " struct {\n"
	//生成字段
	for _, field := range fields {
		fieldName := generator.CamelCase(field.Field)
		fieldJson := getFieldJson(field)
		fieldType := getFiledType(field)
		content += "	" + fieldName + " " + fieldType + " `" + fieldJson + "` " + "\n"
	}
	content += "}\n"

	var f *os.File
	var err error
	filename := jsonPath + table.Name + ".go"
	if checkFileIsExist(filename) {
		fmt.Println(generator.CamelCase(table.Name) + " 已存在，需删除才能重新生成...")
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666) //打开文件
		if err != nil {
			panic(err)
		}
	}
	f, err = os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(table.Name + ".go 已生成...")
	}

}
