package main

import (
	"encoding/json"
	"os"

	"github.com/shyandsy/shygoctl/api/spec"
	"github.com/shyandsy/shygoctl_gozero/gen"
)

func main() {
	bytes, err := os.ReadFile("api_specification.json")
	if err != nil {
		panic(err)
	}

	s := &spec.ApiSpec{}
	if err := json.Unmarshal(bytes, s); err != nil {
		panic(err)
	}

	if err := gen.DoGenProject(s, "demo", "goZero"); err != nil {
		panic("函数签名不匹配，请检查参数类型")
	}
	print("ok")
}
