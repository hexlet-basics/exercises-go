FROM hexletbasics/base-image:latest

RUN apt-get update
RUN apt-get install -y golang-1.16-go
ENV PATH=/usr/lib/go-1.16/bin:$PATH

RUN go get -u golang.org/x/lint/golint
ENV PATH=/root/go/bin:$PATH

WORKDIR /exercises-go

COPY . .

RUN go mod vendor
