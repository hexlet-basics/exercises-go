FROM hexletbasics/base-image:latest

ARG TARGETARCH

RUN apt-get update && apt-get install -yqq wget
RUN wget -q https://dl.google.com/go/go1.19.1.linux-${TARGETARCH}.tar.gz -O - | tar -xz -C /usr/local;
ENV PATH=/usr/local/go/bin:$PATH

# RUN go get -u golang.org/x/lint/golint
# RUN go install github.com/mgechev/revive@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0
ENV PATH=/exercises-go/bin:/root/go/bin:$PATH

WORKDIR /exercises-go

COPY . .

RUN go mod vendor
