# 大衍之数

大衍之数是姓名学中笔画吉凶分析的核心系统，将 1-81 的笔画数对应到不同的吉凶属性和详细解读。

## 81 数理概念

《周易·系辞》云："大衍之数五十，其用四十有九。"在本项目中，大衍之数被引申为 1-81 的笔画数吉凶体系，用于姓名笔画分析。

每个数理包含以下维度：

| 维度 | 字段 | 说明 |
|------|------|------|
| 笔画数 | `Number` | 1-81 |
| 吉凶 | `JiXiong` | 吉/凶/半吉 |
| 女命 | `NvMing` | 女性专属吉凶（为空表示男女通用） |
| 最吉 | `IsMax` | 是否为最吉祥数 |
| 对应卦 | `Gua` | 对应的卦象名称 |
| 天九 | `TianJiu` | 太极数 |
| 意象 | `YiXiang` | 意象描述 |
| 基业 | `Foundation` | 事业基础 |
| 家庭 | `Family` | 家庭状况 |
| 健康 | `Health` | 健康提示 |
| 含义 | `Meaning` | 详细含义解读 |

## Dayan 结构体

```go
type Dayan struct {
    Number     int    // 笔画数（1-81）
    JiXiong    string // 吉凶（吉/凶/半吉）
    NvMing     string // 女命专属吉凶
    IsMax      bool   // 是否最吉
    Gua        string // 对应卦象
    TianJiu    string // 天九
    YiXiang    string // 意象
    Foundation string // 基业
    Family     string // 家庭
    Health     string // 健康
    Meaning    string // 含义
}
```

## API

### 查询数理

```go
// 按笔画数查询（1-81）
dayan, err := yi.GetDayan(21)
if err != nil {
    // err == ErrInvalidDayanNumber（不在1-81范围）
    panic(err)
}
fmt.Printf("笔画: %d\n", dayan.Number)
fmt.Printf("吉凶: %s\n", dayan.JiXiong)
fmt.Printf("天九: %s\n", dayan.TianJiu)
fmt.Printf("意象: %s\n", dayan.YiXiang)
fmt.Printf("基业: %s\n", dayan.Foundation)
fmt.Printf("家庭: %s\n", dayan.Family)
fmt.Printf("健康: %s\n", dayan.Health)
fmt.Printf("含义: %s\n", dayan.Meaning)

// 忽略错误的查询（越界返回零值）
dayan2 := yi.MustGetDayan(100) // 返回 Dayan{}
```

### 吉凶判断

```go
dayan, _ := yi.GetDayan(21)

// 是否吉（吉或半吉）
fmt.Printf("是否吉: %v\n", dayan.IsJi())

// 是否凶
fmt.Printf("是否凶: %v\n", dayan.IsXiong())

// 是否适合女性
fmt.Printf("适合女性: %v\n", dayan.IsSuitableForFemale())

// 是否最吉
fmt.Printf("最吉: %v\n", dayan.IsBest())
```

## 应用场景

### 姓名笔画吉凶分析

大衍之数最常见的应用是姓名笔画分析。将姓名的总笔画数对 81 取余（0 视为 81），查询对应的吉凶：

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func analyzeStroke(strokeCount int) {
    // 笔画数对81取余
    n := strokeCount % 81
    if n == 0 {
        n = 81
    }

    dayan, err := yi.GetDayan(n)
    if err != nil {
        fmt.Printf("无效笔画数: %d\n", n)
        return
    }

    fmt.Printf("笔画数 %d（原%d）:\n", n, strokeCount)
    fmt.Printf("  吉凶: %s\n", dayan.JiXiong)
    fmt.Printf("  意象: %s\n", dayan.YiXiang)
    fmt.Printf("  基业: %s\n", dayan.Foundation)
    fmt.Printf("  家庭: %s\n", dayan.Family)
    fmt.Printf("  健康: %s\n", dayan.Health)
    fmt.Printf("  含义: %s\n", dayan.Meaning)
}

func main() {
    analyzeStroke(21)
    analyzeStroke(34)
    analyzeStroke(81)
}
```

### 结合五行分析

笔画数还可以结合五行系统，获取更丰富的分析：

```go
strokeCount := 21
dayan, _ := yi.GetDayan(strokeCount % 81)

// 笔画数对应的五行
wuxing := yi.GetWuXingByNumber(strokeCount)
fmt.Printf("五行: %s\n", wuxing)  // 如 "阳木"

// 吉凶结合五行
fmt.Printf("笔画%d: %s, 五行%s\n", strokeCount, dayan.JiXiong, wuxing)
```

## 数据说明

> ⚠️ **已知数据缺失**：
>
> - `Gua.JiXiong`：部分卦的吉凶字段为空（源数据不全，非 bug）
> - `Yao.NvMing`：大部分爻的女命字段为空。**空值表示男女通用**，并非数据缺失
> - 只有乾、坤两卦有 `Yong`（用九/用六）和 `YongJiXiong` 字段

这些是源数据的限制，不影响核心功能的使用。`IsJi()` 和 `IsXiong()` 方法会在 JiXiong 为空时返回 `false`。
