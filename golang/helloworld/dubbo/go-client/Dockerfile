FROM golang:latest

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

COPY . $GOPATH/org/apache/dubbogo/helloworld/client
WORKDIR $GOPATH/org/apache/dubbogo/helloworld/client
ADD ./wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

ENV CONF_CONSUMER_FILE_PATH ./profiles/client.yml
ENV APP_LOG_CONF_FILE ./profiles/log.yml

RUN go mod tidy
RUN go build -o ./target/client ./app
