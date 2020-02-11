FROM golang:1.11.5

LABEL maintainer="Radhouane BOUAZIZI <bourados@yahoo.com>"

WORKDIR /app

COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# copy the source from the current directory to the Working Directory
COPY . .

# build the Go app
RUN go build -o main .

# expose port 8080
EXPOSE 8080

# command to run the executable
CMD ["./main"]