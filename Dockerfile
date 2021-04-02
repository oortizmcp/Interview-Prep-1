FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get github.com/oortizmcp/Interview-Prep-1/app
RUN cd /build && git clone https://github.com/oortizmcp/Interview-Prep-1.git

RUN cd /build/Interview-Prep/app && go env -w GO111MODULE=auto && go build 

EXPOSE 8080

ENTRYPOINT [ "/build/Interview-Prep/app/app" ]