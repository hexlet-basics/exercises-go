FROM hexletbasics/base-image:latest

ARG TARGETARCH

RUN apt-get update && apt-get install -yqq wget
RUN wget -q https://dl.google.com/go/go1.25.1.linux-${TARGETARCH}.tar.gz -O - | tar -xz -C /usr/local

ENV PATH=/usr/local/go/bin:$PATH

RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0
ENV PATH=/exercises-go/bin:/root/go/bin:$PATH
ENV GOFLAGS=-buildvcs=false

WORKDIR /exercises-go

COPY go.mod go.sum .

RUN go mod download

COPY . .
