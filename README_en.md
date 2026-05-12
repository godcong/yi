# yi - I Ching (Book of Changes) Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/godcong/yi.svg)](https://pkg.go.dev/github.com/godcong/yi)

A Go library for I Ching (Book of Changes) hexagram calculation and divination, providing complete functionality for hexagram generation, transformation, Five Elements (WuXing), Sexagenary Cycle (JiaZi), Six Relations (LiuQin), Shi-Ying positions, 81 numerology, and hexagram interpretation.

---

## Daily Hexagram Skill

**Recommended usage** — Integrate with your AI assistant for daily divination.

> User says "tell my fortune" → AI collects info → Generates full fortune report (800+ words)

### Installation

```bash
npx skills add https://github.com/godcong/yi --skill daily-hexagram
```

After installation, the AI will **automatically detect and download** the `yi` binary for your platform on first use (via built-in SKILL.md rules).

<details>
<summary>Manual binary installation (optional)</summary>

If automatic installation fails, run the install script manually:

```bash
# macOS / Linux
./scripts/install.sh

# Windows (PowerShell)
.\scripts\install.bat
```

Or compile from source:

```bash
go install github.com/godcong/yi/cmd/divine@latest
# Then copy the compiled binary to the skill's bin/ directory
```

</details>

### Compatible AI Agents

