# 开启 cgo
CGO_ENABLED=1

# 输出目录
OUT_DIR := output
CONFIG_DIR := config

# 确保输出目录存在
.PHONY: prepare
prepare:
	mkdir -p $(OUT_DIR)
	cp -r $(CONFIG_DIR) $(OUT_DIR)/

# 默认编译当前平台
.PHONY: build
build: prepare
	go build -o $(OUT_DIR)/app .

# Linux 平台
.PHONY: linux
linux: linux-amd64 linux-arm64

.PHONY: linux-amd64
linux-amd64: prepare
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-linux-amd64 .

.PHONY: linux-arm64
linux-arm64: prepare
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-linux-arm64 .

# macOS 平台
.PHONY: darwin
darwin: darwin-amd64 darwin-arm64

.PHONY: darwin-amd64
darwin-amd64: prepare
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-darwin-amd64 .

.PHONY: darwin-arm64
darwin-arm64: prepare
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-darwin-arm64 .

# Windows 平台
.PHONY: windows
windows: windows-amd64

.PHONY: windows-amd64
windows-amd64: prepare
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-windows-amd64.exe .

# ARM 平台
.PHONY: arm
arm: linux-arm64 linux-armv7

.PHONY: linux-armv7
linux-armv7: prepare
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 \
	go build -o $(OUT_DIR)/patrol-linux-armv7 .

# 清理
.PHONY: clean
clean:
	rm -rf $(OUT_DIR)

# 编译所有主流平台
.PHONY: all
all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64 linux-armv7

