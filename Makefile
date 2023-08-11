run:
		@go run -tags noaudio .

test:
		@go test -v -tags noaudio ./...

release_linux_amd64:
		@GOOS=linux GOARCH=amd64 \
		make _build_release

release_darwin_arm64:
		@GOOS=darwin GOARCH=arm64 \
		make _build_release

release_darwin_amd64:
		@GOOS=darwin GOARCH=amd64 \
		make _build_release

release_win64:
		@GOOS=windows GOARCH=amd64 BUILD_OUT_FILE_EXT=.exe \
		make _build_release

_build_release:
		@echo "Building for ${GOOS} ${GOARCH} Ext: ${BUILD_OUT_FILE_EXT}/"
		@go build -o build/GoOfLife-$(GOOS)-$(GOARCH)$(BUILD_OUT_FILE_EXT) -tags noaudio .

# Alias for GitHub Actions Build Matrix
github_actions_windows-latest:
		@make release_win64

github_actions_macos-latest:
		@make release_darwin_amd64

github_actions_ubuntu-latest:
		@make release_linux_amd64
