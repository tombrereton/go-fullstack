# Node.js stage for Tailwind CSS
FROM node:latest as tailwindbuilder

WORKDIR /web

COPY static/ ./static/
COPY package*.json .
COPY tailwind.config.js .
RUN ls -l /web
RUN npm install 
RUN npx tailwindcss -i ./static/css/main.css -o ./static/css/styles.css

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
    TEMPLATES_DIR=templates/ \
    STATIC_DIR=./static

WORKDIR /root/

COPY --from=gobuilder /app/cmd/web/main .
COPY --from=tailwindbuilder /web/static/ ./static/
COPY ./templates/ ./templates/

EXPOSE 3001
CMD ["./main"]
