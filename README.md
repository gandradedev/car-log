# Car Log

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)
![technology React](https://img.shields.io/badge/technology-react-61DAFB.svg)

Application for managing vehicle maintenance history and controlling future services. It allows registering vehicles with their maintenance records and automatically alerts when a service is due based on mileage or time interval.

## Technology Stack

- **Backend**: Go 1.22+
- **Frontend**: React + TypeScript + Vite
- **Database**: SQLite (via `modernc.org/sqlite`)

## Repository Structure

```
car-log/
├── backend/                # REST API in Go
│   ├── main.go
│   ├── go.mod
│   └── internal/
│       ├── domain/         # Entities, repository interfaces, custom errors
│       ├── application/    # Use cases and DTOs
│       ├── infrastructure/ # HTTP handlers, SQLite repository, routes
│       └── configuration/  # Database setup
└── frontend/               # React web application
```

## Quick Start

### Prerequisites

- [Go 1.22+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org/)

### Installing Go

1. Download Go from the [official website](https://golang.org/dl/)
   - For macOS: Download the `.pkg` file and follow the installation wizard

2. Verify installation:
   ```bash
   go version
   ```

3. Set up your Go workspace:
   ```bash
   mkdir -p $HOME/go/{bin,src,pkg}
   ```

4. Add the following to your shell profile (`.bashrc`, `.zshrc`, etc.):
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

## Project Setup

1. **Clone the repository**
   ```bash
   git clone git@github.com:gandradedev/car-log.git
   cd car-log
   ```

2. **Configure environment variables**

   ```bash
   cd backend
   cp .env.example .env
   ```

   Edit `backend/.env` if needed:
   ```
   PORT=8080
   DB_PATH=./carlog.db
   ```

   > The `.env` file is ignored by git and should never be committed.

3. **Start the backend**
   ```bash
   make backend-run
   ```

   The API will be available at `http://localhost:8080`.
   Data is automatically persisted in `backend/carlog.db`.

4. **Start the frontend** (in a separate terminal)
   ```bash
   make frontend-install
   make frontend-dev
   ```

   The app will be available at `http://localhost:5173`.

5. **Or run both simultaneously**
   ```bash
   make dev
   ```

## Development

### Available commands

```bash
make help              # List all available commands

# Backend
make backend-run       # Run the Go server
make backend-build     # Compile the binary
make backend-test      # Run all tests
make backend-tidy      # Tidy Go module dependencies
make backend-clean     # Remove compiled binary

# Frontend
make frontend-install  # Install npm dependencies
make frontend-dev      # Start Vite dev server
make frontend-build    # Build for production
make frontend-preview  # Preview the production build

# Combined
make dev               # Run backend and frontend concurrently
```

### Backend

```bash
cd backend

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

### Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start dev server
npm run dev

# Build for production
npm run build
```

## Testing

### Backend

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
