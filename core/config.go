package core

import (
	"log"
	"strconv"
	"strings"

	"github.com/godcong/wego/cache"
	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

const (
	OFF = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

var logList = map[string]int{
	"OFF":   OFF,
	"FATAL": FATAL,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"ALL":   ALL,
}

type Tree toml.Tree

type System struct {
	//debug = true
	Debug bool `toml:"debug"`
	//response_type = 'array'
	ResponseType string `toml:"response_type"`
	//use_cache = true
	//DataType DataType `toml:"data_type"`

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
	GetD(s, d string) string
	Set(k, v string) *Tree
	GetBool(s string) bool
	GetConfig(s string) Config
	GetTree(s string) interface{}
}

var system System
var useCache = false

func ConfigTree(f string) *Tree {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		panic(FileLoadError)
	}
	return (*Tree)(t)
}

func initSystem(v interface{}) {
	v.(*toml.Tree).Unmarshal(&system)
}

func treeLoader() *Tree {
	c := cache.GetCache()
	if UseCache() {
		return c.Get("cache").(*Tree)
	}
	return ConfigTree(c.GetD("cache_path", "config.toml").(string))
}

func GetConfig(path string) Config {
	Debug("GetConfig|path", path)
	c := treeLoader()
	if v, b := c.GetTree(path).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return (*Tree)(nil)
}

func GetRootConfig() Config {
	return treeLoader()
}

func GetSystemConfig() System {
	return system
}

func (t *Tree) GetConfig(s string) Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return nil
}

func (t *Tree) GetTree(s string) interface{} {
	if t == nil {
		return nil
	}
	return (*toml.Tree)(t).Get(s)
}

func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	if v0 := ParseInt(v); v0 == -1 {
		return ""
	} else {
		return strconv.FormatInt(v0, 10)
	}
}

func (t *Tree) GetD(s, d string) string {
	if v := t.Get(s); v != "" {
		return v
	}
	return d
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

func DeployJoin(v ...string) string {
	return strings.Join(v, ".")
}

func (l *Log) LevelInt() (i int) {
	if v, b := logList[strings.ToUpper(l.Level)]; b {
		i = v
	}
	return i
}
