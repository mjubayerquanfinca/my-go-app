# My Go App

A simple Go web application with CI/CD pipeline that includes Docker image building and pushing to GitHub Container Registry.

## Features

- HTTP server running on port 8080
- Automated testing
- Docker containerization
- CI/CD pipeline with GitHub Actions
- Automatic Docker image building and pushing

## Local Development

### Prerequisites

- Go 1.22 or later
- Docker (optional, for containerized development)

### Running Locally

```bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build and run
go build -o my-go-app .
./my-go-app
```

The application will be available at `http://localhost:8080`

### Using Docker

```bash
# Build and run with Docker
docker build -t my-go-app .
docker run -p 8080:8080 my-go-app

# Or use docker-compose
docker-compose up --build
```

## CI/CD Pipeline

The project includes a GitHub Actions workflow that:

1. **Triggers on**: 
   - Push to `main` or `develop` branches
   - Pull requests to `main` branch

2. **Test Job**:
   - Runs Go tests
   - Must pass before building Docker image

3. **Build and Push Job**:
   - Builds Docker image using multi-stage Dockerfile
   - Pushes to GitHub Container Registry (ghcr.io)
   - Tags images based on branch:
     - `main` branch → `latest` tag
     - `develop` branch → `develop` tag
     - Feature branches → `{branch-name}` tag
     - Pull requests → `pr-{number}` tag

### Docker Image Tags

Images are automatically tagged and pushed to: `ghcr.io/{your-username}/my-go-app`

Available tags:
- `latest` - Latest main branch build
- `develop` - Latest develop branch build
- `{branch-name}` - Feature branch builds
- `{branch}-{sha}` - Specific commit builds

### Using the Docker Image

```bash
# Pull and run the latest image
docker pull ghcr.io/{your-username}/my-go-app:latest
docker run -p 8080:8080 ghcr.io/{your-username}/my-go-app:latest
```

## Project Structure

```
.
├── main.go              # Application entry point
├── handler/
│   ├── hello.go         # HTTP handler
│   └── hello_test.go    # Tests
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Local development setup
├── .github/workflows/   # CI/CD configuration
└── README.md           # This file
```

## Security Features

- Multi-stage Docker build for smaller images
- Non-root user execution
- Minimal Alpine Linux base image
- Security scanning in CI/CD pipeline
