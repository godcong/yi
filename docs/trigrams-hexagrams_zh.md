# 八卦与六十四卦

## 八卦基础

八卦是周易的基本符号系统，每卦由三条爻线（3-bit）构成。本项目采用先天八卦序，以二进制编码：

| 卦 | 常量 | 十进制 | 二进制 | 符号 | 象 | 五行 |
|------|------|------|--------|------|------|------|
| 乾 | `Qian` | 0 | 0b000 | ☰ | 天 | 金 |
| 兑 | `Dui` | 1 | 0b001 | ☱ | 泽 | 金 |
| 离 | `Li` | 2 | 0b010 | ☲ | 火 | 火 |
| 震 | `Zhen` | 3 | 0b011 | ☳ | 雷 | 木 |
| 巽 | `Xun` | 4 | 0b100 | ☴ | 风 | 木 |
| 坎 | `Kan` | 5 | 0b101 | ☵ | 水 | 水 |
| 艮 | `Gen` | 6 | 0b110 | ☶ | 山 | 土 |
| 坤 | `Kun` | 7 | 0b111 | ☷ | 地 | 土 |

> **约定**：bit=0 为阳爻（实线），bit=1 为阴爻（虚线）。从低位到高位对应初爻到上爻。

### 先天八卦序与后天八卦序

- **先天八卦序**（伏羲八卦）：乾1、兑2、离3、震4、巽5、坎6、艮7、坤8。取自然对立之序。
- **后天八卦序**（文王八卦）：离9(南)、坎1(北)、震3(东)、兑7(西)等。加入中宫，成九象。

本项目八卦常量 `Qian=0…Kun=7` 采用先天序。

### 查询八卦信息

```go
// 获取八卦名称
fmt.Println(yi.GetBaguaName(yi.Qian))  // "乾"
fmt.Println(yi.GetBaguaName(yi.Kan))   // "坎"

// 获取八卦符号
fmt.Println(yi.GetBaguaSymbol(yi.Qian))  // "☰"
fmt.Println(yi.GetBaguaSymbol(yi.Kan))   // "☵"
```

---

## 六十四卦

六十四卦由上卦（外卦）和下卦（内卦）组合而成。8 × 8 = 64 种组合，每种对应《周易》中的一卦。

上卦为第 4、5、6 爻（四爻、五爻、上爻），下卦为第 1、2、3 爻（初爻、二爻、三爻）。

### Gua 结构体

```go
type Gua struct {
    Xu          int     // 卦序号（1-64）
    Index       string  // 索引键（如 "乾乾"、"坎震"）
    ShangMing   string  // 上卦全名（如 "乾为天"）
    ShangNum    int     // 上卦编号（0-7）
    XiaMing     string  // 下卦全名（如 "坎为水"）
    XiaNum      int     // 下卦编号（0-7）
    JiXiong     string  // 吉凶（吉/凶/半吉）
    GuaName     string  // 卦名单字（如 "乾"、"屯"）
    Ming        string  // 卦名全称（如 "乾为天"、"水雷屯"）
    GuaYi       string  // 卦义（邵雍易学）
    GuaSymbol   string  // Unicode卦象符号（如 ䷀）
    TuanText    string  // 彖辞
    XiangText   string  // 大象辞（如 "天行健，君子以自强不息"）
    Yaos        [6]*Yao // 六爻：初、二、三、四、五、上
    Yong        string  // 用九/用六（仅乾坤两卦）
    YongJiXiong string  // 用九/用六吉凶
}
```

### Yao 结构体

```go
type Yao struct {
    Ci      string // 爻辞
    JiXiong string // 爻吉凶（吉/凶/平）
    NvMing  string // 女命专属吉凶；为空表示男女通用
}
```

> **数据说明**：部分 Gua.JiXiong 字段为空（源数据不全），部分 Yao.NvMing 为空（空值表示男女通用，非缺失）。

### 爻位常量

| 常量 | 值 | 名称 |
|------|------|------|
| `Chu` | 0 | 初爻 |
| `Er` | 1 | 二爻 |
| `San` | 2 | 三爻 |
| `Si` | 3 | 四爻 |
| `Wu` | 4 | 五爻 |
| `Shang` | 5 | 上爻 |

### 查询卦象

```go
// 按卦序查询（1-64）
gua, err := yi.GetGuaByXu(1)
fmt.Printf("第1卦: %s %s\n", gua.GuaSymbol, gua.Ming)

// 按索引查询
gua2, err := yi.GetGuaByIndex("乾乾")
fmt.Printf("乾为天: %s\n", gua2.TuanText)

// 获取爻辞
yao := gua2.GetYao(yi.Chu)
fmt.Printf("初爻: %s (%s)\n", yao.Ci, yao.JiXiong)

// 女命判断
if yao.HasNvMing() {
    fmt.Printf("女命: %s\n", yao.NvMing)
}
```

