# Snap ASPI Go

A production-ready Golang implementation of the ASPI (Indonesia Payment System Administration) API wrapper. Supports Virtual Account and MPM (Merchant Payment Management) services with full endpoint coverage.

---

## 🚀 Quick Start

### 1. Clone the Repository
```sh
git clone <your-repo-url>
cd snap
```

---

## 🛠️ Running Locally (Direct Go)

### 1. **Install Go**
- Requires Go 1.23 or later

### 2. **Prepare Config and Keys**
- Place your config file in `config/config.yaml`
- Place your RSA keys in `keys/private_key.pem` and `keys/public_key.pem`

### 3. **Build the Application**
```sh
go build -o bin/snap-aspi-go ./cmd/snap-bi-go
```

### 4. **Run the Application**
```sh
ASPI_CLIENT_ID=your-client-id ASPI_CLIENT_SECRET=your-client-secret ./bin/snap-aspi-go
```

### 5. **Health Check**
```sh
curl http://localhost:8080/health
```

---

## 🐳 Running with Docker

### 1. **Build the Docker Image**
```sh
docker build -t snap-aspi-go:latest .
```

### 2. **Run the Container**
```sh
docker run --rm -p 8080:8080 \
  -e ASPI_CLIENT_ID=your-client-id \
  -e ASPI_CLIENT_SECRET=your-client-secret \
  -v $(pwd)/config:/app/config \
  -v $(pwd)/keys:/app/keys \
  snap-aspi-go:latest
```

### 3. **Health Check**
```sh
curl http://localhost:8080/health
```

---

## 🐳 Using Docker Compose (Optional)

Create a `docker-compose.yml`:
```yaml
version: '3.8'
services:
  snap-aspi-go:
    build: .
    image: snap-aspi-go:latest
    ports:
      - "8080:8080"
    environment:
      - ASPI_CLIENT_ID=your-client-id
      - ASPI_CLIENT_SECRET=your-client-secret
    volumes:
      - ./config:/app/config
      - ./keys:/app/keys
```

Run:
```sh
docker-compose up --build
```

---

## ⚙️ Configuration & Environment
- **Config file**: `config/config.yaml` (YAML format)
- **Keys**: `keys/private_key.pem`, `keys/public_key.pem`
- **Environment variables**:
  - `ASPI_CLIENT_ID` (required)
  - `ASPI_CLIENT_SECRET` (required)
  - Any config value can be overridden with `ASPI_` prefix

---

## 🧪 Testing

### Run All Tests
```sh
./scripts/run_tests.sh
```

### Run Unit Tests
```sh
go test -v ./test/unit/...
```

### Run Integration Tests
```sh
go test -v ./test/integration/...
```

### Generate Coverage Report
```sh
go test -coverprofile=coverage.out ./test/unit/...
go tool cover -html=coverage.out -o coverage.html
```

---

## 📖 API Endpoints

- **Authentication**: `/api/v1/auth/b2b-token`, `/api/v1/auth/b2b2c-token`
- **Virtual Account**: `/api/v1/virtual-account/*`
- **MPM**: `/api/v1/mpm/*`
- **Health**: `/health`
- **Config**: `/api/v1/utils/config`

---

## 📝 Notes
- For production, ensure your keys and config are secured and not world-readable.
- The application logs in JSON by default.
- Use `GIN_MODE=release` for production (set automatically in Docker).

---

## 📬 Support
For issues or questions, please open an issue in this repository. 