IMPORT_PATH=github.com/ronald-tr/todonotion
BIN_NAME=todonotion

test:
	ginkgo tests

# Usage: make debug.'Test name that i want to run'
debug.%:
	dlv test ./tests -- --ginkgo.focus $*

# Golang build pipeline
go_build_linux_386:
	env GOOS=linux GOARCH=386 go build ${IMPORT_PATH}
	tar -czvf ${BIN_NAME}-linux-386.tar.gz ${BIN_NAME}
	@echo 'Done.'

go_build_mac_arm64:
	env GOOS=darwin GOARCH=arm64 go build ${IMPORT_PATH}
	tar -czvf ${BIN_NAME}-mac-arm64.tar.gz ${BIN_NAME}
	@echo 'Done.'

go_build_windows_386:
	env GOOS=windows GOARCH=386 go build ${IMPORT_PATH}
	tar -czvf ${BIN_NAME}-windows-386.tar.gz ${BIN_NAME}
	@echo 'Done.'

go_build_all: go_build_linux_386 go_build_mac_arm64 go_build_windows_386

# docs
doc-get-theme:
	mkdir -p docs/themes/hugo-geekdoc
	curl -L https://github.com/thegeeklab/hugo-geekdoc/releases/latest/download/hugo-geekdoc.tar.gz | tar -xz -C docs/themes/hugo-geekdoc/ --strip-components=1
