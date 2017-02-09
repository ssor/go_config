package config

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestConfig(t *testing.T) {
	kvs := map[string]interface{}{
		"string":     "8081",
		"int":        10,
		"stringList": []string{"a", "abc"},
		"intList":    []int{1, 2},
	}

	conf, err := LoadConfig("./data.json")
	if err != nil {
		panic(err)
	}

	conf.(ConfigInfo).Dump()

	key := "string"
	conf_value := conf.Get(key)
	kvs_value := kvs[key]
	if conf_value.(string) != kvs_value {
		t.Logf("get key %s failed ", key)
		spew.Dump(kvs_value)
		spew.Dump(conf_value)
		t.FailNow()
	}

	key = "int"
	conf_value = conf.Get(key)
	kvs_value = kvs[key]
	if conf_value.(int) != kvs_value {
		t.Logf("get key %s failed ", key)
		spew.Dump(kvs_value)
		spew.Dump(conf_value)
		t.FailNow()
	}
	key = "stringList"
	conf_value = conf.Get(key)
	kvs_value = kvs[key]
	if string_slice_equal(conf_value.([]string), kvs_value.([]string)) == false {
		t.Logf("get key %s failed ", key)
		spew.Dump(kvs_value)
		spew.Dump(conf_value)
		t.FailNow()
	}

	key = "intList"
	conf_value = conf.Get(key)
	kvs_value = kvs[key]
	if int_slice_equal(conf_value.([]int), kvs_value.([]int)) == false {
		t.Logf("get key %s failed ", key)
		spew.Dump(kvs_value)
		spew.Dump(conf_value)
		t.FailNow()
	}
}

func int_slice_equal(slice1, slice2 []int) bool {
	if slice1 == nil && slice2 == nil {
		return true
	}

	if slice1 == nil && slice2 != nil {
		return false
	}

	if slice2 == nil && slice1 != nil {
		return false
	}

	for _, s := range slice1 {
		if int_in_slice(slice2, s) == false {
			return false
		}
	}

	for _, s := range slice2 {
		if int_in_slice(slice1, s) == false {
			return false
		}
	}

	return true
}

func string_slice_equal(slice1, slice2 []string) bool {
	if slice1 == nil && slice2 == nil {
		return true
	}

	if slice1 == nil && slice2 != nil {
		return false
	}

	if slice2 == nil && slice1 != nil {
		return false
	}

	for _, s := range slice1 {
		if string_in_slice(slice2, s) == false {
			return false
		}
	}

	for _, s := range slice2 {
		if string_in_slice(slice1, s) == false {
			return false
		}
	}

	return true
}

func string_in_slice(slice []string, s string) bool {
	if slice == nil {
		return false
	}

	for _, ss := range slice {
		if ss == s {
			return true
		}
	}
	return false
}

func int_in_slice(slice []int, s int) bool {
	if slice == nil {
		return false
	}

	for _, ss := range slice {
		if ss == s {
			return true
		}
	}
	return false
}