---

## 卦象变换

起卦后，本卦可衍生出四种变换卦象，构成完整的分析体系：

### 变换关系

| 变换 | 常量 | 算法 | 含义 |
|------|------|------|------|
| 本卦 | `Ben` | 起卦原始结果 | 当前状态、事情的起点 |
| 变卦 | `Bian` | 动爻取反（阳↔阴） | 事情的发展趋势 |
| 互卦 | `Hu` | 2-3-4爻为下卦，3-4-5爻为上卦 | 中间过程、内在因素 |
| 错卦 | `Cuo` | 所有爻阴阳取反（bitwise NOT） | 对立面、反向思考 |
| 综卦 | `Zong` | 上下颠倒（reverse bit order） | 换位思考、不同视角 |

### 变换算法详解

**变卦**：找到动爻位置，将该爻取反（0↔1），得到新的上下卦。

```go
// bianYaoTransform: 对指定位置取反
func bianYaoTransform(gua, pos int) int {
    mask := 1 << (2 - uint(pos))
    if gua&mask == 0 {
        return gua | mask   // 阳→阴
    }
    return gua ^ mask       // 阴→阳
}
```

**互卦**：取本卦的第2-4爻为下卦，第3-5爻为上卦。特殊地，若本卦为纯乾或纯坤，互卦取变卦计算。

**错卦**：对上下卦分别按位取反（3-bit NOT）。

```go
shang = ^ben.ShangNum & 0x7  // 按位取反，保留3位
xia   = ^ben.XiaNum & 0x7
```

**综卦**：将上下卦互换，并对每个卦的 bit 顺序反转。

```go
// reverseBits: 反转3位二进制
func reverseBits(n int) int {
    return ((n & 0x4) >> 2) | (n & 0x2) | ((n & 0x1) << 2)
}
shang = reverseBits(ben.XiaNum)
xia   = reverseBits(ben.ShangNum)
```

### 代码示例

```go
zy := yi.DivineByNumber(yi.Zhen, yi.Kan, 2)

ben  := zy.GetGua(yi.Ben)   // 本卦：雷水解
bian := zy.GetGua(yi.Bian)  // 变卦：雷地豫
hu   := zy.GetGua(yi.Hu)    // 互卦：水火既济
cuo  := zy.GetGua(yi.Cuo)   // 错卦：风火家人
zong := zy.GetGua(yi.Zong)  // 综卦：水山蹇

fmt.Printf("本卦: %s %s\n", ben.GuaSymbol, ben.Ming)
fmt.Printf("变卦: %s %s\n", bian.GuaSymbol, bian.Ming)
fmt.Printf("互卦: %s %s\n", hu.GuaSymbol, hu.Ming)
fmt.Printf("错卦: %s %s\n", cuo.GuaSymbol, cuo.Ming)
fmt.Printf("综卦: %s %s\n", zong.GuaSymbol, zong.Ming)
```

---

## ZhouYi 结构体

`ZhouYi` 是起卦的核心结果，包含五种卦象和动爻信息：

```go
type ZhouYi struct {
    // 内部字段，通过方法访问
}

// 获取卦象
zy.GetGua(yi.Ben)   // *Gua
zy.GetGua(yi.Bian)  // *Gua
zy.GetGua(yi.Hu)    // *Gua
zy.GetGua(yi.Cuo)   // *Gua
zy.GetGua(yi.Zong)  // *Gua

// 获取动爻
zy.GetBianYao()       // int (0-5)
zy.GetAllBianYao()    // []int (所有动爻)

// 吉凶判断
zy.IsJi(yi.Male)     // bool
zy.IsJi(yi.Female)   // bool
```

> 英文别名：`ZhouYi` = `IChing`，`Gua` = `Hexagram`，`Yao` = `Line`，`Bagua` = `Trigram`。

---

## 彖辞、象辞与文言

每卦配有彖辞和象辞，乾坤两卦另有文言：

```go
gua, _ := yi.GetGuaByIndex("乾乾")

// 彖辞：论断一卦整体含义
fmt.Printf("彖辞: %s\n", gua.GetTuan())

// 象辞（大象）：从卦象引申出的人生哲理
fmt.Printf("象辞: %s\n", gua.GetXiangText())

// 文言：仅乾坤两卦有
if yi.HasWenYan("乾乾") {
    fmt.Printf("文言: %s\n", yi.GetWenYan("乾乾"))
}

// 批量获取
allTuan := yi.GetAllTuan()   // 全部64卦彖辞
allXiang := yi.GetAllXiang() // 全部64卦象辞
```
