release_win64:
		@echo "Building for Windows 64-bit"
		@GOOS=windows GOARCH=amd64 go build -o build/Release/Win64/GoOfLife.exe cmd/main/main.go

release_darwin_arm64:
		@echo "Building for Darwin ARM64"
		@GOOS=darwin GOARCH=arm64 go build -o build/Release/DarwinARM64/GoOfLife cmd/main/main.go
