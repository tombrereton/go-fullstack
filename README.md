# Go Hot Reload
A Go web application with hot reload for Tailwind

# Run
Go-Hot-Reload uses 2 commands `air` and `npx tailwind`, to rebuild the files

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