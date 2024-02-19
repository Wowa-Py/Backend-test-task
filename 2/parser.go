package parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"log"
	"os"
)

err := godotenv.Load()
if err != nil {
	log.Fatal("Ошибка загрузки файла .env")
}


func main() {
	// Открываем файл для записи данных
	file, err := os.Create("instagram_data.csv")
	if err != nil {
		log.Fatal("Не удалось создать файл:", err)
	}
	defer file.Close()

	// Создаем записчик CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовки столбцов
	headers := []string{"Рейтинг", "Имя", "Ник", "Категория", "Подписчики", "Аудитория", "Подлинность", "Вовлеченность"}
	err = writer.Write(headers)
	if err != nil {
		log.Fatal("Ошибка записи заголовков:", err)
	}

	// URL страницы для парсинга
	url := os.Getenv("URL")

	// Отправляем GET-запрос на страницу
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка при отправке GET-запроса:", err)
	}
	defer response.Body.Close()

	// Загружаем HTML-документ
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Ошибка при загрузке HTML-документа:", err)
	}

	// Парсим данные
	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		// Пропускаем первую строку, так как это заголовки столбцов
		if i == 0 {
			return
		}

		// Извлекаем данные из каждой ячейки строки
		rank := s.Find(".rank span").Text()
		name := s.Find(".contributor__name-content").Text()
		nickname := s.Find(".contributor__name-content").Next().Text()
		category := s.Find(".tag__content").First().Text()
		subscribers := s.Find(".subscribers").Text()
		audience := s.Find(".audience").Text()
		authentic := s.Find(".authentic").Text()
		engagement := s.Find(".engagement").Text()

		// Записываем данные в CSV файл
		data := []string{rank, name, nickname, category, subscribers, audience, authentic, engagement}
		err := writer.Write(data)
		if err != nil {
			log.Fatal("Ошибка записи данных:", err)
		}
	})

	fmt.Println("Парсинг завершен. Данные сохранены в instagram_data.csv")
}
