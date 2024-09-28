FROM golang:1.22 AS builder

ENV APP_HOME /go/src/auth
# ENV MONGODB_URI ${MONGODB_URI}
# ENV FIREBASE_PROJECT_ID ${FIREBASE_PROJECT_ID}

WORKDIR "$APP_HOME"

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o auth

# copy build to a clean image
FROM debian:bullseye-slim

ENV APP_HOME /go/src/auth
ENV MONGODB_URI ${MONGODB_URI}
ENV FIREBASE_PROJECT_ID ${FIREBASE_PROJECT_ID}

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/auth $APP_HOME

CMD ["./auth"]