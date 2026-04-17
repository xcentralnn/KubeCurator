<p align="center">
  <img src="./docs/logo.png" width="220"/>
</p>

<h1 align="center">KubeCurator</h1>

<p align="center">
  Kubernetes-native intelligent control plane.
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

KubeCurator is a Kubernetes-native intelligent control plane designed to optimize workloads, reduce misconfigurations, and improve resource efficiency.

It operates as a set of custom controllers that continuously observe cluster state, analyze system behavior, and apply corrective or optimizing actions.

## Problem

Modern Kubernetes environments face three core challenges:

Unpredictable workload scaling
Frequent misconfigurations in manifests
Resource inefficiency (idle, unused, or improperly scheduled objects)

These issues lead to:

Increased infrastructure cost
Reduced system reliability
Operational complexity and drift

## Solution

KubeCurator introduces an intelligent reconciliation layer:

Machine learning for predictive autoscaling
Policy-based validation for configuration correctness
Heuristic and anomaly detection for resource optimization

## License

MIT License

## Contributing

Contributions are welcome. Please keep changes minimal, focused, and aligned with the project's goals.