.PHONY:build

build:
	wails build -p -x darwin/amd64
	wails build -p -x linux/amd64
	wails build -p -x linux/arm-7
	wails build -p -x windows/amd64

release: build

serve:
	wails serve