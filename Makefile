build:
	@GOOS=linux GOARCH=arm go build .
	@docker build -t voicegpt .

run:
	@docker run -it -v ./responses:/var/opt/responses voicegpt ./voicegpt

setup-piper:
	@chmod +x setup-piper.sh
	@./setup-piper.sh
