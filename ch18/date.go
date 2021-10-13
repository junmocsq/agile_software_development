package ch18

type Date struct {
	day, month, year int
}

func payDay(day, month, year int) Date {
	return Date{
		day:   day,
		month: month,
		year:  year,
	}
}
