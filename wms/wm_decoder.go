package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	expectedCols = []string{
		"Country,Other",
		"TotalCases",
		"NewCases",
		"TotalDeaths",
		"NewDeaths",
		"TotalRecovered",
		"ActiveCases",
		"Serious,Critical",
		"Tot\u00a0Cases/1M pop",
	}
)

type wmProcessor struct {
	out     []wmRow
	updated time.Time
}

type countTuple struct {
	Total int `json:"total"`
	New   int `json:"new"`
}

type wmRow struct {
	Country   string     `json:"country"`
	Cases     countTuple `json:"cases"`
	Deaths    countTuple `json:"deaths"`
	Recovered int        `json:"recovered"`
	Active    int        `json:"active"`
	Serious   int        `json:"serious"`
}

func (p *wmProcessor) Decode(r io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return fmt.Errorf("parse failed: %v", err)
	}

	doc.Find("#main_table_countries").Each(func(i int, s *goquery.Selection) {
		// Sanity check that the headings havent changed.
		var errInner error
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			if i >= len(expectedCols) {
				errInner = fmt.Errorf("new columns were added")
				return
			}
			if got, want := s.Text(), expectedCols[i]; got != want {
				errInner = fmt.Errorf("unexpected column: %q != %q", got, want)
				return
			}
		})
		if errInner != nil {
			err = errInner
			return
		}

		// Iterate through each row.
		s.Find("tbody > tr").Each(func(i int, s *goquery.Selection) {
			var (
				row wmRow
				sub = s.Find("td")
				err error
			)
			row.Country = strings.TrimSpace(sub.Eq(0).Text())
			if row.Country == "Total:" || row.Country == "Bolivia" {
				return
			}

			row.Cases.Total, err = strconv.Atoi(clean(sub.Eq(1).Text()))
			if err != nil && clean(sub.Eq(1).Text()) != "" {
				errInner = fmt.Errorf("parsing cases total: %v", err)
				return
			}
			row.Cases.New, err = strconv.Atoi(clean(sub.Eq(2).Text()))
			if err != nil && clean(sub.Eq(2).Text()) != "" {
				errInner = fmt.Errorf("parsing cases new: %v", err)
				return
			}
			row.Deaths.Total, err = strconv.Atoi(clean(sub.Eq(3).Text()))
			if err != nil && clean(sub.Eq(3).Text()) != "" {
				errInner = fmt.Errorf("parsing total: %v", err)
				return
			}
			row.Deaths.New, err = strconv.Atoi(clean(sub.Eq(4).Text()))
			if err != nil && clean(sub.Eq(4).Text()) != "" {
				errInner = fmt.Errorf("parsing new: %v", err)
				return
			}
			row.Recovered, err = strconv.Atoi(clean(sub.Eq(5).Text()))
			if err != nil && clean(sub.Eq(5).Text()) != "" {
				errInner = fmt.Errorf("parsing recovered: %v", err)
				return
			}
			row.Active, err = strconv.Atoi(clean(sub.Eq(6).Text()))
			if err != nil && clean(sub.Eq(6).Text()) != "" {
				errInner = fmt.Errorf("parsing active: %v", err)
				return
			}
			row.Serious, err = strconv.Atoi(clean(sub.Eq(7).Text()))
			if err != nil && clean(sub.Eq(7).Text()) != "" {
				errInner = fmt.Errorf("parsing serious: %v", err)
				return
			}

			//fmt.Printf("tr %03d: %+v\n", i, row)
			p.out = append(p.out, row)
		})
		if errInner != nil {
			err = errInner
			return
		}
	})

	// Lastly, determine when it was last updated.
	doc.Find("div.content-inner > div").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		if len(t) < 200 && strings.HasPrefix(t, "Last updated: ") {
			t = strings.TrimPrefix(s.Text(), "Last updated: ")
			p.updated, err = time.Parse("January 02, 2006, 15:04 GMT", t)
		}
	})

	return err
}

func clean(in string) string {
	in = strings.Replace(in, ",", "", -1)
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "\t", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	in = strings.Replace(in, "$", "", -1)
	in = strings.Replace(in, "+", "", -1)
	in = strings.Replace(in, "-", "", -1)
	return in
}
