package ch18

// affiliation 联系
type Affiliationer interface {
	CalculateDeductions(pc *Paycheck) float64
}

type UnionAffiliation struct {
	dues int64
}

func NewUnionAffiliation() *UnionAffiliation {
	return &UnionAffiliation{}
}

func (ua *UnionAffiliation) AddServiceCharge(sc *ServiceCharge) {

}
func (ua *UnionAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	return 66.66
}

type NoAffiliation struct {
}

func NewNoAffiliation() *NoAffiliation {
	return &NoAffiliation{}
}
func (ua *NoAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	return 00.01
}
