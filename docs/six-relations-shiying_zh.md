# 六亲与世应

六亲和世应是六爻预测学中的两大核心系统。六亲根据五行生克关系确定爻的属性，世应则确定卦中的核心定位点。

## 六亲定义

六亲是根据卦宫五行与爻支五行的生克关系确定的五个属性：

| 六亲 | 常量 | 关系 | 主事 |
|------|------|------|------|
| 父母 | `LQFuMu` | 生我者 | 庇护、文书、契约 |
| 兄弟 | `LQXiongDi` | 同我者 | 竞争、劫财、同辈 |
| 妻财 | `LQQiCai` | 我克者 | 财富、妻室、物资 |
| 子孙 | `LQZiSun` | 我生者 | 子嗣、福气、医药 |
| 官鬼 | `LQGuanGui` | 克我者 | 官司、祸患、丈夫（女命） |

> "我"指卦宫五行。同一卦宫内的六亲归属相同。

### 五行生克→六亲映射逻辑

```
生我者 → 父母    （水生木，若卦宫为木，则水爻为父母）
同我者 → 兄弟    （木同木，若卦宫为木，则木爻为兄弟）
我克者 → 妻财    （木克土，若卦宫为木，则土爻为妻财）
我生者 → 子孙    （木生火，若卦宫为木，则火爻为子孙）
克我者 → 官鬼    （金克木，若卦宫为木，则金爻为官鬼）
```

### 代码示例

```go
// 获取六亲
lq := yi.GetLiuQin(yi.Wood, yi.Water)
fmt.Printf("木宫水爻: %s\n", lq)  // 父母（水生木，生我者）

lq2 := yi.GetLiuQin(yi.Wood, yi.Fire)
fmt.Printf("木宫火爻: %s\n", lq2)  // 子孙（木生火，我生者）

lq3 := yi.GetLiuQin(yi.Wood, yi.Metal)
fmt.Printf("木宫金爻: %s\n", lq3)  // 官鬼（金克木，克我者）

// 获取六亲描述
fmt.Printf("描述: %s\n", lq.Description())  // "生我者，主庇护文书"
```

## 八宫五行

每个卦宫对应一个五行属性，这是六亲映射的基础：

| 卦宫 | 五行 | 说明 |
|------|------|------|
| 乾宫 | 金 | 乾为天，属金 |
| 兑宫 | 金 | 兑为泽，属金 |
| 离宫 | 火 | 离为火，属火 |
| 震宫 | 木 | 震为雷，属木 |
| 巽宫 | 木 | 巽为风，属木 |
| 坎宫 | 水 | 坎为水，属水 |
| 艮宫 | 土 | 艮为山，属土 |
| 坤宫 | 土 | 坤为地，属土 |

```go
// 查询卦宫五行
wx := yi.GetGuaGongWuXing(yi.Qian)
fmt.Printf("乾宫五行: %s\n", wx)  // 金

wx2 := yi.GetGuaGongWuXing(yi.Zhen)
fmt.Printf("震宫五行: %s\n", wx2) // 木
```

### 获取完整六亲映射表

```go
// 一次性获取某卦宫的五行→六亲映射
mapping := yi.GetLiuQinForGua(yi.Wood)
for wx, lq := range mapping {
    fmt.Printf("%s → %s\n", wx, lq)
}
// 输出:
// 木 → 兄弟
// 火 → 子孙
// 土 → 妻财
// 金 → 官鬼
// 水 → 父母
```

---

## 世应系统

世应用于确定卦中各爻的归属和关系：

- **世爻（Shi）**：代表求测者本人，是卦中的核心定位点
- **应爻（Ying）**：代表所测之人、事、物，与世爻相对应

### 宫位与世应位置

每宫八卦，按以下顺序排列，每种宫位对应固定的世应位置：

| 宫位 | 常量 | 世爻位置 | 应爻位置 |
|------|------|----------|----------|
| 本宫 | `BenGong` | 上爻（第6爻） | 三爻（第3爻） |
| 一世 | `YiShi` | 初爻（第1爻） | 四爻（第4爻） |
| 二世 | `ErShi` | 二爻（第2爻） | 五爻（第5爻） |
| 三世 | `SanShi` | 三爻（第3爻） | 上爻（第6爻） |
| 四世 | `SiShi` | 四爻（第4爻） | 初爻（第1爻） |
| 五世 | `WuShi` | 五爻（第5爻） | 二爻（第2爻） |
| 游魂 | `YouHun` | 四爻（第4爻） | 初爻（第1爻） |
| 归魂 | `GuiHun` | 三爻（第3爻） | 上爻（第6爻） |

> 规律：世应相隔两爻。本宫世六应三，一世逐爻下移，至游魂回到四爻，归魂到三爻。

### 代码示例

```go
gua, _ := yi.GetGuaByIndex("乾乾")

// 获取世应信息
sy := gua.GetShiYing()
fmt.Printf("世爻: %d（%s）\n", sy.ShiPos, sy.Position) // 上爻（本宫）
fmt.Printf("应爻: %d\n", sy.YingPos)                   // 三爻

// 获取卦宫位置
pos := gua.GetGuaPosition()
fmt.Printf("宫位: %s\n", pos)  // 本宫

// 获取归属卦宫
gong := gua.GetGuaGong()
fmt.Printf("卦宫: %s\n", yi.GetBaguaName(gong))  // 乾
```

