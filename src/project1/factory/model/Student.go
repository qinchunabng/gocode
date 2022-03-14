package model

type student struct {
	Name  string
	score float64
}

//因为student结构体首字母是小写，因此只能model包中使用
//通过工厂模式来解决
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
