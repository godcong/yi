# yi - I Ching (Book of Changes) Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/godcong/yi.svg)](https://pkg.go.dev/github.com/godcong/yi)

A Go library for I Ching (Book of Changes) hexagram calculation and divination, providing complete functionality for hexagram generation, transformation, Five Elements (WuXing), Sexagenary Cycle (JiaZi), Six Relations (LiuQin), Shi-Ying positions, 81 numerology, and hexagram interpretation.

---

## Project Architecture

The project uses a clean layered architecture: public types in `core/`, all implementation hidden in `internal/`, root package provides the only public API surface. External code can only import `yi` and `yi/core` — implementation details are fully encapsulated.

```
yi/
├── yi.go              # Public API - re-exports types, provides all entry points
├── yi_test.go         # Integration tests
├── doc.go             # Package declaration
├── core/              # Public types, constants, errors, data stores
│   ├── types.go       # Type definitions (Gua, Yao, ZhouYi, Bagua, WuXing, Sex, etc.)
│   ├── constants.go   # Constants (Bagua, YaoPosition, WuXing, LiuQin, etc.)
│   ├── errors.go      # Error definitions
│   ├── store.go       # Data store declarations
│   └── data_gen.go    # Auto-generated data initialization
├── internal/          # Implementation details (not importable externally)
│   ├── gua/           # Hexagram operations & transformations
│   ├── qigua/         # Divination methods (coins, dayan, meihua, time, lunar)
│   ├── jiegua/        # Hexagram interpretation & formatting
│   ├── wuxing/        # Five Elements operations
│   ├── jiazi/         # Sexagenary Cycle
│   ├── liuqin/        # Six Relations
│   ├── shiying/       # Shi-Ying positions
│   ├── numerology/    # 81 Numerology
│   ├── wenyan/        # Wenyan commentary
│   └── i18n/          # Internationalization (Chinese/English)
├── data/              # JSON data sources
├── cmd/
│   ├── divine/        # CLI tool
│   └── generate/      # Code generator
├── docs/              # Documentation
└── skills/            # AI Agent skills
```

> **Note**: Import `github.com/godcong/yi` for the public API. Types are also available via `github.com/godcong/yi/core`. Implementation packages under `internal/` are not importable by external code.

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
Overview       → Hexagram name / symbol / auspiciousness + core imagery + fortune基调
Details        → Primary / transformed / moving lines / nuclear / inverse / reverse hexagrams
8-Dim Fortune  → Career / Love / Wealth / Exams / Health / Travel / Lawsuit / Home
Do's & Don'ts  → Today's recommended actions / taboos
Guidance       → Action guidance + lucky direction / number / color
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
    zy := yi.DivineByDailyHexagram(2026, 5, 13, "John")
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
zy := yi.DivineByDailyHexagram(2026, 5, 13, "John")
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
| `DivineByDailyHexagram(year, month, day int, personalSeed string) *ZhouYi` | Daily hexagram |
| `DivineByCurrentTime(personalSeed ...string) *ZhouYi` | Current time hexagram |
| `DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult)` | Coin method |
| `DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult)` | Dayan (yarrow) method |
| `DivineByMeihua(upperNum, lowerNum, dongYao int) *ZhouYi` | Plum blossom method |
| `DivineByMeihuaTime(t time.Time, seeds ...string) (*ZhouYi, int, int, int, int)` | Plum blossom time |
| `DivineByTimeGua(params TimeGuaParams, seed ...string) *ZhouYi` | Gregorian time |
| `DivineByTime(year, month, day, hour int, seed ...string) *ZhouYi` | Time-based divination |
| `DivineByLunarTime(lunarYear, lunarMonth, lunarDay, shichenNum int) *ZhouYi` | Lunar time |
| `DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi` | Number-based |
| `Divine(shang, xia Bagua, bianYao ...int) *ZhouYi` | Trigram-based |

#### Interpretation

