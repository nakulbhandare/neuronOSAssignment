# Commander Application

## Overview
A cross-platform command execution application can execute system commands like network ping and getting system info.

## Build Instructions
1. Clone the repo.
2. Run `go build -o commander main.go`.

## API Documentation
- **POST /execute**
```
curl -X POST http://localhost:8080/execute \
-H "Content-Type: application/json" \
-d '{"type": "ping", "payload": "google.com"}'

```

```
curl -X POST http://localhost:8080/execute \
-H "Content-Type: application/json" \
-d '{"type": "sysinfo", "payload": ""}'

```

## Installation Guide
For macOS:
1. Run `installer.sh` to create the `.pkg` installer.
2. Run the Command `./commander`


## Testing
Run the test suite using `go test`.

## Execution Demo

## Demo Video

Here’s a demo video showing the code execution:

[Watch the demo video here](../demo/Demo.mp4)


