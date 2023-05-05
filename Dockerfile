FROM arm64v8/golang:1.20

WORKDIR /var/opt

ADD setup-piper.sh /var/opt/setup-piper.sh
RUN chmod +x /var/opt/setup-piper.sh
RUN ./setup-piper.sh

COPY ./text-to-voice/texttovoice /var/opt/texttovoice
RUN chmod +x /var/opt/texttovoice

CMD sleep