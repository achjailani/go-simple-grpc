FROM golang:1.18

# Install git.
# Git is required for fetching the dependencies.
RUN set -ex \
    && apk --no-cache update \
    && apk add --no-cache git curl openssh-client ca-certificates \
       make perl bash build-base zlib-dev ucl-dev \
    && rm -rf /var/cache/apk/* \

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download & go mod verify

COPY . .

EXPOSE 9002

RUN GOOS=linux go build -v -a -o main .

CMD ["./main"]