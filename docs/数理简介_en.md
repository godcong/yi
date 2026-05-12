# Introduction to Numerology / 数理简介

The core of I Ching is a deductive system from Yin-Yang to hexagrams to numerology. This document traces the numerical chain implemented in this library.

## Yin-Yang and Bit Convention

Everything begins with the two primal forces:

| Concept | Symbol | Bit Representation in This Project |
|---------|--------|-------------------------------------|
| Yang (solid line) | ━━━ | **0** |
| Yin (broken line) | ━ ━ | **1** |

> **Important Convention**: This project encodes Yang as 0 and Yin as 1. This is opposite to the intuition of "Yang=1" but consistent with binary logic: Yin exists = 1, doesn't exist = 0.

This convention runs through the entire library — from the 3-bit encoding of the Eight Trigrams to the upper/lower trigram combinations of the 64 hexagrams.

```go
// Binary representation of the Eight Trigrams
Qian = 0b000  // Qian: three Yang lines
Kun  = 0b111  // Kun: three Yin lines
Kan  = 0b101  // Kan: Yin-Yang-Yin
```

## Four Symbols (Si Xiang)

Adding another layer of Yin-Yang on top of the two forces yields the Four Symbols:

| Four Symbols | Lower | Upper | Meaning |
|--------------|-------|-------|---------|
| Tai Yang (Greater Yang) | 0 (Yang) | 0 (Yang) | Yang + Yang |
| Shao Yin (Lesser Yin) | 0 (Yang) | 1 (Yin) | Yang + Yin |
| Shao Yang (Lesser Yang) | 1 (Yin) | 0 (Yang) | Yin + Yang |
| Tai Yin (Greater Yin) | 1 (Yin) | 1 (Yin) | Yin + Yin |

In the Coin Method, the Four Symbols correspond to four line values: 9=Old Yang (changing), 8=Young Yang, 7=Young Yin, 6=Old Yin (changing).

## Eight Trigrams (Bagua)

Three lines combine to form eight trigrams, each represented by a 3-bit binary number:

| Trigram | Bit | Decimal | Symbol | Image |
|---------|------|---------|--------|-------|
| Qian (乾) | 000 | 0 | ☰ | Heaven |
| Dui (兑) | 001 | 1 | ☱ | Lake |
| Li (离) | 010 | 2 | ☲ | Fire |
| Zhen (震) | 011 | 3 | ☳ | Thunder |
| Xun (巽) | 100 | 4 | ☴ | Wind |
| Kan (坎) | 101 | 5 | ☵ | Water |
| Gen (艮) | 110 | 6 | ☶ | Mountain |
| Kun (坤) | 111 | 7 | ☷ | Earth |

### Earlier Heaven (Fuxi) vs. Later Heaven (King Wen) Arrangement

**Earlier Heaven (Fuxi Bagua)**: Qian and Kun as north-south, representing natural opposition. This project's trigram constants `Qian=0…Kun=7` follow the Earlier Heaven order.

**Later Heaven (King Wen Bagua)**: Adds a "center" position to the Eight Trigrams, forming the **Nine Images**. The numbers change in the Later Heaven arrangement, related to the balance of the "square."

## Sixty-Four Hexagrams

Stacking another layer — upper trigram + lower trigram = 64 hexagrams:

- Upper Trigram (Outer): Lines 4, 5, 6 (4th, 5th, 6th lines)
- Lower Trigram (Inner): Lines 1, 2, 3 (1st, 2nd, 3rd lines)

8 × 8 = 64 combinations, corresponding to the 64 hexagrams of the I Ching.

```go
// Divination by number: upper=Qian(0), lower=Kan(5)
zy := yi.DivineByNumber(0, 5)  // Song (Conflict) - Heaven over Water
fmt.Printf("Hexagram: %s\n", zy.GetGua(yi.Ben).Ming)
```

### Hexagram Transformations

A primary hexagram can derive four transformed hexagrams:

| Transformation | Name | Method | Purpose |
|----------------|------|--------|---------|
| Primary | Ben | Original divination result | Current state |
| Transformed | Bian | Invert changing line | Development trend |
| Nuclear | Hu | Lines 2-3-4 as lower, 3-4-5 as upper | Intermediate process |
| Inverse | Cuo | Flip all Yin/Yang | Opposite perspective |
| Reverse | Zong | Turn upside down (reverse bit order) | Shift viewpoint |

```go
zy := yi.DivineByNumber(0, 0, 0)  // Qian (The Creative), changing at 1st line
fmt.Printf("Primary: %s\n", zy.GetGua(yi.Ben).Ming)     // Qian (The Creative)
fmt.Printf("Transformed: %s\n", zy.GetGua(yi.Bian).Ming) // Gou (Coming to Meet)
fmt.Printf("Nuclear: %s\n", zy.GetGua(yi.Hu).Ming)       // Qian (The Creative)
fmt.Printf("Inverse: %s\n", zy.GetGua(yi.Cuo).Ming)      // Kun (The Receptive)
fmt.Printf("Reverse: %s\n", zy.GetGua(yi.Zong).Ming)     // Qian (The Creative)
```

## Eighty-One Images (Ju)

The Eight Trigrams generate 81 images through upper/lower trigram combinations, called "Ju" (矩). The number 81 corresponds to the 81 chapters of the *Tao Te Ching*.

For details on the 81 images, see [The Creative Method of Bagua, Nine Images, and Laozi](https://zhuanlan.zhihu.com/p/90606767).

This project uses "Dayan Numbers" to name the 81 numerology system, used for analyzing name stroke fortune.

```go
dayan, _ := yi.GetDayan(1)
fmt.Printf("Number 1: %s\n", dayan.JiXiong)  // Auspicious
fmt.Printf("Meaning: %s\n", dayan.Meaning)
```

## Changing Lines and Transformed Hexagrams

Any number modulo 6 (treating 0 as 6) gives the changing line position. Inverting the corresponding line in the original hexagram yields the transformed hexagram, usually predicting the final outcome of the divination.

```go
// DivineByNumber third parameter specifies changing line position (0-5)
zy := yi.DivineByNumber(0, 0, 2)  // Qian, changing at 3rd line
bianPos := zy.GetBianYao()
fmt.Printf("Changing line position: %d\n", bianPos)
```

## Numerology Chain Summary

```
Yin-Yang (0/1)
  ↓ stack
Four Symbols (2-bit)
  ↓ stack
Eight Trigrams (3-bit, 8 types)
  ↓ upper × lower
Sixty-Four Hexagrams (64 types)
  ↓ 8×8 + center
Eighty-One Images / Ju (81 types) ← corresponds to Laozi 81 chapters
```

Each step builds on the previous one, forming a complete numerical system from the simplest Yin-Yang to the most complex 81 images.
