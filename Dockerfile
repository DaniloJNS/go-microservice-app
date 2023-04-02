FROM golang

ARG APP_PATH=/opt/app/
WORKDIR $APP_PATH

# Add the app
COPY . $APP_PATH

RUN go build -o server ./cmd/server/main.go

CMD [ "./server" ]
