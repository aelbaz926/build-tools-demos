# pip Demo Application

A simple REST API using Flask demonstrating pip package management.

## 📋 Prerequisites

- **Python 3.8 or higher** (includes pip)

### Check if installed:
```bash
python3 --version
pip3 --version
```

### Installation (if needed):

**macOS:**
```bash
brew install python3
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install python3 python3-pip python3-venv
```

**Windows:**
- Download from https://www.python.org/downloads/

## 🚀 Running the Demo

### 1. Navigate to the demo directory
```bash
cd pip-demo
```

### 2. Create virtual environment
```bash
python3 -m venv venv
```

### 3. Activate virtual environment

**macOS/Linux:**
```bash
source venv/bin/activate
```

**Windows:**
```bash
venv\Scripts\activate
```

### 4. Install dependencies
```bash
pip install -r requirements.txt
```

### 5. Run the application
```bash
python app.py
```

### 6. Test the endpoints

Open a new terminal and run:
```bash
# Test root endpoint
curl http://localhost:5000/

# Test info endpoint
curl http://localhost:5000/info

# Test external API call
curl http://localhost:5000/external
```

Or open in browser: http://localhost:5000/

### 7. Stop the server
Press `Ctrl+C` in the terminal

### 8. Deactivate virtual environment
```bash
deactivate
```

## 📦 What This Demo Shows

- **requirements.txt** - pip dependencies file
- **Virtual environment** - Isolated Python environment
- **Dependencies** - Flask and requests
- **PyPI** - Python Package Index
- **pip install** - Package installation

## 🔍 Key Files

- `requirements.txt` - List of dependencies
- `app.py` - Flask application code
- `venv/` - Virtual environment (created during setup)

## 📚 pip Commands Reference

```bash
pip install -r requirements.txt    # Install from requirements file
pip install <package>               # Install single package
pip list                            # Show installed packages
pip freeze                          # Show installed packages with versions
pip freeze > requirements.txt       # Generate requirements file
pip show <package>                  # Show package details
pip uninstall <package>             # Remove package
```

## ✅ Expected Output

After running `python app.py`:
```
=== pip Demo Application ===

✅ Server running on http://localhost:5000

Available endpoints:
  GET http://localhost:5000/
  GET http://localhost:5000/info
  GET http://localhost:5000/external

Press Ctrl+C to stop

 * Serving Flask app 'app'
 * Debug mode: on
```

## 🌐 API Endpoints

- `GET /` - Basic info about the application
- `GET /info` - Dependency information
- `GET /external` - Demonstrates external API call using requests

## 💡 Virtual Environment Tips

**Why use virtual environments?**
- Isolate project dependencies
- Avoid version conflicts
- Keep system Python clean
- Easy to recreate environment

**Always activate before working:**
```bash
source venv/bin/activate  # macOS/Linux
venv\Scripts\activate     # Windows
```
