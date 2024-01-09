# Go Hot Reload
A Go web application with hot reload for html and tailwind.

# Live Reloading
The rebuilding of Go code is enabled by air. It's configuration is located in `.air.toml`.

# Hot Reloading
The ability to push updated html to the browser is enabled by the [reload](https://github.com/aarol/reload) packge. This injects some javascript code at the bottom of pages to open a websocket with the server. The server can then push updates to the browser.

# Run
Go-Hot-Reload uses 2 commands `air` and `npx tailwind` to rebuild the files. These commands are run using the npm package `concurrently`. Looks in `package.json` for details.

Steps
1. Install npm modules
```
npm install
```

2. Create `.env` file
```
# .env

IS_DEVELOPMENT=true
TEMPLATES_PATH=templates/
```

2. Start dev server
```
npm run start
```

# Tests
There is an intergration test for the server.

1. Run tests.
```
go test ./...
```