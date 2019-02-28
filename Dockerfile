FROM golang:1.11
ENV GOPATH=/usr/go

COPY . /usr/go/src/github.com/hitesh-goel/loomx
WORKDIR /usr/go/src/github.com/hitesh-goel/loomx
RUN go get -v
RUN go install -v

CMD ["/usr/go/bin/loomx"]
EXPOSE 3000 4000
