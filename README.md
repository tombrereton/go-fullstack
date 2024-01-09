# Go Hot Reload: A Dynamic Go Web Application Template

## Introduction
Go Hot Reload is an open-source web application template using Go, designed to streamline the development process with hot reloading for HTML and Tailwind CSS. It's a perfect starting point for Go web development, particularly for those using `htmx`.

## Features

### Live Reloading
- **Go Code Rebuilding**: Utilizes `air` to automatically rebuild Go code. Configuration for this feature is specified in `.air.toml`.

### Hot Reloading
- **HTML Updates in Browser**: Employs the [reload](https://github.com/aarol/reload) package to dynamically push HTML updates to the browser. It works by injecting JavaScript at the end of pages to establish a WebSocket connection with the server, enabling the server to send updates directly to the browser.

## Getting Started

### Setting Up the Environment
1. **Install npm Modules**: Run `npm install` to set up necessary modules.
2. **Environment Configuration**:
    - Create a `.env` file with the following content:
      ```
      # .env

      IS_DEVELOPMENT=true
      TEMPLATES_PATH=templates/
      ```
3. **Launch the Development Server**: Start the server using `npm run start`. The `concurrently` npm package is used to run `air` and `npx tailwind` commands concurrently for file rebuilding. For more details, see `package.json`.

### Running Tests
- **Integration Testing**: The server includes an integration test.
  - Execute the test using `go test ./...`.

## Contributing to Go Hot Reload

We welcome contributions from the community. Here are some guidelines to help you get started:

1. **Fork the Repository**: Start by forking the repository to your GitHub account.
2. **Clone Locally**: Clone your forked repository to your local machine.
3. **Create a New Branch**: Make your changes in a new git branch.
4. **Commit Your Changes**: Write meaningful commit messages that accurately describe your changes.
5. **Test Your Changes**: Ensure that your changes do not break any existing functionality.
6. **Submit a Pull Request**: Push your branch to your fork on GitHub and submit a pull request.

Thank you for considering to contribute to Go Hot Reload! Your efforts help make this project even better.