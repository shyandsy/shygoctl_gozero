package gen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/shyandsy/shycgoctl_gozero/template"
	"github.com/shyandsy/shygoctl/api/spec"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/util/format"
	"github.com/shyandsy/shygoctl/vars"
)

const (
	configFile = "config"

	jwtTemplate = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
	jwtTransTemplate = ` struct {
		Secret     string
		PrevSecret string
	}
`
)

func genConfig(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}

	authNames := getAuths(api)
	var auths []string
	for _, item := range authNames {
		auths = append(auths, fmt.Sprintf("%s %s", item, jwtTemplate))
	}

	jwtTransNames := getJwtTrans(api)
	var jwtTransList []string
	for _, item := range jwtTransNames {
		jwtTransList = append(jwtTransList, fmt.Sprintf("%s %s", item, jwtTransTemplate))
	}
	authImportStr := fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: template.ConfigTemplate,
		data: map[string]string{
			"authImport": authImportStr,
			"auth":       strings.Join(auths, "\n"),
			"jwtTrans":   strings.Join(jwtTransList, "\n"),
		},
	})
}
