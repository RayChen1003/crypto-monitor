# 🚀 Crypto Monitor

Real-time cryptocurrency price monitoring with Prometheus metrics integration.

## 📊 Features

- 💰 **Bitcoin (BTC)** price tracking
- 💎 **Ethereum (ETH)** price tracking
- 📈 **Prometheus metrics** export
- 🌐 Beautiful web interface
- ⏱️ Auto-refresh every 10 seconds
- 🔄 Live price updates from CoinGecko API

## 🛠️ Tech Stack

- **Go 1.21+**
- **Prometheus Client**
- **CoinGecko API**
- **Kubernetes** (deployment ready)

## 🚀 Quick Start
```bash
# Clone the repository
git clone https://github.com/RayChen1003/crypto-monitor.git
cd crypto-monitor

# Run the application
go run main.go
```

**Access the web interface:** http://localhost:8080

**Prometheus metrics endpoint:** http://localhost:8080/metrics

## 📸 Screenshot

Web interface shows real-time cryptocurrency prices with auto-refresh functionality.

## 📊 Prometheus Metrics

Available metrics:
- `crypto_price_usd{coin="bitcoin"}` - Bitcoin price in USD
- `crypto_price_usd{coin="ethereum"}` - Ethereum price in USD
- `api_calls_total` - Total API calls to CoinGecko

## 🐳 Docker Support
```bash
# Build Docker image
docker build -t crypto-monitor .

# Run container
docker run -p 8080:8080 crypto-monitor
```

## ☸️ Kubernetes Deployment

Ready for deployment to Kubernetes with Prometheus monitoring integration.

## 📚 Learning Journey

This project is part of my DevOps learning journey, covering:
- Go programming
- REST API integration
- Prometheus monitoring
- Container orchestration
- Kubernetes deployment

## 📄 License

MIT License

## 👤 Author

**Ray Chen** - [GitHub](https://github.com/RayChen1003)

DevOps Engineer in Training - 2025

---

⭐ Star this repo if you find it helpful!
