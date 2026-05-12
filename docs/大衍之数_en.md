# Dayan Numerology / 大衍之数

The Dayan Numerology system is the core framework for analyzing stroke count fortune in name studies, mapping numbers 1-81 to different fortune attributes and detailed interpretations.

## 81 Numerology Concepts

The *I Ching · Great Commentary* states: "The numbers of the Great Expansion amount to fifty, of which forty-nine are used." In this project, the Dayan Numbers are extended into a 1-81 stroke count fortune system, used for name stroke analysis.

Each number contains the following dimensions:

| Dimension | Field | Description |
|-----------|-------|-------------|
| Stroke Count | `Number` | 1-81 |
| Fortune | `JiXiong` | Auspicious/Inauspicious/Semi-Auspicious |
| Female Fortune | `NvMing` | Female-specific fortune (empty means gender-neutral) |
| Most Auspicious | `IsMax` | Whether it's the most auspicious number |
| Corresponding Hexagram | `Gua` | Corresponding hexagram name |
| Tian Jiu | `TianJiu` | Taiji number name |
| Imagery | `YiXiang` | Imagery description |
| Foundation | `Foundation` | Career foundation |
| Family | `Family` | Family situation |
| Health | `Health` | Health tips |
| Meaning | `Meaning` | Detailed meaning interpretation |

## Dayan Struct

```go
type Dayan struct {
    Number     int    // Stroke count (1-81)
    JiXiong    string // Fortune (吉/凶/半吉)
    NvMing     string // Female-specific fortune
    IsMax      bool   // Whether most auspicious
    Gua        string // Corresponding hexagram
    TianJiu    string // Tian Jiu (Taiji number)
    YiXiang    string // Imagery
    Foundation string // Foundation
    Family     string // Family
    Health     string // Health
    Meaning    string // Meaning
}
```

## API

### Querying Numerology

```go
// Query by stroke count (1-81)
dayan, err := yi.GetDayan(21)
if err != nil {
    // err == ErrInvalidDayanNumber (out of 1-81 range)
    panic(err)
}
fmt.Printf("Stroke: %d\n", dayan.Number)
fmt.Printf("Fortune: %s\n", dayan.JiXiong)
fmt.Printf("Tian Jiu: %s\n", dayan.TianJiu)
fmt.Printf("Imagery: %s\n", dayan.YiXiang)
fmt.Printf("Foundation: %s\n", dayan.Foundation)
fmt.Printf("Family: %s\n", dayan.Family)
fmt.Printf("Health: %s\n", dayan.Health)
fmt.Printf("Meaning: %s\n", dayan.Meaning)

// Ignore errors (out of bounds returns zero value)
dayan2 := yi.MustGetDayan(100) // returns Dayan{}
```

### Fortune Judgment

```go
dayan, _ := yi.GetDayan(21)

// Is auspicious (吉 or 半吉)
fmt.Printf("Is auspicious: %v\n", dayan.IsJi())

// Is inauspicious
fmt.Printf("Is inauspicious: %v\n", dayan.IsXiong())

// Suitable for female
fmt.Printf("Suitable for female: %v\n", dayan.IsSuitableForFemale())

// Is most auspicious
fmt.Printf("Most auspicious: %v\n", dayan.IsBest())
```

## Application Scenarios

### Name Stroke Fortune Analysis

The most common application of Dayan Numbers is name stroke analysis. Take the total stroke count of a name modulo 81 (treating 0 as 81), then query the corresponding fortune:

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func analyzeStroke(strokeCount int) {
    // Modulo 81
    n := strokeCount % 81
    if n == 0 {
        n = 81
    }

    dayan, err := yi.GetDayan(n)
    if err != nil {
        fmt.Printf("Invalid stroke count: %d\n", n)
        return
    }

    fmt.Printf("Stroke count %d (original %d):\n", n, strokeCount)
    fmt.Printf("  Fortune: %s\n", dayan.JiXiong)
    fmt.Printf("  Imagery: %s\n", dayan.YiXiang)
    fmt.Printf("  Foundation: %s\n", dayan.Foundation)
    fmt.Printf("  Family: %s\n", dayan.Family)
    fmt.Printf("  Health: %s\n", dayan.Health)
    fmt.Printf("  Meaning: %s\n", dayan.Meaning)
}

func main() {
    analyzeStroke(21)
    analyzeStroke(34)
    analyzeStroke(81)
}
```

### Combined with Five Elements Analysis

Stroke counts can also be combined with the Five Elements system for richer analysis:

```go
strokeCount := 21
dayan, _ := yi.GetDayan(strokeCount % 81)

// Five Elements corresponding to stroke count
wuxing := yi.GetWuXingByNumber(strokeCount)
fmt.Printf("Element: %s\n", wuxing)  // e.g., "阳木" (Yang Wood)

// Fortune combined with Five Elements
fmt.Printf("Stroke %d: %s, Element %s\n", strokeCount, dayan.JiXiong, wuxing)
```

## Data Notes

> ⚠️ **Known Data Gaps**:
>
> - `Gua.JiXiong`: Some hexagram fortune fields are empty (incomplete source data, not a bug)
> - `Yao.NvMing`: Most line female fortune fields are empty. **Empty means gender-neutral**, not missing data
> - Only Qian and Kun hexagrams have `Yong` (Yong Jiu / Yong Liu) and `YongJiXiong` fields
>
> These are source data limitations and do not affect core functionality. `IsJi()` and `IsXiong()` methods return `false` when JiXiong is empty.
