FROM golang:1.19-alpine

ENV PORT=4000
ENV PEPPER=tomato
ENV ENV=development
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=db
ENV JWT_SECRET=mysecret
ENV EMAIL_FROM=me@themanan.me

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o bin/ ./cmd/...
EXPOSE $PORT
CMD ["bin/cmd"]


FROM alpine:latest
ARG ENVIRONMENT_NAME
RUN apk add --no-cache libc6-compat 
