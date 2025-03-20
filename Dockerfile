FROM hexletbasics/base-image:latest

RUN apt-get update && apt-get install -y wget build-essential
RUN wget -c https://dl.google.com/go/go1.24.1.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH=/usr/local/go/bin:$PATH

# RUN go get -u golang.org/x/lint/golint
# RUN go install github.com/mgechev/revive@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
ENV PATH=/exercises-go/bin:/root/go/bin:$PATH

WORKDIR /exercises-go

COPY . .

RUN go mod vendor
