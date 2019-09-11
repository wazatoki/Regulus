PROJECT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))/
PROJECT_NAME := regulus

.PHONY: build
build: clean jsBuild
	cd $(PROJECT_DIR)golang/src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(PROJECT_DIR)build/$(PROJECT_NAME)_linux_amd64
	cd $(PROJECT_DIR)golang/src && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(PROJECT_DIR)build/$(PROJECT_NAME)_win_amd64.exe
	cd $(PROJECT_DIR)golang/src && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(PROJECT_DIR)build/$(PROJECT_NAME)_darwin_amd64
	cp -r $(PROJECT_DIR)golang/resources/* $(PROJECT_DIR)build/resources/

.PHONY: jsBuild
jsBuild:
	cd $(PROJECT_DIR)js && npm run build
	cp -r $(PROJECT_DIR)js/dist/* $(PROJECT_DIR)build/resources/

.PHONY: clean
clean:
	cd $(PROJECT_DIR)golang/src && go clean
	rm -rf $(PROJECT_DIR)build
	mkdir -p $(PROJECT_DIR)build/resources

.PHONY: dev
dev:
	rm -f $(PROJECT_DIR)build/$(PROJECT_NAME)_linux_amd64
	cd $(PROJECT_DIR)golang/src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(PROJECT_DIR)build/$(PROJECT_NAME)_linux_amd64
	cd $(PROJECT_DIR)build/ && ./$(PROJECT_NAME)_linux_amd64

.PHONY: run
run: build
	cd $(PROJECT_DIR)build/ && ./$(PROJECT_NAME)_linux_amd64
