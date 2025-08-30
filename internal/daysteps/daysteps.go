package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("Размер данных не соответствует размеру пакета")
	}

	numberOfSteps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования")
	}

	times, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка при парсинге времени")
	}
	return numberOfSteps, times, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	numberSteps, t, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if numberSteps == 0 {
		return ""
	}

	var distanceInM float64
	distanceInM = float64(numberSteps) * stepLength

	destanceInKm := distanceInM / float64(mInKm)
	// func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// }
	caloriesBurned, err := spentcalories.WalkingSpentCalories(numberSteps, weight, height, t)
	if err != nil {
		return ""
	}
	returnString := fmt.Sprintf("Количество шагов: %d.\n Дистанция составила %.2f км. \n Вы сожгли %.2f ккал. \n", numberSteps, destanceInKm, caloriesBurned)
	return returnString
}
