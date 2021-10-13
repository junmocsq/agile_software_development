package ch18

type TimeCardTransaction struct {
	itsDate  int64
	itsHours float64
	itsEmpid int
}

func NewTimeCardTransaction(date int64, hours float64, empid int) *TimeCardTransaction {
	return &TimeCardTransaction{
		itsDate:  date,
		itsHours: hours,
		itsEmpid: empid,
	}
}

func (tct *TimeCardTransaction) Execute() {
	e := NewPayrollDatabase().GetEmployee(tct.itsEmpid)
	if e != nil {
		pc := e.GetClassification()
		if hc, ok := pc.(*HourlyClassification); ok {
			hc.AddTimeCard(NewTimeCard(tct.itsDate, tct.itsHours))
		} else {
			panic("Tried to add timecard to non-hourly employee!")
		}
	} else {
		panic("No such employee.")
	}
}

type TimeCard struct {
	itsDate  int64
	itsHours float64
}

func NewTimeCard(date int64, hours float64) *TimeCard {
	return &TimeCard{
		itsDate:  date,
		itsHours: hours,
	}
}

func (tc *TimeCard) GetDate() int64 {
	return tc.itsDate
}
func (tc *TimeCard) GetHours() float64 {
	return tc.itsHours
}
