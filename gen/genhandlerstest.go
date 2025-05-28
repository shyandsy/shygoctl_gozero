package gen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/shyandsy/shygoctl/api/spec"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/util"
	"github.com/shyandsy/shygoctl/util/format"
	"github.com/shyandsy/shygoctl/util/pathx"
	"github.com/shyandsy/shygoctl_gozero/template"
)

func genHandlerTest(dir, rootPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	handler := getHandlerName(route)
	handlerPath := getHandlerFolderPath(group, route)
	pkgName := handlerPath[strings.LastIndex(handlerPath, "/")+1:]
	logicName := defaultLogicPackage
	if handlerPath != handlerDir {
		handler = strings.Title(handler)
		logicName = pkgName
	}
	filename, err := format.FileNamingFormat(cfg.NamingFormat, handler)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          getHandlerFolderPath(group, route),
		filename:        filename + "_test.go",
		templateName:    "handlerTestTemplate",
		category:        category,
		templateFile:    handlerTestTemplateFile,
		builtinTemplate: template.HandlerTestTemplate,
		data: map[string]any{
			"PkgName":        pkgName,
			"ImportPackages": genHandlerTestImports(group, route, rootPkg),
			"HandlerName":    handler,
			"RequestType":    util.Title(route.RequestTypeName()),
			"ResponseType":   util.Title(route.ResponseTypeName()),
			"LogicName":      logicName,
			"LogicType":      strings.Title(getLogicName(route)),
			"Call":           strings.Title(strings.TrimSuffix(handler, "Handler")),
			"HasResp":        len(route.ResponseTypeName()) > 0,
			"HasRequest":     len(route.RequestTypeName()) > 0,
			"HasDoc":         len(route.JoinedDoc()) > 0,
			"Doc":            getDoc(route.JoinedDoc()),
		},
	})
}

func genHandlersTest(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	for _, group := range api.Service.Groups {
		for _, route := range group.Routes {
			if err := genHandlerTest(dir, rootPkg, cfg, group, route); err != nil {
				return err
			}
		}
	}

	return nil
}

func genHandlerTestImports(group spec.Group, route spec.Route, parentPkg string) string {
	imports := []string{
		//fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, getLogicFolderPath(group, route))),
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, contextDir)),
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, configDir)),
	}
	if len(route.RequestTypeName()) > 0 {
		imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, typesDir)))
	}

	return strings.Join(imports, "\n\t")
}
