package main

import (
	"encoding/json"
	"fmt"
	"os"
	"plugin"

	"github.com/shyandsy/shygoctl/api/spec"
)

func main() {
	bytes, err := os.ReadFile("test_api_spec.json")
	if err != nil {
		panic(err)
	}

	s := &spec.ApiSpec{}
	if err := json.Unmarshal(bytes, s); err != nil {
		panic(err)
	}

	// load plugin
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(fmt.Sprintf("插件加载失败: %v", err))
	}

	// 查找方法符号
	//DoGenProject
	method, err := p.Lookup("DoGenProject")
	if err != nil {
		panic(fmt.Sprintf("符号查找失败: %v", err))
	}

	doGenProject, ok := method.(func(api *spec.ApiSpec, dir, style string) error)
	if !ok {
		panic("函数签名不匹配，请检查参数类型")
	}

	if err := doGenProject(s, "demo", "goZero"); err != nil {
		panic("函数签名不匹配，请检查参数类型")
	}
	print("ok")
}
