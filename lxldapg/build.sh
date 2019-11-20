#!/bin/bash

source ./proxy.conf

if [ "$PRproxy" == "yes" ]; then
echo
echo "###########################"
echo "#   Proxy configuration   #"
echo "###########################"

if [ "$PRurl" == "" ]; then
        echo "Proxy url:"
        read PRurl
fi

if [ "$PRport" == "" ]; then
        echo "Proxy port:"
        read PRport
fi

if [ "$PRusername" == "" ]; then
	echo "Proxy User Name:"
	read PRusername
fi

if [ "$PRpassword" == "" ]; then
	echo "Proxy password for user $PRusername:"
	read -s PRpassword
fi

cat <<EOF >/home/admin/.docker/config.json
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://$PRusername:$PRpassword@$PRurl:$PRport",
     "httpsProxy": "http://$PRusername:$PRpassword@$PRurl:$PRport",
   }
 }
}
EOF

cat <<EOF >./config/00proxy.apt 
Acquire::http::proxy "http://$PRusername:$PRpassword@$PRurl:$PRport/";
Acquire::https::proxy "http:///$PRusername:$PRpassword@$PRurl:$PRport";
EOF

cat <<EOF >./config/.gitconfig 
[https]
        proxy = http://$PRusername:$PRpassword@$PRurl:$PRport
        sslVerify = false
        proxyAuthMethod = basic
[http]
        proxy = http://$PRusername:$PRpassword@$PRurl:$PRport
        sslVerify = false
        proxyAuthMethod = basic
EOF
fi

cat <<EOF >./Dockerfile
# syntax = docker/dockerfile:experimental

FROM golang:1.12-stretch AS builder
EOF

if [ "$PRproxy" == "yes" ]; then
cat <<EOF >>./Dockerfile
ENV https_proxy=http://$PRusername:$PRpassword@$PRurl:$PRport/
ENV http_proxy=http://$PRusername:$PRpassword@$PRurl:$PRport/
ENV GIT_HTTP_PROXY_AUTHMETHOD basic
EOF
fi

cat <<EOF >>./Dockerfile
COPY config/00proxy.apt /etc/apt/apt.conf.d/00proxy
RUN apt-get update && apt-get -y install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev
RUN apt-get -y --no-install-recommends install libqt*5-dev qt*5-dev qml-module-qtquick-* qt*5-doc-html
RUN mkdir -p /go/src/lxldapg
COPY lxldapg.* /go/src/lxldapg/
RUN cd /go/src/lxldapg
COPY config/.gitconfig /r/oot/
RUN mkdir -p /go/src/golang.org/x/
RUN cd /go/src/golang.org/x/
RUN git clone https://github.com/golang/sys
RUN cd /go/src/lxldapg
RUN go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...
RUN go get -u -v github.com/mitchellh/go-homedir
RUN go get -u -v github.com/spf13/viper
RUN go get -u -v gopkg.in/ldap.v3
RUN export QT_PKG_CONFIG=true && \$(go env GOPATH)/bin/qtsetup -test=false
RUN cd /go/src/lxldapg && export QT_PKG_CONFIG=true && go build lxldapg.go

FROM debian:stretch

COPY config/00proxy.apt /etc/apt/apt.conf.d/00proxy
RUN apt-get update && apt-get -y --no-install-recommends install libqt*5 ldap-utils openssl ca-certificates
RUN rm /etc/apt/apt.conf.d/00proxy

COPY config/.lxldapg.toml /root/.lxldapg.toml
COPY config/.ldaprc /root/.ldaprc
COPY certs/*.pem /root/
COPY --from=builder /go/src/lxldapg/lxldapg /root/

CMD [ "bash" ]
EOF

echo
echo "###########################"
echo "#      Docker Build       #"
echo "###########################"
echo

read -p "Press enter to continue"

docker build -t lxldap:1.0 .

if [ "$PRproxy" == "yes" ]; then
	echo
	echo "###########################"
	echo "#     Proxy Cleanning     #"
	echo "###########################"
	echo

	read -p "Press enter to continue"

	rm /home/admin/.docker/config.json
	rm ./config/00proxy.apt
	rm ./config/.gitconfig
	rm ./Dockerfile
fi
