FROM hexletbasics/base-image:latest

RUN apt-get update
RUN apt-get install wget
RUN wget -c https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH=/usr/local/go/bin:$PATH

RUN go get -u golang.org/x/lint/golint
ENV PATH=/exercises-go/bin:/root/go/bin:$PATH

WORKDIR /exercises-go

COPY . .

# RUN go mod vendor