This Skill follows the [Agent Skills open standard](https://agentskills.io/) and works with all supporting AI agents:

| Agent | Install Command |
|-------|----------------|
| Claude Code | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a claude-code` |
| Trae | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a trae` |
| Cursor | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a cursor` |
| Codex | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a codex` |
| Goose | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a goose` |
| Gemini CLI | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a gemini-cli` |
| Roo Code | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a roo` |
| Windsurf | `npx skills add https://github.com/godcong/yi --skill daily-hexagram -a windsurf` |
| Others | Omit `-a` flag, interactively select target agent |

See [npx skills docs](https://www.npmjs.com/package/skills) for the full list.

### Usage

Simply talk to your AI assistant:

| You say | What AI does |
|---------|-------------|
| "tell my fortune", "daily horoscope" | Collects name + birthday → Daily hexagram (same person, same day = same hexagram) |
| "try another one" | Changes time and generates a new hexagram |
| "use coins" | Switches to coin method |
| "use yarrow stalks" | Switches to Dayan (yarrow) method |
| "plum blossom divination" | Switches to Meihua method |

**First use**: AI will ask your name and birthday. After that, it remembers. Same person on the same day always gets the same hexagram.

### Report Contents

```
Overview       — Hexagram name / symbol / auspiciousness + core imagery + fortune基调
Details        — Primary / transformed / moving lines / nuclear / inverse / reverse hexagrams
8-Dim Fortune  — Career / Love / Wealth / Exams / Health / Travel / Lawsuit / Home
Do's & Don'ts  — Today's recommended actions / taboos
Guidance       — Action guidance + lucky direction / number / color
```

Full report is at least 800 words.

### Supported Methods

| Method | CLI | Description |
|--------|-----|-------------|
| Daily | `daily` | Date-based only, same person same day fixed |
| Time | `time` | Includes hour, different times yield different hexagrams |
| Coins | `coins` | Six coin tosses, most classic |
| Dayan | `dayan` | Yarrow stalk method, most ancient |
| Meihua | `meihua` | Plum blossom time-based |
| Number | `number` | Custom upper/lower trigram numbers |

### Skill Directory Structure

```
daily-hexagram/
├── SKILL.md              # AI instructions (data-driven + AI interpretation)
├── bin/yi                # Divination program (compiled binary)
├── references/
│   └── data-format.md    # JSON output structure reference
├── scripts/
│   ├── install.sh        # Install script
│   └── install.bat
└── user_profile.json     # User info (auto-created)
```

---

## CLI Tool

### Installation

```bash
go install github.com/godcong/yi/cmd/divine@latest
```

### Usage

```bash
# Daily hexagram (requires -seed for personalization)
yi -method daily -seed "John" -format json

# Time-based hexagram
yi -method time -seed "John" -sex male -format json

# Coin method
yi -method coins -coins-seed 42

# Plum blossom method
yi -method meihua -seed "John" -year 2026 -month 5 -day 12

# Dayan (yarrow) method
yi -method dayan -dayan-seed 999

# Number-based hexagram
yi -method number -ben 3 -bian 5 -dong 2
```

### All Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-method` | `time` | Method: `time` / `daily` / `coins` / `meihua` / `dayan` / `number` |
| `-seed` | `""` | User identifier (name/ID), required for `daily` and `time` |
| `-format` | `text` | Output format: `text` / `json` |
| `-lang` | `zh` | Output language: `zh` (Chinese) / `en` (English) |
| `-sex` | `male` | Gender: `male` / `female` |
| `-year` | current year | Gregorian year |
| `-month` | current month | Gregorian month |
| `-day` | current day | Gregorian day |
| `-hour` | current hour | Hour 0-23 |
| `-coins-seed` | 0 | Coin method seed (0=random) |
| `-dayan-seed` | 0 | Dayan method seed |
| `-ben` | -1 | Number method upper trigram (0-7) |
| `-bian` | -1 | Number method lower trigram (0-7) |
| `-dong` | 0 | Number method moving line (0-5) |
| `-version` | — | Print version |

---

## Go Library

### Installation

```bash
go get github.com/godcong/yi
```

### Quick Start

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // Daily hexagram (same person, same day = same hexagram)
    zy, _ := yi.DivineByDailyHexagram("John")
    result := yi.JieGua(zy, yi.Male)
    fmt.Println(yi.FormatJieGua(result))

    // Coin method
    zy2, coins := yi.DivineByCoins(42)
    fmt.Printf("Primary: %s, Transformed: %s\n", zy2.GetGua(yi.Ben).Ming, zy2.GetGua(yi.Bian).Ming)

    // Dayan (yarrow) method
    zy3, _ := yi.DivineByDayan(999)

    // Plum blossom method
    zy4 := yi.DivineByMeihua(3, 5, 2)

    // Time-based hexagram
    zy5 := yi.DivineByTimeGua(yi.TimeGuaParams{
        Year: 2024, Month: 6, Day: 15, Hour: 10,
    })

    // Quick hex from current time
    zy6 := yi.DivineByCurrentTime()
    _ = zy6
}
```

### Daily Hexagram Function

```go
// Same person, same day = same hexagram
zy, err := yi.DivineByDailyHexagram("John")
if err != nil {
    panic(err)
}
result := yi.JieGua(zy, yi.Male)

// result.JieDu contains pre-built 8-dimension fortune + do's/don'ts
fmt.Println(result.JieDu.ShiYe)   // Career fortune
fmt.Println(result.JieDu.CoreImage) // Core imagery
fmt.Println(result.JieDu.Yi)      // Do's
fmt.Println(result.JieDu.Ji)      // Don'ts

// result.WuXingInfo contains Five Elements lucky attributes
fmt.Println(result.WuXingInfo.LuckyNumber) // Lucky number
fmt.Println(result.WuXingInfo.LuckyColor)  // Lucky color
```

### API Overview

#### Divination Functions

| Function | Description |
|----------|-------------|
| `DivineByDailyHexagram(seed string) (*ZhouYi, error)` | Daily hexagram |
| `DivineByCurrentTime(personalSeed ...string) *ZhouYi` | Current time hexagram |
| `DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult)` | Coin method |
| `DivineByCoinValues(values [6]int) (*ZhouYi, error)` | Coin method manual |
| `DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult)` | Dayan (yarrow) method |
| `DivineByDayanValues(values [6]int) (*ZhouYi, error)` | Dayan method manual |
| `DivineByMeihua(shang, xia, bian int) *ZhouYi` | Plum blossom method |
| `DivineByMeihuaTime(t time.Time, personalSeed ...string)` | Plum blossom time |
| `DivineByTimeGua(params TimeGuaParams, personalSeed ...string) *ZhouYi` | Gregorian time |
| `DivineByLunarTime(y, m, d, h int) *ZhouYi` | Lunar time |
| `DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi` | Number-based |

#### Interpretation

| Function | Description |
|----------|-------------|
| `JieGua(zy *ZhouYi, sex Sex) *JieGuaResult` | Main interpretation function |
| `FormatJieGua(result *JieGuaResult) string` | Formatted output |

#### Interpretation Result Structure

`JieGuaResult` contains:

| Field | Description |
|-------|-------------|
| `BenGuaInfo` / `BianGuaInfo` | Primary / Transformed hexagram {Ming, GuaName, Symbol, GuaYi, TuanText, XiangText, JiXiong, ...} |
| `HuGuaInfo` / `CuoGuaInfo` / `ZongGuaInfo` | Nuclear / Inverse / Reverse hexagrams |
| `DongYaoPos` / `DongYaoText` / `DongYaoJiXiong` | Moving line position / text / auspiciousness |
| `IsJi` / `JiXiongReason` | Overall auspiciousness / reason |
| `FenXi` | 8-dimension interpretation [{Category, Content, JiXiong, Source}] |
| `JieDu` | Pre-built interpretation {CoreImage, ShiYe, AiQing, ..., Yi, Ji} |
| `WuXingInfo` | Five Elements lucky info {WuXing, Direction, LuckyNumber, LuckyColor} |

#### Hexagram Lookup

| Function/Method | Description |
|-----------------|-------------|
| `GetGuaByIndex(index string) (*Gua, error)` | Lookup by index (e.g., "乾乾") |
| `GetGuaByXu(xu int) (*Gua, error)` | Lookup by sequence number (1-64) |
| `ZhouYi.GetGua(guaType int) *Gua` | Get primary / transformed / nuclear / inverse / reverse |
| `ZhouYi.IsJi(sex Sex) bool` | Check auspiciousness |
| `Gua.GetYao(pos YaoPosition) *Yao` | Get specific line |
| `Gua.GetShiYing() *ShiYingInfo` | Get Shi-Ying info |
| `Gua.GetGuaGong() Bagua` | Get palace归属 |

#### Five Elements & Six Relations

| Function/Method | Description |
|-----------------|-------------|
| `GetWuXingByNumber(n int) string` | Stroke count → Five Elements |
| `WuXing.Sheng()` / `.Ke()` / `.BeiSheng()` / `.BeiKe()` | Generating / overcoming relations |
| `GetLiuQin(guaGongWX, yaoWX WuXing) LiuQin` | Calculate Six Relations |

#### 81 Numerology

| Function/Method | Description |
|-----------------|-------------|
| `GetDayan(n int) (*Dayan, error)` | Lookup numerology (1-81) |
| `Dayan.IsJi()` / `.IsXiong()` / `.IsBest()` | Auspiciousness check |

### Core Types

| Type | Description |
|------|-------------|
| `Gua` (alias `Hexagram`) | Hexagram: name / meaning / Tuan / Xiang / six lines |
| `Yao` (alias `Line`) | Line: line text / auspiciousness / female fate |
| `ZhouYi` (alias `IChing`) | I Ching: five hexagram types + moving lines |
| `Bagua` (alias `Trigram`) | Eight Trigrams: 0=Qian...7=Kun |
| `WuXing` | Five Elements: Wood / Fire / Earth / Metal / Water |
| `Sex` | Gender: Male / Female |
| `LiuQin` | Six Relations: Parents / Siblings / Wealth / Offspring / Official |
| `Dayan` | Dayan numerology: 1-81 auspiciousness details |
| `JieGuaResult` | Interpretation result (includes JieDu + WuXingInfo) |
| `GuaJieDu` | Pre-built interpretation (8 dimensions + do's/don'ts) |
| `WuXingInfo` | Five Elements lucky attributes |

### Hexagram Type Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `Ben` | 0 | Primary hexagram |
| `Bian` | 1 | Transformed hexagram |
| `Hu` | 2 | Nuclear hexagram (lines 2-3-4 lower, 3-4-5 upper) |
| `Cuo` | 3 | Inverse hexagram (all yin/yang flipped) |
| `Zong` | 4 | Reverse hexagram (upside down) |

### Eight Trigrams Constants

| Constant | Value | WuXing | Name |
|----------|-------|--------|------|
| `Qian` | 0 | Metal | Qian ☰ (Heaven) |
| `Dui` | 1 | Metal | Dui ☱ (Lake) |
| `Li` | 2 | Fire | Li ☲ (Fire) |
| `Zhen` | 3 | Wood | Zhen ☳ (Thunder) |
| `Xun` | 4 | Wood | Xun ☴ (Wind) |
| `Kan` | 5 | Water | Kan ☵ (Water) |
| `Gen` | 6 | Earth | Gen ☶ (Mountain) |
| `Kun` | 7 | Earth | Kun ☷ (Earth) |

> **Convention**: This project marks Yang as 0, Yin as 1 (bit representation), which is opposite to intuition.

---

## Data Generation

Data is stored in JSON files under `data/`:

- `data/gua.json` — 64 hexagrams
- `data/tuan.json` — Tuan (Judgment) texts
- `data/xiang.json` — Xiang (Image) texts
- `data/wenyan.json` — Wenyan (Commentary) texts
- `data/jiazi.json` — Sexagenary Cycle (60 JiaZi)
- `data/jiegua.json` — Hexagram interpretations (8-dimension fortune + do's/don'ts + core imagery)
- `data/guagong.json` — Palace归属

Regenerate `data.gen.go`:

```bash
# Delete old file first, otherwise go generate will skip
rm data.gen.go
go generate ./...
```

**Note**: `go generate` only generates when `data.gen.go` doesn't exist. Delete it first after modifying data.

---

## Documentation

- [Introduction to Numerology](docs/数理简介.md) — The numerical chain from Yin-Yang to 81 numerology
- [Divination Methods](docs/起卦方法.md) — Four divination methods explained
- [Eight Trigrams & 64 Hexagrams](docs/八卦与六十四卦.md) — Hexagram system and transformations
- [Dayan Numerology](docs/大衍之数.md) — 81 numerology auspiciousness
- [Six Relations & Shi-Ying](docs/六亲与世应.md) — Six Relations system and Shi-Ying calculation

---

## License

MIT License
