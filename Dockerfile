FROM docker.io/library/golang:1.21.1-alpine AS build
WORKDIR /app
RUN apk update && apk upgrade
COPY . /app
RUN go build -o ./cmd/bin/slack_bot ./cmd/slack_bot

FROM docker.io/library/alpine AS final
WORKDIR /app/cmd/bin
COPY --from=build --chmod=755 /app/cmd/bin/slack_bot /app/cmd/bin/slack_bot
RUN apk update && apk upgrade && \
    mkdir -p /app/configs && \
    adduser -D app
USER app
ENV PATH /app/cmd/bin:${PATH}
CMD ["/app/cmd/bin/slack_bot"]