| Function | Description |
|----------|-------------|
| `JieGua(zy *ZhouYi, sex Sex) *JieGuaResult` | Main interpretation function |
| `JieGuaWithLang(zy *ZhouYi, sex Sex, lang Language) *JieGuaResult` | Interpretation with language |
| `FormatJieGua(result *JieGuaResult) string` | Formatted output (Chinese) |
| `FormatJieGuaWithLang(result *JieGuaResult, lang Language) string` | Formatted output with language |

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

#### Hexagram Lookup & Judgment

| Function | Description |
|----------|-------------|
| `GetGuaByIndex(index string) (*Gua, error)` | Lookup by index (e.g., "乾乾") |
| `GetGuaByXu(xu int) (*Gua, error)` | Lookup by sequence number (1-64) |
| `IsJi(zy *ZhouYi, sex Sex) bool` | Check auspiciousness (standalone function) |
| `FilterYao(zy *ZhouYi, sex Sex, filters ...string) bool` | Filter by auspiciousness |
| `GetShiYing(g *Gua) *ShiYingInfo` | Get Shi-Ying info (standalone function) |
| `GetGuaGong(g *Gua) Bagua` | Get palace attribution (standalone function) |
| `GetGuaPosition(g *Gua) GuaPosition` | Get hexagram position |

#### Five Elements & Six Relations

| Function | Description |
|----------|-------------|
| `GetWuXingByBagua(bagua Bagua) WuXing` | Trigram → Five Elements |
| `GetLiuQin(guaGongWX, yaoWX WuXing) LiuQin` | Calculate Six Relations |
| `Sheng(wx WuXing) WuXing` | Generating relation |
| `Ke(wx WuXing) WuXing` | Overcoming relation |
| `BeiSheng(wx WuXing) WuXing` | Being generated |
| `BeiKe(wx WuXing) WuXing` | Being overcome |

#### 81 Numerology

| Function | Description |
|----------|-------------|
| `GetDayan(n int) (*Dayan, error)` | Lookup numerology (1-81) |
| `MustGetDayan(n int) Dayan` | Lookup numerology (panics on invalid) |
| `Dayan.IsJi()` / `.IsXiong()` / `.IsBest()` | Auspiciousness check |

#### Wenyan Commentary

| Function | Description |
|----------|-------------|
| `GetWenYan(g *Gua) *WenYanData` | Get Wenyan commentary |
| `HasWenYan(g *Gua) bool` | Check if Wenyan exists |

#### Internationalization

| Function | Description |
|----------|-------------|
| `JieGuaWithLang(zy, sex, lang)` | Interpretation with language |
| `FormatJieGuaWithLang(result, lang)` | Formatted output with language |
| `TranslateJiXiong(jx, lang)` | Translate auspiciousness |
| `TranslateCategory(cat, lang)` | Translate category |
| `TranslateYaoPos(pos, lang)` | Translate line position |
| `TranslateBaguaName(name, lang)` | Translate trigram name |
| `TranslateWuXing(name, lang)` | Translate Five Elements |

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
| `Language` | Output language: LangZH / LangEN |

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
- `data/guagong.json` — Palace attribution

Regenerate `core/data_gen.go`:

```bash
go run ./cmd/generate/
```

The generator reads JSON files from `data/` and outputs `core/data_gen.go`.

---

## Documentation

- [Introduction to Numerology](docs/numerology.md) — The numerical chain from Yin-Yang to 81 numerology
- [Divination Methods](docs/divination-methods.md) — Four divination methods explained
- [Eight Trigrams & 64 Hexagrams](docs/trigrams-hexagrams.md) — Hexagram system and transformations
- [Dayan Numerology](docs/dayan-numerology.md) — 81 numerology auspiciousness
- [Six Relations & Shi-Ying](docs/six-relations-shiying.md) — Six Relations system and Shi-Ying calculation

---

## License

MIT License
