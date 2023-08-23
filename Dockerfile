FROM golang:1.20.4-alpine3.17

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .

COPY *.go ./

# Build
RUN go build -o /counter

EXPOSE 9000

# Run
CMD [ "/counter" ]