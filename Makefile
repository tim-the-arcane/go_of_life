release_win64:
		@echo "Building for Windows amd64"
		@GOOS=windows GOARCH=amd64 go build -o build/Release/Win64/GoOfLife.exe -tags noaudio cmd/main/main.go

release_darwin_amd64:
		@echo "Building for Darwin amd64"
		@GOOS=darwin GOARCH=amd64 go build -o build/Release/DarwinARM64/GoOfLife -tags noaudio cmd/main/main.go

release_linux_amd64:
		@echo "Building for Linux amd64"
		@GOOS=linux GOARCH=amd64 go build -o build/Release/LinuxAMD64/GoOfLife -tags noaudio cmd/main/main.go

# Alias for GitHub Actions Build Matrix
github_actions_windows-latest:
		@make release_win64

github_actions_macos-latest:
		@make release_darwin_amd64

github_actions_ubuntu-latest:
		@make release_linux_amd64
