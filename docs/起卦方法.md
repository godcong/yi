# 起卦方法

本库实现了四种经典起卦方法：铜钱法、蓍草法（大衍法）、梅花易数、时间起卦。每种方法都支持随机模拟和手动输入两种模式。

## 1. 铜钱法（三钱法）

### 原理

铜钱法是最常用的起卦方式。取 3 枚铜钱，抛掷 6 次，每次记录阴面（背）数量：

| 阴面数 | 爻值 | 名称 | 是否动爻 |
|--------|------|------|----------|
| 0（三阳） | 9 | 老阳 | ✅ 动爻 |
| 1 | 8 | 少阳 | ❌ |
| 2 | 7 | 少阴 | ❌ |
| 3（三阴） | 6 | 老阴 | ✅ 动爻 |

老阳(9)和老阴(6)为动爻，动爻取反即得变卦。

爻值到卦象的映射：阳爻(8/9)=0，阴爻(6/7)=1。六次投掷从初爻到上爻构成完整卦象。

### Go 代码示例

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 方式一：传入随机种子，自动模拟
    zy, coins := yi.DivineByCoins(42)
    for i, c := range coins {
        fmt.Printf("第%d次: %s(值=%d, 动爻=%v)\n",
            i+1, c, c.YaoValue(), c.IsChanging())
    }

    ben := zy.GetGua(yi.Ben)
    bian := zy.GetGua(yi.Bian)
    fmt.Printf("本卦: %s %s\n", ben.GuaSymbol, ben.Ming)
    fmt.Printf("变卦: %s %s\n", bian.GuaSymbol, bian.Ming)

    // 方式二：手动输入6个爻值（6/7/8/9）
    // 例如：老阳、少阳、少阴、老阴、少阳、少阳
    zy2, err := yi.DivineByCoinValues([6]int{9, 8, 7, 6, 8, 8})
    if err != nil {
        panic(err)
    }
    fmt.Printf("手动起卦: %s\n", zy2.GetGua(yi.Ben).Ming)

    // 方式三：使用当前时间作为种子
    zy3, coins3 := yi.DivineByCoinsRand()
    _ = zy3
    _ = coins3

    // 辅助：模拟单次投掷（输入阴面数0-3）
    result := yi.CoinThrowSimulate(2) // 2个阴面 → 少阴(7)
    fmt.Printf("模拟: %s\n", result)
}
```

### 输出说明

- `DivineByCoins` 返回 `(*ZhouYi, [6]CoinResult)`
- `CoinResult` 实现了 `fmt.Stringer`，打印为中文名（老阴/少阴/少阳/老阳）
- `CoinResult.IsChanging()` 判断是否为动爻
- `CoinResult.IsYang()` 判断是否为阳爻
- `CoinResult.YaoValue()` 返回爻值 6/7/8/9

---

## 2. 蓍草法（大衍法）

### 原理

蓍草法是最古老的起卦方法，源自《周易·系辞》：

> "大衍之数五十，其用四十有九。分二以象两，卦一以象三，揲之以四以象四时，归奇于扐以象闰，五岁再闰，故再扐而后挂。"

每次得一爻需重复 3 变：

1. 49 根蓍草，随机分为两堆（"分二以象两"）
2. 从右堆取 1 根挂起（"挂一"）
3. 左堆以 4 根一组数尽，取余数 1-4
4. 右堆（不含挂的 1 根）同样以 4 根一组数尽，取余数 1-4
5. 挂的 1 根 + 左余 + 右余 = 挂扐数（必为 5 或 9）
6. 49 - 挂扐数 = 本次剩余
7. 重复 3 次后，剩余数映射到爻值：

| 剩余数 | 爻值 | 名称 | 是否动爻 |
|--------|------|------|----------|
| 36 | 9 | 老阳 | ✅ 动爻 |
| 40 | 8 | 少阳 | ❌ |
| 44 | 7 | 少阴 | ❌ |
| 37 | 6 | 老阴 | ✅ 动爻 |

6 次共 18 变得到完整卦象。

### Go 代码示例

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 方式一：传入随机种子
    zy, results := yi.DivineByDayan(999)
    for i, r := range results {
        fmt.Printf("第%d爻: 剩余=%d, 爻值=%d, 动爻=%v\n",
            i+1, r.Remaining, r.YaoValue, r.IsChanging)
        // 每爻3变详情
        for s, step := range r.Steps {
            fmt.Printf("  变%d: 总=%d, 左=%d, 右=%d, 挂扐=%d, 余=%d\n",
                s+1, step.Total, step.Left, step.Right, step.Remainder, step.Final)
        }
    }

    ben := zy.GetGua(yi.Ben)
    fmt.Printf("本卦: %s %s\n", ben.GuaSymbol, ben.Ming)

    // 方式二：手动输入6个爻值
    zy2, err := yi.DivineByDayanValues([6]int{7, 7, 7, 7, 7, 7})
    if err != nil {
        panic(err)
    }
    fmt.Printf("手动大衍: %s\n", zy2.GetGua(yi.Ben).Ming)

    // 方式三：当前时间随机
    zy3, results3 := yi.DivineByDayanRand()
    _ = zy3
    _ = results3
}
```

### 输出说明

- `DivineByDayan` 返回 `(*ZhouYi, [6]DayanResult)`
- `DayanResult.Steps` 记录每爻的 3 变详情
- `DayanResult.Remaining` 最终剩余数（36/40/44/37）
- `DayanResult.YaoValue` 爻值（6/7/8/9）
- `DayanResult.IsChanging` 是否为动爻

