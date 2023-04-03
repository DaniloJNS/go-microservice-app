FROM golang

ARG APP_PATH=/opt/app/
WORKDIR $APP_PATH

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

# Add the app
COPY . $APP_PATH
