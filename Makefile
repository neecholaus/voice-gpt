build:
	@GOOS=linux GOARCH=arm go build .
	@docker build -t voicegpt .

setup-piper:
	@chmod +x setup-piper.sh
	@./setup-piper.sh