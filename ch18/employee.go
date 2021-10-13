package ch18

type Employee struct {
	itsEmpid   int
	itsName    string
	itsAddress string

	itsClassification PaymentClassificationer
	itsSchedule       PaymentScheduler
	itsMethod         PaymentMethodor
	itsAffiliation    Affiliationer
}

func NewEmployee(empid int, name, address string) *Employee {
	return &Employee{
		itsEmpid:   empid,
		itsName:    name,
		itsAddress: address,
	}
}

func (e *Employee) SetClassification(c PaymentClassificationer) {
	e.itsClassification = c
}
func (e *Employee) GetClassification() PaymentClassificationer {
	return e.itsClassification
}
func (e *Employee) SetSchedule(s PaymentScheduler) {
	e.itsSchedule = s
}
func (e *Employee) GetSchedule() PaymentScheduler {
	return e.itsSchedule
}
func (e *Employee) SetMethod(m PaymentMethodor) {
	e.itsMethod = m
}
func (e *Employee) SetName(name string) {
	e.itsName = name
}
func (e *Employee) SetAddress(address string) {
	e.itsAddress = address
}
func (e *Employee) SetAffiliations(a Affiliationer) {
	e.itsAffiliation = a
}
func (e *Employee) GetAffiliations() Affiliationer {
	return e.itsAffiliation
}

func (e *Employee) IsPayDate(payDate *Date) bool {
	return e.itsSchedule.IsPayDate(payDate)
}

func (e *Employee) GetPayPeriodStartDate(payPeriodStartDate *Date) Date {
	return e.itsSchedule.GetPayPeriodStartDate(payPeriodStartDate)
}

func (e *Employee) PayDay(pc *Paycheck) {
	//payDate := pc.GetPayPeriodEndDate()
	grossPay := e.itsClassification.CalculatePay(pc)
	deductions := e.itsAffiliation.CalculateDeductions(pc)

	netPay := grossPay - deductions
	pc.SetGrossPay(grossPay)
	pc.SetDeductions(deductions)
	pc.SetNetPay(netPay)
	e.itsMethod.Pay(pc)
}

type Paycheck struct {
	itsGrossPay, itsDeductions, itsNetPay float64
}

func (pc *Paycheck) GetPayPeriodEndDate() Date {
	return Date{
		day:   3,
		month: 12,
		year:  2021,
	}
}

func (pc *Paycheck) SetGrossPay(grossPay float64) {
	pc.itsGrossPay = grossPay
}
func (pc *Paycheck) SetDeductions(deductions float64) {
	pc.itsDeductions = deductions
}
func (pc *Paycheck) SetNetPay(netPay float64) {
	pc.itsNetPay = netPay
}
