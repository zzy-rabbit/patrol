# 开启cgo
CGO_ENABLED=1

# 默认编译当前平台
.PHONY: build
build:
	go build -o app .

# Linux平台
.PHONY: linux
linux: linux-amd64 linux-arm64

.PHONY: linux-amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o patrol-linux-amd64 .

.PHONY: linux-arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build -o patrol-linux-arm64 .

# macOS平台
.PHONY: darwin
darwin: darwin-amd64 darwin-arm64

.PHONY: darwin-amd64
darwin-amd64:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o patrol-darwin-amd64 .

.PHONY: darwin-arm64
darwin-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o patrol-darwin-arm64 .

# Windows平台
.PHONY: windows
windows: windows-amd64

.PHONY: windows-amd64
windows-amd64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o patrol-windows-amd64.exe .

# ARM平台（包括ARMv7和ARM64）
.PHONY: arm
arm: linux-arm64 linux-armv7

.PHONY: linux-armv7
linux-armv7:
	GOOS=linux GOARCH=arm CGO_ENABLED=1 GOARM=7 go build -o patrol-linux-armv7 .

# 清理
.PHONY: clean
clean:
	rm -f app patrol-* *.exe

# 编译所有主流平台
.PHONY: all
all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64 linux-armv7
