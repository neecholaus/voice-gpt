build:
	@GOOS=linux GOARCH=arm64 go build -o profinbox-arm64 .

setup-piper:
	@chmod +x provision.sh
	@./setup-piper.sh