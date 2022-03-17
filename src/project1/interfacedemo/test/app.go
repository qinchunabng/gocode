package test

type AInterface interface {
	A()
}

type Hero struct {
	Name string
	Age  int
}

// HeroSlice 声明一个Hero切片类型
type HeroSlice []Hero

// Len 实现Interface的接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 决定使用标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	//按照年龄的从小到达排序
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}
