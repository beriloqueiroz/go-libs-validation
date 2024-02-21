package validation

import (
	"regexp"
	"slices"
	"strconv"
)

func Cpf(cpfIn string) bool {
	invalidCPFs := []string{"11111111111", "22222222222", "33333333333", "44444444444", "55555555555", "66666666666", "77777777777", "88888888888", "99999999999", "00000000000"}
	regex := regexp.MustCompile(`\\D`)
	cpf := regex.ReplaceAllString(cpfIn, "")
	if len(cpf) == 11 && !slices.Contains(invalidCPFs, cpf) {
		digits := cpf[9:]
		first, err := firstDigit(cpf)
		if err != nil {
			return false
		}
		second, err := secondDigit(cpf)
		if err != nil {
			return false
		}
		er := strconv.Itoa(first) + strconv.Itoa(second)
		return digits == er
	} else {
		return false
	}
}

func secondDigit(cpf string) (int, error) {
	multiply := 11
	sum := 0
	var rest int
	for rest = 0; rest < 10; rest++ {
		value, err := strconv.Atoi(string(cpf[rest]))
		if err != nil {
			return -1, err
		}
		sum += multiply * value
		multiply--
	}
	rest = 11 - sum%11
	if rest >= 10 {
		rest = 0
	}
	return rest, nil
}

func firstDigit(cpf string) (int, error) {
	multiply := 10
	sum := 0
	var rest int
	for rest = 0; rest < 9; rest++ {
		value, err := strconv.Atoi(string(cpf[rest]))
		if err != nil {
			return -1, err
		}
		sum += multiply * value
		multiply--
	}
	rest = 11 - sum%11
	if rest >= 10 {
		rest = 0
	}
	return rest, nil
}
