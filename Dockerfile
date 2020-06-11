FROM melodyn/base-image:latest

# go

RUN go get -u golang.org/x/lint/golint

WORKDIR /exercises-go

COPY --from=melodyn/base-image:latest /tmp/basics/common/* ./
COPY . .
