package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	yi "yi"
)

// Injected by goreleaser via -ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	showVersion := flag.Bool("version", false, "Print version and exit")
	method := flag.String("method", "time", "Divination method: time, daily, coins, meihua, dayan, number")
	seed := flag.String("seed", "", "Personal seed for unique result (user ID, name, etc.)")
	sex := flag.String("sex", "male", "Sex for interpretation: male or female")
	format := flag.String("format", "text", "Output format: text or json")
	lang := flag.String("lang", "zh", "Output language: zh (Chinese) or en (English)")
	benGua := flag.Int("ben", -1, "Ben gua number (0-7, for method=number)")
	bianGua := flag.Int("bian", -1, "Bian gua number (0-7, for method=number)")
	dongYao := flag.Int("dong", 0, "Dong yao position (0-5, for method=number)")
	coinsSeed := flag.Int64("coins-seed", 0, "Seed for coins method (0 = random)")
	dayanSeed := flag.Int64("dayan-seed", 0, "Seed for dayan method")
	year := flag.Int("year", 0, "Year (for method=time/meihua, default: now)")
	month := flag.Int("month", 0, "Month (for method=time/meihua, default: now)")
	day := flag.Int("day", 0, "Day (for method=time/meihua, default: now)")
	hour := flag.Int("hour", -1, "Hour 0-23 (for method=time/meihua, default: now)")

	flag.Parse()

	if *showVersion {
		fmt.Printf("yi %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	var langVal yi.Language
	switch strings.ToLower(*lang) {
	case "en", "english":
		langVal = yi.LangEN
	default:
		langVal = yi.LangZH
	}

	var zy *yi.ZhouYi

	var sexVal yi.Sex
	switch strings.ToLower(*sex) {
	case "female", "f":
		sexVal = yi.Female
	default:
		sexVal = yi.Male
	}

	now := time.Now()
	y := *year
	m := *month
	d := *day
	h := *hour
	if y == 0 {
		y = now.Year()
	}
	if m == 0 {
		m = int(now.Month())
	}
	if d == 0 {
		d = now.Day()
	}
	if h < 0 {
		h = now.Hour()
	}

	switch *method {
	case "coins":
		zy, _ = yi.DivineByCoins(*coinsSeed)

	case "dayan":
		zy, _ = yi.DivineByDayan(*dayanSeed)

	case "meihua":
		t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		if *seed != "" {
			zy, _, _, _, _ = yi.DivineByMeihuaTime(t, *seed)
		} else {
			zy, _, _, _, _ = yi.DivineByMeihuaTime(t)
		}

	case "number":
		if *benGua < 0 || *bianGua < 0 {
			fmt.Fprintln(os.Stderr, "method=number requires -ben and -bian flags (0-7)")
			os.Exit(1)
		}
		zy = yi.DivineByNumber(*benGua, *bianGua, *dongYao)

	default: // "time"
		params := yi.TimeGuaParams{Year: y, Month: m, Day: d, Hour: h}
		if *seed != "" {
			zy = yi.DivineByTimeGua(params, *seed)
		} else {
			zy = yi.DivineByTimeGua(params)
		}

	case "daily":
		if *seed == "" {
			fmt.Fprintln(os.Stderr, "method=daily requires -seed (your name, ID, or any stable identifier)")
			os.Exit(1)
		}
		zy = yi.DivineByDailyHexagram(y, m, d, *seed)
	}

	if zy == nil {
		fmt.Fprintln(os.Stderr, "divination failed: nil result")
		os.Exit(1)
	}

	result := yi.JieGuaWithLang(zy, sexVal, langVal)

	switch *format {
	case "json":
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	default:
		fmt.Println(yi.FormatJieGuaWithLang(result, langVal))
	}
}
