FROM golang:1.11.4

ENV APP_DIR $GOPATH/src/api
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./api)
ADD . $APP_DIR

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0
ENV APIRUNMODE prod
# Compile the binary and statically link
RUN cd $APP_DIR && go mod vendor

RUN cd $APP_DIR && go build

EXPOSE 8080
