# Data Format Reference / 数据格式参考

## gua.json Structure

```json
{
  "index": "乾乾",
  "ming": "乾为天",
  "king_wen_order": 1,
  "tuan_text": "大哉乾元...",
  "xiang_text": "天行健..."
}
```

## jiegua.json Structure (📡 Program Direct Output)

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

**Lookup key**: Use `index` field (e.g., "乾乾"), not `ming` (e.g., "乾为天").

## JieGuaResult JSON Fields

Full structure returned by `JieGua()`:

| Field | Type | Description |
|-------|------|-------------|
| BenGuaInfo | object | Primary hexagram {Ming, GuaName, Symbol, GuaYi, TuanText, XiangText, JiXiong, GuaGong, Position, ShiYao, YingYao, ShangNum} |
| BianGuaInfo | object | Transformed hexagram (same fields as above) |
| HuGuaInfo | object | Nuclear hexagram {Ming, GuaName, GuaYi} |
| CuoGuaInfo | object | Inverse hexagram (same as HuGuaInfo) |
| ZongGuaInfo | object | Reverse hexagram (same as HuGuaInfo) |
| DongYaoPos | int | Moving line position 0-5 (1st to 6th line) |
| DongYaoText | string | Moving line text |
| DongYaoJiXiong | string | Moving line fortune |
| IsJi | bool | Overall auspiciousness judgment |
| JiXiongReason | string | Reason for judgment |
| FenXi | array | 8-dimension interpretation {Category, Content, JiXiong, Source} |
| JieDu | object | 📡 Pre-built interpretation (see below) |
| WuXingInfo | object | 📡 Five Elements lucky attributes (see below) |

## GuaJieDu Fields (📡 Program Direct Output)

| Field | Description |
|-------|-------------|
| CoreImage | Core imagery, one-sentence summary of hexagram essence |
| ShiYe | Career fortune (60-100 char paragraph) |
| AiQing | Love fortune |
| CaiYun | Wealth fortune |
| KaoShi | Exam fortune |
| JianKang | Health fortune |
| ChuXing | Travel fortune |
| GuanSi | Lawsuit fortune |
| JiaZhai | Home fortune |
| Yi | Do's list (4-6 items) |
| Ji | Don'ts list (4-6 items) |

## WuXingInfo Fields (📡 Program Direct Output)

| Field | Description |
|-------|-------------|
| WuXing | Metal/Wood/Water/Fire/Earth (determined by primary hexagram upper trigram ShangNum) |
| Direction | Lucky direction |
| LuckyNumber | Lucky number |
| LuckyColor | Lucky color |

**Five Elements → Direction/Number/Color Mapping**:

| Element | Number | Direction | Color |
|---------|--------|-----------|-------|
| Water | 1, 6 | North | Black, Blue |
| Fire | 2, 7 | South | Red, Purple |
| Wood | 3, 8 | East | Green, Cyan |
| Metal | 4, 9 | West | White, Silver |
| Earth | 5, 10 | Northeast, Southwest | Yellow, Brown |

## GuaFenXi Structure

```go
type GuaFenXi struct {
    Category FenXiCategory // Career/Love/Wealth/Exams/Health/Travel/Lawsuit/Home
    Content  string
    JiXiong  string        // Auspicious/Inauspicious/Neutral
    Source   string        // "Interpretation" | "Hexagram Meaning" | "Line Text" | "Meaning+Line"
}
```

Source priority: `"Interpretation"` > `"Meaning+Line"` > `"Meaning"` > `"Line"`

## FormatJieGua Three-Section Output

1. **Hexagram Interpretation** — Primary/Transformed/Nuclear/Inverse/Reverse + Moving Line
2. **Detailed Explanations** — 8-dimension categorized interpretation
3. **Overall Judgment** — Overall fortune + reason

## Eight Trigrams Mapping

| Trigram | Index | Earthly Branch | Element |
|---------|-------|----------------|---------|
| Qian (乾) | 0 | Xu-Hai | Metal |
| Dui (兑) | 1 | You | Metal |
| Li (离) | 2 | Wu | Fire |
| Zhen (震) | 3 | Mao | Wood |
| Xun (巽) | 4 | Chen-Si | Wood |
| Kan (坎) | 5 | Zi | Water |
| Gen (艮) | 6 | Chou-Yin | Earth |
| Kun (坤) | 7 | Wei-Shen | Earth |

**Key**: Trigram index 0-7 (not 1-8), ShangNum uses this index value.
