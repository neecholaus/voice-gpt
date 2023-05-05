#!/bin/bash

cd /var/opt;

apt-get -y update;
apt-get -y install curl vim htop;

mkdir /var/opt/piper;
curl -L -o /var/opt/piper/piper_arm64.tar.gz https://github.com/rhasspy/piper/releases/download/v0.0.2/piper_arm64.tar.gz
tar -xf /var/opt/piper/piper_arm64.tar.gz

mkdir /var/opt/piper-voices;
curl -L -o /var/opt/piper-voices/en-us-libritts-high.tar.gz https://github.com/rhasspy/piper/releases/download/v0.0.2/voice-en-us-libritts-high.tar.gz
tar -xf /var/opt/piper-voices/en-us-libritts-high.tar.gz -C /var/opt/piper-voices

mkdir /var/opt/responses