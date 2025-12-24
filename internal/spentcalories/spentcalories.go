package spentcalories

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/daysteps"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
var (
	errConvSteps = errors.New("ошибка преобразования шагов")
	errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
	errConvTime = errors.New("ошибка конвертации времени")
	errIncomData = errors.New("неверные входящие данные")
	numberOfSteps int
	activityType string
	activityTime time.Duration
)
	parseData := strings.Split(data, ",")
	if len(parseData) == 3 { // проверка наличия трех элементов в слайсе
		numberOfSteps, err := strconv.Atoi(parseData[0]) // конвертация строки в число
			if err != nil { // проверка на ошибку конвертации
				return 0, "", 0, errConvSteps
			}
			if numberOfSteps <= 0 { // проверка числа шагов на отрицательность и равенство 0
				return 0, "", 0, errNuberOfStepsNeg
			}
			activityType = parseData[1]
			activityTime, err = time.ParseDuration(parseData[2]) // преобразуем вторую строку в time.Duration
			if err != nil { // проверка на ошибку преобразования
				return 0, "", 0, errConvTime
			}
	} else {
		return 0, "", 0, errIncomData
	}
	return numberOfSteps, activityType, activityTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height*walkingCaloriesCoefficient // рассчитываем длину шага
	distM := float64(steps)*stepLength // дистанция в метрах
	distKm := distM/mInKm 
	return distKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	var distKm float64
	if duration >= 0 {
		distKm = distance(steps, height)
	} else {
		return  0
	}
	averageSpeed := distKm/duration.Hours()
	return  averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var (
		trainingInfo string
		badTypeTraining = errors.New("неизвестный тип тренировки")
	numberOfSteps, activityType, activityTime, err := parseTraining(data)
	if err != nil {
		log.Printf("%v", err)
	}
	switch activityType {
	case "Ходьба":
		dist := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, activityTime)
		calories, _ := WalkingSpentCalories(numberOfSteps, weight, height, activityTime)
		trainingInfo = "Тип тренировки: "+activityType+" Длительность: "+strconv.FormatFloat(activityTime.Hours(),'f', 2, 64)+" ч. Дистанция: "+strconv.FormatFloat(dist,'f', 2, 64)+" км. Скорость: "+strconv.FormatFloat(averageSpeed,'f', 2, 64)+" км/ч Сожгли калорий: "+strconv.FormatFloat(calories,'f', 2, 64)
	case "Бег":
		dist := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, activityTime)
		calories, _ := WalkingSpentCalories(numberOfSteps, weight, height, activityTime)
		trainingInfo = "Тип тренировки: "+activityType+" Длительность: "+strconv.FormatFloat(activityTime.Hours(),'f', 2, 64)+" ч. Дистанция: "+strconv.FormatFloat(dist,'f', 2, 64)+" км. Скорость: "+strconv.FormatFloat(averageSpeed,'f', 2, 64)+" км/ч Сожгли калорий: "+strconv.FormatFloat(calories,'f', 2, 64)
	default: 
		err = badTypeTraining
		return _, err
	}
return trainingInfo, err
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	var (
		errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
		errWeight = errors.New("вес должен быть больше 0")
		errHeight = errors.New("рост должен быть больше 0")
	)
	if steps <= 0 {
		return 0, errNuberOfStepsNeg
	}
	if weight <= 0{
		return 0, errWeight
	}
	if height <= 0 {
		return 0, errHeight
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runSpentcalories := (weight * averageSpeed * durationInMinutes) / minInH
	return runSpentcalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	var (
		errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
		errWeight = errors.New("вес должен быть больше 0")
		errHeight = errors.New("рост должен быть больше 0")
	)
	if steps <= 0 {
		return 0, errNuberOfStepsNeg
	}
	if weight <= 0{
		return 0, errWeight
	}
	if height <= 0 {
		return 0, errHeight
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkSpentcalories := ((weight * averageSpeed * durationInMinutes) / minInH)*walkingCaloriesCoefficient
	return walkSpentcalories, nil
}
