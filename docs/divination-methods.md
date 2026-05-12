# Divination Methods / 起卦方法

This library implements four classic divination methods: Coin Method, Yarrow Stalk Method (Dayan), Plum Blossom Method, and Time Method. Each method supports both random simulation and manual input modes.

---

## 1. Coin Method (Three Coins)

### Principle

The Coin Method is the most commonly used divination method. Take 3 coins, toss them 6 times, and record the number of yin sides (tails) each time:

| Yin Sides | Line Value | Name | Changing? |
|-----------|------------|------|-----------|
| 0 (three yang) | 9 | Old Yang | ✅ Changing |
| 1 | 8 | Young Yang | ❌ |
| 2 | 7 | Young Yin | ❌ |
| 3 (three yin) | 6 | Old Yin | ✅ Changing |

Old Yang (9) and Old Yin (6) are changing lines. A changing line inverts to produce the transformed hexagram.

Line value to hexagram mapping: Yang line (8/9) = 0, Yin line (6/7) = 1. Six tosses from the 1st to 6th line form the complete hexagram.

### Go Code Example

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // Method 1: Pass a random seed, auto-simulate
    zy, coins := yi.DivineByCoins(42)
    for i, c := range coins {
        fmt.Printf("Toss %d: %s (value=%d, changing=%v)\n",
            i+1, c, c.YaoValue(), c.IsChanging())
    }

    ben := zy.GetGua(yi.Ben)
    bian := zy.GetGua(yi.Bian)
    fmt.Printf("Primary: %s %s\n", ben.GuaSymbol, ben.Ming)
    fmt.Printf("Transformed: %s %s\n", bian.GuaSymbol, bian.Ming)

    // Method 2: Manual input of 6 line values (6/7/8/9)
    // Example: Old Yang, Young Yang, Young Yin, Old Yin, Young Yang, Young Yang
    zy2, err := yi.DivineByCoinValues([6]int{9, 8, 7, 6, 8, 8})
    if err != nil {
        panic(err)
    }
    fmt.Printf("Manual: %s\n", zy2.GetGua(yi.Ben).Ming)

    // Method 3: Use current time as seed
    zy3, coins3 := yi.DivineByCoinsRand()
    _ = zy3
    _ = coins3

    // Helper: Simulate a single toss (input yin count 0-3)
    result := yi.CoinThrowSimulate(2) // 2 yin → Young Yin (7)
    fmt.Printf("Simulated: %s\n", result)
}
```

### Output Notes

- `DivineByCoins` returns `(*ZhouYi, [6]CoinResult)`
- `CoinResult` implements `fmt.Stringer`, prints as Chinese name (老阴/少阴/少阳/老阳)
- `CoinResult.IsChanging()` checks if it's a changing line
- `CoinResult.IsYang()` checks if it's a yang line
- `CoinResult.YaoValue()` returns line value 6/7/8/9

---

## 2. Yarrow Stalk Method (Dayan Method)

### Principle

The Yarrow Stalk Method is the oldest divination method, from the *I Ching · Great Commentary*:

> "The numbers of the Great Expansion amount to fifty, of which forty-nine are used. Divide them into two to represent the two primal forces. Set one aside to represent the three powers. Count by fours to represent the four seasons. Place the remainder to represent the intercalary month. There are two intercalary months in five years, so place the remainder again and then proceed."

Each line requires 3 transformations:

1. 49 yarrow stalks, randomly divided into two piles ("represent the two primal forces")
2. Take 1 stalk from the right pile and set aside ("set one aside")
3. Count the left pile in groups of 4, take remainder 1-4
4. Count the right pile (excluding the 1 set aside) in groups of 4, take remainder 1-4
5. Set-aside 1 + left remainder + right remainder = remainder count (always 5 or 9)
6. 49 - remainder count = remaining for this round
7. After 3 repetitions, the remaining number maps to line value:

| Remaining | Line Value | Name | Changing? |
|-----------|------------|------|-----------|
| 36 | 9 | Old Yang | ✅ Changing |
| 40 | 8 | Young Yang | ❌ |
| 44 | 7 | Young Yin | ❌ |
| 37 | 6 | Old Yin | ✅ Changing |

6 lines × 3 transformations = 18 total transformations for a complete hexagram.

### Go Code Example

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // Method 1: Pass a random seed
    zy, results := yi.DivineByDayan(999)
    for i, r := range results {
        fmt.Printf("Line %d: remaining=%d, value=%d, changing=%v\n",
            i+1, r.Remaining, r.YaoValue, r.IsChanging)
        // 3 transformations per line
        for s, step := range r.Steps {
            fmt.Printf("  Transform %d: total=%d, left=%d, right=%d, remainder=%d, final=%d\n",
                s+1, step.Total, step.Left, step.Right, step.Remainder, step.Final)
        }
    }

    ben := zy.GetGua(yi.Ben)
    fmt.Printf("Primary: %s %s\n", ben.GuaSymbol, ben.Ming)

    // Method 2: Manual input of 6 line values
    zy2, err := yi.DivineByDayanValues([6]int{7, 7, 7, 7, 7, 7})
    if err != nil {
        panic(err)
    }
    fmt.Printf("Manual Dayan: %s\n", zy2.GetGua(yi.Ben).Ming)

    // Method 3: Current time random
    zy3, results3 := yi.DivineByDayanRand()
    _ = zy3
    _ = results3
}
```

### Output Notes

- `DivineByDayan` returns `(*ZhouYi, [6]DayanResult)`
- `DayanResult.Steps` records 3 transformations per line
- `DayanResult.Remaining` final remaining count (36/40/44/37)
- `DayanResult.YaoValue` line value (6/7/8/9)
- `DayanResult.IsChanging` whether it's a changing line

