package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    cryptoPrice = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "crypto_price_usd",
            Help: "Cryptocurrency price in USD",
        },
        []string{"coin"},
    )
    
    apiCalls = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "api_calls_total",
            Help: "Total API calls to CoinGecko",
        },
    )
)

type CoinGeckoResponse struct {
    Bitcoin struct {
        USD float64 `json:"usd"`
    } `json:"bitcoin"`
    Ethereum struct {
        USD float64 `json:"usd"`
    } `json:"ethereum"`
}

func init() {
    prometheus.MustRegister(cryptoPrice)
    prometheus.MustRegister(apiCalls)
}

func fetchPrices() {
    apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"
    
    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    
    var data CoinGeckoResponse
    if err := json.Unmarshal(body, &data); err != nil {
        fmt.Printf("Parse error: %v\n", err)
        return
    }
    
    cryptoPrice.WithLabelValues("bitcoin").Set(data.Bitcoin.USD)
    cryptoPrice.WithLabelValues("ethereum").Set(data.Ethereum.USD)
    apiCalls.Inc()
    
    fmt.Printf("BTC: $%.2f | ETH: $%.2f\n", data.Bitcoin.USD, data.Ethereum.USD)
}

func main() {
    fmt.Println("Crypto Monitor Starting...")
    fmt.Println("==================================================")
    
    fetchPrices()
    
    ticker := time.NewTicker(10 * time.Second)
    go func() {
        for range ticker.C {
            fetchPrices()
        }
    }()
    
    http.Handle("/metrics", promhttp.Handler())
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Crypto Monitor</title>
    <meta http-equiv="refresh" content="10">
    <style>
        body { 
            font-family: Arial, sans-serif; 
            padding: 50px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            margin: 0;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: rgba(255,255,255,0.1);
            padding: 40px;
            border-radius: 20px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.3);
        }
        h1 { 
            text-align: center; 
            font-size: 48px;
            margin-bottom: 40px;
        }
        .price-card { 
            background: rgba(255,255,255,0.2);
            padding: 30px;
            margin: 20px 0;
            border-radius: 15px;
            text-align: center;
            transition: transform 0.3s;
        }
        .price-card:hover {
            transform: translateY(-5px);
        }
        .coin-name { 
            font-size: 24px; 
            margin-bottom: 10px;
            font-weight: 600;
        }
        .price { 
            font-size: 56px; 
            font-weight: bold; 
            margin: 10px 0;
        }
        .bitcoin { color: #f7931a; }
        .ethereum { color: #627eea; }
        .info { 
            text-align: center; 
            margin-top: 30px; 
            opacity: 0.8; 
        }
        .info a {
            color: white;
            text-decoration: none;
            border-bottom: 1px dashed white;
        }
        .info a:hover {
            border-bottom: 1px solid white;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Crypto Monitor</h1>
        <div class="price-card">
            <div class="coin-name">💰 Bitcoin (BTC)</div>
            <div class="price bitcoin" id="btc">Loading...</div>
        </div>
        <div class="price-card">
            <div class="coin-name">💎 Ethereum (ETH)</div>
            <div class="price ethereum" id="eth">Loading...</div>
        </div>
        <div class="info">
            <p>⏱️ Auto-refresh every 10 seconds</p>
            <p>📊 <a href="/metrics">Prometheus Metrics</a></p>
        </div>
    </div>
    <script>
        function updatePrices() {
            fetch('https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd')
                .then(res => res.json())
                .then(data => {
                    document.getElementById('btc').textContent = '$' + data.bitcoin.usd.toLocaleString('en-US', {minimumFractionDigits: 2, maximumFractionDigits: 2});
                    document.getElementById('eth').textContent = '$' + data.ethereum.usd.toLocaleString('en-US', {minimumFractionDigits: 2, maximumFractionDigits: 2});
                })
                .catch(err => console.error('Error:', err));
        }
        
        updatePrices();
        setInterval(updatePrices, 10000);
    </script>
</body>
</html>`
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprint(w, html)
    })
    
    fmt.Println("\nServer Started!")
    fmt.Println("   Web:     http://localhost:8080")
    fmt.Println("   Metrics: http://localhost:8080/metrics")
    fmt.Println("\nPrice Updates:")
    
    http.ListenAndServe(":8080", nil)
}
