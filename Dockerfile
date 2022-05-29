FROM golang:1.16-alpine3.15

COPY . /app

WORKDIR /app

# Builds an executable from the project.
RUN go build -o math_cli_app ./src
