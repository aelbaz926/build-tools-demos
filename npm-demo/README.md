# npm Demo Application

A simple REST API using Express.js demonstrating npm package management.

## 📋 Prerequisites

- **Node.js 16 or higher** (includes npm)

### Check if installed:
```bash
node -v
npm -v
```

### Installation (if needed):

**macOS:**
```bash
brew install node
```

**Ubuntu/Debian:**
```bash
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
```

**Windows:**
- Download from https://nodejs.org/

## 🚀 Running the Demo

### 1. Navigate to the demo directory
```bash
cd npm-demo
```

### 2. Install dependencies
```bash
npm install
```

### 3. Start the application
```bash
npm start
```

### 4. Test the endpoints

Open a new terminal and run:
```bash
# Test root endpoint
curl http://localhost:3000/

# Test info endpoint
curl http://localhost:3000/info

# Test external API call
curl http://localhost:3000/external
```

Or open in browser: http://localhost:3000/

### 5. Stop the server
Press `Ctrl+C` in the terminal

## 📦 What This Demo Shows

- **package.json** - npm configuration file
- **Dependencies** - Express and Axios
- **npm scripts** - Custom commands
- **package-lock.json** - Dependency locking
- **node_modules/** - Installed packages
- **Semantic versioning** - Version ranges (^)

## 🔍 Key Files

- `package.json` - Project configuration and dependencies
- `package-lock.json` - Locked dependency versions (created after install)
- `index.js` - Application code
- `node_modules/` - Installed packages (created after install)

## 📚 npm Commands Reference

```bash
npm install              # Install all dependencies
npm install <package>    # Add new dependency
npm update               # Update dependencies
npm list                 # Show dependency tree
npm audit                # Check for vulnerabilities
npm run <script>         # Run custom script
```

## ✅ Expected Output

After running `npm start`:
```
=== npm Demo Application ===

✅ Server running on http://localhost:3000

Available endpoints:
  GET http://localhost:3000/
  GET http://localhost:3000/info
  GET http://localhost:3000/external

Press Ctrl+C to stop
```

## 🌐 API Endpoints

- `GET /` - Basic info about the application
- `GET /info` - Dependency information
- `GET /external` - Demonstrates external API call using axios
