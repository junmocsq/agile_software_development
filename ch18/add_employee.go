package ch18

type AddEmployeeTransactioner interface {
	GetClassification() PaymentClassificationer
	GetSchedule() PaymentScheduler
	Transaction
}

type AddEmployeeTransaction struct {
	itsEmpid   int
	itsName    string
	itsAddress string
}

func NewAddEmployeeTransaction(empid int, name, address string) *AddEmployeeTransaction {
	return &AddEmployeeTransaction{
		itsEmpid:   empid,
		itsName:    name,
		itsAddress: address,
	}
}

func EmployeeTransactionExecute(a AddEmployeeTransactioner, aet *AddEmployeeTransaction) {
	pc := a.GetClassification()
	ps := a.GetSchedule()
	pm := NewHoldMethod()

	e := NewEmployee(aet.itsEmpid, aet.itsName, aet.itsAddress)
	e.SetClassification(pc)
	e.SetSchedule(ps)
	e.SetMethod(pm)
	NewPayrollDatabase().AddEmployee(aet.itsEmpid, e)
}

var _ Transaction = &AddHourlyEmployee{}
var _ AddEmployeeTransactioner = &AddHourlyEmployee{}

type AddHourlyEmployee struct {
	AddEmployeeTransaction
	itsHourlyRate int
}

func NewAddHourlyEmployee(empid int, name, address string, hourlyRate int) AddEmployeeTransactioner {
	return &AddHourlyEmployee{
		AddEmployeeTransaction{
			itsEmpid:   empid,
			itsName:    name,
			itsAddress: address,
		},
		hourlyRate,
	}
}
func (ase *AddHourlyEmployee) GetClassification() PaymentClassificationer {
	return NewHourlyClassification()
}
func (ase *AddHourlyEmployee) GetSchedule() PaymentScheduler {
	return NewWeeklySchedule()
}
func (ase *AddHourlyEmployee) Execute() {
	EmployeeTransactionExecute(ase, &ase.AddEmployeeTransaction)
}

var _ Transaction = &AddSalariedEmployee{}
var _ AddEmployeeTransactioner = &AddSalariedEmployee{}

// 销售凭证
type AddSalariedEmployee struct {
	AddEmployeeTransaction
	itsSalary float64
}

func NewAddSalariedEmployee(empid int, name, address string, salary float64) AddEmployeeTransactioner {
	return &AddSalariedEmployee{
		AddEmployeeTransaction{
			itsEmpid:   empid,
			itsName:    name,
			itsAddress: address,
		},
		salary,
	}
}
func (ase *AddSalariedEmployee) GetClassification() PaymentClassificationer {
	return NewSalariedClassification(ase.itsSalary)
}
func (ase *AddSalariedEmployee) GetSchedule() PaymentScheduler {
	return NewMonthlySchedule()
}
func (ase *AddSalariedEmployee) Execute() {
	EmployeeTransactionExecute(ase, &ase.AddEmployeeTransaction)
}

var _ Transaction = &AddCommissionedEmployee{}
var _ AddEmployeeTransactioner = &AddCommissionedEmployee{}

// AddCommissionedEmployee 雇员
type AddCommissionedEmployee struct {
	AddEmployeeTransaction
	itsSalary         float64
	itsCommissionRate float64
}

func NewAddCommissionedEmployee(empid int, name, address string, salary float64, commissionRate float64) AddEmployeeTransactioner {
	return &AddCommissionedEmployee{
		AddEmployeeTransaction{
			itsEmpid:   empid,
			itsName:    name,
			itsAddress: address,
		},
		salary,
		commissionRate,
	}
}

func (ase *AddCommissionedEmployee) GetClassification() PaymentClassificationer {
	return NewCommissionedClassification()
}
func (ase *AddCommissionedEmployee) GetSchedule() PaymentScheduler {
	return NewBiweeklySchedule()
}
func (ase *AddCommissionedEmployee) Execute() {
	EmployeeTransactionExecute(ase, &ase.AddEmployeeTransaction)
}
