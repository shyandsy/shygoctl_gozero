build:
	go build -gcflags="all=-N -l" -o build/shygoctl_gozero main.go

build-test:
	go build -ldflags="-s -w" main.go
	#$(if $(shell command -v upx || which upxll), upx goctl)