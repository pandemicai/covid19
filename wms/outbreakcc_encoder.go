package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	pdt = time.FixedZone("PDT", -7*60*60)
	edt = time.FixedZone("EDT", -4*60*60)
	num = message.NewPrinter(language.English)
)

const (
	sumKeyWorld = "Worldwide"
	sumKeyChina = "China"
	sumKeyRest  = "Rest of the World"
)

type liveUpdateData struct {
	UpdateTime string             `json:"update_time"`
	Summary    []summaryRow       `json:"summary"`
	Breakdown  []countryBreakdown `json:"world"`
}

func (d *liveUpdateData) setUpdatedTime(t time.Time) {
	utcPrefix := t.UTC().Format("3PM Jan 2 2006")
	pt := t.In(pdt).Format("3PM") + " PDT"
	et := t.In(edt).Format("3PM") + " EDT"
	d.UpdateTime = fmt.Sprintf("%s (%s, %s)", utcPrefix, et, pt)
}

type summaryRow struct {
	Row       string `json:"row"`
	Icon      string `json:"icon"`
	Confirmed string `json:"confirmed"`
	Death     string `json:"death"`
	Suspected string `json:"suspected"`
	Active    string `json:"active"`
	Serious   string `json:"serious"`
	Recovered string `json:"recovered"`
}

func updateString(currentVal string, add int) string {
	var (
		v   int
		err error
	)
	if currentVal != "" {
		if v, err = strconv.Atoi(strings.Replace(currentVal, ",", "", -1)); err != nil {
			return "ERROR: " + err.Error()
		}
	}
	return num.Sprintf("%d", v+add)
}

func (s *summaryRow) Add(r wmRow) {
	s.Confirmed = updateString(s.Confirmed, r.Cases.Total)
	s.Death = updateString(s.Death, r.Deaths.Total)
	s.Suspected = ""
	s.Active = updateString(s.Active, r.Active)
	s.Serious = updateString(s.Serious, r.Serious)
	s.Recovered = updateString(s.Recovered, r.Recovered)
}

type countryBreakdown struct {
	Name        string `json:"country"`
	NameFR      string `json:"country_fr"`
	NameCN      string `json:"country_cn"`
	NameJP      string `json:"country_jp"`
	URL         string `json:"url"`
	Source      string `json:"source"`
	CountryCode string `json:"country_code"`
	Icon        string `json:"icon"`
	OfficialURL string `json:"official_url"`

	New       string `json:"new"`
	Confirmed string `json:"confirmed"`
	Suspected string `json:"suspected"`
	Death     string `json:"death"`
	NewDeath  string `json:"death_new"`
	Active    string `json:"active"`
	Cured     string `json:"cured"`
	Serious   string `json:"serious"`

	Unused1 string `json:"cities"`
}

func load(path string) (*liveUpdateData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out liveUpdateData
	return &out, json.NewDecoder(f).Decode(&out)
}

func (p *wmProcessor) computeSummaries() ([]summaryRow, error) {
	var (
		total = summaryRow{Row: sumKeyWorld, Icon: "world.png"}
		china = summaryRow{Row: sumKeyChina, Icon: "CN.png"}
		rest  = summaryRow{Row: sumKeyRest, Icon: ""}
	)

	for _, country := range p.out {
		switch country.Country {
		case "China":
			total.Add(country)
			china.Add(country)
		default:
			total.Add(country)
			rest.Add(country)
		}
	}
	return []summaryRow{total, china, rest}, nil
}

func findCountryIdx(name string, data *liveUpdateData) int {
	for i, row := range data.Breakdown {
		if row.Name == name {
			return i
		}
	}
	return -1
}

func mapCountryName(worldmeterName string) string {
	switch worldmeterName {
	case "S. Korea":
		return "South Korea"
	case "UK":
		return "United Kingdom"
	case "Czechia":
		return "Czech Republic"
	case "Macao":
		return "Macau"
	case "Vatican City":
		return "Vatican"
	}

	return worldmeterName
}

func (p *wmProcessor) upsertCountryData(data *liveUpdateData) error {
	for _, row := range p.out {
		idx := findCountryIdx(mapCountryName(row.Country), data)
		if idx < 0 {
			data.Breakdown = append(data.Breakdown, countryBreakdown{Name: row.Country})
			idx = len(data.Breakdown) - 1
		}

		data.Breakdown[idx].Active = num.Sprintf("%v", row.Active)
		data.Breakdown[idx].Confirmed = num.Sprintf("%v", row.Cases.Total)
		data.Breakdown[idx].Cured = num.Sprintf("%v", row.Recovered)
		data.Breakdown[idx].Death = num.Sprintf("%v", row.Deaths.Total)
		data.Breakdown[idx].New = num.Sprintf("%v", row.Cases.New)
		data.Breakdown[idx].NewDeath = num.Sprintf("%v", row.Deaths.New)
		data.Breakdown[idx].Serious = num.Sprintf("%v", row.Serious)
		data.Breakdown[idx].Suspected = ""
	}
	return nil
}

func (p *wmProcessor) Encode(path string) error {
	data, err := load(path)
	if err != nil {
		return fmt.Errorf("failed load: %v", err)
	}
	data.setUpdatedTime(p.updated)
	// fmt.Printf("Worldmeter was last updated %v\n", p.updated)

	if data.Summary, err = p.computeSummaries(); err != nil {
		return err
	}
	if err := p.upsertCountryData(data); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	return e.Encode(data)
}