---

## 3. 梅花易数法

### 原理

梅花易数为宋代邵雍所创，以数起卦，灵活便捷。核心公式：

```
上卦 = 上数 mod 8（从1起算，0视为8）
下卦 = 下数 mod 8（从1起算，0视为8）
动爻 = (上数 + 下数 + 变数) mod 6（从1起算，0视为6）
```

数可以是任意正整数——字数、年龄、金额、方位等均可。

常见用法：
- **报数起卦**：直接用报出的两个数作为上卦数和下卦数
- **时间起卦**：(年+月+日) mod 8 = 上卦，(年+月+日+时) mod 8 = 下卦
- **字数起卦**：总字数分上下两半
- **方位起卦**：方位数分上下卦

### Go 代码示例

```go
package main

import (
    "fmt"
    "time"
    "github.com/godcong/yi"
)

func main() {
    // 报数起卦：上卦=3, 下卦=5, 动爻=2
    zy := yi.DivineByMeihua(3, 5, 2)
    fmt.Printf("梅花起卦: %s\n", zy.GetGua(yi.Ben).Ming)
    fmt.Printf("变卦: %s\n", zy.GetGua(yi.Bian).Ming)

    // 时间梅花起卦：自动计算年月日时
    t := time.Date(2024, 6, 15, 10, 0, 0, 0, time.Local)
    zy2, year, month, day, hour := yi.DivineByMeihuaTime(t)
    fmt.Printf("时间: %d年%d月%d日%d时\n", year, month, day, hour)
    fmt.Printf("梅花时间卦: %s\n", zy2.GetGua(yi.Ben).Ming)

    // 动爻为-1时不指定动爻
    zy3 := yi.DivineByMeihua(8, 3, -1)
    fmt.Printf("无动爻: %s\n", zy3.GetGua(yi.Ben).Ming)
}
```

### 输出说明

- `DivineByMeihua` 返回 `*ZhouYi`，第三参数 `bian` 为动爻位置（0-5），传 -1 不指定
- `DivineByMeihuaTime` 返回 `(*ZhouYi, year, month, day, hour)`，时辰自动转换为 1-12

---

## 4. 时间起卦法

### 原理

时间起卦是最常用的自动起卦方法，只需提供年月日时即可：

```
上卦 = (年 + 月 + 日) mod 8
下卦 = (年 + 月 + 日 + 时) mod 8
动爻 = (年 + 月 + 日 + 时) mod 6
```

所有数取农历或公历均可，只要前后一致。

时辰对照（24小时制→十二时辰）：

| 时辰 | 时间范围 | 编号 |
|------|----------|------|
| 子时 | 23:00-01:00 | 1 |
| 丑时 | 01:00-03:00 | 2 |
| 寅时 | 03:00-05:00 | 3 |
| 卯时 | 05:00-07:00 | 4 |
| 辰时 | 07:00-09:00 | 5 |
| 巳时 | 09:00-11:00 | 6 |
| 午时 | 11:00-13:00 | 7 |
| 未时 | 13:00-15:00 | 8 |
| 申时 | 15:00-17:00 | 9 |
| 酉时 | 17:00-19:00 | 10 |
| 戌时 | 19:00-21:00 | 11 |
| 亥时 | 21:00-23:00 | 12 |

### Go 代码示例

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 公历时间起卦
    zy := yi.DivineByTimeGua(yi.TimeGuaParams{
        Year:  2024,
        Month: 6,
        Day:   15,
        Hour:  10,
    })
    ben := zy.GetGua(yi.Ben)
    bian := zy.GetGua(yi.Bian)
    fmt.Printf("时间起卦: 本卦=%s, 变卦=%s\n", ben.Ming, bian.Ming)

    // 农历时间起卦（月/日/时辰均为农历编号）
    zy2 := yi.DivineByLunarTime(2024, 5, 10, 4) // 农历甲辰年五月初十卯时
    fmt.Printf("农历起卦: %s\n", zy2.GetGua(yi.Ben).Ming)

    // 当前时间快速起卦
    zy3 := yi.DivineByCurrentTime()
    fmt.Printf("即时卦: %s\n", zy3.GetGua(yi.Ben).Ming)
}
```

### 输出说明

- `DivineByTimeGua` 接受 `TimeGuaParams{Year, Month, Day, Hour}`，Hour 为 0-23（自动转换时辰）
- `DivineByLunarTime` 接受农历参数，`shichen` 为 1-12 时辰编号
- `DivineByCurrentTime` 是 `DivineByTimeGua` 的快捷方式

---

## 方法对比

| 方法 | 输入 | 爻值来源 | 动爻判定 | 适用场景 |
|------|------|----------|----------|----------|
| 铜钱法 | 随机种子 或 6个爻值 | 抛铜钱模拟 | 老6/老9 | 传统占卜 |
| 蓍草法 | 随机种子 或 6个爻值 | 蓍草分堆模拟 | 老6/老9 | 正统古法 |
| 梅花易数 | 两个数+动爻 或 时间 | 数 mod 8/6 | 指定或自动 | 便捷占卜 |
| 时间起卦 | 年月日时 | 时间数 mod 8/6 | 自动计算 | 即时占卜 |

所有方法的最终结果都是 `*ZhouYi`，可进一步调用解卦、六亲、世应等功能。
