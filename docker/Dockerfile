FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# COPY *.go ./
COPY . .

# RUN cd cmd/main
RUN go build cmd/main/main.go

EXPOSE 4040

CMD [ "./main" ]