### 在起卦结果中使用

```go
zy, _ := yi.DivineByCoins(42)
ben := zy.GetGua(yi.Ben)

// 完整六亲世应分析
sy := ben.GetShiYing()
gongWX := yi.GetGuaGongWuXing(ben.GetGuaGong())
liuQinMap := yi.GetLiuQinForGua(gongWX)

fmt.Printf("卦宫: %s宫（%s）\n", yi.GetBaguaName(ben.GetGuaGong()), gongWX)
fmt.Printf("宫位: %s\n", sy.Position)
fmt.Printf("世爻: 第%d爻, 应爻: 第%d爻\n", sy.ShiPos+1, sy.YingPos+1)

// 输出每爻的六亲
for i := yi.YaoPosition(0); i < yi.YaoCount; i++ {
    yao := ben.GetYao(i)
    if yao != nil {
        // 实际使用时，爻的五行由纳甲法确定
        fmt.Printf("第%d爻: %s\n", i+1, yao.Ci)
    }
}
```

---

## 六十甲子与纳甲法基础

纳甲法将天干地支与六爻结合，是确定爻支五行的前提。

### 天干与地支

```go
// 天干
fmt.Println(yi.GetTianGanName(yi.TGJia))   // "甲"
fmt.Println(yi.GetTianGanName(yi.TGGui))   // "癸"

// 地支
fmt.Println(yi.GetDiZhiName(yi.DZZi))      // "子"
fmt.Println(yi.GetDiZhiName(yi.DZHai))     // "亥"
```

### 天干地支五行

```go
// 天干五行
fmt.Printf("甲: %s\n", yi.GetTianGanWuXing(yi.TGJia))   // 木
fmt.Printf("丙: %s\n", yi.GetTianGanWuXing(yi.TGBing))  // 火
fmt.Printf("庚: %s\n", yi.GetTianGanWuXing(yi.TGGeng))  // 金

// 地支五行
fmt.Printf("寅: %s\n", yi.GetDiZhiWuXing(yi.DZYin))   // 木
fmt.Printf("午: %s\n", yi.GetDiZhiWuXing(yi.DZWu))    // 火
fmt.Printf("辰: %s\n", yi.GetDiZhiWuXing(yi.DZChen))  // 土
```

### 六十甲子查询

```go
// 按序号查询（1-60）
jiazi, err := yi.GetJiaZi(1)
fmt.Printf("第1甲子: %s, 五行: %s\n", jiazi.Name, jiazi.WuXing)
// 甲子, 海中金

// 按干支组合查询
jiazi2, err := yi.GetJiaZiByGanZhi(yi.TGJia, yi.DZZi)
fmt.Printf("甲子: %s\n", jiazi2.Name)

// 阴阳属性
fmt.Printf("甲的阴阳: %s\n", yi.GetTianGanYinYang(yi.TGJia))  // 阳
fmt.Printf("丑的阴阳: %s\n", yi.GetDiZhiYinYang(yi.DZChou))   // 阴
```

### JiaZiInfo 结构体

```go
type JiaZiInfo struct {
    Index   int    // 1-60，在六十甲子中的位置
    Name    string // 全称（如 "甲子"、"乙丑"）
    TianGan string // 天干名称
    DiZhi   string // 地支名称
    WuXing  string // 干支五行属性（如 "海中金"）
    YinYang string // 天干阴阳属性
}
```

> **注意**：`GetJiaZiByGanZhi` 要求干支必须奇偶相同（甲子可，甲丑不可），否则返回 `ErrInvalidJiaZiCombo`。

---

## 完整分析示例

将六亲、世应、甲子结合进行完整分析：

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 1. 起卦
    zy, _ := yi.DivineByCoins(42)
    ben := zy.GetGua(yi.Ben)

    // 2. 基本信息
    fmt.Printf("本卦: %s %s\n", ben.GuaSymbol, ben.Ming)

    // 3. 卦宫与宫位
    gong := ben.GetGuaGong()
    gongName := yi.GetBaguaName(gong)
    gongWX := yi.GetGuaGongWuXing(gong)
    sy := ben.GetShiYing()

    fmt.Printf("卦宫: %s宫（五行%s）\n", gongName, gongWX)
    fmt.Printf("宫位: %s\n", sy.Position)
    fmt.Printf("世爻: 第%d爻  应爻: 第%d爻\n", sy.ShiPos+1, sy.YingPos+1)

    // 4. 六亲映射表
    liuQinMap := yi.GetLiuQinForGua(gongWX)
    fmt.Printf("\n六亲映射:\n")
    for wx, lq := range liuQinMap {
        fmt.Printf("  %s → %s（%s）\n", wx, lq, lq.Description())
    }

    // 5. 吉凶判断
    isJi := zy.IsJi(yi.Male)
    fmt.Printf("\n吉凶: %v\n", isJi)
}
```
