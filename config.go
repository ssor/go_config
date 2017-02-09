package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	json "github.com/laktak/hjson-go"
)

func LoadConfig(file string) (IConfigInfo, error) {
	bs, err := getFileContent(file)
	if err != nil {
		return nil, err
	}

	return ParseConfig(bs)
}

func ParseConfig(data []byte) (IConfigInfo, error) {
	var conf ConfigInfo
	err := json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	err = conf.validate_data_type()
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (conf ConfigInfo) validate_data_type() error {
	for key, value := range conf {
		switch value.(type) {
		case float64:
			conf[key] = int(value.(float64))
		case []interface{}:
			new_value, err := reset_data_type(value.([]interface{}))
			if err != nil {
				return err
			}
			conf[key] = new_value
		}
	}
	return nil
}

func reset_data_type(in []interface{}) (interface{}, error) {
	if in == nil || len(in) <= 0 {
		return nil, nil
	}
	switch in[0].(type) {
	case string:
		out := []string{}
		for _, v := range in {
			if sv, ok := v.(string); ok {
				out = append(out, sv)
			} else {
				return nil, fmt.Errorf("data type should be string for %s", v)
			}
		}
		return out, nil
	case float64:
		out := []int{}
		for _, v := range in {
			if iv, ok := v.(float64); ok {
				out = append(out, int(iv))
			} else {
				return nil, fmt.Errorf("data type should be int for %s", v)
			}
		}
		return out, nil
	}
	return nil, nil
}

type IConfigInfo interface {
	Get(string) interface{}
	Set(string, interface{})
	String() string
}

type ConfigInfo map[string]interface{}

func (conf ConfigInfo) Dump() {
	for key, value := range conf {
		fmt.Println("key -> ", key)
		spew.Dump(value)
	}
}
func (conf ConfigInfo) String() string {
	r := ""
	for k, v := range conf {
		r = fmt.Sprintf("%s\r\n%s -> %s", r, k, v)
	}
	return r
}

func (conf ConfigInfo) Set(key string, val interface{}) {
	conf[key] = val
}

func (conf ConfigInfo) Get(key string) interface{} {
	if v, exists := conf[(key)]; exists {
		return v
	}
	return nil
}

// get string from text file
func getFileContent(file string) ([]byte, error) {
	if !isFileExist(file) {
		return nil, os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return nil, e
	}
	return b, nil
}

// exists returns whether the given file or directory exists or not
func isFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
