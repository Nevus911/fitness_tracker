package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	//"github.com/Yandex-Practicum/tracker/internal/daysteps"
	//"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	// объявляем необходимые ошибки и переменные
	errBadNumElements = errors.New("неверное количество элементов")
	errConvSteps = errors.New("ошибка преобразования шагов")
	errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
	errConvTime = errors.New("ошибка конвертации времени")
	errTimeNeg = errors.New("время не может быть отрицательным или 0")
	numberOfSteps int
	activityType string
	activityTime time.Duration
)
	parseData := strings.Split(data, ",")
	if len(parseData) != 3 { // проверка наличия трех элементов в слайсе
		return 0, "", 0, errBadNumElements
	}
		numberOfSteps, err := strconv.Atoi(parseData[0]) // конвертация строки в число
			if err != nil { // проверка на ошибку конвертации возвращаем нули и ошибку
				return 0, "", 0, errConvSteps
			}
			if numberOfSteps <= 0 { // проверка числа шагов на отрицательность и равенство 0 возвращаем нули и ошибку
				return 0, "", 0, errNuberOfStepsNeg
			}
			activityType = parseData[1] // присваиваем переменной значение второго элемента слайса
			activityTime, err = time.ParseDuration(parseData[2]) // преобразуем вторую строку в time.Duration
			if err != nil { // проверка на ошибку преобразования возвращаем нули и ошибку
				return 0, "", 0, errConvTime
			}
			if activityTime <= 0 {
				return 0, "", 0, errTimeNeg // проверка времени на 0 и отрицательность возвращаем нули и ошибку
			}
	return numberOfSteps, activityType, activityTime, nil // возвращаем полученные данные и nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height*stepLengthCoefficient // рассчитываем длину шага
	distM := float64(steps)*stepLength // дистанция в метрах
	distKm := distM/mInKm // дистанция в километрах
	return distKm // функция возвращаем дистанцию в километрах
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 { // проверяем переданное время на 0 и отрицательность возвращаем 0
		return 0
	}
	averageSpeed := distance(steps, height)/duration.Hours() // рассчет и возвращение средней скорости
	return  averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// объявляем необходимые переменные и ошибки
	var (
		trainingInfo string
		badTypeTraining = errors.New("неизвестный тип тренировки")
	)
	// передаем в parseTraining data и возвращаем нужные значения и ошибки
	numberOfSteps, activityType, activityTime, err := parseTraining(data)
	if err != nil {
		log.Println(err) // при возвращении ошибки пишем ее в лог
	}
	// проверяем тип тренировки возвращаем из функций meanSpeed и Walking.../Running... необходимые данные, формируем строки
	switch activityType {
	case "Ходьба":
		dist := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, activityTime)
		calories, _ := WalkingSpentCalories(numberOfSteps, weight, height, activityTime)
		trainingInfo = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, float64(activityTime.Hours()), dist, averageSpeed, calories)
	case "Бег":
		dist := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, activityTime)
		calories, _ := RunningSpentCalories(numberOfSteps, weight, height, activityTime)
		trainingInfo = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, float64(activityTime.Hours()), dist, averageSpeed, calories)
	default : 
	return "", badTypeTraining
	}
	return trainingInfo, nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// объявляем необходимые ошибки
	var (
		errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
		errWeight = errors.New("вес должен быть больше 0")
		errHeight = errors.New("рост должен быть больше 0")
		errDuration = errors.New("время должно быть больше 0")
	)
	// проверка шагов, веса, роста и времени на 0 и отрицательность
	if steps <= 0 {
		return 0, errNuberOfStepsNeg
	}
	if weight <= 0{
		return 0, errWeight
	}
	if height <= 0 {
		return 0, errHeight
	}
	if duration <= 0 {
		return 0, errDuration
	}
	// расчет калорий при беге
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runSpentcalories := (weight * averageSpeed * durationInMinutes) / minInH
	return runSpentcalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// объявляем необходимые ошибки
	var (
		errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
		errWeight = errors.New("вес должен быть больше 0")
		errHeight = errors.New("рост должен быть больше 0")
		errDuration = errors.New("время должно быть больше 0")
	)
	// проверка шагов, веса, роста и времени на 0 и отрицательность
	if steps <= 0 {
		return 0, errNuberOfStepsNeg
	}
	if weight <= 0{
		return 0, errWeight
	}
	if height <= 0 {
		return 0, errHeight
	}
	if duration <= 0 {
		return 0, errDuration
	}
	// расчет калорий при беге 
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkSpentcalories := ((weight * averageSpeed * durationInMinutes) / minInH)*walkingCaloriesCoefficient
	return walkSpentcalories, nil
}
