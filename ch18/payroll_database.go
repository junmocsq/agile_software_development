package ch18

var gPayrollDatabase *PayrollDatabase = &PayrollDatabase{
	dict:    make(map[int]*Employee),
	dictMem: make(map[int]*Employee),
}

type PayrollDatabase struct {
	dict    map[int]*Employee
	dictMem map[int]*Employee
}

func NewPayrollDatabase() *PayrollDatabase {
	return gPayrollDatabase
}
func (p *PayrollDatabase) GetEmployee(empID int) *Employee {
	return p.dict[empID]
}

func (p *PayrollDatabase) AddEmployee(empID int, e *Employee) {
	p.dict[empID] = e
}
func (p *PayrollDatabase) DeleteEmployee(empid int) {
	delete(p.dict, empid)
}
func (p *PayrollDatabase) GetUnionMember(memberId int) *Employee {
	return p.dictMem[memberId]
}
func (p *PayrollDatabase) clear() {
	p.dict = nil
}
