# Go Modules Demo Application

A simple REST API using Gorilla Mux demonstrating Go modules dependency management.

## 📋 Prerequisites

- **Go 1.19 or higher**

### Check if installed:
```bash
go version
```

### Installation (if needed):

**macOS:**
```bash
brew install go
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**Windows:**
- Download from https://go.dev/dl/

## 🚀 Running the Demo

### 1. Navigate to the demo directory
```bash
cd go-demo
```

### 2. Download dependencies (auto-detects from imports)
```bash
go mod tidy
```

### 3. Run the application
```bash
go run main.go
```

### 4. Test the endpoints

Open a new terminal and run:
```bash
# Test root endpoint
curl http://localhost:8080/

# Test info endpoint
curl http://localhost:8080/info

# Test health endpoint
curl http://localhost:8080/health
```

Or open in browser: http://localhost:8080/

### 5. Build executable (optional)
```bash
go build -o go-demo
./go-demo
```

### 6. Stop the server
Press `Ctrl+C` in the terminal

## 📦 What This Demo Shows

- **go.mod** - Go module definition file
- **go.sum** - Cryptographic checksums for dependencies
- **Auto-detection** - Go automatically adds imports to go.mod
- **Gorilla Mux** - Popular HTTP router dependency
- **Module proxy** - Downloads from proxy.golang.org
- **Minimal version selection** - Go's unique approach

## 🔍 Key Files

- `go.mod` - Module definition and dependencies
- `go.sum` - Dependency checksums (auto-generated)
- `main.go` - Application code

## 📚 Go Module Commands Reference

```bash
go mod init <module-path>    # Initialize new module
go mod tidy                   # Add missing and remove unused modules
go get <package>              # Add or update dependency
go mod download               # Download modules to local cache
go mod verify                 # Verify dependencies
go list -m all                # List all dependencies
go mod vendor                 # Copy dependencies to vendor/
go build                      # Build executable
go run main.go                # Run without building
```

## ✅ Expected Output

After running `go run main.go`:
```
=== Go Modules Demo Application ===

✅ Server running on http://localhost:8080

Available endpoints:
  GET http://localhost:8080/
  GET http://localhost:8080/info
  GET http://localhost:8080/health

Press Ctrl+C to stop
```

## 🌐 API Endpoints

- `GET /` - Basic info about the application
- `GET /info` - Dependency information
- `GET /health` - Health check endpoint

## 💡 Go Modules Tips

**Automatic dependency management:**
- Just import packages in your code
- Run `go mod tidy` to download them
- No manual editing of go.mod needed!

**Version format:**
- Uses semantic versioning
- Example: `v1.8.1`
- Can specify exact versions or ranges

**Module proxy:**
- Downloads from proxy.golang.org by default
- Caches modules for faster builds
- Provides security and availability
