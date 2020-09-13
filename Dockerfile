# We specify the base image we need for our
# go application
FROM golang:1.12.0 AS builder

# We create an /app directory within our
# image that will hold our application source
# files
EXPOSE 8080

RUN mkdir /application
# We copy everything in the root directory
# into our /app directory
ADD . /application
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /application
# we run go build to compile the binary
# executable of our Go program
RUN CGO_ENABLED=0 GOOS=linux go build -o main .



FROM alpine:latest AS production


COPY --from=builder /application .
# Our start command which kicks off
# our newly created binary executable


CMD ["./main"]