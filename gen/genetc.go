package gen

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/shyandsy/shygoctl/api/spec"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/util/format"
	"github.com/shyandsy/shygoctl_gozero/template"
)

const (
	defaultPort = 8888
	etcDir      = "etc"
)

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, api.Service.Name)
	if err != nil {
		return err
	}

	service := api.Service
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("%s.yaml", filename),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: template.EtcTemplate,
		data: map[string]string{
			"serviceName": service.Name,
			"host":        host,
			"port":        port,
		},
	})
}
