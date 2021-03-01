package controllerfactory

import (
	"../../httpserver"
	"../../utils/basicutils"
	"../../utils/logger"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var rootRouter = httpserver.HttpServletRouter{}

func GetRouter() *httpserver.HttpServletRouter {
	return &rootRouter
}

func BuildController(apis []interface{}) {
	for _, api := range apis {
		packageName, funcName := getFuncDir(api)
		funcList := basicutils.ListPublicFuncInPackage(packageName)
		for _, f := range funcList {
			if strings.Compare(f.Name.Name, funcName) == 0 && f.Doc != nil && f.Doc.List != nil && f.Type != nil && f.Type.Params != nil {
				execute := httpserver.HttpServletExecute{}
				for _, comment := range f.Doc.List {
					commentContent := strings.TrimSpace(string([]byte(comment.Text)[2:]))
					index := strings.Index(commentContent, "=")
					if index < 0 {
						continue
					}
					key := strings.TrimSpace(commentContent[:index])
					value := strings.TrimSpace(commentContent[index+1:])
					if strings.Compare(key, "@method") == 0 {
						execute.Method = value
					} else if strings.Compare(key, "@path") == 0 {
						execute.Path = value
					}
				}

				if len(execute.Method) == 0 || len(execute.Path) == 0 {
					continue
				} else if len(execute.ContentType) == 0 {
					execute.ContentType = httpserver.HttpContentTypeJson
				}

				execute.Handler = api
				logger.Info("Register API: %s", packageName, funcName)
				rootRouter.SaveChild(&execute)
			}
		}
	}
}

var rootDir = ""

func getRootDir() string {
	if len(rootDir) < 1 {
		str, _ := os.Getwd()
		str = strings.ReplaceAll(str, "\\", "/")
		str = strings.ReplaceAll(str, ":", "_")
		rootDir = "_/" + str + "/"
	}
	return rootDir
}

func getFuncDir(f interface{}) (string, string) {
	rootPath := getRootDir()
	funcPath := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcPath = strings.ReplaceAll(funcPath, rootPath, "")
	funcPath = strings.ReplaceAll(funcPath, "/", ".")
	index := strings.LastIndex(funcPath, ".")
	packageName := string([]byte(funcPath)[:index])
	funcName := string([]byte(funcPath)[index+1:])
	return packageName, funcName
}
