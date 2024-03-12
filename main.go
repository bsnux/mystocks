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

type Ticket struct {
	Ticket string `yaml:"ticket"`
	Stocks int    `yaml:"stocks"`
}

type Tickets struct {
	Tickets []Ticket `yaml:"tickets"`
}

var wg sync.WaitGroup

func getTicket(t Ticket, tbl table.Table) error {
	defer wg.Done()

	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s?.tsrc=fin-srch", t.Ticket)
	// Header required here to get real time price
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.24")

	res, err := soup.Get(url)
	if err != nil {
		return err
	}
	doc := soup.HTMLParse(res)
	div := doc.Find("fin-streamer", "data-symbol", t.Ticket)
	val := div.Attrs()["value"]

	parseVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}

	p := message.NewPrinter(language.English)
	total := p.Sprintf("$%.2f", parseVal*float64(t.Stocks))
	stockPrice := p.Sprintf("$%.2f", parseVal)
	tbl.AddRow(t.Ticket, stockPrice, t.Stocks, total)

	return nil
}

func main() {
	var ticket Tickets
	yamlFile, err := os.ReadFile("tickets.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &ticket)
	if err != nil {
		panic(err)
	}

	headerFmt := color.New(color.FgGreen).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Ticket", "Value", "Stocks", "Total")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	wg.Add(len(ticket.Tickets))
	for _, i := range ticket.Tickets {
		go getTicket(i, tbl)
	}

	wg.Wait()

	tbl.Print()
}
