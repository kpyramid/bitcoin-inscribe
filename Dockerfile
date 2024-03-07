FROM golang:latest as builder
COPY . /src
WORKDIR /src
RUN apt-get update && apt-get install -y unzip git && wget https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip && \
    mv protoc-24.4-linux-x86_64.zip /opt/  && cd /opt/ && unzip protoc-24.4-linux-x86_64.zip && ln -s /opt/bin/protoc /usr/bin/ && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN git clone https://github.com/googleapis/googleapis.git && cp -rf googleapis /usr/local/
RUN cd /src && PROTO_INCLUDE=/usr/local/googleapis make proto && go mod tidy && go build -o build/main main.go
RUN cd /src/init && go build -o ../build/init init.go

FROM golang:latest
COPY --from=builder /src/build /app
COPY --from=builder init/rare_data.json /app/
COPY --from=builder init/recursion_data.json /app/
WORKDIR /app
CMD ["/app/main"]