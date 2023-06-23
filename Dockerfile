FROM arm64v8/golang:1.20

WORKDIR /var/opt

ADD setup-piper.sh /var/opt/setup-piper.sh
RUN chmod +x /var/opt/setup-piper.sh
RUN ./setup-piper.sh

COPY ./voicegpt /var/opt/voicegpt
RUN chmod +x /var/opt/voicegpt

COPY ./.env /var/opt/.env
COPY ./audio-in /var/opt/audio-in

CMD sleep