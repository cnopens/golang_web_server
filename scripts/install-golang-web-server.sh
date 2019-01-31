#!/bin/bash 
GOLANG_WEB_SERVER_VERSION="0.1.0" 
#wget https://github.com/yangtinngting/golang_web_server/releases/download/v${GOLANG_WEB_SERVER_VERSION}/golangWebServer-${GOLANG_WEB_SERVER_VERSION}.darwin.amd64.tar.gz 
#tar -xzvf golangWebServer-${GOLANG_WEB_SERVER_VERSION}.darwin.amd64.tar.gz 
cd golang_web_server/ 

# if you just want to start prometheus as root 
# ./prometheus --config.file=prometheus.yml  

# create user 
useradd --no-create-home --shell /bin/false golang_web_server 

# copy binaries   
cp golangWebServer /usr/local/bin 
chown golang_web_server:golang_web_server  /usr/local/bin/golangWebServer  

# setup systemd 
echo '[Unit]
Description=Golang Web Server
Wants=network-online.target
After=network-online.target

[Service]
User=golang_web_server
Group=golang_web_server
Type=simple
ExecStart=/usr/local/bin/golangWebServer 

[Install]
WantedBy=multi-user.target' > /etc/systemd/system/golang_web_server.service

systemctl daemon-reload
systemctl enable golang_web_server
systemctl start golang_web_server




