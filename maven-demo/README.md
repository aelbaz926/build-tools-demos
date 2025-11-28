# Maven Demo Application

A simple Java application demonstrating Maven build tool and dependency management.

## 📋 Prerequisites

- **Java JDK 11 or higher**
- **Maven 3.6 or higher**

### Check if installed:
```bash
java -version
mvn -version
```

### Installation (if needed):

**macOS:**
```bash
brew install openjdk@11
brew install maven
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install openjdk-11-jdk maven
```

**Windows:**
- Download JDK from https://adoptium.net/
- Download Maven from https://maven.apache.org/download.cgi
- Add to PATH

## 🚀 Running the Demo

### 1. Navigate to the demo directory
```bash
cd maven-demo
```

### 2. Clean and compile
```bash
mvn clean compile
```

### 3. Run the application
```bash
mvn exec:java -Dexec.mainClass="com.example.demo.App"
```

### 4. Package as JAR
```bash
mvn package
```

### 5. Run the JAR
```bash
java -jar target/maven-demo-1.0.0.jar
```

## 📦 What This Demo Shows

- **pom.xml** - Maven configuration file
- **Dependencies** - Apache Commons Lang and Gson
- **Build lifecycle** - clean, compile, package
- **Maven Central** - Automatic dependency download
- **JAR creation** - Packaging application

## 🔍 Key Files

- `pom.xml` - Project configuration and dependencies
- `src/main/java/` - Java source code
- `target/` - Build output (created after build)
- `~/.m2/repository/` - Local Maven repository

## 📚 Maven Commands Reference

```bash
mvn clean              # Remove target directory
mvn compile            # Compile source code
mvn test               # Run tests
mvn package            # Create JAR file
mvn install            # Install to local repository
mvn dependency:tree    # Show dependency tree
```

## ✅ Expected Output

```
=== Maven Demo Application ===

Original: hello maven
Capitalized: Hello maven

Project Info (JSON):
{"tool":"Maven","language":"Java","version":"1.0.0","dependencies":2}

✅ Maven demo completed successfully!
```
