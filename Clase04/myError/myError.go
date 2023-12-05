package myError

import (
	"errors"
	"fmt"
)

var (
	ErrorTax          = &ErrTax{"the salary entered does not reach the taxable minimum"}
	ErrorLess         = &ErrLess{"salary is less than 10000"}
	ErrorTax2         = errors.New("the salary entered does not reach the taxable minimum")
	MinimumHoursError = &ErrSalary{"the worker cannot have worked less than 80 hours per month"}
)

type ErrTax struct {
	Msg string
}
type ErrLess struct {
	Msg string
}
type ErrSalary struct {
	Msg string
}

func (te *ErrTax) Error() string {
	return te.Msg
}

func (le *ErrLess) Error() string {
	return le.Msg
}
func (se *ErrSalary) Error() string {
	return se.Msg
}

func IsPayTax(salary int) (err error) {

	if salary < 150000 {
		err = ErrorTax
		//err = ErrorTax2
		//err = fmt.Errorf("minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		return
	}
	return
}

func IsLess(salary int) (err error) {
	if salary < 10000 {
		err = ErrorLess

		return
	}
	return
}

func CalculateSalaryLessTax(hour int, vph float64) (salary float64, err error) {

	if hour < 80 {
		err = MinimumHoursError

		err = fmt.Errorf("%w - something else", err)
		return
	}
	salary = CalculateSalary(hour, vph)
	if salary >= 150000 {
		salary -= salary * 0.1
	}

	return
}

func CalculateSalary(hour int, vph float64) (salary float64) {

	salary = vph * float64(hour)

	return

}
