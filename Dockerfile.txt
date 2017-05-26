FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
CMD ["/app/helloworld"]
