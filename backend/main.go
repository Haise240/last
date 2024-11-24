package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Tour struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Duration     int            `json:"duration"`
	Price        float64        `json:"price"`
	Days         []Day          `json:"days"`
	ImageURL     sql.NullString `json:"image_url"`
	DisplayOrder int            `json:"display_order"`
}

type Day struct {
	ID        int    `json:"id"`
	TourID    int    `json:"tour_id"`
	DayNumber int    `json:"day_number"`
	Details   string `json:"details"`
}

// Message represents a message from the contact form
type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Phone   string `json:"phone"`
}

type Image struct {
	ID       uint   `gorm:"primaryKey"`                 // ID изображения
	Filepath string `gorm:"type:varchar(255);not null"` // Путь к файлу изображения
}

func main() {
	// Подключение к базе данных с использованием GORM
	connStr := "user=postgres dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Роуты для туров
	router.HandleFunc("/api/tours", handleTours(db)).Methods("GET", "POST")
	router.HandleFunc("/api/tours/{id}", handleTour(db)).Methods("GET", "PUT", "DELETE")

	// Роуты для сообщений
	router.HandleFunc("/api/messages", handleMessages(db)).Methods("POST", "GET")
	router.HandleFunc("/api/messages/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		deleteMessage(db, w, r)
	}).Methods("DELETE")

	// Роуты для галереи изображений
	router.HandleFunc("/api/gallery/upload", func(w http.ResponseWriter, r *http.Request) {
		uploadGalleryImage(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/api/gallery", getGalleryImages).Methods("GET")
	router.HandleFunc("/api/delete-image", deleteImageHandler).Methods("DELETE")

	// Статические файлы
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Применение CORS middleware
	corsEnabledMux := enableCors(router)

	// Логирование текущей рабочей директории
	log.Printf("Current working directory: %s", getCurrentDirectory())

	// Запуск сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsEnabledMux))
}

// CORS middleware for allowing requests from other origins
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ТУРЫ

func handleTours(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTours(db, w, r)
		case http.MethodPost:
			createTourHandler(db)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// handleTour handles GET, PUT, and DELETE requests for a single tour
func handleTour(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		switch r.Method {
		case http.MethodGet:
			getTour(db, w, r, id)
		case http.MethodPut:
			updateTourHandler(db, w, r, id) // Передаем w, r и id
		case http.MethodDelete:
			deleteTour(db, w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getTours(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Определяем переменную для хранения результатов
	var tours []Tour

	// Выполняем запрос с помощью Raw
	if err := db.Raw("SELECT id, name, description, duration, price, days, image_url, display_order FROM tours ORDER BY display_order ASC").Scan(&tours).Error; err != nil {
		http.Error(w, "Ошибка выполнения запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Обрабатываем поле Days и формируем правильные URL для изображений
	for i := range tours {
		var daysData []byte

		// Обработка поля Days
		if len(tours[i].Days) == 0 {
			tours[i].Days = []Day{}
		} else {
			if err := json.Unmarshal(daysData, &tours[i].Days); err != nil {
				http.Error(w, "Ошибка декодирования данных: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Удаляем начальную точку из пути к изображению, если она присутствует
		if tours[i].ImageURL.Valid {
			imagePath := strings.TrimPrefix(tours[i].ImageURL.String, ".")
			// Формирование правильного URL для изображений
			tours[i].ImageURL.String = "http://" + r.Host + imagePath
		}
	}

	// Устанавливаем заголовок и кодируем ответ в JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tours); err != nil {
		http.Error(w, "Ошибка кодирования ответа в JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func createTourHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ограничение размера запроса (например, до 10 МБ)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Получаем поля формы
		name := r.FormValue("name")
		description := r.FormValue("description")
		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			http.Error(w, "Invalid duration", http.StatusBadRequest)
			return
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}

		// Получаем файл изображения из формы
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Image upload failed", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Сохраняем изображение на диск
		imagePath, err := saveImageTour(file, header, "static/uploads")
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}

		// Начинаем транзакцию для атомарных операций
		tx := db.Begin()

		// Проверка на наличие ошибок при начале транзакции
		if tx.Error != nil {
			http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
			return
		}
		imageURL := sql.NullString{
			String: imagePath,
			Valid:  imagePath != "",
		}

		// Вставляем основной тур с изображением
		tour := Tour{
			Name:        name,
			Description: description,
			Duration:    duration,
			Price:       price,
			ImageURL:    imageURL, // Здесь мы используем imagePath
		}
		if err := tx.Create(&tour).Error; err != nil { // Используем Create для вставки
			tx.Rollback() // Откатываем транзакцию в случае ошибки
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Обрабатываем дни тура, переданные как отдельные поля формы
		i := 0
		for {
			// Получаем данные дня
			dayNumber := r.FormValue(fmt.Sprintf("days[%d][dayNumber]", i))
			details := r.FormValue(fmt.Sprintf("days[%d][details]", i))

			// Если день отсутствует, выходим из цикла
			if dayNumber == "" || details == "" {
				break
			}

			// Преобразуем день в число
			dayNum, err := strconv.Atoi(dayNumber)
			if err != nil {
				tx.Rollback() // Откатываем транзакцию в случае ошибки
				http.Error(w, "Invalid day number", http.StatusBadRequest)
				return
			}

			// Вставляем день в базу данных
			if err = tx.Exec(
				"INSERT INTO tour_days (tour_id, day_number, details) VALUES (?, ?, ?)",
				tour.ID, dayNum, details,
			).Error; err != nil {
				tx.Rollback() // Откатываем транзакцию в случае ошибки
				http.Error(w, "Error inserting day data: "+err.Error(), http.StatusInternalServerError)
				return
			}

			i++ // Переходим к следующему дню
		}

		// Коммитим транзакцию
		if err := tx.Commit().Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Tour created successfully")
	}
}

func getTour(db *gorm.DB, w http.ResponseWriter, r *http.Request, id string) {
	var tour Tour

	// Получаем основной тур по ID
	if err := db.First(&tour, id).Error; err != nil {
		http.Error(w, "Tour not found", http.StatusNotFound)
		return
	}

	// Извлекаем дни тура
	var days []Day
	if err := db.Table("tour_days").Where("tour_id = ?", tour.ID).Order("day_number").Find(&days).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tour.Days = days

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tour); err != nil {
		http.Error(w, "Ошибка кодирования ответа в JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func updateTourHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request, id string) {
	log.Println("Starting updateTourHandler")

	// Попытка конвертировать строковый ID в uint
	tourID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Invalid tour ID: %s", id)
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}
	uintTourID := uint(tourID)
	log.Printf("Parsed tour ID: %d", uintTourID)

	// Проверка типа контента
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") {
		log.Printf("Invalid Content-Type: %s", contentType)
		http.Error(w, "Content-Type must be multipart/form-data", http.StatusBadRequest)
		return
	}

	// Парсинг формы
	if err := r.ParseMultipartForm(10 << 20); err != nil { // Лимит 10MB
		log.Printf("Failed to parse multipart form: %v", err)
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Чтение полей формы
	name := r.FormValue("name")
	description := r.FormValue("description")
	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		log.Printf("Invalid duration value: %v", err)
		http.Error(w, "Invalid duration value", http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Printf("Invalid price value: %v", err)
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}
	log.Printf("Received data - Name: %s, Description: %s, Duration: %d, Price: %.2f", name, description, duration, price)

	// Обработка изображения
	var imagePath string
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		uploadDir := filepath.Join("static", "uploads")
		imagePath, err = saveImageTour(file, header, uploadDir)
		if err != nil {
			log.Printf("Failed to save image: %v", err)
			http.Error(w, "Failed to save image: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Image saved to path: %s", imagePath)
	} else if err != http.ErrMissingFile {
		log.Printf("Failed to process image: %v", err)
		http.Error(w, "Failed to process image", http.StatusBadRequest)
		return
	} else {
		log.Println("No new image uploaded")
	}

	// Чтение и проверка данных о днях тура
	daysData := r.FormValue("days")
	var tourDays []Day
	if daysData != "" {
		if err := json.Unmarshal([]byte(daysData), &tourDays); err != nil {
			log.Printf("Invalid days format: %v", err)
			http.Error(w, "Invalid days format", http.StatusBadRequest)
			return
		}
	}

	// Начало транзакции
	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("Failed to begin transaction: %v", tx.Error)
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Получение существующего тура
	var tour Tour
	if err := tx.First(&tour, uintTourID).Error; err != nil {
		log.Printf("Tour not found: %d", uintTourID)
		tx.Rollback()
		http.Error(w, "Tour not found", http.StatusNotFound)
		return
	}
	log.Printf("Updating tour: %+v", tour)

	// Обновление полей тура
	tour.Name = name
	tour.Description = description
	tour.Duration = duration
	tour.Price = price
	if imagePath != "" {
		tour.ImageURL = sql.NullString{
			String: "/" + imagePath,
			Valid:  true,
		}
	}

	// Сохранение обновленного тура
	if err := tx.Save(&tour).Error; err != nil {
		log.Printf("Failed to save updated tour: %v", err)
		tx.Rollback()
		http.Error(w, "Failed to update tour", http.StatusInternalServerError)
		return
	}
	log.Println("Tour updated successfully")

	// Обработка и обновление дней тура
	if daysData != "" {
		// Получаем текущие дни тура из таблицы "tour_days"
		var existingDays []Day
		if err := tx.Table("tour_days").Where("tour_id = ?", uintTourID).Find(&existingDays).Error; err != nil {
			log.Printf("Failed to fetch existing days: %v", err)
			tx.Rollback()
			http.Error(w, "Failed to fetch existing tour days", http.StatusInternalServerError)
			return
		}

		// Сбор существующих дней в карту по ID
		existingIDs := make(map[uint]Day)
		for _, existingDay := range existingDays {
			existingIDs[uint(existingDay.ID)] = existingDay
		}

		// Обновляем или добавляем новые дни
		for _, newDay := range tourDays {
			newDay.TourID = int(uintTourID)
			if newDay.ID == 0 {
				// Новый день
				if err := tx.Table("tour_days").Create(&newDay).Error; err != nil {
					log.Printf("Failed to create day: %v", err)
					tx.Rollback()
					http.Error(w, "Failed to add tour day", http.StatusInternalServerError)
					return
				}
			} else {
				// Обновление существующего дня
				if _, found := existingIDs[uint(newDay.ID)]; found {
					if err := tx.Table("tour_days").Save(&newDay).Error; err != nil {
						log.Printf("Failed to update day: %v", err)
						tx.Rollback()
						http.Error(w, "Failed to update tour day", http.StatusInternalServerError)
						return
					}
					// Удаляем обновлённый день из карты, чтобы в конце удалить только лишние
					delete(existingIDs, uint(newDay.ID))
				}
			}
			log.Printf("Day processed: %+v", newDay)
		}

		// Удаляем дни, которые отсутствуют в новом списке
		for _, obsoleteDay := range existingIDs {
			if err := tx.Table("tour_days").Delete(&obsoleteDay).Error; err != nil {
				log.Printf("Failed to delete day: %v", err)
				tx.Rollback()
				http.Error(w, "Failed to update tour days", http.StatusInternalServerError)
				return
			}
		}
	}

	// Коммит транзакции
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Tour updated and transaction committed successfully")

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Tour updated successfully")
}

func deleteTour(db *gorm.DB, w http.ResponseWriter, r *http.Request, id string) {
	// Конвертируем ID тура в целое число
	tourID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	// Удаляем тур из базы данных
	if err := db.Delete(&Tour{}, tourID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func saveImageTour(file multipart.File, header *multipart.FileHeader, uploadDir string) (string, error) {
	// Создаем директорию, если она не существует
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Формируем уникальное имя файла, чтобы избежать конфликтов
	uniqueFileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	imagePath := filepath.Join(uploadDir, uniqueFileName)

	// Создаем файл
	out, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Копируем содержимое загруженного файла в созданный файл
	if _, err := io.Copy(out, file); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Возвращаем относительный путь к изображению для использования в URL
	relativePath := "/" + filepath.ToSlash(imagePath)
	// Преобразуем в UNIX-совместимый путь
	return relativePath, nil
}

//СООБЩЕНИЯ

func handleMessages(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postMessage(db, w, r) // Используем GORM
		case http.MethodGet:
			getMessages(db, w, r) // Используем GORM
		case http.MethodDelete:
			deleteMessage(db, w, r) // Используем GORM
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// Пример запроса на получение всех сообщений с использованием GORM
func getMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var messages []Message

	// Используем метод GORM для выполнения SQL-запроса
	if err := db.Find(&messages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Ошибка кодирования в JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func postMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var msg Message

	// Декодируем тело запроса в структуру Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("Error decoding message:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received message: %+v\n", msg)

	// Вставляем сообщение в базу данных с помощью GORM
	if err := db.Create(&msg).Error; err != nil {
		log.Println("Error inserting message:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
}

func deleteMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		http.Error(w, "Message ID is required", http.StatusBadRequest)
		return
	}

	// Конвертируем идентификатор в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	// Удаляем сообщение из базы данных
	if err := db.Delete(&Message{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//ГАЛЕРЕЯ

func saveImageGallery(file multipart.File, header *multipart.FileHeader) (string, error) {
	defer file.Close()

	// Путь для сохранения файлов в "static/gallery"
	uploadPath := "static/gallery"

	// Проверяем, существует ли директория, если нет — создаем
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err = os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory: %s\n", err)
			return "", err
		}
	}

	// Создаем путь для файла
	filePath := fmt.Sprintf("%s/%s", uploadPath, header.Filename)

	// Открываем файл для записи
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file: %s\n", err)
		return "", err
	}
	defer out.Close()

	// Копируем файл
	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error copying file: %s\n", err)
		return "", err
	}

	log.Printf("Gallery image saved successfully: %s\n", filePath)
	return filePath, nil
}

func uploadGalleryImage(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Ограничиваем размер запроса (например, до 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Image upload failed", http.StatusBadRequest)
		return
	}

	// Получаем файл изображения из формы
	file, header, err := r.FormFile("image")
	if err != nil {
		log.Println("Error getting form file:", err)
		http.Error(w, "Image upload failed", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Сохраняем изображение в галерею
	imagePath, err := saveImageGallery(file, header)
	if err != nil {
		log.Println("Error saving image to gallery:", err)
		http.Error(w, "Failed to save gallery image", http.StatusInternalServerError)
		return
	}

	// Создаем запись изображения для таблицы images
	imageRecord := Image{
		Filepath: imagePath,
	}

	// Вставляем запись в базу данных с использованием метода Create
	if err := db.Create(&imageRecord).Error; err != nil {
		log.Printf("Error inserting image record into database: %s\n", err)
		http.Error(w, "Failed to save image record in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Изображение загружено успешно")
}

// Обработчик для получения списка изображений галереи
func getGalleryImages(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	files, err := os.ReadDir("static/gallery")
	if err != nil {
		http.Error(w, "Unable to read gallery directory", http.StatusInternalServerError)
		return
	}

	imageURLs := []string{}
	for _, file := range files {
		if !file.IsDir() {
			imageURLs = append(imageURLs, "http://"+host+"/static/gallery/"+file.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imageURLs)
	println("Галерея отдана!")
}

// deleteImageHandler обрабатывает запросы на удаление изображений
func deleteImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Разбираем тело запроса
	var request struct {
		ImagePath string `json:"imagePath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неверный формат запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Удаляем префикс URL
	imagePath := strings.TrimPrefix(request.ImagePath, "http://localhost:8080/static/gallery/")

	// Полный путь к файлу
	fullPath := filepath.Join("static", "gallery", imagePath)
	log.Printf("Attempting to delete file at: %s", fullPath)

	// Удаление файла
	if err := os.Remove(fullPath); err != nil {
		http.Error(w, "Не удалось удалить изображение: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Успешное удаление, но без содержимого в ответе
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	return dir
}
