# Quick Start Guide

## 🎯 Test All Demos Quickly

### Maven Demo (Java)
```bash
cd maven-demo
mvn clean compile
mvn exec:java -Dexec.mainClass="com.example.demo.App"
```

### npm Demo (Node.js)
```bash
cd npm-demo
npm install
npm start
# In another terminal: curl http://localhost:3000/
# Press Ctrl+C to stop
```

### pip Demo (Python)
```bash
cd pip-demo
python3 -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
python app.py
# In another terminal: curl http://localhost:5000/
# Press Ctrl+C to stop, then: deactivate
```

## 📋 Prerequisites Check

Run these commands to verify you have everything installed:

```bash
# Java & Maven
java -version    # Should be 11+
mvn -version     # Should be 3.6+

# Node.js & npm
node -v          # Should be 16+
npm -v           # Should be 8+

# Python & pip
python3 --version  # Should be 3.8+
pip3 --version
```

## 🚀 Installation Commands

### macOS (using Homebrew)
```bash
brew install openjdk@11 maven node python3
```

### Ubuntu/Debian
```bash
sudo apt update
sudo apt install openjdk-11-jdk maven nodejs npm python3 python3-pip python3-venv
```

## 📁 Repository Structure

```
build-tools-demos/
├── maven-demo/          # Java + Maven
│   ├── pom.xml
│   ├── src/main/java/
│   └── README.md
├── npm-demo/            # Node.js + npm
│   ├── package.json
│   ├── index.js
│   └── README.md
├── pip-demo/            # Python + pip
│   ├── requirements.txt
│   ├── app.py
│   └── README.md
└── README.md
```

## 💡 Tips

- Each demo has its own detailed README with step-by-step instructions
- Maven demo: Simple console application
- npm demo: REST API on port 3000
- pip demo: REST API on port 5000
- All demos show dependency management in action
