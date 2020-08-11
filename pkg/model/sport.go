package model

import "fmt"

type sport int

// Sport тип данных для спорта
type Sport interface {
	fmt.Stringer
	CalExpr() int
	disabler()
}

// запрет на реализацию этого интерфейса извне данного пакета
func (s sport) disabler() {}

var (
	// Soccer - вид спорта
	Soccer Sport = sport(0)

	// Football - вид спорта
	Football Sport = sport(1)

	// Baseball - вид спорта
	Baseball Sport = sport(2)
)

// метод интерфейса Stringer
func (s sport) String() string {
	switch s {
	case Soccer:
		return "soccer"
	case Football:
		return "football"
	case Baseball:
		return "baseball"
	}
	return "-"
}

func (s sport) CalExpr() int {
	switch s {
	case Soccer:
		return 0
	case Football:
		return 1
	case Baseball:
		return 2
	}
	return -1
}

// NewSport конструктор для спорта
func NewSport(strSport string) Sport {
	var s Sport
	switch strSport {
	case "soccer":
		s = Soccer
	case "football":
		s = Football
	case "baseball":
		s = Baseball
	}
	// default value
	return s
}
