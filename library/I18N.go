package library

import (
	"fmt"
	"github.com/lifenglin/micro-library/helper"
	"golang.org/x/net/context"
	"errors"
	"strings"
)

var HaveGetParamList = false
var ParamsList map[string]interface{}

func init() {
	err := getParamList()
	if nil == err {
		HaveGetParamList = true
	}
}

func getParamList() error {
	hlp := new(helper.Helper)
	ParamsListMap, err := GetVocabularyParamsByKey(context.Background(), hlp, "vocabulary_ii8n")
	for _, one := range ParamsListMap {
		ParamsList = one
		break
	}
	return err
}

func LiteralLang(key string, language string, argv map[string]string) (string, error) {
	if !HaveGetParamList {
		err := getParamList()
		if nil != err {
			return "", err
		}

		HaveGetParamList = true
	}

	if "" == language {
		language = "cn"
	}

	if languageData, ok := ParamsList[language]; ok {
		data, err := parse(languageData, key)
		if nil == err {
			data = placeholderReplace(data, argv)
			return data, nil
		}

		if language == "en" {
			return "", err
		}
	}

	if languageData, ok := ParamsList["en"]; ok {
		data, err := parse(languageData, key)
		if nil == err {
			data = placeholderReplace(data, argv)
			return data, nil
		} else {
			return "", err
		}
	}

	return "", errors.New("not find data")
}

func parse(languageData interface{}, key string) (string, error) {
	if paramsMap, ok := languageData.(map[string]interface{}); ok {
		if param, ok := paramsMap[key]; ok {
			if paramString, ok := param.(string); ok {
				return paramString, nil
			} else {
				return "", fmt.Errorf("param is not string, is %T", param)
			}
		} else {
			return "", fmt.Errorf("not find key: %s", key)
		}
	}
	
	return "", errors.New("parse paramsMap failed")
}

func placeholderReplace(target string, argv map[string]string) string {
	for key, item := range argv {
		oldString := fmt.Sprintf("{%s}", key)
		target = strings.Replace(target, oldString, item, -1)
	}
	return target
}