# yi - Go 语言周易占卜库

[![Go Reference](https://pkg.go.dev/badge/github.com/godcong/yi.svg)](https://pkg.go.dev/github.com/godcong/yi)

周易六十四卦计算与查询库，提供完整的起卦、卦象变换、五行生克、六十甲子、六亲世应、81 数理、解卦等功能。

---

## 🌟 每日一卦 (Daily Hexagram Skill)

**推荐使用方式** — 集成 AI 助手的每日占卜体验，一句话完成起卦。

> 用户说「算一卦」→ AI 自动收集信息 → 生成完整运势报告（800+ 字）

### 安装

#### 方式一：npx 一键安装（推荐）

```bash
npx skills add daily-hexagram
```

自动下载 Skill 包 + 对应平台二进制，无需手动配置。

#### 方式二：安装脚本

```bash
# macOS / Linux
curl -fsSL https://raw.githubusercontent.com/godcong/yi/master/skill/scripts/install.sh | bash

# Windows (PowerShell)
iwr -Uri https://raw.githubusercontent.com/godcong/yi/master/skill/scripts/install.sh | iex
```

#### 方式三：Go Install（开发者）

```bash
go install github.com/godcong/yi/cmd/divine@latest
```

手动复制 Skill 文件：
```bash
cp -r skill/ ~/.qclaw/skills/daily-hexagram/
cp $(which divine) ~/.qclaw/skills/daily-hexagram/bin/yi
```

### 使用

在 AI 助手中，直接对话即可：

| 你说 | AI 做的事 |
|------|----------|
| 「算一卦」「今日运势」 | 收集姓名+生日 → 每日卦（同人同日同卦） |
| 「换个卦看看」 | 换时辰重新起一卦 |
| 「用铜钱算」 | 切换到铜钱法起卦 |
| 「用蓍草算」 | 切换到大衍法起卦 |
| 「梅花易数算一卦」 | 切换到梅花易数 |

**首次使用** AI 会询问姓名和生日，之后自动记住。同一个人同一天永远是同一个卦。

### 输出报告内容

```
🎯 卦象概览     — 卦名/符号/吉凶 + 核心意象 + 运势基调
🔮 卦象详解     — 本卦/变卦/动爻/互卦/错卦/综卦解读
💫 八维度运势   — 事业/爱情/财运/考试/健康/出行/官司/家宅
📋 宜忌指南     — 今日宜做之事 / 今日禁忌
🌟 今日指引     — 综合行动指引 + 幸运方位/数字/颜色
```

完整报告不少于 800 字。

### 支持的方法

| 方法 | CLI | 说明 |
|------|-----|------|
| 每日卦 | `daily` | 仅用日期，同人同日固定 |
| 时间卦 | `time` | 含时辰，不同时间不同卦 |
| 铜钱法 | `coins` | 六次投掷，最经典 |
| 大衍法 | `dayan` | 蓍草五十策，最古老 |
| 梅花易数 | `meihua` | 时辰推演 |
| 数字起卦 | `number` | 自定义上下卦数字 |

### Skill 目录结构

```
daily-hexagram/
├── SKILL.md              # AI 指令（数据直出 + AI 串联分工）
├── bin/yi                # 占卜程序（编译好的二进制）
├── references/
│   └── data-format.md    # JSON 输出结构参考
├── scripts/
│   ├── install.sh        # 安装脚本
│   └── install.bat
└── user_profile.json     # 用户信息（自动创建）
```

---

## 🖥️ CLI 命令行工具

### 安装

```bash
go install github.com/godcong/yi/cmd/divine@latest
```

### 用法

```bash
# 每日一卦（需要 -seed 确保个性化）
yi -method daily -seed "张三" -format json

# 时间起卦
yi -method time -seed "张三" -sex female -format json

# 铜钱法
yi -method coins -coins-seed 42

# 梅花易数
yi -method meihua -seed "张三" -year 2026 -month 5 -day 12

# 大衍法
yi -method dayan -dayan-seed 999

# 数字起卦
yi -method number -ben 3 -bian 5 -dong 2
```

### 完整参数

| 参数 | 默认 | 说明 |
|------|------|------|
| `-method` | `time` | 起卦方法：`time` / `daily` / `coins` / `meihua` / `dayan` / `number` |
| `-seed` | `""` | 用户标识（姓名/ID），`daily` 和 `time` 必传 |
| `-format` | `text` | 输出格式：`text` / `json` |
| `-sex` | `male` | 性别：`male` / `female` |
| `-year` | 当前年 | 公历年 |
| `-month` | 当前月 | 公历月 |
| `-day` | 当前日 | 公历日 |
| `-hour` | 当前时 | 小时 0-23 |
| `-coins-seed` | 0 | 铜钱法种子（0=随机） |
| `-dayan-seed` | 0 | 大衍法种子 |
| `-ben` | -1 | 数字起卦上卦（0-7） |
| `-bian` | -1 | 数字起卦下卦（0-7） |
| `-dong` | 0 | 数字起卦动爻（0-5） |
| `-version` | — | 打印版本号 |

---

## 📚 Go 函数库

### 安装

```bash
go get github.com/godcong/yi
```

### 快速开始

```go
package main

import (
    "fmt"
    "github.com/godcong/yi"
)

func main() {
    // 每日一卦（同人同日同卦）
    zy, _ := yi.DivineByDailyHexagram("张三")
    result := yi.JieGua(zy, yi.Male)
    fmt.Println(yi.FormatJieGua(result))

    // 铜钱法
    zy2, coins := yi.DivineByCoins(42)
    fmt.Printf("本卦: %s, 变卦: %s\n", zy2.GetGua(yi.Ben).Ming, zy2.GetGua(yi.Bian).Ming)

    // 蓍草法
    zy3, _ := yi.DivineByDayan(999)

    // 梅花易数
    zy4 := yi.DivineByMeihua(3, 5, 2)

    // 时间起卦
    zy5 := yi.DivineByTimeGua(yi.TimeGuaParams{
        Year: 2024, Month: 6, Day: 15, Hour: 10,
    })

    // 当前时间快速起卦
    zy6 := yi.DivineByCurrentTime()
    _ = zy6
}
```

### 新增: 每日卦函数

```go
// 同人同日同卦 — 每天固定
zy, err := yi.DivineByDailyHexagram("张三")
if err != nil {
    panic(err)
}
result := yi.JieGua(zy, yi.Male)

// result.JieDu 包含预置的 8 维度运势 + 宜忌
fmt.Println(result.JieDu.ShiYe)   // 事业运势
fmt.Println(result.JieDu.CoreImage) // 核心意象
fmt.Println(result.JieDu.Yi)      // 宜
fmt.Println(result.JieDu.Ji)      // 忌

// result.WuXingInfo 包含五行幸运元素
fmt.Println(result.WuXingInfo.LuckyNumber) // 幸运数字
fmt.Println(result.WuXingInfo.LuckyColor)  // 幸运颜色
```

### API 总览

#### 起卦函数

| 函数 | 说明 |
|------|------|
| `DivineByDailyHexagram(seed string) (*ZhouYi, error)` | 🆕 每日一卦 |
| `DivineByCurrentTime(personalSeed ...string) *ZhouYi` | 当前时间起卦 |
| `DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult)` | 铜钱法 |
| `DivineByCoinValues(values [6]int) (*ZhouYi, error)` | 铜钱法手动 |
| `DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult)` | 大衍法 |
| `DivineByDayanValues(values [6]int) (*ZhouYi, error)` | 大衍法手动 |
| `DivineByMeihua(shang, xia, bian int) *ZhouYi` | 梅花易数 |
| `DivineByMeihuaTime(t time.Time, personalSeed ...string)` | 梅花时间 |
| `DivineByTimeGua(params TimeGuaParams, personalSeed ...string) *ZhouYi` | 公历时间 |
| `DivineByLunarTime(y, m, d, h int) *ZhouYi` | 农历时间 |
| `DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi` | 数字起卦 |

#### 解卦

| 函数 | 说明 |
|------|------|
| `JieGua(zy *ZhouYi, sex Sex) *JieGuaResult` | 解卦主函数 |
| `FormatJieGua(result *JieGuaResult) string` | 格式化输出 |

#### 解卦结果结构

`JieGuaResult` 包含：

| 字段 | 说明 |
|------|------|
| `BenGuaInfo` / `BianGuaInfo` | 本卦/变卦 {Ming, GuaName, Symbol, GuaYi, TuanText, XiangText, JiXiong, ...} |
| `HuGuaInfo` / `CuoGuaInfo` / `ZongGuaInfo` | 互卦/错卦/综卦 |
| `DongYaoPos` / `DongYaoText` / `DongYaoJiXiong` | 动爻位置/爻辞/吉凶 |
| `IsJi` / `JiXiongReason` | 综合吉凶/理由 |
| `FenXi` | 八维度解读 [{Category, Content, JiXiong, Source}] |
| `JieDu` | 🆕 预置释义 {CoreImage, ShiYe, AiQing, ..., Yi, Ji} |
| `WuXingInfo` | 🆕 五行幸运元素 {WuXing, Direction, LuckyNumber, LuckyColor} |

#### 卦象查询

| 函数/方法 | 说明 |
|------|------|
| `GetGuaByIndex(index string) (*Gua, error)` | 按索引查卦（如 "乾乾"） |
| `GetGuaByXu(xu int) (*Gua, error)` | 按卦序查卦（1-64） |
| `ZhouYi.GetGua(guaType int) *Gua` | 获取本/变/互/错/综卦 |
| `ZhouYi.IsJi(sex Sex) bool` | 判断吉凶 |
| `Gua.GetYao(pos YaoPosition) *Yao` | 获取指定爻 |
| `Gua.GetShiYing() *ShiYingInfo` | 获取世应信息 |
| `Gua.GetGuaGong() Bagua` | 归属卦宫 |

#### 五行与六亲

| 函数/方法 | 说明 |
|------|------|
| `GetWuXingByNumber(n int) string` | 笔画数→五行 |
| `WuXing.Sheng()` / `.Ke()` / `.BeiSheng()` / `.BeiKe()` | 生克关系 |
| `GetLiuQin(guaGongWX, yaoWX WuXing) LiuQin` | 计算六亲 |

#### 81 数理

| 函数/方法 | 说明 |
|------|------|
| `GetDayan(n int) (*Dayan, error)` | 查询数理（1-81） |
| `Dayan.IsJi()` / `.IsXiong()` / `.IsBest()` | 吉凶判断 |

### 核心类型

| 类型 | 说明 |
|------|------|
| `Gua` (别名 `Hexagram`) | 卦象：含卦名/卦义/彖辞/象辞/六爻 |
| `Yao` (别名 `Line`) | 爻：含爻辞/吉凶/女命 |
| `ZhouYi` (别名 `IChing`) | 周易：含五种卦象 + 动爻 |
| `Bagua` (别名 `Trigram`) | 八卦：0=乾…7=坤 |
| `WuXing` | 五行：木/火/土/金/水 |
| `Sex` | 性别：Male/Female |
| `LiuQin` | 六亲：父母/兄弟/妻财/子孙/官鬼 |
| `Dayan` | 大衍数理：1-81 的吉凶详情 |
| `JieGuaResult` | 解卦结果（含 JieDu + WuXingInfo） |
| `GuaJieDu` | 🆕 预置释义（8 维度 + 宜忌） |
| `WuXingInfo` | 🆕 五行幸运元素 |

### 卦象类型常量

| 常量 | 值 | 说明 |
|------|------|------|
| `Ben` | 0 | 本卦 |
| `Bian` | 1 | 变卦 |
| `Hu` | 2 | 互卦（2-3-4爻为下卦，3-4-5爻为上卦） |
| `Cuo` | 3 | 错卦（阴阳全反） |
| `Zong` | 4 | 综卦（上下颠倒） |

### 八卦常量

| 常量 | 值 | 五行 | 名称 |
|------|------|------|------|
| `Qian` | 0 | 金 | 乾 ☰ |
| `Dui` | 1 | 金 | 兑 ☱ |
| `Li` | 2 | 火 | 离 ☲ |
| `Zhen` | 3 | 木 | 震 ☳ |
| `Xun` | 4 | 木 | 巽 ☴ |
| `Kan` | 5 | 水 | 坎 ☵ |
| `Gen` | 6 | 土 | 艮 ☶ |
| `Kun` | 7 | 土 | 坤 ☷ |

> **约定**：项目记阳为 0、阴为 1（bit 表示），与直觉相反。

---

## 数据生成

数据存储在 `data/` 目录下的 JSON 文件中：

- `data/gua.json` — 64 卦
- `data/tuan.json` — 彗辞
- `data/xiang.json` — 象辞
- `data/wenyan.json` — 文言
- `data/jiazi.json` — 六十甲子
- `data/jiegua.json` — 🆕 解卦释义（8维度运势+宜忌+核心意象）
- `data/guagong.json` — 卦宫

运行以下命令重新生成 `data.gen.go`：

```bash
# 先删除旧文件，否则 go generate 会跳过
rm data.gen.go
go generate ./...
```

**注意**：`go generate` 只在 `data.gen.go` 不存在时才生成，修改数据后需先删除再运行。

---

## 文档

- [数理简介](docs/数理简介.md) — 阴阳→八卦→六十四卦→八十一象的数理链路
- [起卦方法](docs/起卦方法.md) — 四种起卦法详解
- [八卦与六十四卦](docs/八卦与六十四卦.md) — 卦象体系与变换
- [大衍之数](docs/大衍之数.md) — 81 数理吉凶
- [六亲与世应](docs/六亲与世应.md) — 六亲系统与世应推算

---

## 许可

MIT License