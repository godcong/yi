# 数据格式参考

## gua.json 结构

```json
{
  "index": "乾乾",
  "ming": "乾为天",
  "king_wen_order": 1,
  "tuan_text": "大哉乾元...",
  "xiang_text": "天行健..."
}
```

## jiegua.json 结构（📡 程序直出）

```json
{
  "index": "乾乾",
  "core_image": "乾为天，纯阳刚健之象。自强不息，唯不骄不躁方得长久",
  "shiyi": "事业运势旺盛...",
  "aiqing": "感情方面阳刚主导...",
  "caiyun": "财运亨通...",
  "kaoshi": "学业考试运势极佳...",
  "jiankang": "身体健康...",
  "chuxing": "出行运势吉...",
  "guansi": "官司诉讼运势平...",
  "jiazhai": "家宅运势大吉...",
  "yi": ["积极进取", "拓展业务", "建立威信", "独立决策", "把握时机"],
  "ji": ["骄傲自满", "独断专行", "忽视细节", "过度劳累", "与人冲突"]
}
```

**查找键**：用 `index` 字段（如"乾乾"），不是 `ming`（如"乾为天"）。

## JieGuaResult JSON 字段

`JieGua()` 返回的完整结构：

| 字段 | 类型 | 说明 |
|------|------|------|
| BenGuaInfo | object | 本卦 {Ming, GuaName, Symbol, GuaYi, TuanText, XiangText, JiXiong, GuaGong, Position, ShiYao, YingYao, ShangNum} |
| BianGuaInfo | object | 变卦（字段同上） |
| HuGuaInfo | object | 互卦 {Ming, GuaName, GuaYi} |
| CuoGuaInfo | object | 错卦（同HuGuaInfo） |
| ZongGuaInfo | object | 综卦（同HuGuaInfo） |
| DongYaoPos | int | 动爻位置 0-5（初爻到上爻） |
| DongYaoText | string | 动爻爻辞 |
| DongYaoJiXiong | string | 动爻吉凶 |
| IsJi | bool | 综合吉凶判断 |
| JiXiongReason | string | 吉凶理由 |
| FenXi | array | 八维度解读 {Category, Content, JiXiong, Source} |
| JieDu | object | 📡 预置释义（见下） |
| WuXingInfo | object | 📡 五行幸运元素（见下） |

## GuaJieDu 字段（📡 程序直出）

| 字段 | 说明 |
|------|------|
| CoreImage | 核心意象，一句话概括卦象本质 |
| ShiYe | 事业运势（60-100字段落） |
| AiQing | 爱情运势 |
| CaiYun | 财运 |
| KaoShi | 考试 |
| JianKang | 健康 |
| ChuXing | 出行 |
| GuanSi | 官司 |
| JiaZhai | 家宅 |
| Yi | 宜列表（4-6项） |
| Ji | 忌列表（4-6项） |

## WuXingInfo 字段（📡 程序直出）

| 字段 | 说明 |
|------|------|
| WuXing | 金/木/水/火/土（由本卦上卦 ShangNum 决定） |
| Direction | 幸运方位 |
| LuckyNumber | 幸运数字 |
| LuckyColor | 幸运颜色 |

**五行→方位/数字/颜色映射**：

| 五行 | 数字 | 方位 | 颜色 |
|------|------|------|------|
| 水 | 1, 6 | 北 | 黑、蓝 |
| 火 | 2, 7 | 南 | 红、紫 |
| 木 | 3, 8 | 东 | 青、绿 |
| 金 | 4, 9 | 西 | 白、银 |
| 土 | 5, 10 | 东北、西南 | 黄、棕 |

## GuaFenXi 结构

```go
type GuaFenXi struct {
    Category FenXiCategory // 事业/爱情/财运/考试/健康/出行/官司/家宅
    Content  string
    JiXiong  string        // 吉/凶/平
    Source   string        // "释义" | "卦义" | "爻辞" | "卦义+爻辞"
}
```

Source 优先级：`"释义"` > `"卦义+爻辞"` > `"卦义"` > `"爻辞"`

## FormatJieGua 三段式输出

1. **【解卦】卦象解读** — 本卦/变卦/互卦/错卦/综卦 + 动爻
2. **【释义】解读内容的释义** — 八维度分类详解
3. **【综合判断】** — 整体吉凶 + 理由

## 八卦对应

| 卦 | 序号 | 地支 | 五行 |
|----|------|------|------|
| 乾 | 0 | 戌亥 | 金 |
| 兑 | 1 | 酉 | 金 |
| 离 | 2 | 午 | 火 |
| 震 | 3 | 卯 | 木 |
| 巽 | 4 | 辰巳 | 木 |
| 坎 | 5 | 子 | 水 |
| 艮 | 6 | 丑寅 | 土 |
| 坤 | 7 | 未申 | 土 |

**关键**：八卦序号 0-7（不是 1-8），ShangNum 用这个序号值。