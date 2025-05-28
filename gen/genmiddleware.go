package gen

import (
	_ "embed"
	"strings"

	"github.com/shyandsy/shygoctl/api/spec"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/util/format"
	"github.com/shyandsy/shygoctl_gozero/template"
)

func genMiddleware(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	middlewares := getMiddleware(api)
	for _, item := range middlewares {
		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware") + "_middleware"
		filename, err := format.FileNamingFormat(cfg.NamingFormat, middlewareFilename)
		if err != nil {
			return err
		}

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		err = genFile(fileGenConfig{
			dir:             dir,
			subdir:          middlewareDir,
			filename:        filename + ".go",
			templateName:    "contextTemplate",
			category:        category,
			templateFile:    middlewareImplementCodeFile,
			builtinTemplate: template.MiddlewareImplementCode,
			data: map[string]string{
				"name": strings.Title(name),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
