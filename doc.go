// Package yi is an I Ching (Zhou Yi) hexagram calculation library.
//
// This package implements the calculation and query functions for the 64 hexagrams of I Ching,
// used for numerology analysis.
//
// # Core Concepts
//
// ## Bagua (Eight Trigrams)
//
// The Bagua are the foundation of I Ching, each composed of three lines:
//   - Qian (乾) ☰ Heaven
//   - Dui (兑)  ☱ Lake
//   - Li (离)   ☲ Fire
//   - Zhen (震) ☳ Thunder
//   - Xun (巽)  ☴ Wind
//   - Kan (坎)  ☵ Water
//   - Gen (艮)  ☶ Mountain
//   - Kun (坤)  ☷ Earth
//
// ## 64 Gua (Hexagrams)
//
// Formed by combining two Bagua (upper and lower), creating 64 combinations.
//
// ## Yao (Lines)
//
// The basic unit of a hexagram, divided into:
//   - Yang line ⚊
//   - Yin line ⚋
//
// The six lines from bottom to top are: Chu (First), Er (Second), San (Third), Si (Fourth), Wu (Fifth), Shang (Top)
//
// ## Divination
//
// Supports multiple divination methods:
//   - By number: DivineByNumber()
//   - By time: DivineByTime()
//
// ## Hexagram Transformations
//
//   - Ben (Original): The original hexagram
//   - Bian (Changed): The hexagram after changing lines transformation
//   - Hu (Mutual): The hexagram formed by lines 2-3-4 and 3-4-5
//   - Cuo (Opposite): The hexagram with all Yin/Yang reversed
//   - Zong (Reversed): The hexagram turned upside down
//
// # Naming Conventions
//
// This package follows these naming conventions:
//
// 1. I Ching specific terms use Pinyin:
//    - Gua (hexagram), Yao (line), Bagua (eight trigrams)
//    - Dayan (great expansion numbers), WuXing (five elements)
//    - JiXiong (fortune), Ben/Bian/Hu/Cuo/Zong (transformations)
//
// 2. General operations use English:
//    - Get, Divine, Filter, Is
//    - Calculate, Transform, Parse
//
// 3. English aliases are provided for convenience:
//    - Hexagram = Gua
//    - Line = Yao
//    - Trigram = Bagua
//    - IChing = ZhouYi
//
// # Usage Examples
//
// Basic divination:
//
//	// Divination by number (upper trigram 3, lower trigram 5, changing line 2)
//	yi := yi.DivineByNumber(3, 5, 2)
//
//	// Get the original hexagram
//	ben := yi.GetGua(yi.Ben)
//	fmt.Println(ben.Ming) // Print hexagram name
//
//	// Check if it is auspicious
//	if yi.IsJi(yi.Female) {
//	    fmt.Println("Auspicious")
//	}
//
// Query Dayan numbers:
//
//	// Get fortune for stroke count 21
//	dayan, err := yi.GetDayan(21)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(dayan.JiXiong) // Output: 吉 (Auspicious)
//
// # References
//
//   - I Ching (Book of Changes)
//   - "I Ching Casual Talks" by Nan Huai-Chin
//   - "Zhou Yi Zheng Yi" by Kong Ying-Da
//
package yi
