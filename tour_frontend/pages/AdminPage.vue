<template>
  <div class="container">
    <h1>Админ Панель</h1>

    <!-- Раздел для отображения сообщений -->
    <h2>Сообщения от пользователей</h2>
    <ul class="message-list">
      <li v-for="message in messages" :key="message.id" class="message-item">
        <strong>{{ message.name }} ({{ message.email }})</strong>
        <p>{{ message.message }}</p>
        <p><strong>Телефон:</strong> {{ message.phone }}</p>
        <button @click="deleteMessage(message.id)" class="btn btn-delete">Удалить</button>
      </li>
    </ul>

        <!-- Раздел для загрузки фотографий в галерею -->
    <div>

      <h2>Загрузить изображение в галерею</h2>
      <form @submit.prevent="uploadGalleryImage" enctype="multipart/form-data">
        <input type="file" name="galleryImage" @change="onGalleryImageChange" required>
        <button type="submit" class="btn">Загрузить изображение</button>
      </form>

    <div>
      <h2>Галерея</h2>
  <button @click="toggleGallery" class="btn">
    {{ galleryVisible ? 'Свернуть' : 'Развернуть' }}
  </button>
  
  <div v-show="galleryVisible" class="gallery">
    <div v-for="image in galleryImages" :key="image" class="gallery-item">
      <img :src="image" alt="Gallery Image" />
      <button @click="deleteImage(image)">Удалить</button>
    </div>
  </div>
    </div>

  
    </div>




    <!-- Форма для добавления или редактирования тура -->
    <div>
      <h2>{{ editingTour ? 'Редактировать тур' : 'Добавить тур' }}</h2>
      <form @submit.prevent="saveTour" class="admin-form" enctype="multipart/form-data">
        <input v-model="tourForm.name" placeholder="Название тура" required class="form-input" />
        <textarea v-model="tourForm.description" placeholder="Описание тура" class="form-input"></textarea>
        <input type="number" v-model="tourForm.duration" placeholder="Длительность (дни)" class="form-input" />
        <input type="number" v-model="tourForm.price" placeholder="Цена тура" step="0.01" required class="form-input" />
        <input type="file" name="image" @change="uploadFile">
        <div v-for="(day, index) in tourForm.days" :key="index" class="day-entry">
          <h3>День {{ day.dayNumber }}</h3>
          <textarea v-model="day.details" placeholder="Описание дня" class="form-input"></textarea>  <!-- Изменено на day.details -->
          <button @click="removeDay(index)" type="button" class="btn btn-remove">Удалить день</button>
        </div>


        <button @click="addDay" type="button" class="btn btn-add-day">Добавить день</button>

        <button type="submit" class="btn">{{ editingTour ? 'Обновить' : 'Добавить' }} тур</button>
        <button v-if="editingTour" @click="cancelEditTour" class="btn btn-cancel">Отменить</button>
      </form>
    </div>

    <!-- Список существующих туров -->
    <div>
      <h2>Существующие туры</h2>
      <ul class="tour-list">
        <li v-for="tour in tours" :key="tour.id" class="tour-item">
          <h3>{{ tour.name }}</h3>
          <p>{{ tour.description }}</p>
          <p><strong>Длительность:</strong> {{ tour.duration }} дней</p>
          <p><strong>Цена:</strong> {{ tour.price }} ₽</p>
          <button @click="editTour(tour)" class="btn btn-edit">Редактировать</button>
          <button @click="deleteTour(tour.id)" class="btn btn-delete">Удалить</button>
        </li>
      </ul>
    </div>
    
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      tours: [],
      messages: [],
      galleryImages: [],  // Массив для хранения ссылок на изображения галереи
      galleryVisible: true, // По умолчанию галерея видима
      galleryForm: {
        galleryImage: null  // Поле для загружаемого изображения
      },
      currentTourId: null,
      editingTour: false,
      selectedTourId: null,
      tourForm: {
        id: null,
        name: '',
        description: '',
        duration: null,
        price: null,
        days: [],
        image: null
      }
    };

  },



  methods: {

    toggleGallery() {
    this.galleryVisible = !this.galleryVisible;
  },
    // Метод для загрузки файла в галерею
    onGalleryImageChange(event) {
      const file = event.target.files[0];
        if (file) {
          this.galleryForm.galleryImage = file;
          }
    },

    async uploadGalleryImage() {
      try {
        const formData = new FormData();
        formData.append('image', this.galleryForm.galleryImage);  // Изменено 'galleryImage' на 'image'

        await axios.post('/api/gallery/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
          }
        });

        // После загрузки обновляем список изображений
        await this.fetchGalleryImages();
          alert('Изображение успешно загружено!');
        } catch (error) {
          console.error('Error uploading gallery image:', error);
          alert('Ошибка при загрузке изображения.');
          }
    },

    async fetchGalleryImages() {
      try {
        const response = await fetch('/api/gallery');
        const data = await response.json();
          console.log("Fetched images:", data); 
          this.galleryImages = data; 
      } catch (error) {
          console.error('Ошибка при загрузке изображений:', error);
      }
    },


    async deleteImage(image) {
      try {
        const response = await fetch('/api/delete-image', {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ imagePath: image }) // Отправляем путь к изображению
        });
        
        if (response.ok) {
          // Если запрос прошел успешно, обновляем список изображений
          this.galleryImages = this.galleryImages.filter(img => img !== image);
          console.log('Изображение удалено');
        } else {
          console.error('Ошибка при удалении изображения');
        }
      } catch (error) {
          console.error('Ошибка при выполнении запроса на удаление:', error);
      }
    },




    uploadFile(event) {
      const file = event.target.files[0];
      if (file) {
      this.tourForm.image = file; // Сохраняем файл в tourForm.image
      }
    },


    async fetchMessages() {
      try {
        const response = await axios.get('/api/messages');
        this.messages = response.data;
      } catch (error) {
        console.error('Error fetching messages:', error);
      }
    },


    async deleteMessage(id) {
      try {
        await axios.delete(`/api/messages/${id}`);
        await this.fetchMessages();
      } catch (error) {
        console.error('Error deleting message:', error);
      }
    },


    setCurrentTour(tourId) {
      this.currentTourId = tourId;
      this.selectedTourId = tourId;
    },

    async fetchTours() {
      try {
        const response = await axios.get('/api/tours');
        this.tours = response.data;
      } catch (error) {
        console.error('Error fetching tours:', error);
      }
    },

    resetTourForm() {
      this.tourForm = {
        id: null,
        name: '',
        description: '',
        duration: null,
        price: null,
        days: [] // Сброс информации по дням
      };
    },

    editTour(tour) {
      this.editingTour = true;
      this.tourForm = { ...tour, days: tour.days.map(day => ({ ...day, details: day.details || '' })) };
      this.setCurrentTour(tour.id);
    },

    cancelEditTour() {
      this.editingTour = false;
      this.resetTourForm();
    },

    async deleteTour(id) {
      try {
      // Подтверждение перед удалением
      const confirmed = confirm('Вы уверены, что хотите удалить этот тур?');
      if (!confirmed) return;

      // Отправляем DELETE запрос на бэкенд
      await axios.delete(`/api/tours/${id}`, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      // После успешного удаления обновляем список туров
      this.tours = this.tours.filter(tour => tour.id !== id);
      console.log(`Тур с id ${id} был удален.`);
      } catch (error) {
        console.error('Ошибка при удалении тура:', error);
        alert('Произошла ошибка при удалении тура. Попробуйте снова.');
        }
    },


    // Добавление нового дня в тур
    addDay() {
      const newDayNumber = this.tourForm.days.length + 1;  // Новый день будет иметь порядковый номер, увеличенный на 1
      this.tourForm.days.push({ dayNumber: newDayNumber, description: '' });
    },


    // Удаление дня из тура
    removeDay(index) {
      this.tourForm.days.splice(index, 1);
    },

    
    async saveTour() {
  try {
    // Создаем объект FormData для отправки данных формы
    const formData = new FormData();

    // Добавляем основные поля тура
    formData.append('name', this.tourForm.name);
    formData.append('description', this.tourForm.description);
    formData.append('duration', this.tourForm.duration);
    formData.append('price', this.tourForm.price);

    // Добавляем дни тура, проверяя на пустые данные
    if (this.tourForm.days && this.tourForm.days.length > 0) {
      this.tourForm.days.forEach((day, index) => {
        if (day.details) { // Проверка, что день содержит описание
          formData.append(`days[${index}][dayNumber]`, day.dayNumber || index + 1);
          formData.append(`days[${index}][details]`, day.details);
        }
      });
    }

    // Если изображение загружено, добавляем его в formData
    if (this.tourForm.image) {
      formData.append('image', this.tourForm.image);
    }

    // Проверяем, редактируем ли существующий тур или создаем новый
    if (this.editingTour && this.tourForm.id) {
      // Обновляем существующий тур через PUT запрос
      await axios.put(`api/tours/${this.tourForm.id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });
    } else {
      // Создаем новый тур через POST запрос
      await axios.post('/api/tours', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });
    }

    // После успешного сохранения обновляем список туров
    await this.fetchTours();

    // Сбрасываем форму после сохранения
    this.cancelEditTour();

  } catch (error) {
    // Ловим и обрабатываем ошибки
    console.error('Ошибка при сохранении тура:', error);
    alert('Произошла ошибка при сохранении тура. Пожалуйста, проверьте данные и попробуйте снова.');
  }
},




  },
  
  mounted() {
    this.fetchTours();
    this.fetchMessages();
    this.fetchGalleryImages();
  }
};
</script>


<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Roboto', sans-serif;
  color: #333;
  line-height: 1.6;
  background-color: #f4f4f4;
}

.container {
  width: 90%;
  margin: 0 auto;
  max-width: 1200px;
}

/* Форма администратора */
.admin-form {
  margin-bottom: 20px;
}

/* Галерея */
.gallery {
  display: grid;
  grid-template-columns: repeat(5, 1fr); /* 5 изображений в ряд */
  gap: 10px; /* Отступы между изображениями */
}

.gallery-item {
  position: relative;
}

.gallery-item img {
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.gallery-item button {
  position: absolute;
  top: 5px;
  right: 5px;
  padding: 5px;
  background-color: rgba(255, 0, 0, 0.7);
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 4px;
  font-size: 12px;
}

.gallery-item button:hover {
  background-color: rgba(255, 0, 0, 1);
}

/* Формы */
.form-input {
  display: block;
  margin-bottom: 10px;
  padding: 12px;
  width: 100%;
  max-width: 400px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

/* Кнопки */
button {
  padding: 12px 25px;
  margin:10px;
  background-color: #42b983;
  color: white;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
  border: none;
}

button:hover {
  background-color: #369972;
}


.btn-cancel {
  background-color: #e74c3c;
}

.btn-cancel:hover {
  background-color: #c0392b;
}

.btn-delete {
  background-color: #e74c3c;
}

.btn-delete:hover {
  background-color: #c0392b;
}

/* Адаптивность */
@media (max-width: 1024px) {
  .gallery {
    grid-template-columns: repeat(3, 1fr); /* 3 изображения в ряд для планшетов */
  }
}

@media (max-width: 768px) {
  .gallery {
    grid-template-columns: repeat(2, 1fr); /* 2 изображения в ряд для смартфонов */
  }

  button {
    font-size: 14px; /* Уменьшение размера шрифта для кнопок */
    padding: 10px 20px;
  }
}

@media (max-width: 480px) {
  .gallery {
    grid-template-columns: 1fr; /* 1 изображение в ряд для маленьких экранов */
  }

  button {
    font-size: 12px; /* Еще меньше шрифт для мобильных устройств */
    padding: 8px 16px;
  }

  .form-input {
    max-width: 100%; /* Ширина формы на всю ширину экрана */
  }
}

/* Стили для списков туров и сообщений */
ul {
  list-style-type: none;
  padding: 0;
}

.tour-item {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.tour-list {
  margin-top: 30px;
}

.message-list {
  margin-top: 30px;
}

.message-item {
  margin-bottom: 20px;
  padding: 15px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.message-item strong {
  display: block;
  margin-bottom: 5px;
}
</style>

