# Node.js stage for Tailwind CSS
FROM node:latest as tailwindbuilder

WORKDIR /node

COPY web ./web
COPY package*.json .
COPY tailwind.config.js .
RUN npm install 
RUN npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/outpus.css --minify

# Go build stage
FROM golang:latest as gobuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

ENV PORT=3001

WORKDIR /root/

COPY --from=gobuilder /app/cmd/main .
COPY --from=tailwindbuilder /node/web/ ./web/

EXPOSE 3001
CMD ["./main"]
