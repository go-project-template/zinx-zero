package auto

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	IDType      string           `json:"id_type"`
	StructName  string           `json:"structName"`
	PackageName string           `json:"packageName"`
	Fields      []Field          `json:"fields"`
	Structs     []AutoCodeStruct `json:"structs"`
}

type Field struct {
	FieldName string `json:"fieldName"`
	FieldDesc string `json:"fieldDesc"`
	FieldType string `json:"fieldType"`
	FieldJson string `json:"fieldJson"`
}

func (a AutoCodeStruct) GetFieldType(val string) string {
	switch val {
	case "i64":
		return "int64"
	case "ai64":
		return "[]int64"
	case "aai64":
		return "[][]int64"
	case "i":
		return "int32"
	case "c":
		return "string"
	case "ai":
		return "[]int32"
	case "ac":
		return "[]string"
	case "aac":
		return "[][]string"
	case "aai":
		return "[][]int32"
	case "f":
		return "float64"
	case "af":
		return "[]float64"
	case "aaf":
		return "[][]float64"
	}
	return ""
}