---

## 3. Plum Blossom Method (Meihua)

### Principle

The Plum Blossom Method was created by Shao Yong in the Song Dynasty. It uses numbers to generate hexagrams. Core formula:

```
Upper Trigram = Upper Number mod 8 (1-based, 0 treated as 8)
Lower Trigram = Lower Number mod 8 (1-based, 0 treated as 8)
Changing Line = (Upper + Lower + Variation) mod 6 (1-based, 0 treated as 6)
```

Numbers can be any positive integers — word count, age, amount, direction, etc.

Common usages:
- **Number Method**: Directly use two reported numbers as upper and lower trigram numbers
- **Time Method**: (year+month+day) mod 8 = upper, (year+month+day+hour) mod 8 = lower
- **Word Count Method**: Split total word count into upper and lower halves
- **Direction Method**: Direction numbers as upper and lower trigrams

### Go Code Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/godcong/yi"
)

func main() {
    // Number divination: upper=3, lower=5, changing=2
    zy := yi.DivineByMeihua(3, 5, 2)
    fmt.Printf("Meihua: %s\n", zy.GetGua(yi.Ben).Ming)
    fmt.Printf("Transformed: %s\n", zy.GetGua(yi.Bian).Ming)

    // Time-based Meihua: auto-calculate year/month/day/hour
    t := time.Date(2024, 6, 15, 10, 0, 0, 0, time.Local)
    zy2, year, month, day, hour := yi.DivineByMeihuaTime(t)
    fmt.Printf("Time: %d year %d month %d day %d hour\n", year, month, day, hour)
    fmt.Printf("Meihua Time Hexagram: %s\n", zy2.GetGua(yi.Ben).Ming)

    // Changing line -1 means no specific changing line
    zy3 := yi.DivineByMeihua(8, 3, -1)
    fmt.Printf("No changing line: %s\n", zy3.GetGua(yi.Ben).Ming)
}
```

### Output Notes

- `DivineByMeihua` returns `*ZhouYi`, third parameter `bian` is changing line position (0-5), pass -1 to not specify
- `DivineByMeihuaTime` returns `(*ZhouYi, year, month, day, hour)`, hour auto-converted to 1-12

---

## 4. Time Method

### Principle

The Time Method is the most commonly used automatic divination method. Just provide year, month, day, and hour:

```
Upper Trigram = (Year + Month + Day) mod 8
Lower Trigram = (Year + Month + Day + Hour) mod 8
Changing Line = (Year + Month + Day + Hour) mod 6
```

Numbers can be from either the Gregorian or Lunar calendar, as long as they're consistent.

Chinese Hour Mapping (24-hour → 12 double-hours):

| Hour | Time Range | Number |
|------|------------|--------|
| Zi (子) | 23:00-01:00 | 1 |
| Chou (丑) | 01:00-03:00 | 2 |
| Yin (寅) | 03:00-05:00 | 3 |
| Mao (卯) | 05:00-07:00 | 4 |
| Chen (辰) | 07:00-09:00 | 5 |
| Si (巳) | 09:00-11:00 | 6 |
| Wu (午) | 11:00-13:00 | 7 |
| Wei (未) | 13:00-15:00 | 8 |
| Shen (申) | 15:00-17:00 | 9 |
| You (酉) | 17:00-19:00 | 10 |
| Xu (戌) | 19:00-21:00 | 11 |
| Hai (亥) | 21:00-23:00 | 12 |

### Go Code Example

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // Gregorian time divination
    zy := yi.DivineByTimeGua(yi.TimeGuaParams{
        Year:  2024,
        Month: 6,
        Day:   15,
        Hour:  10,
    })
    ben := zy.GetGua(yi.Ben)
    bian := zy.GetGua(yi.Bian)
    fmt.Printf("Time divination: Primary=%s, Transformed=%s\n", ben.Ming, bian.Ming)

    // Lunar time divination (month/day/hour are lunar calendar numbers)
    zy2 := yi.DivineByLunarTime(2024, 5, 10, 4) // Lunar JiaChen year, 5th month, 10th day, Mao hour
    fmt.Printf("Lunar divination: %s\n", zy2.GetGua(yi.Ben).Ming)

    // Quick divination from current time
    zy3 := yi.DivineByCurrentTime()
    fmt.Printf("Instant hexagram: %s\n", zy3.GetGua(yi.Ben).Ming)
}
```

### Output Notes

- `DivineByTimeGua` accepts `TimeGuaParams{Year, Month, Day, Hour}`, Hour is 0-23 (auto-converted to Chinese hour)
- `DivineByLunarTime` accepts lunar parameters, `shichen` is 1-12 hour number
- `DivineByCurrentTime` is a shortcut for `DivineByTimeGua`

---

## Method Comparison

| Method | Input | Line Value Source | Changing Line | Use Case |
|--------|-------|-------------------|---------------|----------|
| Coin Method | Random seed or 6 line values | Coin toss simulation | Old 6/Old 9 | Traditional divination |
| Yarrow Stalk | Random seed or 6 line values | Yarrow stalk division simulation | Old 6/Old 9 | Classical orthodox method |
| Plum Blossom | Two numbers + changing line or time | Number mod 8/6 | Specified or auto | Convenient divination |
| Time Method | Year/Month/Day/Hour | Time number mod 8/6 | Auto-calculated | Instant divination |

All methods ultimately return `*ZhouYi`, which can then be used for interpretation, Six Relations, Shi-Ying, and other functions.
