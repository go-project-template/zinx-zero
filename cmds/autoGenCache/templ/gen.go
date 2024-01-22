package templ

import (
	"autoGenCache/model"
	"autoGenCache/util/pathx"
	"autoGenCache/util/stringx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"text/template"
	"time"
	"unicode"
)

func GenAll(val *model.RedisCacheKey) {
	if val.Prefix1 == "" || val.RetStruct == "" {
		logx.Must(errors.New("Prefix1|RetStruct can't be empty"))
	}
	checkStartByLetter(val.Prefix1)
	checkStartByLetter(val.Prefix2)
	checkStartByLetter(val.Prefix3)
	checkStartByLetter(val.Prefix4)
	checkStartByLetter(val.Prefix5)
	checkStartByLetter(val.Prefix6)
	checkStartByLetter(val.ArgInt64)
	checkStartByLetter(val.ArgString)
	checkStartByLetter(val.RetStruct)

	dir := "cachex/" + val.Prefix1
	err := pathx.MkdirIfNotExist(dir)
	logx.Must(err)
	var fileName = stringx.From(val.Prefix2)
	addFileName := func(prefix string) {
		if fileName.Source() == "" {
			logx.Must(errors.New("before fileName cann't empty"))
		}
		if prefix != "" {
			fileName = stringx.From(fmt.Sprintf("%v_%v", fileName.Source(), prefix))
		}
	}
	addFileName(val.Prefix3)
	addFileName(val.Prefix4)
	addFileName(val.Prefix5)
	addFileName(val.Prefix6)
	if fileName.IsEmptyOrSpace() {
		fileName = stringx.From(val.Prefix1)
	}

	if val.Expiry <= 0 || time.Duration(val.Expiry)*time.Second > time.Hour*24*30 {
		logx.Must(errors.New("Expiry must > 0 && <= 30 day"))
	}

	if val.NotFountExpiry <= 0 || time.Duration(val.NotFountExpiry)*time.Second > time.Hour {
		logx.Must(errors.New("NotFountExpiry must > 0 && <= 1 hour"))
	}

	genCache(val, dir, fileName)
	genCacheGen(val, dir, fileName)
	genTrace()
	genCachedData()
}

func genCache(val *model.RedisCacheKey, dir string, fileName stringx.String) {
	temp, err := template.New("cacheTpl").Parse(Cache)
	logx.Must(err)
	file, err := pathx.CreateIfNotExist(fmt.Sprintf(dir+"/%v.go", fileName.Untitle()))
	logx.Must(err)
	err = temp.Execute(file, map[string]any{
		"pkg":                   val.Prefix1,
		"upperStartCamelObject": fileName.ToCamel(),
		"lowerStartCamelObject": fileName.Untitle(),
		"args":                  genArgs(val),
		"expiry":                val.Expiry,
		"notFoundExpiry":        val.NotFountExpiry,
	})
	logx.Must(err)
}

func genCacheGen(val *model.RedisCacheKey, dir string, fileName stringx.String) {
	var cacheKeyField = stringx.From(val.Prefix1).ToCamel()
	var cacheKeyVal = val.Prefix1
	addCacheKey := func(prefix string) {
		if prefix != "" {
			cacheKeyField += stringx.From(prefix).ToCamel()
			cacheKeyVal = fmt.Sprintf("%v:%v", cacheKeyVal, prefix)
		}
	}
	addCacheKey(val.Prefix2)
	addCacheKey(val.Prefix3)
	addCacheKey(val.Prefix4)
	addCacheKey(val.Prefix5)
	addCacheKey(val.Prefix6)
	cacheKeyField = fmt.Sprintf(`cache%vPrefix`, cacheKeyField)

	cacheKey := fmt.Sprintf(`%v = "cache:%v"`, cacheKeyField, cacheKeyVal)
	temp, err := template.New("cacheTpl").Parse(CacheGen)
	logx.Must(err)
	file, err := pathx.CreateIfNotExist(fmt.Sprintf(dir+"/%v_gen.go", fileName.Untitle()))
	logx.Must(err)
	err = temp.Execute(file, map[string]any{
		"pkg":                   val.Prefix1,
		"cacheKey":              cacheKey,
		"upperStartCamelObject": fileName.ToCamel(),
		"lowerStartCamelObject": fileName.Untitle(),
		"args":                  genArgs(val),
		"args_use":              genArgsUse(val),
		"ret_struct":            val.RetStruct,
		"cacheKeyField":         cacheKeyField,
	})
	logx.Must(err)
}

func genTrace() {
	temp, err := template.New("cacheTpl").Parse(Trace)
	logx.Must(err)
	file, err := pathx.CreateIfNotExist(fmt.Sprintf("cachex/trace.go"))
	logx.Must(err)
	err = temp.Execute(file, map[string]any{})
	logx.Must(err)
}

func genCachedData() {
	temp, err := template.New("cacheTpl").Parse(CachedData)
	logx.Must(err)
	file, err := pathx.CreateIfNotExist(fmt.Sprintf("cachex/cachedData.go"))
	logx.Must(err)
	err = temp.Execute(file, map[string]any{})
	logx.Must(err)
}

func genArgs(val *model.RedisCacheKey) string {
	res := addArg(val.ArgInt64, " int64", "")
	res = addArg(val.ArgString, " string", res)
	return res
}

func genArgsUse(val *model.RedisCacheKey) string {
	res := addArg(val.ArgInt64, "", "")
	res = addArg(val.ArgString, "", res)
	return res
}

func addArg(_sqlArg string, end, res string) string {
	if _sqlArg != "" {
		for _, arg := range strings.Split(_sqlArg, ",") {
			checkStartByLetter(arg)
			// check repeat arg name
			for _, _checkContains := range strings.Split(res, ",") {
				if strings.Contains(_checkContains, arg) {
					logx.Must(errors.New("repeat arg name"))
				}
			}
			// add arg
			if res == "" {
				res += fmt.Sprintf("%v", arg)
			} else {
				res += fmt.Sprintf(", %v", arg)
			}
		}
		res += end
	}
	return res
}

func checkStartByLetter(text string) {
	if len(text) == 0 {
		return
	}
	// check start by letter
	r := rune(text[0])
	if !unicode.IsUpper(r) && !unicode.IsLower(r) {
		logx.Must(errors.New("must start by letter"))
	}
}
