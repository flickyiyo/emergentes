FROM golang
ENV LOG_FILE_LOCATION=/LOGFILE
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 go build ./app.go
RUN ls
FROM alpine:latest
ENV LOG_FILE_LOCATION=/LOGFILE
RUN touch /LOGFILE
COPY --from=0 /app .
RUN pwd
EXPOSE 8080
ENTRYPOINT ["./app"]