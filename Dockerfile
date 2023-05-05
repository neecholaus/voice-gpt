FROM arm64v8/golang:1.20

WORKDIR /var/opt

RUN apt-get -y update;
RUN apt-get -y install curl vim htop;

RUN mkdir /var/opt/piper
RUN curl -L -o /var/opt/piper/piper_arm64.tar.gz https://github.com/rhasspy/piper/releases/download/v0.0.2/piper_arm64.tar.gz
RUN tar -xf /var/opt/piper/piper_arm64.tar.gz

RUN mkdir /var/opt/piper-voices;
RUN curl -L -o /var/opt/piper-voices/en-us-libritts-high.tar.gz https://github.com/rhasspy/piper/releases/download/v0.0.2/voice-en-us-libritts-high.tar.gz
RUN tar -xf /var/opt/piper-voices/en-us-libritts-high.tar.gz -C /var/opt/piper-voices

RUN mkdir /var/opt/responses
RUN echo "echo 'this is a test.' | \
    /var/opt/piper/piper --model /var/opt/piper-voices/en-us-libritts-high.onnx --f a.wav --d /var/opt/responses" > /var/opt/test.sh -s 13

ADD ./profinbox-arm64 /var/opt/profinbox

CMD sleep