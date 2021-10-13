package ch18

type PaymentClassificationer interface {
	CalculatePay(paycheck *Paycheck) float64
}

type SalariedClassification struct {
	itsSalary float64
}

func NewSalariedClassification(salary float64) *SalariedClassification {
	return &SalariedClassification{
		itsSalary: salary,
	}
}

func (scf *SalariedClassification) CalculatePay(paycheck *Paycheck) float64 {
	return 88.88
}

type CommissionedClassification struct {
	itsSalary float64
}

func NewCommissionedClassification() *CommissionedClassification {
	return &CommissionedClassification{}
}
func (scf *CommissionedClassification) CalculatePay(paycheck *Paycheck) float64 {
	return 88.88
}
func (cc *CommissionedClassification) SalesReceipt(sr *SalesReceipt) {

}

type HourlyClassification struct {
	itsSalary float64
}

func NewHourlyClassification() *HourlyClassification {
	return &HourlyClassification{}
}
func (scf *HourlyClassification) CalculatePay(paycheck *Paycheck) float64 {
	return 88.88
}
func (hc *HourlyClassification) AddTimeCard(tc *TimeCard) {

}
