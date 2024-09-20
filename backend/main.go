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

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Tour represents a tour in the database
type Tour struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Price       float64        `json:"price"`
	Days        []Day          `json:"days"`      // Новое поле для дней тура
	ImageURL    sql.NullString `json:"image_url"` // Используем sql.NullString

}

type Day struct {
	DayNumber int    `json:"dayNumber"`
	Details   string `json:"details"`
	ID        int    `json:"id"`
	details   string `json:"text"`
}

// Message represents a message from the contact form
type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Phone   string `json:"phone"`
}

func main() {
	// Connect to the database
	connStr := "user=postgres dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create routes using gorilla/mux
	router := mux.NewRouter()

	router.HandleFunc("/api/tours", handleTours(db)).Methods("GET", "POST")
	router.HandleFunc("/api/tours/{id}", handleTour(db)).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/api/messages", handleMessages(db)).Methods("POST", "GET")
	router.HandleFunc("/api/messages/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		deleteMessage(db, w, r)
	}).Methods("DELETE")

	// Новый маршрут для загрузки изображений в галерею
	router.HandleFunc("/api/gallery/upload", func(w http.ResponseWriter, r *http.Request) {
		uploadGalleryImage(w, r, db)
	}).Methods("POST")

	// Новый маршрут для получения списка изображений
	router.HandleFunc("/api/gallery", getGalleryImages).Methods("GET")
	router.HandleFunc("/api/delete-image", deleteImageHandler).Methods("DELETE")

	// Настройка сервера для отдачи статических файлов
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Apply CORS middleware
	corsEnabledMux := enableCors(router)
	log.Printf("Current working directory: %s", getCurrentDirectory())
	// Start server
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

//ТУРЫ

