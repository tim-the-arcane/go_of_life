release_linux_amd64:
		@echo "Building for Linux amd64"
		@GOOS=linux GOARCH=amd64 \
		make build_release

release_darwin_arm64:
		@echo "Building for Darwin arm64"
		@GOOS=darwin GOARCH=arm64 \
		make build_release

release_darwin_amd64:
		@echo "Building for Darwin amd64"
		@GOOS=darwin GOARCH=amd64 \
		make build_release

release_win64:
		@echo "Building for Windows amd64"
		@GOOS=windows GOARCH=amd64 BUILD_OUT_FILE_EXT=.exe \
		make build_release

build_release:
		@echo "Building for ${GOOS} ${GOARCH} Ext: ${BUILD_OUT_FILE_EXT}/"
		@go build -o build/GoOfLife-$(GOOS)-$(GOARCH)$(BUILD_OUT_FILE_EXT) -tags noaudio cmd/main/main.go

# Alias for GitHub Actions Build Matrix
github_actions_windows-latest:
		@make release_win64

github_actions_macos-latest:
		@make release_darwin_amd64

github_actions_ubuntu-latest:
		@make release_linux_amd64
