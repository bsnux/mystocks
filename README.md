# My Stocks

Simple CLI tool for displaying the real time value of your current stocks based on the tickers
and number of stocks.

# Configuration

Open the `tickers.yml` file and add your stocks tickers and number of stocks.

YAML configuration file example:

```yaml
tickers:
  - ticker: AAPL
    stocks: 1
  - ticker: MSFT
    stocks: 2
  - ticker: GOOG
    stocks: 3
  - ticker: TSLA
    stocks: 4
```

# Usage

```
./mystocks
```

The following screenshot displays the result:

![Stocks value](table.png)

# Compiling

```
make build
```
