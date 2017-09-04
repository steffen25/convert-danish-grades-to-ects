package grades

import (
	"time"
	"errors"
	"strconv"
)

var grades = map[int]int{
	12: 12,
	10: 10,
	7: 7,
	4: 4,
	2: 2,
	0: 0,
	3: 3,
}

func ReadString(s string, pos int, length int) (string, error) {
	if pos < 0 || pos+length > len(s) {
		return "", errors.New("slice bounds out of range")
	}

	return s[pos-1:pos+length], nil
}

func ValidateDate(date string) bool {
	_, err := time.Parse("20060102", date)

	if err != nil {
		return false
	}

	return true
}

func ValidateGrade(grade string) bool {
	i, err := strconv.Atoi(grade)
	if err != nil {
		return false
	}
	if _, ok := grades[i]; !ok {
		return false
	}

	return true
}

func ConvertToECTS(danishGrade string) string {
	switch danishGrade {
		case "12":
			return "A"
		case "10":
			return "B"
		case "7":
			return "C"
		case "4":
			return "D"
		case "2":
			return "E"
		case "0":
			return "Fx"
		case "3":
			return "F"
		default:
			panic("unrecognized grade")
	}
}
