# Six Relations & Shi-Ying / 六亲与世应

Six Relations (LiuQin) and Shi-Ying are two core systems in six-line divination. Six Relations determine line attributes based on Five Elements generating/overcoming relationships, while Shi-Ying determines the core positioning points in the hexagram.

## Six Relations Definition

Six Relations are five attributes determined by the generating/overcoming relationship between the palace's Five Elements and the line's Five Elements:

| Six Relations | Constant | Relationship | Governs |
|---------------|----------|--------------|---------|
| Parents (父母) | `LQFuMu` | Generates me | Protection, documents, contracts |
| Siblings (兄弟) | `LQXiongDi` | Same as me | Competition, wealth loss, peers |
| Wealth (妻财) | `LQQiCai` | I overcome | Wealth, spouse, resources |
| Offspring (子孙) | `LQZiSun` | I generate | Children, fortune, medicine |
| Official/Ghost (官鬼) | `LQGuanGui` | Overcomes me | Lawsuits, disasters, husband (female) |

> "Me" refers to the palace's Five Elements. All six relations within the same palace share the same attribution.

### Five Elements → Six Relations Mapping Logic

```
Generates me → Parents    (Water generates Wood; if palace is Wood, Water line = Parents)
Same as me → Siblings     (Wood same as Wood; if palace is Wood, Wood line = Siblings)
I overcome → Wealth       (Wood overcomes Earth; if palace is Wood, Earth line = Wealth)
I generate → Offspring    (Wood generates Fire; if palace is Wood, Fire line = Offspring)
Overcomes me → Official   (Metal overcomes Wood; if palace is Wood, Metal line = Official)
```

### Code Example

```go
// Get Six Relations
lq := yi.GetLiuQin(yi.Wood, yi.Water)
fmt.Printf("Wood palace Water line: %s\n", lq)  // Parents (Water generates Wood)

lq2 := yi.GetLiuQin(yi.Wood, yi.Fire)
fmt.Printf("Wood palace Fire line: %s\n", lq2)  // Offspring (Wood generates Fire)

lq3 := yi.GetLiuQin(yi.Wood, yi.Metal)
fmt.Printf("Wood palace Metal line: %s\n", lq3)  // Official (Metal overcomes Wood)

// Get Six Relations description
fmt.Printf("Description: %s\n", lq.Description())  // "Generates me, governs protection and documents"
```

## Eight Palaces Five Elements

Each palace corresponds to a Five Elements attribute, which is the basis for Six Relations mapping:

| Palace | Element | Description |
|--------|---------|-------------|
| Qian Palace | Metal | Qian is Heaven, belongs to Metal |
| Dui Palace | Metal | Dui is Lake, belongs to Metal |
| Li Palace | Fire | Li is Fire, belongs to Fire |
| Zhen Palace | Wood | Zhen is Thunder, belongs to Wood |
| Xun Palace | Wood | Xun is Wind, belongs to Wood |
| Kan Palace | Water | Kan is Water, belongs to Water |
| Gen Palace | Earth | Gen is Mountain, belongs to Earth |
| Kun Palace | Earth | Kun is Earth, belongs to Earth |

```go
// Query palace Five Elements
wx := yi.GetGuaGongWuXing(yi.Qian)
fmt.Printf("Qian palace element: %s\n", wx)  // Metal

wx2 := yi.GetGuaGongWuXing(yi.Zhen)
fmt.Printf("Zhen palace element: %s\n", wx2) // Wood
```

### Get Complete Six Relations Mapping

```go
// Get all element → Six Relations mapping for a palace
mapping := yi.GetLiuQinForGua(yi.Wood)
for wx, lq := range mapping {
    fmt.Printf("%s → %s\n", wx, lq)
}
// Output:
// Wood → Siblings
// Fire → Offspring
// Earth → Wealth
// Metal → Official
// Water → Parents
```

---

## Shi-Ying System

Shi-Ying is used to determine the attribution and relationships of lines in the hexagram:

- **Shi Line (世爻)**: Represents the querent, the core positioning point in the hexagram
- **Ying Line (应爻)**: Represents the person/event/thing being queried, corresponding to the Shi line

### Palace Positions and Shi-Ying Locations

Each palace has eight hexagrams, arranged in the following order, with fixed Shi-Ying positions:

| Palace Position | Constant | Shi Line | Ying Line |
|-----------------|----------|----------|-----------|
| Main Palace | `BenGong` | 6th line | 3rd line |
| First Shi | `YiShi` | 1st line | 4th line |
| Second Shi | `ErShi` | 2nd line | 5th line |
| Third Shi | `SanShi` | 3rd line | 6th line |
| Fourth Shi | `SiShi` | 4th line | 1st line |
| Fifth Shi | `WuShi` | 5th line | 2nd line |
| Wandering Soul | `YouHun` | 4th line | 1st line |
| Returning Soul | `GuiHun` | 3rd line | 6th line |

> Pattern: Shi and Ying are always separated by two lines. Main palace: Shi at 6, Ying at 3. First Shi moves down line by line. Wandering Soul returns to 4th, Returning Soul to 3rd.

### Code Example

```go
gua, _ := yi.GetGuaByIndex("乾乾")

// Get Shi-Ying info
sy := gua.GetShiYing()
fmt.Printf("Shi line: %d (%s)\n", sy.ShiPos, sy.Position) // 6th line (Main Palace)
fmt.Printf("Ying line: %d\n", sy.YingPos)                   // 3rd line

// Get palace position
pos := gua.GetGuaPosition()
fmt.Printf("Palace position: %s\n", pos)  // Main Palace

// Get palace归属
gong := gua.GetGuaGong()
fmt.Printf("Palace: %s\n", yi.GetBaguaName(gong))  // Qian
```

