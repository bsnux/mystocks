package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"sync"
)

type Ticker struct {
	Ticker string `yaml:"ticker"`
	Stocks int    `yaml:"stocks"`
}

type Tickers struct {
	Tickers []Ticker `yaml:"tickers"`
}

var wg sync.WaitGroup

func getTicker(t Ticker, tbl table.Table) error {
	defer wg.Done()

	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s?.tsrc=fin-srch", t.Ticker)
	// Header required here to get real time price
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.24")

	res, err := soup.Get(url)
	if err != nil {
		return err
	}
	doc := soup.HTMLParse(res)
	div := doc.Find("fin-streamer", "data-symbol", t.Ticker)
	val := div.Attrs()["value"]

	parseVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}

	p := message.NewPrinter(language.English)
	total := p.Sprintf("$%.2f", parseVal*float64(t.Stocks))
	stockPrice := p.Sprintf("$%.2f", parseVal)
	tbl.AddRow(t.Ticker, stockPrice, t.Stocks, total)

	return nil
}

func main() {
	var ticker Tickers
	yamlFile, err := os.ReadFile("tickers.yml")
	if err != nil {
		fmt.Printf("Error: %s. Make sure you created the required tickers.yml file\n", err.Error())
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, &ticker)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	headerFmt := color.New(color.FgGreen).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Ticker", "Value", "Stocks", "Total")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	wg.Add(len(ticker.Tickers))
	for _, i := range ticker.Tickers {
		go getTicker(i, tbl)
	}

	wg.Wait()

	tbl.Print()
}
