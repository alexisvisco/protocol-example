FROM jaegertracing/protobuf

ARG buf_version

RUN wget https://github.com/bufbuild/buf/releases/download/v$buf_version/buf-Linux-x86_64
RUN apk add --update nodejs npm make git

COPY --from=golang:1.17.6-alpine3.15 /usr/local/go/ /usr/local/go/


RUN mv buf-Linux-x86_64 buf && chmod 777 buf && mv buf /bin/buf

RUN npm install -g twirpscript@0.0.50

#COPY compile.js /usr/lib/node_modules/twirpscript/dist/compile.js


ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

RUN go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
    go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2


RUN mkdir /workspace
WORKDIR /workspace

COPY buf.yaml ./
COPY buf.gen.yaml ./
COPY buf.lock ./

RUN buf build ; exit 0

ENTRYPOINT ["/bin/bash", "-c"]
