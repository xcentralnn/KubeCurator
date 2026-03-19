<p align="center">
  <img src="./docs/logo-alertless.png" width="220"/>
</p>

<h1 align="center">Alertless</h1>

<p align="center">
  Less noise. More signal.
</p>

![GitHub stars](https://img.shields.io/github/stars/xcentralnn/Alertless?style=social)
![GitHub forks](https://img.shields.io/github/forks/xcentralnn/Alertless?style=social)
![License](https://img.shields.io/badge/license-Apache--2.0-green)
![Go](https://img.shields.io/badge/Go-1.20+-blue?logo=go)
![Prometheus](https://img.shields.io/badge/Prometheus-Alertmanager-orange?logo=prometheus)
![Webhook](https://img.shields.io/badge/Webhook-Driven-lightgrey)
![Status](https://img.shields.io/badge/status-active-success)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)
---

## Overview

Alertless is a lightweight alert deduplication and noise reduction service designed for Prometheus Alertmanager.

It acts as an intermediary between Alertmanager and notification systems such as Slack or Telegram, ensuring that only meaningful alerts are delivered to engineers.

## Problem

Modern distributed systems generate large volumes of alerts:

- Repeated alerts overwhelm notification channels  
- Alert fatigue leads to ignored incidents  
- Increased noise reduces operational efficiency  

## Solution

Alertless reduces alert noise by applying simple but effective filtering mechanisms:

- Deduplicates repeated alerts within a configurable time window  
- Groups similar alerts into a single signal  
- Prioritizes actionable notifications  

## Architecture


Alertmanager → Alertless → Notification (Slack / Telegram / Email)


## Features

- Time-based alert deduplication  
- Lightweight and dependency-free design  
- Easy integration via Alertmanager webhook  
- Minimal resource footprint  

Planned:

- Notification integrations (Slack, Telegram)  
- Alert grouping and batching  
- Configurable rules via YAML  
- Redis-backed distributed deduplication  

## Quick Start

### Clone repository

```bash
git clone https://github.com/xcentralnn/Alertless.git
cd Alertless
Run service
go run ./cmd/server
```
The service will be available at:
```
http://localhost:8080
Test Webhook
curl -X POST http://localhost:8080/webhook \
-H "Content-Type: application/json" \
-d '{
  "alerts": [
    {
      "labels": {
        "alertname": "HighCPU",
        "instance": "pod-1",
        "severity": "warning"
      },
      "annotations": {
        "summary": "CPU usage high"
      }
    }
  ]
}'
```
### Alertmanager Integration

Example configuration:
```
receivers:
  - name: alertless
    webhook_configs:
      - url: http://alertless:8080/webhook
```

### How It Works

Alertless processes incoming alerts as follows:
Receives alerts via webhook
-> Generates a unique key based on alert labels
-> Applies a deduplication window
-> Forwards only non-duplicate alerts
<br>e.g: </br>
```
go run ./cmd/server/
2026/03/19 14:59:59 Alertless running on :8080
2026/03/19 15:00:55 Sending alert: HighCPU|pod-1|
2026/03/19 15:00:55 ALERT: HighCPU | CPU usage high
2026/03/19 15:01:00 Duplicate skipped: HighCPU|pod-1|
```

## Roadmap

Slack and Telegram integration

Alert grouping and batching

Configuration support (YAML / environment variables)

Distributed mode with Redis

## License

Apache License 2.0

## Contributing

Contributions are welcome. Please keep changes minimal, focused, and aligned with the project's goals.