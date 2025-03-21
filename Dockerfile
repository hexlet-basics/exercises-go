FROM hexletbasics/base-image:latest

RUN apt-get update && apt-get install -yqq \
    wget \
    build-essential && \
    rm -rf /var/lib/apt/lists/*

RUN wget -c -q https://dl.google.com/go/go1.19.1.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH=/usr/local/go/bin:$PATH

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.6
ENV PATH=/exercises-go/bin:/root/go/bin:$PATH

WORKDIR /exercises-go

COPY . .

RUN go mod vendor
