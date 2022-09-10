FROM golang:alpine AS dev-build

RUN apk update && apk add git

# directory we'll use for building app binary
WORKDIR /app

ARG TARGET=main
ARG GOOS=linux
ARG CGO_ENABLED=0

ENV GO111MODULE=on
ENV CGO_ENABLED=$CGO_ENABLED
ENV GOOS=$GOOS

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN make build

# Production image
FROM scratch AS prod

ARG TARGET=main
ENV APP_NAME=$TARGET

WORKDIR /app
COPY --from=builder /app/$APP_NAME ./$APP_NAME
# COPY --from=builder /app/configs/$APP_NAME.env ./configs/.env

ENTRYPOINT ["/app/$APP_NAME}"]