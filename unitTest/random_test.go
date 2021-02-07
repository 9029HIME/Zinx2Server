package unitTest

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"testing"
	config2 "zinx2server/config"
)

func TestYaml(t *testing.T) {
	data, err := ioutil.ReadFile("../config/application.yaml")
	if err != nil {
		fmt.Println("读配置文件错误：", err)
		return
	}
	content := make(map[string]interface{})
	yaml.Unmarshal(data, content)
	server := content["Server"].(map[interface{}]interface{})
	for k, v := range server {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
		fmt.Printf("%T\n", v)
		fmt.Println()
	}
}

func TestReflectName(t *testing.T) {
	config := new(config2.ServerConfig)
	typ := reflect.TypeOf(config)
	fieldSet := make(map[string]int)
	if typ.Kind() == reflect.Ptr {
		// 获取的是指针
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldSet[fieldName] = 1
	}
	for k, v := range fieldSet {
		fmt.Println("K:", k)
		fmt.Println("V:", v)
		fmt.Println()
	}
}

func TestAutoConfig(t *testing.T) {
	//config := config2.Init("")
	config := config2.Init("config/application.yaml")
	fmt.Println(config)
}