func handleTours(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTours(db, w, r)
		case http.MethodPost:
			createTourHandler(db)(w, r) // Передаем db в createTourHandler
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// handleTour handles GET, PUT, and DELETE requests for a single tour
func handleTour(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		switch r.Method {
		case http.MethodGet:
			getTour(db, w, r, id)
		case http.MethodPut:
			updateTour(db, w, r, id)
		case http.MethodDelete:
			deleteTour(db, w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getTours(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, duration, price, days, image_url FROM tours")
	if err != nil {
		http.Error(w, "Ошибка выполнения запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tours := []Tour{}
	for rows.Next() {
		var tour Tour
		var daysData []byte

		if err := rows.Scan(&tour.ID, &tour.Name, &tour.Description, &tour.Duration, &tour.Price, &daysData, &tour.ImageURL); err != nil {
			http.Error(w, "Ошибка чтения данных из базы: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if len(daysData) == 0 {
			tour.Days = []Day{}
		} else {
			if err := json.Unmarshal(daysData, &tour.Days); err != nil {
				http.Error(w, "Ошибка декодирования данных: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Удаляем начальную точку из пути к изображению, если она присутствует
		imagePath := strings.TrimPrefix(tour.ImageURL.String, ".")

		// Формирование правильного URL для изображений
		if tour.ImageURL.Valid {
			tour.ImageURL.String = "http://" + r.Host + imagePath
		}

		tours = append(tours, tour)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Ошибка при итерации по строкам: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tours); err != nil {
		http.Error(w, "Ошибка кодирования ответа в JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func createTourHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ограничение размера запроса (например, до 10 МБ)
		r.ParseMultipartForm(10 << 20)

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

		// Сохраняем изображение на диск
		imagePath, err := saveImageTour(file, header, "./static/uploads")
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}

		// Получаем данные о днях из формы
		daysData := r.FormValue("days") // Предполагаем, что данные о днях передаются как JSON строка

		// Начинаем транзакцию для атомарных операций
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Вставляем основной тур с изображением
		var tourID int
		err = tx.QueryRow(
			"INSERT INTO tours (name, description, duration, price, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			name, description, duration, price, imagePath,
		).Scan(&tourID)

		if err != nil {
			tx.Rollback() // Откатываем транзакцию в случае ошибки
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Обрабатываем и вставляем дни
		var days []struct {
			DayNumber int    `json:"day_number"`
			Details   string `json:"details"`
		}
		if err := json.Unmarshal([]byte(daysData), &days); err != nil {
			tx.Rollback() // Откатываем транзакцию в случае ошибки
			http.Error(w, "Invalid days data: "+err.Error(), http.StatusBadRequest)
			return
		}

		for i, day := range days {
			day.DayNumber = i + 1 // Начинаем нумерацию с 1
			_, err := tx.Exec(
				"INSERT INTO tour_days (tour_id, day_number, details) VALUES ($1, $2, $3)",
				tourID, day.DayNumber, day.Details,
			)
			if err != nil {
				tx.Rollback() // Откатываем транзакцию в случае ошибки
				http.Error(w, "Error inserting day data: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Коммитим транзакцию
		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Tour created successfully")
	}
}

func getTour(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) {
	var tour Tour
	row := db.QueryRow("SELECT id, name, description, duration, price FROM tours WHERE id=$1", id)
	if err := row.Scan(&tour.ID, &tour.Name, &tour.Description, &tour.Duration, &tour.Price); err != nil {
		http.Error(w, "Tour not found", http.StatusNotFound)
		return
	}

	// Извлекаем дни тура
	rows, err := db.Query("SELECT day_number, details FROM tour_days WHERE tour_id = $1 ORDER BY day_number", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var days []Day
	for rows.Next() {
		var day Day
		if err := rows.Scan(&day.DayNumber, &day.Details); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		days = append(days, day)
	}
	tour.Days = days

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tour)
}

// Update a tour
func updateTour(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) {
	r.ParseMultipartForm(10 << 20) // Ограничение размера на 10 МБ

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

	// Получаем информацию по дням тура в формате JSON
	days := r.FormValue("days")
	var tourDays []Day
	err = json.Unmarshal([]byte(days), &tourDays)
	if err != nil {
		http.Error(w, "Invalid days format", http.StatusBadRequest)
		return
	}

	// Проверяем, загрузил ли пользователь новое изображение
	file, header, err := r.FormFile("image")
	var imagePath string
	if err == nil {
		// Сохраняем новое изображение на диск
		imagePath, err = saveImageTour(file, header, "./static/uploads")
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	}

	// Обновляем тур с новым изображением
	var query string
	if imagePath != "" {
		query = "UPDATE tours SET name=$1, description=$2, duration=$3, price=$4, image_url=$5 WHERE id=$6"
		_, err = db.Exec(query, name, description, duration, price, imagePath, id)
	} else {
		query = "UPDATE tours SET name=$1, description=$2, duration=$3, price=$4 WHERE id=$5"
		_, err = db.Exec(query, name, description, duration, price, id)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Удаляем существующие дни тура
	_, err = db.Exec("DELETE FROM tour_days WHERE tour_id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete existing tour days", http.StatusInternalServerError)
		return
	}

	// Добавляем новые дни тура
	for _, day := range tourDays {
		_, err = db.Exec("INSERT INTO tour_days (tour_id, day_number, details) VALUES ($1, $2, $3)", id, day.DayNumber, day.Details)
		if err != nil {
			http.Error(w, "Failed to insert tour days", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete a tour
func deleteTour(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) {
	// Выполняем запрос на удаление тура
	result, err := db.Exec("DELETE FROM tours WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Ошибка при удалении тура: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Проверяем, был ли удалён тур
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Ошибка при проверке результата удаления: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Тур с таким ID не найден", http.StatusNotFound)
		return
	}

	// Отправляем успешный ответ, если удаление прошло успешно
	w.WriteHeader(http.StatusNoContent)
}

func saveImageTour(file multipart.File, header *multipart.FileHeader, uploadPath string) (string, error) {
	defer file.Close()

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

	log.Printf("File saved successfully: %s\n", filePath)
	return filePath, nil
}

//СООБЩЕНИЯ

// handleMessages handles POST, GET, and DELETE requests for messages
func handleMessages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postMessage(db, w, r)
		case http.MethodGet:
			getMessages(db, w, r)
		case http.MethodDelete:
			deleteMessage(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// Пример запроса на получение всех сообщений
func getMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email, message, phone FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Name, &msg.Email, &msg.Message, &msg.Phone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// Post a message
func postMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("Error decoding message:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received message: %+v\n", msg)
	_, err := db.Exec("INSERT INTO messages (name, email, message, phone) VALUES ($1, $2, $3, $4)", msg.Name, msg.Email, msg.Message, msg.Phone)
	if err != nil {
		log.Println("Error inserting message:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
	_, err = db.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
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

func uploadGalleryImage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	// Записываем путь изображения в таблицу images
	_, err = db.Exec("INSERT INTO images (filepath) VALUES ($1)", imagePath)
	if err != nil {
		log.Printf("Error inserting image record into database: %s\n", err)
		http.Error(w, "Failed to save image record in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Image uploaded successfully")
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
	println("все отдал!")
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
