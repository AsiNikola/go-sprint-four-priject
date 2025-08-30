package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, fmt.Errorf("Размер данных не соответствует размеру пакета")
	}
	numSteps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования")
	}
	if numSteps <= 0 {
		return 0, "", 0, fmt.Errorf("Неверное количество")
	}
	t, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка при парсинге времени")
	}
	if t <= 0 {
		return 0, "", 0, fmt.Errorf("Неверная продолжительность")
	}
	return numSteps, slice[1], t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	strideLength := height * stepLengthCoefficient
	distanceInM := float64(steps) * strideLength
	distanceInKm := distanceInM / mInKm
	return distanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	averageSpeed := dist / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	numSteps, str, t, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	pattern := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	switch str {
	case "Ходьба":
		caloriesWalking, err := WalkingSpentCalories(numSteps, weight, height, t)
		if err != nil {
			log.Println(err)
		}
		speedWalking := meanSpeed(numSteps, height, t)
		ditanceWalking := distance(numSteps, height)

		return fmt.Sprintf(pattern, str, t.Hours(), ditanceWalking, speedWalking, caloriesWalking), nil
	case "Бег":
		caloriesRunning, err := RunningSpentCalories(numSteps, weight, height, t)
		if err != nil {
			log.Println(err)
		}
		speedRunning := meanSpeed(numSteps, height, t)
		ditanceRunning := distance(numSteps, height)

		return fmt.Sprintf(pattern, str, t.Hours(), ditanceRunning, speedRunning, caloriesRunning), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, fmt.Errorf("Входные данные времени некорректны")
	}
	if steps <= 0 {
		return 0, fmt.Errorf("Входные данные количества шагов некорректны")

	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("Входные данные параметров некорректны")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * averageSpeed * weight) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, fmt.Errorf("Входные данные времени некорректны")
	}
	if steps <= 0 {
		return 0, fmt.Errorf("Входные данные количества шагов некорректны")

	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("Входные данные параметров некорректны")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * averageSpeed * weight) / minInH

	caloriesBurned := calories * walkingCaloriesCoefficient
	return caloriesBurned, nil
}
