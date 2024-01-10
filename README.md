# Go Hot Reload: A Go Web Application Template

[![Go](https://github.com/tombrereton/go-hot-reload/actions/workflows/go.yml/badge.svg)](https://github.com/tombrereton/go-hot-reload/actions/workflows/go.yml)

## Introduction

Go Hot Reload is an open-source starter template, designed to streamline the development process with hot reloading for Go, HTML and Tailwind CSS. It's a perfect starting point for server side web development, particularly for those using `htmx`.

## Features

- 🔄 **Live Reloading**: Automatic Go rebuild using [`air`](https://github.com/cosmtrek/air). Configured in `.air.toml`.
- 🔥 **Hot Reloading**: Dynamic HTML updates with [`reload`](https://github.com/aarol/reload). Uses `concurrently` npm package for running `air` and `npx tailwind` together. Details in `package.json`.
- 📄 **HTML Templates**: [`chi`](https://github.com/go-chi/chi) router for serving HTML templates.
- 💅 **Tailwind**: CSS framework for responsive, customizable UI components with [Tailwind](https://tailwindcss.com/).
- 🐳 **Docker Ready**: Containerized environment for easy deployment and scalability. Includes multi-stage builds for efficient image size (approximately 18 MB).
- 🚀 **HTMX**: Enhance your HTML with AJAX, WebSockets, and more using [HTMX](https://htmx.org/), enabling rich interactions with minimal JavaScript.

## Getting Started

### Setting Up the Environment

1. **Install npm Modules**: Run `npm install` to set up necessary modules.
2. **Environment Configuration**: Create a `.env` file with the following content:
   ```
   PORT=3000
   IS_DEVELOPMENT=true
   TEMPLATES_DIR=web/templates/
   STATIC_DIR=web/static
   ```
3. **Launch the Development Server**: Start the server using `npm run dev`.
4. **Hot Reload**: Change the tailwind classes in `templates/landing.html` to see hot reload in action.

### Running Tests

The server includes an integration test.

- Execute tests using `go test ./...`.
- Or using npm command `npm run test`

### Docker

1. **Build Image**: `npm run docker:build`
2. **Run Image**: `npm run docker:run`

### Commands

Full list of commands for `npm run` are:

```json
"watch:tailwind": "npx tailwindcss -i ./static/css/main.css -o ./static/css/styles.css --watch",
"watch:go": "air",
"dev": "concurrently \"npm run watch:tailwind\" \"npm run watch:go\"",
"test": "go test ./...",
"docker:build": "docker build -f build/Dockerfile -t go-hot-reload .",
"docker:run": "docker run --rm -p 3001:3001 --name go-hot-reload-container go-hot-reload",
"docker:it": "docker run -it --rm go-hot-reload sh"
```

## Contributing to Go Hot Reload

Contributions from the community are welcome. Here are some guidelines to help you get started:

1. **Fork the Repository**: Start by forking the repository to your GitHub account.
2. **Clone Locally**: Clone your forked repository to your local machine.
3. **Create a New Branch**: Make your changes in a new git branch.
4. **Commit Your Changes**: Write meaningful commit messages that accurately describe your changes.
5. **Test Your Changes**: Ensure that your changes do not break any existing functionality.
6. **Submit a Pull Request**: Push your branch to your fork on GitHub and submit a pull request.

Thank you for considering to contribute to Go Hot Reload! Your efforts help make this project even better.
