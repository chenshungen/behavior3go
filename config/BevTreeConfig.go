package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
)

//编辑器地址@http://editor.behavior3.com/#/editor
//节点json类型
type BTNodeCfg struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Category    string                 `json:"category"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Children    []string               `json:"children"`
	Child       string                 `json:"child"`
	Parameters  map[string]interface{} `json:"parameters"`
	Properties  map[string]interface{} `json:"properties"`
}

func (this *BTNodeCfg) GetProperty(name string) float64 {
	v, ok := this.Properties[name]
	if !ok {
		panic("GetProperty err ,no vlaue:" + name)
	}
	f64, fok := v.(float64)
	if !fok {
		fmt.Println("GetProperty err ,format not fload64:", name, v)
		panic("GetProperty err ,format not fload64:" + name)
	}
	return f64
}

func (this *BTNodeCfg) GetPropertyAsInt(name string) int {
	v := this.GetProperty(name)
	i := int(v)
	return i
}
func (this *BTNodeCfg) GetPropertyAsInt64(name string) int64 {
	v := this.GetProperty(name)
	i := int64(v)
	return i
}
func (this *BTNodeCfg) GetPropertyAsBool(name string) bool {
	v, ok := this.Properties[name]
	if !ok {
		//panic("GetProperty err ,no vlaue:" + name)
		return false
	}

	b, fok := v.(bool)
	if !fok {
		if str, sok := v.(string); sok {
			return str == "true"
		}
		fmt.Println("GetProperty err ,format not bool:", name, v)
		panic("GetProperty err ,format not bool:" + name)
	}
	return b
}

func (this *BTNodeCfg) GetPropertyAsString(name string) string {
	v, ok := this.Properties[name]
	if !ok {
		panic("GetProperty err ,no vlaue:" + name)
	}

	strCast := func(v interface{}) string {
		switch v := v.(type) {
		case nil:
			return ""
		case string:
			return v
		case []byte:
			return string(v)
		case int, int8, int16, int32, int64:
			return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
		case float32:
			return strconv.FormatFloat(float64(v), 'f', 3, 32)
		case float64:
			return strconv.FormatFloat(float64(v), 'f', 3, 64)
		case bool:
			return strconv.FormatBool(v)
		default:
			return ""
		}
	}

	return strCast(v)
}

//树json类型
type BTTreeCfg struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Root        string                 `json:"root"`
	Properties  map[string]interface{} `json:"properties"`
	Nodes       map[string]BTNodeCfg   `json:"nodes"`
}

//加载
func LoadTreeCfg(path string) (*BTTreeCfg, bool) {

	var tree BTTreeCfg
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("fail:", err)
		return nil, false
	}
	err = json.Unmarshal(file, &tree)
	if err != nil {
		fmt.Println("fail, ummarshal:", err, len(file))
		return nil, false
	}

	//fmt.Println("load tree:", tree.Title, " nodes:", len(tree.Nodes))
	return &tree, true
}
