package ch18

type PaymentScheduler interface {
	IsPayDate(payDate *Date) bool
	GetPayPeriodStartDate(payPeriodStartDate *Date) Date
}

type MonthlySchedule struct {
}

func NewMonthlySchedule() *MonthlySchedule {
	return &MonthlySchedule{}
}
func (ms *MonthlySchedule) IsPayDate(payDate *Date) bool {
	return true
}
func (ms *MonthlySchedule) GetPayPeriodStartDate(payPeriodStartDate *Date) Date {
	return Date{
		day:   1,
		month: 2,
		year:  3,
	}
}

type WeeklySchedule struct {
}

func NewWeeklySchedule() *WeeklySchedule {
	return &WeeklySchedule{}
}
func (ms *WeeklySchedule) IsPayDate(payDate *Date) bool {
	return true
}
func (ms *WeeklySchedule) GetPayPeriodStartDate(payPeriodStartDate *Date) Date {
	return Date{
		day:   1,
		month: 2,
		year:  3,
	}
}

type BiweeklySchedule struct {
}

func NewBiweeklySchedule() *BiweeklySchedule {
	return &BiweeklySchedule{}
}
func (ms *BiweeklySchedule) IsPayDate(payDate *Date) bool {
	return true
}
func (ms *BiweeklySchedule) GetPayPeriodStartDate(payPeriodStartDate *Date) Date {
	return Date{
		day:   1,
		month: 2,
		year:  3,
	}
}
