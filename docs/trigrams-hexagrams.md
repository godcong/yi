# Eight Trigrams & Sixty-Four Hexagrams / 八卦与六十四卦

## Eight Trigrams Basics

The Eight Trigrams (Bagua) are the basic symbol system of the I Ching. Each trigram consists of three lines (3-bit). This project uses the Earlier Heaven (Fuxi) order, encoded in binary:

| Trigram | Constant | Decimal | Binary | Symbol | Image | Element |
|---------|----------|---------|--------|--------|-------|---------|
| Qian (乾) | `Qian` | 0 | 0b000 | ☰ | Heaven | Metal |
| Dui (兑) | `Dui` | 1 | 0b001 | ☱ | Lake | Metal |
| Li (离) | `Li` | 2 | 0b010 | ☲ | Fire | Fire |
| Zhen (震) | `Zhen` | 3 | 0b011 | ☳ | Thunder | Wood |
| Xun (巽) | `Xun` | 4 | 0b100 | ☴ | Wind | Wood |
| Kan (坎) | `Kan` | 5 | 0b101 | ☵ | Water | Water |
| Gen (艮) | `Gen` | 6 | 0b110 | ☶ | Mountain | Earth |
| Kun (坤) | `Kun` | 7 | 0b111 | ☷ | Earth | Earth |

> **Convention**: bit=0 is Yang line (solid), bit=1 is Yin line (broken). From LSB to MSB corresponds to 1st through 6th line.

### Earlier Heaven vs. Later Heaven Order

- **Earlier Heaven (Fuxi Bagua)**: Qian 1, Dui 2, Li 3, Zhen 4, Xun 5, Kan 6, Gen 7, Kun 8. Represents natural opposition.
- **Later Heaven (King Wen Bagua)**: Li 9 (South), Kan 1 (North), Zhen 3 (East), Dui 7 (West), etc. Adds center palace, forming Nine Images.

This project's trigram constants `Qian=0…Kun=7` use the Earlier Heaven order.

### Querying Trigram Information

```go
// Get trigram name
fmt.Println(yi.GetBaguaName(yi.Qian))  // "乾"
fmt.Println(yi.GetBaguaName(yi.Kan))   // "坎"

// Get trigram symbol
fmt.Println(yi.GetBaguaSymbol(yi.Qian))  // "☰"
fmt.Println(yi.GetBaguaSymbol(yi.Kan))   // "☵"
```

---

## Sixty-Four Hexagrams

The 64 hexagrams are formed by combining an upper trigram (outer) and a lower trigram (inner). 8 × 8 = 64 combinations, each corresponding to one hexagram in the I Ching.

The upper trigram is lines 4, 5, 6 (4th, 5th, 6th lines), and the lower trigram is lines 1, 2, 3 (1st, 2nd, 3rd lines).

### Gua Struct

```go
type Gua struct {
    Xu          int     // Hexagram sequence number (1-64)
    Index       string  // Index key (e.g., "乾乾", "坎震")
    ShangMing   string  // Upper trigram full name (e.g., "乾为天")
    ShangNum    int     // Upper trigram number (0-7)
    XiaMing     string  // Lower trigram full name (e.g., "坎为水")
    XiaNum      int     // Lower trigram number (0-7)
    JiXiong     string  // Fortune (吉/凶/半吉 = Auspicious/Inauspicious/Semi-Auspicious)
    GuaName     string  // Hexagram single character name (e.g., "乾", "屯")
    Ming        string  // Hexagram full name (e.g., "乾为天", "水雷屯")
    GuaYi       string  // Hexagram meaning (Shao Yong I Ching)
    GuaSymbol   string  // Unicode hexagram symbol (e.g., ䷀)
    TuanText    string  // Tuan (Judgment) text
    XiangText   string  // Xiang (Image) text (e.g., "天行健，君子以自强不息")
    Yaos        [6]*Yao // Six lines: 1st, 2nd, 3rd, 4th, 5th, 6th
    Yong        string  // Yong Jiu / Yong Liu (only for Qian and Kun)
    YongJiXiong string  // Yong Jiu / Yong Liu fortune
}
```

### Yao Struct

```go
type Yao struct {
    Ci      string // Line text
    JiXiong string // Line fortune (吉/凶/平 = Auspicious/Inauspicious/Neutral)
    NvMing  string // Female-specific fortune; empty means gender-neutral
}
```

> **Data Note**: Some Gua.JiXiong fields are empty (incomplete source data), some Yao.NvMing are empty (empty means gender-neutral, not missing).

### Line Position Constants

| Constant | Value | Name |
|----------|-------|------|
| `Chu` | 0 | 1st Line (初爻) |
| `Er` | 1 | 2nd Line (二爻) |
| `San` | 2 | 3rd Line (三爻) |
| `Si` | 3 | 4th Line (四爻) |
| `Wu` | 4 | 5th Line (五爻) |
| `Shang` | 5 | 6th Line (上爻) |

### Querying Hexagrams

