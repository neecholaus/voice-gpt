build:
	@cd ./text-to-voice; GOOS=linux GOARCH=arm64 go build -o texttovoice .

setup-piper:
	@chmod +x setup-piper.sh
	@./setup-piper.sh