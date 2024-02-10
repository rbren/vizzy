FROM node:18 as react-build

WORKDIR /app
COPY ./app/package.json ./app/package-lock.json ./
RUN npm install

COPY ./app .
RUN npm run build


FROM golang:1.21 as go-build

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=go-build /go/src/app/server .
COPY --from=react-build /app/dist ./app/dist

ENV PORT 8080
EXPOSE 8080

# Run the Go server
CMD ["./server"]