### Using in Divination Results

```go
zy, _ := yi.DivineByCoins(42)
ben := zy.GetGua(yi.Ben)

// Complete Six Relations Shi-Ying analysis
sy := ben.GetShiYing()
gongWX := yi.GetGuaGongWuXing(ben.GetGuaGong())
liuQinMap := yi.GetLiuQinForGua(gongWX)

fmt.Printf("Palace: %s Palace (Element %s)\n", yi.GetBaguaName(ben.GetGuaGong()), gongWX)
fmt.Printf("Position: %s\n", sy.Position)
fmt.Printf("Shi line: %d, Ying line: %d\n", sy.ShiPos+1, sy.YingPos+1)

// Output Six Relations for each line
for i := yi.YaoPosition(0); i < yi.YaoCount; i++ {
    yao := ben.GetYao(i)
    if yao != nil {
        // In practice, line's element is determined by Na Jia method
        fmt.Printf("Line %d: %s\n", i+1, yao.Ci)
    }
}
```

---

## Sexagenary Cycle (JiaZi) and Na Jia Basics

The Na Jia method combines Heavenly Stems and Earthly Branches with the six lines, which is the prerequisite for determining line branch Five Elements.

### Heavenly Stems and Earthly Branches

```go
// Heavenly Stems
fmt.Println(yi.GetTianGanName(yi.TGJia))   // "甲" (Jia)
fmt.Println(yi.GetTianGanName(yi.TGGui))   // "癸" (Gui)

// Earthly Branches
fmt.Println(yi.GetDiZhiName(yi.DZZi))      // "子" (Zi)
fmt.Println(yi.GetDiZhiName(yi.DZHai))     // "亥" (Hai)
```

### Heavenly Stems and Earthly Branches Five Elements

```go
// Heavenly Stems Five Elements
fmt.Printf("Jia: %s\n", yi.GetTianGanWuXing(yi.TGJia))   // Wood
fmt.Printf("Bing: %s\n", yi.GetTianGanWuXing(yi.TGBing))  // Fire
fmt.Printf("Geng: %s\n", yi.GetTianGanWuXing(yi.TGGeng))  // Metal

// Earthly Branches Five Elements
fmt.Printf("Yin: %s\n", yi.GetDiZhiWuXing(yi.DZYin))   // Wood
fmt.Printf("Wu: %s\n", yi.GetDiZhiWuXing(yi.DZWu))    // Fire
fmt.Printf("Chen: %s\n", yi.GetDiZhiWuXing(yi.DZChen))  // Earth
```

### Sexagenary Cycle Query

```go
// Query by sequence number (1-60)
jiazi, err := yi.GetJiaZi(1)
fmt.Printf("1st JiaZi: %s, Element: %s\n", jiazi.Name, jiazi.WuXing)
// JiaZi, Sea Metal

// Query by Stem-Branch combination
jiazi2, err := yi.GetJiaZiByGanZhi(yi.TGJia, yi.DZZi)
fmt.Printf("JiaZi: %s\n", jiazi2.Name)

// Yin-Yang attribute
fmt.Printf("Jia Yin-Yang: %s\n", yi.GetTianGanYinYang(yi.TGJia))  // Yang
fmt.Printf("Chou Yin-Yang: %s\n", yi.GetDiZhiYinYang(yi.DZChou))   // Yin
```

### JiaZiInfo Struct

```go
type JiaZiInfo struct {
    Index   int    // 1-60, position in Sexagenary Cycle
    Name    string // Full name (e.g., "甲子", "乙丑")
    TianGan string // Heavenly Stem name
    DiZhi   string // Earthly Branch name
    WuXing  string // Stem-Branch Five Elements (e.g., "海中金" Sea Metal)
    YinYang string // Heavenly Stem Yin-Yang attribute
}
```

> **Note**: `GetJiaZiByGanZhi` requires Stem and Branch to have matching parity (JiaZi is valid, JiaChou is not), otherwise returns `ErrInvalidJiaZiCombo`.

---

## Complete Analysis Example

Combining Six Relations, Shi-Ying, and JiaZi for a complete analysis:

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 1. Divination
    zy, _ := yi.DivineByCoins(42)
    ben := zy.GetGua(yi.Ben)

    // 2. Basic info
    fmt.Printf("Primary: %s %s\n", ben.GuaSymbol, ben.Ming)

    // 3. Palace and position
    gong := ben.GetGuaGong()
    gongName := yi.GetBaguaName(gong)
    gongWX := yi.GetGuaGongWuXing(gong)
    sy := ben.GetShiYing()

    fmt.Printf("Palace: %s Palace (Element %s)\n", gongName, gongWX)
    fmt.Printf("Position: %s\n", sy.Position)
    fmt.Printf("Shi line: %d  Ying line: %d\n", sy.ShiPos+1, sy.YingPos+1)

    // 4. Six Relations mapping
    liuQinMap := yi.GetLiuQinForGua(gongWX)
    fmt.Printf("\nSix Relations Mapping:\n")
    for wx, lq := range liuQinMap {
        fmt.Printf("  %s → %s (%s)\n", wx, lq, lq.Description())
    }

    // 5. Fortune judgment
    isJi := zy.IsJi(yi.Male)
    fmt.Printf("\nFortune: %v\n", isJi)
}
```
