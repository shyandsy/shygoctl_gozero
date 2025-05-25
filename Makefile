build-plugin:
	CGO_ENABLED=1 go build -buildmode=plugin -gcflags="all=-N -l" -o gozero_template_plugin.so plugin.go

build-test:
	go build -ldflags="-s -w" main.go
	#$(if $(shell command -v upx || which upxll), upx goctl)