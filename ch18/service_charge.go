package ch18

type ServiceChargeTransaction struct {
	itsDate     int64
	itsCharge   float64
	itsMemberID int
}

func NewServiceChargeTransaction(memberId int, date int64, charge float64) *ServiceChargeTransaction {
	return &ServiceChargeTransaction{
		itsDate:     date,
		itsCharge:   charge,
		itsMemberID: memberId,
	}
}

func (sct *ServiceChargeTransaction) Execute() {
	e := NewPayrollDatabase().GetUnionMember(sct.itsMemberID)
	af := e.GetAffiliations()
	if uaf, ok := af.(*UnionAffiliation); ok {
		uaf.AddServiceCharge(NewServiceCharge(sct.itsDate, sct.itsCharge))
	}
}

type ServiceCharge struct {
	itsDate   int64
	itsCharge float64
}

func NewServiceCharge(date int64, charge float64) *ServiceCharge {
	return &ServiceCharge{
		itsDate:   date,
		itsCharge: charge,
	}
}

func (sc *ServiceCharge) GetDate() int64 {
	return sc.itsDate
}
func (sc *ServiceCharge) GetCharge() float64 {
	return sc.itsCharge
}
