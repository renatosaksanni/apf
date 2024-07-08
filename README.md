[![GoDoc](https://godoc.org/github.com/renatosaksanni/apf?status.svg)](https://godoc.org/github.com/renatosaksanni/apf)
[![codecov](https://codecov.io/gh/renatosaksanni/apf/branch/main/graph/badge.svg)](https://codecov.io/gh/renatosaksanni/apf)
![Build Status](https://img.shields.io/github/actions/workflow/status/renatosaksanni/apf/ci.yml?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/renatosaksanni/apf)](https://goreportcard.com/report/github.com/renatosaksanni/apf)
![License](https://img.shields.io/github/license/renatosaksanni/apf.svg)
![Issues](https://img.shields.io/github/issues/renatosaksanni/apf.svg)

# APF - API Payment Transaction Forecasting

APF (API Payment Forecasting) is a project that demonstrates how to use Facebook Prophet and GARCH models for forecasting transaction volumes in an API payment system. This project includes fetching real-time transaction data from Alpha Vantage and visualizing the forecast results.

# Setup
### Prerequisites
- Python 3.x
- Golang
- pip (Python package installer)

### Install Dependencies
1. Clone the repository:
```sh
git clone https://github.com/renatosaksanni/apf.git
cd apf
```
2. Install Python packages:
```sh
pip install prophet arch matplotlib pandas requests
```
3. Install Go packages:
```sh
go mod tidy
```
### Configuration
1. Set up the Alpha Vantage API key:
```sh
export ALPHA_VANTAGE_API_KEY=your_api_key
```
# Usage
### Fetch Real-time Data
To fetch real-time data from Alpha Vantage:
```sh
go run cmd/main.go --fetch --symbol=AAPL
```
### Generate Forecast
To generate a forecast using Prophet:
```sh
go run cmd/main.go --model=prophet --data=data/real_time_data.csv --periods=30
```
To generate a forecast using GARCH:
```sh
go run cmd/main.go --model=garch --data=data/real_time_data.csv --periods=30
```
### Visualize Forecast
To visualize the forecast data:
```sh
python3 internal/infra/forecasting/visualize_forecast.py prophet
python3 internal/infra/forecasting/visualize_forecast.py garch
```
### Running Tests
To run the unit tests:
```sh
go test ./...
```
# Continuous Integration (CI)
This project uses GitHub Actions for Continuous Integration. The CI workflow is defined in .github/workflows/ci.yml and includes steps to set up Python and Go, install dependencies, and run tests.
```yaml
name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout repository
      uses: actions/checkout@v2

    - name: Set up Python
      uses: actions/setup-python@v2
      with:
        python-version: '3.x'

    - name: Install Python dependencies
      run: |
        pip install prophet arch matplotlib pandas requests

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.16'

    - name: Verify Go installation
      run: go version

    - name: Install Go dependencies
      run: go mod tidy

    - name: Run tests
      run: go test ./...
```      
# Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.
# License
This project is licensed under the MIT License - see the LICENSE file for details.
