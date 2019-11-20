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

FROM golang:1.12-alpine AS builder
EOF

if [ "$PRproxy" == "yes" ]; then
cat <<EOF >>./Dockerfile
ENV HTTPS_PROXY http://$PRusername:$PRpassword@$PRurl:$PRport/
ENV HTTP_PROXY http://$PRusername:$PRpassword@$PRurl:$PRport/
EOF
fi

cat <<EOF >>./Dockerfile
RUN apk add --no-cache \
                                git
RUN mkdir -p /go/src/lxldap/cmd /lxldap/lib
COPY *.go /go/src/lxldap/
COPY cmd/* /go/src/lxldap/cmd/
COPY lib/* /go/src/lxldap/lib/
COPY config/.gitconfig /root/
RUN go get -u -v github.com/gorilla/mux
RUN go get -u -v github.com/mitchellh/go-homedir
RUN go get -u -v github.com/bradfitz/slice
RUN go get -u -v github.com/olekukonko/tablewriter
RUN go get -u -v github.com/rifflock/lfshook
RUN go get -u -v github.com/sirupsen/logrus
RUN go get -u -v github.com/spf13/cobra
RUN go get -u -v github.com/spf13/viper
RUN go get -u -v gopkg.in/ldap.v3
RUN cd /go/src/lxldap && go build lxldap.go

FROM alpine:latest

# Create User
ENV MYUSER=lxldap
ENV MYUID=500
ENV MYGID=500
ENV PATH=\${PATH}:/home/\${MYUSER}

RUN addgroup -S -g \${MYGID} \${MYUSER} && \
    adduser -S \${MYUSER} -u \${MYUID} -g \${MYGID} -s /bin/bash -G users

COPY config/.lxldap.toml /home/\${MYUSER}/.lxldap.toml
COPY certs/*.pem /home/\${MYUSER}/
COPY --from=builder /go/src/lxldap /home/\${MYUSER}/
RUN mkdir /home/\${MYUSER}/logs
RUN chown -R lxldap:lxldap /home/\${MYUSER}

USER \${MYUSER}
WORKDIR /home/\${MYUSER}

EXPOSE 4433

CMD ["lxldap","service","start"]
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
