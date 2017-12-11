FROM golang:1.9-alpine AS build-env
WORKDIR /go/src/github.com/arbarlow/kuboid/mock_prom
COPY . /go/src/github.com/arbarlow/kuboid/mock_prom
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get
RUN go build -o mock_prom .


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /go/src/github.com/arbarlow/kuboid/mock_prom/mock_prom /usr/bin/mock_prom
CMD ["mock_prom"]
