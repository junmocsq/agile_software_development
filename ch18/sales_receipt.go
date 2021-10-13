package ch18

type SalesReceiptTransaction struct {
	itsDate   int64
	itsAmount float64
	itsEmpid  int
}

func NewSalesReceiptTransaction(date int64, amount float64, empid int) *SalesReceiptTransaction {
	return &SalesReceiptTransaction{
		itsDate:   date,
		itsAmount: amount,
		itsEmpid:  empid,
	}
}

func (srt *SalesReceiptTransaction) Execute() {
	e := NewPayrollDatabase().GetEmployee(srt.itsEmpid)
	if e != nil {
		pc := e.GetClassification()
		if cc, ok := pc.(*CommissionedClassification); ok {
			cc.SalesReceipt(NewSalesReceipt(srt.itsDate, srt.itsAmount))
		} else {
			panic("Tried to add SalesReceipt to non-commissioned employee!")
		}
	} else {
		panic("No such employee.")
	}
}

type SalesReceipt struct {
	itsDate   int64
	itsAmount float64
}

func NewSalesReceipt(date int64, amount float64) *SalesReceipt {
	return &SalesReceipt{
		itsDate:   date,
		itsAmount: amount,
	}
}

func (sr *SalesReceipt) GetDate() int64 {
	return sr.itsDate
}

func (sr *SalesReceipt) GetAmount() float64 {
	return sr.itsAmount
}
