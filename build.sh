#!/bin/sh

if [ -d "./out" ]; then
    rm -rf ./out
fi
mkdir out

# Get the version from the go.mod file
version="0.2.1"

env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o    "out/adf-cli_${version}_linux_arm64"
env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o      "out/adf-cli_${version}_linux_arm"
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o    "out/adf-cli_${version}_linux_amd64"
env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o   "out/adf-cli_${version}_darwin_amd64"
env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o   "out/adf-cli_${version}_darwin_arm64"
env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o  "out/adf-cli_${version}_windows_amd64.exe"
env GOOS=windows GOARCH=arm go build -ldflags "-s -w" -o    "out/adf-cli_${version}_windows_arm.exe"
env GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o  "out/adf-cli_${version}_windows_arm64.exe"

cd out

# Rename all executables to adf-cli or adf-cli.exe and zip them with their original name
for file in *; do
    if [ -f "$file" ]; then
        if [[ "$file" == *"linux"* ]]; then
            mv "$file" "adf-cli"
            zip "${file}.zip" "adf-cli"
            rm -rf "adf-cli"
        elif [[ "$file" == *"darwin"* ]]; then
            mv "$file" "adf-cli"
            zip "${file}.zip" "adf-cli"
            rm -rf "adf-cli"
        elif [[ "$file" == *"windows"* ]]; then
            mv "$file" "adf-cli.exe"
            # Get the name of the file without the '.exe' extension
            zip_name=$(echo "$file" | tr -d ".exe")
            zip "${zip_name}.zip" "adf-cli.exe"
            rm -rf "adf-cli.exe"
        fi
    fi
done
