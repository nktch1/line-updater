package model

import "fmt"

type sport int
type Sport interface {
	fmt.Stringer
	CalExpr() int
	disabler()
}

func (s sport) disabler() {}

var (
	Soccer   Sport = sport(0)
	Football Sport = sport(1)
	Baseball Sport = sport(2)
)

func (s sport) String() string {
	switch s {
	case Soccer:
		return "SOCCER"
	case Football:
		return "FOOTBALL"
	case Baseball:
		return "BASEBALL"
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

func NewSport(strSport string) Sport {
	var s Sport
	switch strSport {
	case "SOCCER":
		s = Soccer
	case "FOOTBALL":
		s = Football
	case "BASEBALL":
		s = Baseball
	}
	// default value
	return s
}
