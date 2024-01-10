# Node.js stage for Tailwind CSS
FROM node:latest as tailwindbuilder

WORKDIR /node

COPY web ./web
COPY package*.json .
COPY tailwind.config.js .
RUN npm install 
RUN npx tailwindcss -i ./web/static/css/main.css -o ./web/static/css/styles.css --minify

# Go build stage
FROM golang:latest as gobuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/web
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

ENV PORT=3001 \
    IS_DEVELOPMENT=false \
    TEMPLATES_DIR=web/templates/ \
    STATIC_DIR=./web/static

WORKDIR /root/

COPY --from=gobuilder /app/cmd/web/main .
COPY --from=tailwindbuilder /node/web/ ./web/

EXPOSE 3001
CMD ["./main"]
