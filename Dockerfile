FROM golang:1.13-alpine
LABEL author=mat.woess@gmail.com
WORKDIR /src
COPY main.go go.* ./
RUN CGO_ENABLED=0 go build -o /usr/myapp
EXPOSE 8888
CMD ["/usr/myapp"]
