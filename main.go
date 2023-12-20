package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Data structure to hold chart data
type ChartData struct {
	Name   string  `json:"name"`
	Damage float64 `json:"damage"`
	Color  string  `json:"color"`
}

type Data struct {
	Title     string      `json:"title"`
	ChartData []ChartData `json:"chartData"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
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

	publicPath := "./public"

	// Create a file server for the specified directory
	fileServer := http.FileServer(http.Dir(publicPath))

	// Register the file server to handle requests under "/public/"
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/character", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		charName := queryParams.Get("name")
		enemyHitRaw, _ := strconv.ParseInt(queryParams.Get("enemy"), 10, 8)
		enemyHit := int16(enemyHitRaw)

		psychubeAmp1 := queryParams.Get("amp1")
		psychubeAmp2 := queryParams.Get("amp2")
		psychubeAmp3 := queryParams.Get("amp3")
		psychubeAmp4 := queryParams.Get("amp4")
		psychubeAmp5 := queryParams.Get("amp5")

		// Buff
		buff := DmgCal.Buff{}
		anAnLee := queryParams.Get("ananlee")
		sonetto := queryParams.Get("sonetto")
		necrologist := queryParams.Get("necrologist")
		if anAnLee != "" {
			value, _ := strconv.ParseInt(anAnLee, 10, 8)
			buff.AnAnLee = int16(value)
		}
		if sonetto != "" {
			value, _ := strconv.ParseInt(sonetto, 10, 8)
			buff.Sonetto = int16(value)
		}
		if necrologist == "true" {
			buff.Necrologist = true
		}

		// Debuff
		debuff := DmgCal.Debuff{}
		bkornblume := queryParams.Get("bkornblume")
		babyblueskill1 := queryParams.Get("babyblueskill1")
		babyblueskill2 := queryParams.Get("babyblueskill2")
		confusion := queryParams.Get("confusion")
		babyteeth := queryParams.Get("babyteeth")
		senseweakness := queryParams.Get("senseweakness")
		if bkornblume != "" {
			value, _ := strconv.ParseInt(bkornblume, 10, 8)
			debuff.Bkornblume = int16(value)
		}
		if babyblueskill1 != "" {
			value, _ := strconv.ParseInt(babyblueskill1, 10, 8)
			debuff.BabyBlueSkill1 = int16(value)
		}
		if babyblueskill2 != "" {
			value, _ := strconv.ParseInt(babyblueskill2, 10, 8)
			debuff.BabyBlueSkill2 = int16(value)
		}
		if confusion != "" {
			value, _ := strconv.ParseInt(confusion, 10, 8)
			debuff.Confusion = int16(value)
		}
		if babyteeth == "true" {
			debuff.BabyTeeth = true
		}
		if senseweakness == "true" {
			debuff.SenseWeakness = true
		}

		afflatusAdvantageBool := false
		afflatusAdvantage := queryParams.Get("afflatusadvantage")
		if afflatusAdvantage == "true" {
			afflatusAdvantageBool = true
		}

		data := Data{
			Title:     cases.Title(language.English).String(charName),
			ChartData: []ChartData{},
		}

		calculatorFunc := DmgCal.Calculator[DmgCal.CharacterIndex(charName)]
		amps := []psychube.Amplification{}
		if psychubeAmp1 == "true" {
			amps = append(amps, psychube.Amp1)
		}
		if psychubeAmp2 == "true" {
			amps = append(amps, psychube.Amp2)
		}
		if psychubeAmp3 == "true" {
			amps = append(amps, psychube.Amp3)
		}
		if psychubeAmp4 == "true" {
			amps = append(amps, psychube.Amp4)
		}
		if psychubeAmp5 == "true" {
			amps = append(amps, psychube.Amp5)
		}

		hasBoundenDuty := false
		for _, amp := range amps {
			responseDamage := calculatorFunc(DmgCal.CalParams{
				EnemyHit:          enemyHit,
				PsychubeAmp:       amp,
				ResonanceIndex:    0,
				EnemyDef:          600.0,
				Buff:              buff,
				Debuff:            debuff,
				AfflatusAdvantage: afflatusAdvantageBool,
			})
			for _, res := range responseDamage {
				if hasBoundenDuty && strings.Contains(res.Name, "His Bounden Duty") {
					continue
				}
				chartData := ChartData{
					Name:   res.Name,
					Damage: res.Damage,
					Color:  getColor(res.Name),
				}
				if strings.Contains(res.Name, "His Bounden Duty") {
					hasBoundenDuty = true
				}
				data.ChartData = append(data.ChartData, chartData)
			}
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
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

func getColor(psyChubeName string) string {
	if strings.Contains(psyChubeName, "Brave") {
		return "#FF6E40"
	}
	if strings.Contains(psyChubeName, "Thunder") {
		return "#D50000"
	}
	if strings.Contains(psyChubeName, "Lux") {
		return "#80CBC4"
	}
	if strings.Contains(psyChubeName, "His") {
		return "#F1F8E9"
	}
	if strings.Contains(psyChubeName, "Yearning") {
		return "#FFE57F"
	}
	if strings.Contains(psyChubeName, "Hop") {
		return "#E0E0E0"
	}
	if strings.Contains(psyChubeName, "Blas") {
		return "#8D6E63"
	}
	return "#3D5AFE"
}
