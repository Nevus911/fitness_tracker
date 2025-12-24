package daysteps

import (
	"errors"
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
	var (
		errIncomData = errors.New("неверные входящие данные")
		errConvSteps = errors.New("ошибка преобразования шагов")
		errNuberOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
		errConvTime = errors.New("ошибка конвертации времени")
		numberOfSteps int
		walkingTime time.Duration
	)
	parseData := strings.Split(data, ",") // разделяем строку на слайс строк
	if len(parseData) == 2 { // проверка наличия двух элементов в слайсе
		numberOfSteps, err := strconv.Atoi(parseData[0]) // конвертация строки в число
			if err != nil { // проверка на ошибку конвертации
				return 0, 0, errConvSteps
			}
			if numberOfSteps <= 0 { // проверка числа шагов на отрицательность и равенство 0
				return 0, 0, errNuberOfStepsNeg
			}
		walkingTime, err = time.ParseDuration(parseData[1]) // преобразуем вторую строку в time.Duration
			if err != nil { // проверка на ошибку преобразования
				return 0, 0, errConvTime
			}
	} else {
		return 0, 0, errIncomData //если длина слайса не равна 2 возвращаем нули и ошибку
	}
return numberOfSteps, walkingTime, nil // после преобразований и проверок функция parsePackage возвращает шаги, время прогулки
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var (
		distM float64
		distKm float64
	)
	numberOfSteps, walkingTime, err := parsePackage(data) // вызываем функцию и получаем данные из parsePackage
		if err != nil {
			fmt.Println("", err) // при возникновении ошибки возвращаем пустую строку и ошибку
		}
		if numberOfSteps <= 0 {
			fmt.Println("") // при отрицательном количестве шагов возвращаем пустую строку
		} else {
			// рассчет дистанции в км
			distM = float64(numberOfSteps)*stepLength
			distKm = distM/mInKm
	}
	// возвращаем из spentcalories.WalkingSpentCalories количество ккал
	kKal, _ := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, walkingTime)
	actionInfo := "Количество шагов: "+strconv.Itoa(numberOfSteps)+". Дистанция составила "+strconv.FormatFloat(distKm, 'f', 2, 64)+" км Вы сожгли "+strconv.FormatFloat(kKal,'f', 2, 64)
	return actionInfo
}
