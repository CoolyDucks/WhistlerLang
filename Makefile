# Makefile for WhistlerLang 1.2 (think about it, you are not Windows User)
# Builds for....wait a second 

SRC=source/*.go
BUILD_DIR=build

all: linux-amd64 linux-386 linux-arm64 freebsd-amd64 openbsd-x86 netbsd-amd64 windows-amd64 windows-x86 mac-x64 mac-arm android-arm

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

linux-amd64: $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/WhistlerLang-linux-amd64 $(SRC)

linux-386: $(BUILD_DIR)
	GOOS=linux GOARCH=386 go build -o $(BUILD_DIR)/WhistlerLang-linux-386 $(SRC)

linux-arm64: $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/WhistlerLang-linux-arm64 $(SRC)

freebsd-amd64: $(BUILD_DIR)
	GOOS=freebsd GOARCH=amd64 go build -o $(BUILD_DIR)/WhistlerLang-freebsd-amd64 $(SRC)

openbsd-x86: $(BUILD_DIR)
	GOOS=openbsd GOARCH=386 go build -o $(BUILD_DIR)/WhistlerLang-openbsd-x86 $(SRC)

netbsd-amd64: $(BUILD_DIR)
	GOOS=netbsd GOARCH=amd64 go build -o $(BUILD_DIR)/WhistlerLang-netbsd-amd64 $(SRC)

windows-amd64: $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/WhistlerLang-windows-amd64.exe $(SRC)

windows-x86: $(BUILD_DIR)
	GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/WhistlerLang-windows-x86.exe $(SRC)

mac-x64: $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/WhistlerLang-mac-x64 $(SRC)

mac-arm: $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/WhistlerLang-mac-arm $(SRC)

android-arm: $(BUILD_DIR)
	# Requires NDK installed and arm-linux-androideabi-gcc available
	GOOS=android GOARCH=arm CGO_ENABLED=1 CC=arm-linux-androideabi-gcc go build -o $(BUILD_DIR)/WhistlerLang-android-arm $(SRC)

clean:
	rm -rf $(BUILD_DIR)/*

.PHONY: all clean linux-amd64 linux-386 linux-arm64 freebsd-amd64 openbsd-x86 netbsd-amd64 windows-amd64 windows-x86 mac-x64 mac-arm android-arm
