package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

type ServerConfig struct {
	ServerName string
	Host       string
	Port       string
	IPVersion  string
	// TODO 暂时无效
	MaxConnection int
	Valid         bool
}

func Init(configPath string) *ServerConfig {
	config := &ServerConfig{
		ServerName:    "server-0",
		Host:          "localhost",
		Port:          "6789",
		IPVersion:     "tcp4",
		MaxConnection: int(^uint32(0) >> 1),
	}
	config.reload(configPath)
	return config
}

// 用来加载新的配置，以覆盖初始配置
func (config *ServerConfig) reload(configPath string) {
	if configPath == "" {
		return
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("读配置文件错误：", err)
		return
	}
	yamlData := make(map[string]interface{})
	yaml.Unmarshal(data, yamlData)
	serverContent, ok := yamlData["Server"]
	if !ok {
		log.Println("错误的配置文件内容")
		return
	}
	configData := serverContent.(map[interface{}]interface{})
	config.fullfill(configData)
}

func (config *ServerConfig) fullfill(configData map[interface{}]interface{}) {
	names := config.getFieldNames()
	for k, v := range configData {
		if v == nil {
			continue
		}
		key := k.(string)
		_, ok := names[key]
		if !ok {
			continue
		}
		config.setValue(key, v)
	}
}

func (config *ServerConfig) getFieldNames() map[string]int {
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
	return fieldSet
}

func (config *ServerConfig) setValue(fieldName string, value interface{}) {
	typ := reflect.ValueOf(config)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field := typ.FieldByName(fieldName)
	switch value.(type) {
	case int:
		field.SetString(strconv.Itoa(value.(int)))
	case string:
		field.SetString(value.(string))
	case bool:
		field.SetBool(value.(bool))
	}
}
