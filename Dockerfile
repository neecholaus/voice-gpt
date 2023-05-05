FROM arm64v8/golang:1.20

WORKDIR /var/opt

ADD setup-piper.sh /var/opt/provision.sh
RUN chmod +x /var/opt/setup-piper.sh
RUN ./setup-piper.sh

COPY profinbox-arm64 /var/opt/profinbox

CMD sleep