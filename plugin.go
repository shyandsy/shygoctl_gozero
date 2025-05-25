package main

import (
	"github.com/shyandsy/shycgoctl_gozero/gen"
	"github.com/shyandsy/shygoctl/api/spec"
)

func DoGenProject(api *spec.ApiSpec, dir, style string) error {
	if err := gen.DoGenProject(api, dir, style); err != nil {
		return err
	}
	return nil
}
