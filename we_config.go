package wego

import (
	"log"
	"strconv"

	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

type Tree toml.Tree

type System struct {
	//debug = true
	Debug bool `toml:"debug"`
	//response_type = 'array'
	ResponseType string `toml:"response_type"`
	//use_cache = true
	UseCache bool `toml:"use_cache"`
	Log      Log
}

type Log struct {
	//level = 'debug'
	Level string
	//file = 'logs/wechat.log'
	File string
}

//type config struct {
//	Content *Tree
//}

type Config interface {
	Get(s string) string
	Set(k, v string) *Tree
	GetBool(s string) bool
	GetConfig(s string) Config
	GetTree(s string) interface{}
}

var useCache = false

func ConfigTree() *Tree {
	t, e := toml.LoadFile(*f)
	if e != nil {
		log.Println("filepath: " + *f)
		log.Println(e.Error())
		panic(FileLoadError)
	}
	return (*Tree)(t)
}

//
//func TransToMap(tree *toml.Tree, name string) map[string]Tree {
//	cm := make(map[string]Tree)
//	mp := tree.Get(name).(*toml.Tree).ToMap()
//	for key, value := range mp {
//		cm[key] = value.(map[string]Tree
//	}
//	return cm
//}

func treeLoader() *Tree {
	if system.UseCache {
		return configCache
	}
	return ConfigTree()
}

func GetConfig(s string) Config {
	c := treeLoader()
	if v, b := c.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return nil
}

func GetRootConfig() Config {
	return treeLoader()
}

func (t *Tree) GetConfig(s string) Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return nil
}

func (t *Tree) GetTree(s string) interface{} {
	return (*toml.Tree)(t).Get(s)
}

func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	if ParseInt(v) == -1 {
		return ""
	}
	return strconv.Itoa(ParseInt(v))
}

func (t *Tree) Set(k, v string) *Tree {
	tt := (*toml.Tree)(t)
	tt.Set(k, v)
	return t
}

func (t *Tree) GetBool(s string) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return false
}

func CacheOn() {
	useCache = true
}

func CacheOff() {
	useCache = false
}

func UseCache() bool {
	return useCache
}