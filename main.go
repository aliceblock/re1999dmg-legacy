package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sort"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Data structure to hold chart data
type ChartData struct {
	Name   string  `json:"name"`
	Damage float64 `json:"damage"`
}

type Data struct {
	Title     string      `json:"title"`
	ChartData []ChartData `json:"chartData"`
}

func main() {
	templatesDir, err := filepath.Abs("templates")
	if err != nil {
		panic(err)
	}

	htmlFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.html"))
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{}
		err := tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/character", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		charName := queryParams.Get("name")

		data := Data{
			Title:     cases.Title(language.English).String(charName),
			ChartData: []ChartData{},
		}

		calculatorFunc := DmgCal.Calculator[DmgCal.CharacterIndex(charName)]
		responseDamage := calculatorFunc(DmgCal.CalParams{
			EnemyHit:       1,
			PsychubeAmp:    psychube.Amp1,
			ResonanceIndex: 0,
			EnemyDef:       600.0,
			Buff:           DmgCal.Buff{},
			Debuff: DmgCal.Debuff{
				SenseWeakness: true,
				Bkornblume:    1,
			},
		})
		for _, v := range responseDamage {
			data.ChartData = append(data.ChartData, ChartData(v))
		}

		sort.Slice(data.ChartData, func(i, j int) bool {
			return data.ChartData[i].Damage > data.ChartData[j].Damage
		})

		err := tmpl.ExecuteTemplate(w, "psychube_chart.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	http.ListenAndServe(":8088", nil)
}
