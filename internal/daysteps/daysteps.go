package daysteps

import (
	"errors"
	"fmt"
	"log"
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
var (
	// возможные ошибки
	errIncomData = errors.New("неверные входящие данные")
	errConvSteps = errors.New("ошибка преобразования шагов")
	errNubmerOfStepsNeg = errors.New("количество шагов 0 или отрицательно")
	errConvTime = errors.New("ошибка конвертации времени или время отрицательное")
	errTimeNeg = errors.New("нулевая или отрицательная продолжительность")
)
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	var (
		// внутренние переменные функции
		numberOfSteps int
		walkingTime time.Duration
	)
	parseData := strings.Split(data, ",") // парсим строку на слайс строк
	if len(parseData) != 2 {
		return 0, 0, errIncomData // отсекаем все значения если их больше двух
	}
		numberOfSteps, err := strconv.Atoi(parseData[0]) // конвертация строки в число
			if err != nil{ // проверка на ошибку конвертации
				return 0, 0, errConvSteps
			}
			if numberOfSteps <=0 {
				return 0, 0, errNubmerOfStepsNeg //проверка на 0 и отриц количество шагов
			}
		walkingTime, err = time.ParseDuration(parseData[1]) // преобразуем вторую строку из полученного слайса в time.Duration
			if err != nil { // проверка на ошибку преобразования
				return 0, 0, errConvTime
			}
			if walkingTime <= 0 {
				return 0, 0, errTimeNeg // проверка времени на нулевое значение и отрицательность
			}
return numberOfSteps, walkingTime, nil // после преобразований и проверок функция parsePackage возвращает шаги, время прогулки и ошибки
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var (
		distM float64
		distKm float64
	)
	numberOfSteps, walkingTime, err := parsePackage(data) // передаем data в функцию parsePackage и получаем шаги время и ошибку
		if err != nil {
			log.Println(err) // при возвращении ошибки из parsePackage пишем ее в лог и возвращаем пустую строку
			return ""
		}
		// дистанция в метрах
		distM = float64(numberOfSteps)*stepLength
		// перевод дистанции в километры 
		distKm = distM/mInKm
	// возвращаем из spentcalories.WalkingSpentCalories количество ккал
	kKal, _ := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, walkingTime)
	// формируем нужную информацию в строку и возвращаем ее
	actionInfo := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",numberOfSteps, distKm, kKal)
	return actionInfo
}
