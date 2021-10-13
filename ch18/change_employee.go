package ch18

type ChangeEmployeeTransactioner interface {
	Change(e *Employee)
}
type ChangeEmployeeTransaction struct {
	itsEmpid int
}

func NewChangeEmployeeTransaction(empid int) *ChangeEmployeeTransaction {
	return &ChangeEmployeeTransaction{
		itsEmpid: empid,
	}
}

func (cet *ChangeEmployeeTransaction) execute(cetr ChangeEmployeeTransactioner) {
	e := NewPayrollDatabase().GetEmployee(cet.itsEmpid)
	if e != nil {
		cetr.Change(e)
	}
}

type ChangeNameTransaction struct {
	ChangeEmployeeTransaction
	name string
}

func NewChangeNameTransaction(empid int, name string) *ChangeNameTransaction {
	return &ChangeNameTransaction{
		ChangeEmployeeTransaction{
			itsEmpid: empid,
		},
		name,
	}
}
func (cet *ChangeNameTransaction) Change(e *Employee) {
	e.SetName(cet.name)
}

func (cet *ChangeNameTransaction) Execute() {
	cet.execute(cet)
}

// ChangeAddressTransaction ---------------------Address--------------------------------
type ChangeAddressTransaction struct {
	ChangeEmployeeTransaction
	itsAddress string
}

func NewChangeAddressTransaction(empid int, address string) *ChangeAddressTransaction {
	return &ChangeAddressTransaction{
		ChangeEmployeeTransaction{
			itsEmpid: empid,
		},
		address,
	}
}
func (cet *ChangeAddressTransaction) Change(e *Employee) {
	e.SetAddress(cet.itsAddress)
}

func (cet *ChangeAddressTransaction) Execute() {
	cet.execute(cet)
}

// ChangeClassificationTransaction ---------------------Classification--------------------------------
type ChangeClassificationTransactioner interface {
	GetClassification() PaymentClassificationer
	GetSchedule() PaymentScheduler
}
type ChangeClassificationTransaction struct {
	ChangeEmployeeTransaction
}

func (ccft *ChangeClassificationTransaction) change(e *Employee, cetr ChangeClassificationTransactioner) {
	e.SetClassification(cetr.GetClassification())
	e.SetSchedule(cetr.GetSchedule())
}

type ChangeHourlyTransaction struct {
	ChangeClassificationTransaction
	itsHourlyRate float64
}

func NewChangeHourlyTransaction(empid int, hoursRate float64) *ChangeHourlyTransaction {
	return &ChangeHourlyTransaction{
		ChangeClassificationTransaction{ChangeEmployeeTransaction{itsEmpid: empid}},
		hoursRate,
	}
}
func (cht *ChangeHourlyTransaction) GetClassification() PaymentClassificationer {
	return NewHourlyClassification()
}
func (cht *ChangeHourlyTransaction) GetSchedule() PaymentScheduler {
	return NewWeeklySchedule()
}
func (cht *ChangeHourlyTransaction) Change(e *Employee) {
	cht.change(e, cht)
}
func (cht *ChangeHourlyTransaction) Execute() {
	cht.execute(cht)
}

type ChangeSalariedTransaction struct {
	ChangeClassificationTransaction
	itsSalary float64
}

func NewChangeSalariedTransaction(empid int, salary float64) *ChangeSalariedTransaction {
	return &ChangeSalariedTransaction{
		ChangeClassificationTransaction{ChangeEmployeeTransaction{itsEmpid: empid}},
		salary,
	}
}
func (cst *ChangeSalariedTransaction) GetClassification() PaymentClassificationer {
	return NewSalariedClassification(cst.itsSalary)
}
func (cst *ChangeSalariedTransaction) GetSchedule() PaymentScheduler {
	return NewMonthlySchedule()
}
func (cst *ChangeSalariedTransaction) Change(e *Employee) {
	cst.change(e, cst)
}
func (cst *ChangeSalariedTransaction) Execute() {
	cst.execute(cst)
}

type ChangeCommissionedTransaction struct {
	ChangeClassificationTransaction
	itsSalary float64
}

func NewChangeCommissionedTransaction(empid int, salary float64) *ChangeCommissionedTransaction {
	return &ChangeCommissionedTransaction{
		ChangeClassificationTransaction{ChangeEmployeeTransaction{itsEmpid: empid}},
		salary,
	}
}
func (cct *ChangeCommissionedTransaction) GetClassification() PaymentClassificationer {
	return NewCommissionedClassification()
}
func (cct *ChangeCommissionedTransaction) GetSchedule() PaymentScheduler {
	return NewBiweeklySchedule()
}
func (cct *ChangeCommissionedTransaction) Change(e *Employee) {
	cct.change(e, cct)
}
func (cct *ChangeCommissionedTransaction) Execute() {
	cct.execute(cct)
}

type ChangeMethodTransaction struct {
}

type ChangeAffiliationTransaction struct {
}
