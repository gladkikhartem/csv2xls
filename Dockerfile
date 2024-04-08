FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && apt-get install -y gnumeric
WORKDIR /
COPY app.run .
CMD ["/app.run"]