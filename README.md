# yi - Go 语言周易占卜库

[![Go Reference](https://pkg.go.dev/badge/github.com/godcong/yi.svg)](https://pkg.go.dev/github.com/godcong/yi)

周易六十四卦计算与查询库，提供完整的起卦、卦象变换、五行生克、六十甲子、六亲世应、81 数理、解卦等功能。

## 功能特性

- **四种起卦法**：铜钱法、蓍草法（大衍法）、梅花易数、时间起卦
- **卦象变换**：本卦、变卦、互卦、错卦、综卦五种关系
- **解卦系统**：综合本卦/变卦/互卦/动爻/彖象辞输出结构化解读
- **五行生克**：五行相生、相克、被生、被克关系计算
- **六十甲子**：天干地支、纳音五行、阴阳属性
- **六亲世应**：卦宫五行→六亲映射，世应位置推算
- **81 数理**：大衍之数吉凶分析，含天九/意象/基业/家庭/健康
- **彖象文言**：64 卦彖辞、大象辞，乾坤文言

## 安装

```bash
go get github.com/godcong/yi
```

## 快速开始

### 铜钱法起卦

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 方式一：随机种子
    zy, coins := yi.DivineByCoins(42)
    fmt.Printf("六次投掷: %v\n", coins)
    fmt.Printf("本卦: %s\n", zy.GetGua(yi.Ben).Ming)
    fmt.Printf("变卦: %s\n", zy.GetGua(yi.Bian).Ming)

    // 方式二：手动输入6个爻值（6=老阴, 7=少阴, 8=少阳, 9=老阳）
    zy2, err := yi.DivineByCoinValues([6]int{9, 8, 7, 6, 9, 8})
    if err != nil {
        panic(err)
    }
    fmt.Printf("手动起卦: %s\n", zy2.GetGua(yi.Ben).Ming)

    // 方式三：当前时间随机
    zy3, coins3 := yi.DivineByCoinsRand()
    _ = zy3 // use result
    _ = coins3
}
```

### 蓍草法（大衍法）起卦

```go
zy, results := yi.DivineByDayan(999)
fmt.Printf("本卦: %s\n", zy.GetGua(yi.Ben).Ming)
for i, r := range results {
    fmt.Printf("第%d爻: 值=%d, 变=%v\n", i+1, r.YaoValue, r.IsChanging)
}

// 手动输入爻值
zy2, err := yi.DivineByDayanValues([6]int{7, 7, 7, 7, 7, 7})
if err != nil {
    panic(err)
}
```

### 梅花易数起卦

```go
// 报数起卦：上卦=3, 下卦=5, 动爻=2
zy := yi.DivineByMeihua(3, 5, 2)
fmt.Printf("梅花起卦: %s\n", zy.GetGua(yi.Ben).Ming)

// 时间梅花起卦
zy2, year, month, day, hour := yi.DivineByMeihuaTime(time.Now())
fmt.Printf("时间: %d年%d月%d日%d时\n", year, month, day, hour)
fmt.Printf("梅花时间起卦: %s\n", zy2.GetGua(yi.Ben).Ming)
```

### 时间起卦法

```go
// 公历时间起卦
zy := yi.DivineByTimeGua(yi.TimeGuaParams{
    Year:  2024,
    Month: 6,
    Day:   15,
    Hour:  10,
})
fmt.Printf("时间起卦: %s\n", zy.GetGua(yi.Ben).Ming)

// 农历时间起卦
zy2 := yi.DivineByLunarTime(2024, 5, 10, 4)
fmt.Printf("农历起卦: %s\n", zy2.GetGua(yi.Ben).Ming)

// 当前时间快速起卦
zy3 := yi.DivineByCurrentTime()
_ = zy3
```

### 解卦

```go
zy, _ := yi.DivineByCoins(42)
result := yi.JieGua(zy, yi.Male)
fmt.Println(yi.FormatJieGua(result))
```

### 81 数理查询

```go
dayan, err := yi.GetDayan(21)
if err != nil {
    panic(err)
}
fmt.Printf("笔画数: %d\n", dayan.Number)
fmt.Printf("吉凶: %s\n", dayan.JiXiong)
fmt.Printf("天九: %s\n", dayan.TianJiu)
fmt.Printf("意象: %s\n", dayan.YiXiang)
fmt.Printf("基业: %s\n", dayan.Foundation)
fmt.Printf("含义: %s\n", dayan.Meaning)
fmt.Printf("是否吉: %v\n", dayan.IsJi())
fmt.Printf("是否最吉: %v\n", dayan.IsBest())
```

### 五行生克

```go
// 根据笔画数获取五行
wx := yi.GetWuXingByNumber(7)
fmt.Printf("笔画7的五行: %s\n", wx) // 阴金

// 五行相生：木→火→土→金→水→木
fmt.Printf("木生: %s\n", yi.Wood.Sheng())   // 火
fmt.Printf("火生: %s\n", yi.Fire.Sheng())   // 土

// 五行相克：木→土→水→火→金→木
fmt.Printf("木克: %s\n", yi.Wood.Ke())      // 土
fmt.Printf("金克: %s\n", yi.Metal.Ke())     // 木

// 被生/被克
fmt.Printf("木被生: %s\n", yi.Wood.BeiSheng()) // 水
fmt.Printf("木被克: %s\n", yi.Wood.BeiKe())    // 金
```

### 卦象查询

```go
// 按卦序查询（1-64）
gua, err := yi.GetGuaByXu(1)
if err != nil {
    panic(err)
}
fmt.Printf("第1卦: %s %s\n", gua.GuaSymbol, gua.Ming)

// 按索引查询
gua2, err := yi.GetGuaByIndex("乾乾")
fmt.Printf("乾为天: %s\n", gua2.TuanText)

// 获取爻辞
yao := gua2.GetYao(yi.Chu) // 初爻
fmt.Printf("初爻爻辞: %s\n", yao.Ci)
```

## API 总览

### 起卦函数

| 函数 | 说明 |
|------|------|
| `DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi` | 按数起卦（核心函数） |
| `DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult)` | 铜钱法起卦 |
| `DivineByCoinValues(values [6]int) (*ZhouYi, error)` | 铜钱法手动输入 |
| `DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult)` | 蓍草法起卦 |
| `DivineByDayanValues(values [6]int) (*ZhouYi, error)` | 蓍草法手动输入 |
| `DivineByMeihua(shang, xia, bian int) *ZhouYi` | 梅花易数起卦 |
| `DivineByMeihuaTime(t time.Time) (*ZhouYi, int, int, int, int)` | 梅花时间起卦 |
| `DivineByTimeGua(params TimeGuaParams) *ZhouYi` | 公历时间起卦 |
| `DivineByLunarTime(y, m, d, h int) *ZhouYi` | 农历时间起卦 |
| `DivineByCurrentTime() *ZhouYi` | 当前时间快速起卦 |

### 解卦函数

| 函数 | 说明 |
|------|------|
| `JieGua(zy *ZhouYi, sex Sex) *JieGuaResult` | 解卦主函数 |
| `FormatJieGua(result *JieGuaResult) string` | 格式化解卦结果 |

### 卦象查询

| 函数/方法 | 说明 |
|------|------|
| `GetGuaByIndex(index string) (*Gua, error)` | 按索引查卦（如 "乾乾"） |
| `GetGuaByXu(xu int) (*Gua, error)` | 按卦序查卦（1-64） |
| `GetBaguaName(bg Bagua) string` | 获取八卦名称 |
| `GetBaguaSymbol(bg Bagua) string` | 获取八卦符号 |
| `ZhouYi.GetGua(guaType int) *Gua` | 获取本/变/互/错/综卦 |
| `ZhouYi.IsJi(sex Sex) bool` | 判断吉凶 |
| `Gua.GetYao(pos YaoPosition) *Yao` | 获取指定爻 |
| `Gua.GetShiYing() *ShiYingInfo` | 获取世应信息 |
| `Gua.GetGuaGong() Bagua` | 获取归属卦宫 |
| `Gua.GetGuaPosition() GuaPosition` | 获取宫位 |

### 五行与六亲

| 函数/方法 | 说明 |
|------|------|
| `GetWuXingByNumber(n int) string` | 笔画数→五行 |
| `WuXing.Sheng() WuXing` | 我生 |
| `WuXing.Ke() WuXing` | 我克 |
| `WuXing.BeiSheng() WuXing` | 生我 |
| `WuXing.BeiKe() WuXing` | 克我 |
| `GetLiuQin(guaGongWX, yaoWX WuXing) LiuQin` | 计算六亲 |
| `GetGuaGongWuXing(bagua Bagua) WuXing` | 卦宫→五行 |
| `GetLiuQinForGua(guaGongWX WuXing) map[WuXing]LiuQin` | 卦宫六亲映射表 |

### 甲子干支

| 函数 | 说明 |
|------|------|
| `GetJiaZi(index int) (*JiaZiInfo, error)` | 按序号查甲子（1-60） |
| `GetJiaZiByGanZhi(gan TianGan, zhi DiZhi) (*JiaZiInfo, error)` | 按干支查甲子 |
| `GetTianGanName(gan TianGan) string` | 天干名称 |
| `GetDiZhiName(zhi DiZhi) string` | 地支名称 |
| `GetTianGanWuXing(gan TianGan) WuXing` | 天干五行 |
| `GetDiZhiWuXing(zhi DiZhi) WuXing` | 地支五行 |

### 81 数理

| 函数/方法 | 说明 |
|------|------|
| `GetDayan(n int) (*Dayan, error)` | 查询数理（1-81） |
| `MustGetDayan(n int) Dayan` | 查询数理（忽略错误） |
| `Dayan.IsJi() bool` | 是否吉 |
| `Dayan.IsXiong() bool` | 是否凶 |
| `Dayan.IsBest() bool` | 是否最吉 |
| `Dayan.IsSuitableForFemale() bool` | 是否适合女性 |

### 彖象文言

| 函数/方法 | 说明 |
|------|------|
| `Gua.GetTuan() string` | 获取彖辞 |
| `GetAllTuan() []struct{Index, Text string}` | 全部彖辞 |
| `Gua.GetXiang() string` | 获取象辞（大象） |
| `GetAllXiang() []struct{Index, Text string}` | 全部象辞 |
| `GetWenYan(index string) string` | 获取文言（仅乾坤） |
| `HasWenYan(index string) bool` | 是否有文言 |

## 核心类型

| 类型 | 说明 |
|------|------|
| `Gua` (别名 `Hexagram`) | 卦象：含卦名/卦义/彖辞/象辞/六爻 |
| `Yao` (别名 `Line`) | 爻：含爻辞/吉凶/女命 |
| `ZhouYi` (别名 `IChing`) | 周易：含五种卦象 + 动爻 |
| `Bagua` (别名 `Trigram`) | 八卦：0=乾…7=坤 |
| `WuXing` | 五行：木/火/土/金/水 |
| `YinYang` | 阴阳 |
| `Sex` | 性别：Male/Female |
| `LiuQin` | 六亲：父母/兄弟/妻财/子孙/官鬼 |
| `GuaPosition` | 宫位：本宫/一世/…/游魂/归魂 |
| `Dayan` | 大衍数理：1-81 的吉凶详情 |
| `TianGan` | 天干：甲乙丙丁… |
| `DiZhi` | 地支：子丑寅卯… |
| `JiaZiInfo` | 甲子信息 |
| `JieGuaResult` | 解卦结果 |
| `GuaInfo` | 卦象解读信息 |

## 卦象类型常量

| 常量 | 值 | 说明 |
|------|------|------|
| `Ben` | 0 | 本卦（原始卦象） |
| `Bian` | 1 | 变卦（动爻变化后） |
| `Hu` | 2 | 互卦（2-3-4爻为下卦，3-4-5爻为上卦） |
| `Cuo` | 3 | 错卦（阴阳全反） |
| `Zong` | 4 | 综卦（上下颠倒） |

## 八卦常量

| 常量 | 值 | 二进制 | 名称 |
|------|------|------|------|
| `Qian` | 0 | 0b000 | 乾 ☰ |
| `Dui` | 1 | 0b001 | 兑 ☱ |
| `Li` | 2 | 0b010 | 离 ☲ |
| `Zhen` | 3 | 0b011 | 震 ☳ |
| `Xun` | 4 | 0b100 | 巽 ☴ |
| `Kan` | 5 | 0b101 | 坎 ☵ |
| `Gen` | 6 | 0b110 | 艮 ☶ |
| `Kun` | 7 | 0b111 | 坤 ☷ |

> **约定**：项目记阳为 0、阴为 1（bit 表示），与直觉相反。

## 数据生成

数据存储在 `data/` 目录下的 JSON 文件中：

- `data/gua.json` — 64 卦数据
- `data/tuan.json` — 彗辞数据
- `data/xiang.json` — 象辞数据
- `data/wenyan.json` — 文言数据
- `data/jiazi.json` — 六十甲子数据
- `data/guagong.json` — 卦宫数据

运行以下命令重新生成 `data_generated.go`：

```bash
go generate ./...
```

## 文档

- [数理简介](docs/数理简介.md) — 阴阳→八卦→六十四卦→八十一象的数理链路
- [起卦方法](docs/起卦方法.md) — 四种起卦法详解
- [八卦与六十四卦](docs/八卦与六十四卦.md) — 卦象体系与变换
- [大衍之数](docs/大衍之数.md) — 81 数理吉凶
- [六亲与世应](docs/六亲与世应.md) — 六亲系统与世应推算

## 许可

MIT License
