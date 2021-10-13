package ch18

var _ Transaction = &DeleteEmployeeTransaction{}

type DeleteEmployeeTransaction struct {
	itsEmpid int
}

func NewDeleteEmployeeTransaction(empid int) *DeleteEmployeeTransaction {
	return &DeleteEmployeeTransaction{itsEmpid: empid}
}

func (det *DeleteEmployeeTransaction) Execute() {
	NewPayrollDatabase().DeleteEmployee(det.itsEmpid)
}
