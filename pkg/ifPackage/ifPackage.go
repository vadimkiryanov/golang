package ifPackage

import "errors"

// Функции Условныи оператор if Обработка ошибок
func EnterTheClub(age int) (string, error) {
	if age >= 18 && age < 100 {
		return "Вход разрешен", nil

	} else if age < 18 {
		return "Вход запрещен", nil
	}

	return "Вход тоже запрещен", errors.New("do not enter the club")
}
