package ch18

import "testing"

func TestChangeName(t *testing.T) {
	e := NewEmployee(1, "陆雪琪", "小竹峰望月台")
	NewPayrollDatabase().AddEmployee(1, e)
	t.Log(NewPayrollDatabase().GetEmployee(1))
	cnt := NewChangeNameTransaction(1, "张小凡")
	cnt.Execute()
	t.Log(NewPayrollDatabase().GetEmployee(1))

	cat := NewChangeAddressTransaction(1, "大竹峰后山")
	cat.Execute()
	t.Log(NewPayrollDatabase().GetEmployee(1))
}

func TestChangeHourlyTransaction(t *testing.T) {
	empid := 2
	e := NewAddCommissionedEmployee(2, "卢云", "天地立心", 2500, 3.2)
	e.Execute()
	cht := NewChangeHourlyTransaction(empid, 27.52)
	cht.Execute()
	t.Log(NewPayrollDatabase().GetEmployee(2))

}
