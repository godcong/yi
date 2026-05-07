// Package yi 周易卦象计算库
//
// 本包实现了周易六十四卦的计算和查询功能，用于姓名学分析。
//
// # 核心概念
//
// ## 八卦 (Bagua / Trigram)
//
// 八卦是周易的基础，由三个爻组成：
//   - 乾 (Qian) ☰ 天
//   - 兑 (Dui)  ☱ 泽
//   - 离 (Li)   ☲ 火
//   - 震 (Zhen) ☳ 雷
//   - 巽 (Xun)  ☴ 风
//   - 坎 (Kan)  ☵ 水
//   - 艮 (Gen)  ☶ 山
//   - 坤 (Kun)  ☷ 地
//
// ## 六十四卦 (64 Gua / Hexagrams)
//
// 由两个八卦上下组合而成，共64种组合。
//
// ## 爻 (Yao / Line)
//
// 卦的基本组成单位，分为：
//   - 阳爻 ⚊
//   - 阴爻 ⚋
//
// 六爻从下往上依次为：初、二、三、四、五、上
//
// ## 起卦 (Divination)
//
// 支持多种起卦方式：
//   - 按数起卦: DivineByNumber()
//   - 按时间起卦: DivineByTime()
//
// ## 卦象变换
//
//   - 本卦 (Ben): 原始卦象
//   - 变卦 (Bian): 动爻变化后的卦象
//   - 互卦 (Hu): 取二三四、三四五爻组成的卦
//   - 错卦 (Cuo): 阴阳全反的卦象
//   - 综卦 (Zong): 上下颠倒的卦象
//
// # 命名规范
//
// 本包采用以下命名规范：
//
// 1. 周易专有术语使用拼音：
//    - Gua (卦), Yao (爻), Bagua (八卦)
//    - Dayan (大衍之数), Wuxing (五行)
//    - JiXiong (吉凶), Ben/Bian/Hu/Cuo/Zong (本变互错综)
//
// 2. 通用操作使用英文：
//    - Get, Divine, Filter, Is
//    - Calculate, Transform, Parse
//
// 3. 提供英文别名便于理解：
//    - Hexagram = Gua
//    - Line = Yao
//    - Trigram = Bagua
//    - IChing = ZhouYi
//
// # 使用示例
//
// 基本起卦：
//
//	// 按数起卦 (上卦3, 下卦5, 变爻2)
//	yi := yi.DivineByNumber(3, 5, 2)
//
//	// 获取本卦
//	ben := yi.GetGua(yi.Ben)
//	fmt.Println(ben.Ming) // 输出卦名
//
//	// 判断是否吉
//	if yi.IsJi(yi.Female) {
//	    fmt.Println("吉")
//	}
//
// 查询大衍之数：
//
//	// 获取笔画数21的吉凶
//	dayan, err := yi.GetDayan(21)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(dayan.JiXiong) // 输出: 吉
//
// # 参考资料
//
//   - 《周易》
//   - 《易经杂说》南怀瑾
//   - 《周易正义》孔颖达
//
package yi
