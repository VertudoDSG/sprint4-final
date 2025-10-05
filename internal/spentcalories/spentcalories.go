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
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	//Алгоритм реализации функции:
	// Разделить строку на слайс строк.
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов,
	// вид активности и продолжительность.
	parsedData := strings.Split(data, ",")
	if len(parsedData) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect Input")
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, "", 0, err
	} else if steps <= 0 {
		return 0, "", 0, fmt.Errorf("Количество шагов должно быть больше 0")
	}
	// Преобразовать третий элемент слайса в time.Duration.
	// В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	activTime, err := time.ParseDuration(parsedData[2])
	if err != nil {
		return 0, "", 0, err
	} else if activTime <= 0 {
		return 0, "", 0, fmt.Errorf("Колличество минут должно быть больше 0")
	}
	// Если всё прошло без ошибок, верните количество шагов, вид активности,
	// продолжительность и nil (для ошибки).
	return steps, parsedData[1], activTime, nil
}

func distance(steps int, height float64) float64 {
	// Для вычисления дистанции: рассчитайте длину шага.
	// Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	// Соответствующая константа уже определена в пакете. умножьте пройденное количество шагов на длину шага.
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
	// Обратите внимание, что целочисленную переменную steps необходимо будет привести к другому числовому типу.
	distance := height * stepLengthCoefficient * float64(steps) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Функция принимает количество шагов steps, рост пользователя height и
	// продолжительность активности duration  и возвращает среднюю скорость.
	// Алгоритм реализации функции:
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	// Вычислить дистанцию с помощью distance().
	distance := distance(steps, height)
	// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах.
	// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Получить значения из строки данных с помощью функции parseTraining()
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Рассчитать дистанцию и среднюю скорость
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	// Проверить тип тренировки с помощью switch
	switch trainingType {
	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			duration.Hours(), dist, speed, calories), nil

	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			duration.Hours(), dist, speed, calories), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Алгоритм реализации функции:
	//Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect Input")
	}

	//Рассчитать среднюю скорость с помощью meanSpeed().
	maenSpeed := meanSpeed(steps, height, duration)
	//Рассчитать и вернуть количество калорий. Для этого:
	//Переведите продолжительность в минуты с помощью функции из пакета time.u
	spentCalories := maenSpeed * weight * duration.Minutes() / minInH
	return spentCalories, nil
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Алгоритм реализации функции:
	//Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и
	//  соответствующую ошибку.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect Input")
	}
	//Рассчитать среднюю скорость с помощью meanSpeed().
	maenSpeed := meanSpeed(steps, height, duration)
	// Рассчитать количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	spentCalories := maenSpeed * weight * duration.Minutes() / minInH * walkingCaloriesCoefficient
	return spentCalories, nil
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient.
	// Соответствующая константа объявлена в пакете. Вернуть полученное значение.
}
