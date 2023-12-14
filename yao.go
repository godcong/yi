package yi

type YaoIndex int

const (
	YaoChu   YaoIndex = iota
	YaoEr    YaoIndex = iota
	YaoSan   YaoIndex = iota
	YaoSi    YaoIndex = iota
	YaoWu    YaoIndex = iota
	YaoShang YaoIndex = iota
	YaoMax   YaoIndex = iota
)

type Yao struct {
	Index   YaoIndex
	JiXiong string
	YaoYue  string
	YaoYi   string
}
