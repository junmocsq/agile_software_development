package ch18

type PaymentMethodor interface {
	Pay(pc *Paycheck)
}
type DirectMethod struct {
	bank    string
	account string
}

func NewDirectMethod() *DirectMethod {
	return &DirectMethod{}
}
func (dm *DirectMethod) Pay(pc *Paycheck) {

}

type HoldMethod struct {
}

func NewHoldMethod() *HoldMethod {
	return &HoldMethod{}
}
func (dm *HoldMethod) Pay(pc *Paycheck) {

}

type MailMethod struct {
	address string
}

func NewMailMethod() *MailMethod {
	return &MailMethod{}
}
func (dm *MailMethod) Pay(pc *Paycheck) {

}
