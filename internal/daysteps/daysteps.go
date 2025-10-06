package daysteps

import (
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

//Алгоритм реализации функции:

func parsePackage(data string) (int, time.Duration, error) {
	//Разделить строку на слайс строк.
	//Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	parsedData := strings.Split(data, ",")
	if len(parsedData) != 2 {
		return 0, 0, fmt.Errorf("incorrect Input")
	}

	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	// Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.
	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, 0, err
	} else if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше 0")
	}

	// Преобразовать второй элемент слайса в time.Duration.
	// В пакете time есть метод для парсинга строки в time.Duration. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	// Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	//
	activTime, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return 0, 0, err
	} else if activTime <= 0 {
		return 0, 0, fmt.Errorf("Колличество минут должно быть больше 0")
	}
	return steps, activTime, nil

}

//Алгоритм реализации функции:

func DayActionInfo(data string, weight, height float64) string {

	// Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage().
	// В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	// Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.
	steps, activTime, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага.
	// Константа stepLength (длина шага) уже определена в коде.

	distance := float64(steps) * stepLength / mInKm

	// Вычислить количество калорий, потраченных на прогулке.
	// Функция для вычисления калорий WalkingSpentCalories() будет определена в пакете spentcalories,
	// которую вы тоже реализуете.
	// Сформировать строку, которую будете возвращать, пример которой был представлен выше.
	activSpentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, activTime)
	if err != nil {
		log.Println(err)
		return ""
	}

	outputFormat := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
	output := fmt.Sprintf(outputFormat, steps, distance, activSpentCalories)

	return output
}
