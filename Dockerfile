FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=$GOPATH/src/github.com/jairvercosa/auth-go/

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
ADD main.go Gopkg.toml Gopkg.lock $SRC_DIR
ADD jwt-key jwt-key.pub /app/

RUN cd $SRC_DIR; dep ensure --vendor-only; go build -o auth-go; cp auth-go /app/
ENV PORT=8080
ENV HOST=0.0.0.0
ENTRYPOINT ["./auth-go"]