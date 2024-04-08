FROM golang:1.20

WORKDIR /app

RUN go mod init github.com/BrentGrammer/goapi

# COPY go.mod .
COPY main.go .

RUN go mod tidy

CMD ["go","run","main.go"]

# create a file bin which is the executable. will run inside the /app workdir
# RUN go build -o bin .

# ENTRYPOINT [ "/app/bin" ]

### To run: docker build . -t go-container:latest
### From https://www.youtube.com/watch?v=C5y-14YFs_8