```go
// Query by sequence number (1-64)
gua, err := yi.GetGuaByXu(1)
fmt.Printf("Hexagram 1: %s %s\n", gua.GuaSymbol, gua.Ming)

// Query by index
gua2, err := yi.GetGuaByIndex("乾乾")
fmt.Printf("Qian: %s\n", gua2.TuanText)

// Get line text
yao := gua2.GetYao(yi.Chu)
fmt.Printf("1st line: %s (%s)\n", yao.Ci, yao.JiXiong)

// Female fortune check
if yao.HasNvMing() {
    fmt.Printf("Female: %s\n", yao.NvMing)
}
```

---

## Hexagram Transformations

After divination, the primary hexagram can derive four transformed hexagrams, forming a complete analysis system:

### Transformation Relationships

| Transformation | Constant | Algorithm | Meaning |
|----------------|----------|-----------|---------|
| Primary | `Ben` | Original divination result | Current state, starting point |
| Transformed | `Bian` | Invert changing line (Yang↔Yin) | Development trend |
| Nuclear | `Hu` | Lines 2-3-4 as lower, 3-4-5 as upper | Intermediate process, inner factors |
| Inverse | `Cuo` | All lines inverted (bitwise NOT) | Opposite perspective, reverse thinking |
| Reverse | `Zong` | Turn upside down (reverse bit order) | Shift viewpoint, different angle |

### Transformation Algorithm Details

**Transformed**: Find the changing line position, invert that line (0↔1), get new upper/lower trigrams.

```go
// bianYaoTransform: invert at specified position
func bianYaoTransform(gua, pos int) int {
    mask := 1 << (2 - uint(pos))
    if gua&mask == 0 {
        return gua | mask   // Yang → Yin
    }
    return gua ^ mask       // Yin → Yang
}
```

**Nuclear**: Take lines 2-4 of primary as lower trigram, lines 3-5 as upper trigram. Special case: if primary is pure Qian or pure Kun, nuclear uses transformed calculation.

**Inverse**: Bitwise NOT on upper and lower trigrams separately (3-bit NOT).

```go
shang = ^ben.ShangNum & 0x7  // Bitwise NOT, keep 3 bits
xia   = ^ben.XiaNum & 0x7
```

**Reverse**: Swap upper and lower trigrams, and reverse bit order of each trigram.

```go
// reverseBits: reverse 3-bit binary
func reverseBits(n int) int {
    return ((n & 0x4) >> 2) | (n & 0x2) | ((n & 0x1) << 2)
}
shang = reverseBits(ben.XiaNum)
xia   = reverseBits(ben.ShangNum)
```

### Code Example

```go
zy := yi.DivineByNumber(yi.Zhen, yi.Kan, 2)

ben  := zy.GetGua(yi.Ben)   // Primary: Lei Shui Jie (Deliverance)
bian := zy.GetGua(yi.Bian)  // Transformed: Lei Di Yu (Enthusiasm)
hu   := zy.GetGua(yi.Hu)    // Nuclear: Shui Huo Ji Ji (After Completion)
cuo  := zy.GetGua(yi.Cuo)   // Inverse: Feng Huo Jia Ren (The Family)
zong := zy.GetGua(yi.Zong)  // Reverse: Shui Shan Jian (Obstruction)

fmt.Printf("Primary: %s %s\n", ben.GuaSymbol, ben.Ming)
fmt.Printf("Transformed: %s %s\n", bian.GuaSymbol, bian.Ming)
fmt.Printf("Nuclear: %s %s\n", hu.GuaSymbol, hu.Ming)
fmt.Printf("Inverse: %s %s\n", cuo.GuaSymbol, cuo.Ming)
fmt.Printf("Reverse: %s %s\n", zong.GuaSymbol, zong.Ming)
```

---

## ZhouYi Struct

`ZhouYi` is the core result of divination, containing five hexagram types and changing line information:

```go
type ZhouYi struct {
    // Internal fields, accessed via methods
}

// Get hexagrams
zy.GetGua(yi.Ben)   // *Gua
zy.GetGua(yi.Bian)  // *Gua
zy.GetGua(yi.Hu)    // *Gua
zy.GetGua(yi.Cuo)   // *Gua
zy.GetGua(yi.Zong)  // *Gua

// Get changing lines
zy.GetBianYao()       // int (0-5)
zy.GetAllBianYao()    // []int (all changing lines)

// Fortune judgment
zy.IsJi(yi.Male)     // bool
zy.IsJi(yi.Female)   // bool
```

> English aliases: `ZhouYi` = `IChing`, `Gua` = `Hexagram`, `Yao` = `Line`, `Bagua` = `Trigram`.

---

## Tuan, Xiang, and Wenyan

Each hexagram has Tuan (Judgment) and Xiang (Image) texts. Qian and Kun additionally have Wenyan (Commentary):

```go
gua, _ := yi.GetGuaByIndex("乾乾")

// Tuan: explains the overall meaning of a hexagram
fmt.Printf("Tuan: %s\n", gua.GetTuan())

// Xiang (Great Image): life philosophy derived from hexagram imagery
fmt.Printf("Xiang: %s\n", gua.GetXiangText())

// Wenyan: only for Qian and Kun
if yi.HasWenYan("乾乾") {
    fmt.Printf("Wenyan: %s\n", yi.GetWenYan("乾乾"))
}

// Batch retrieval
allTuan := yi.GetAllTuan()   // All 64 hexagram Tuan texts
allXiang := yi.GetAllXiang() // All 64 hexagram Xiang texts
```
