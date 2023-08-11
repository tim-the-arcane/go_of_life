# Go of Life

A trippy interpretation of Conway's Game of Life

## Download

I provide pre-built binaries for Windows, Linux and macOS for the amd64 architecture. For other OS/arch combinations please check ["Building from source"](#building-from-source).

To download a pre-built binary go to the [latest release](https://github.com/tim-the-arcane/go_of_life/releases/latest) and download the correct binary for your system.

## Usage

### Linux

Make the file executable and start it

### MacOS

MacOS doesn't recognize me as certified developer and stops the execution. 

You have to allow the binary by first making it executable and then running it once:

```bash
chmod +x GoOfLife-darwin-amd64
./GoOfLife-darwin-amd64
```

After this you have to open the system settings of MacOS and go to the security tab. There should be a section about the recently run program. Click on "Run anyway".

```bash
chmod +x GoOfLife-darwin-amd64
./GoOfLife-darwin-amd64
```

### Windows

IDKTBH... Tell me if you find out! :D

## Building from Source

### Requirements

1. installed and onfigured golang toolchain
2. the requirements of [raylib-go](https://github.com/gen2brain/raylib-go) (on which this projects depends) for your environment

### Start build

Run the following command:

```bash
GOOS=[your target os] GOARCH=[your target arch] go build . -tags noaudio
```

### Examples

#### Build for Apple Silicon

```bash
GOOS=darwin GOARCH=arm64 go build . -tags noaudio
```

#### Build for Windows x86

```bash
GOOS=windows GOARCH=386 go build . -tags noaudio
```
