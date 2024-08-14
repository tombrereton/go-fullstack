# Go Hot Reload: A Go Web Application Template

[![Go](https://github.com/tombrereton/go-hot-reload/actions/workflows/go.yml/badge.svg)](https://github.com/tombrereton/go-hot-reload/actions/workflows/go.yml)

## Introduction

Go Hot Reload is an open-source starter template, designed to streamline the development process with hot reloading for Go, Templ (HTML), htmx and Tailwind CSS. It's a perfect starting point for server side web development, particularly for those using `htmx`.

## Motivation
I want a starter project to quickly prototype ideas. Things that are important:
- Good developer experience for running and testing locally
- Zero config setup for running integration tests against datastore
- Full offline local development support
- True hot reloading i.e. the browser automatically refreshes when UI code is changed
- Lightweight deployment artifact to maximise cheap or free hosting solutions
- Integrates with cheap or free data storage solutions e.g. Supabase or Dynamo DB
- Infrastructure as Code for free hosting solution

## Hosting

I recommend hosting for free on https://fly.io/.

The first 3 `shared-cpu-1x@256MB` VMs are free and you can deploy docker images easily.

## Features

- 🔄 **Live Reloading**: Automatic Go rebuild using [`air`](https://github.com/cosmtrek/air). Configured in `.air.toml`.
- 📄 **Templ**: Templ is a component and typed html templating system - it plays nice with htmx.
- 💅 **Tailwind**: CSS framework for responsive, customizable UI components with [Tailwind](https://tailwindcss.com/).
- 🐳 **Docker Ready**: Containerized environment for easy deployment and scalability. Includes multi-stage builds for efficient image size (approximately 18 MB).
- 🚀 **HTMX**: Enhance your HTML with AJAX, WebSockets, and more using [HTMX](https://htmx.org/), enabling rich interactions with minimal JavaScript.

## Getting Started

### Setting Up the Environment

1. **Install npm Modules**: Run `npm install` to set up necessary modules.
2. **Environment Configuration**: Create a `.env` file with the following content:
   ```
   PORT=4000
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
"watch:tailwind": "npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch",
"watch:templ": "templ generate -path web/view  --watch --proxy='http://localhost:4000'",
"watch:go": "air",
"dev": "concurrently \"npm run watch:tailwind\" \"npm run watch:go\" \"npm run watch:templ\"",
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
