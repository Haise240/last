<template>
  <div v-if="!loading && tour">
    <header class="page-header">
      <h1 class="tour-title">{{ tour.name }}</h1>
    </header>
    <div class="tour-page">
      <p class="tour-description">{{ tour.description }}</p>

      <div class="tour-info">
        <p>Длительность: <span class="info-value">{{ tour.duration }} дня</span></p> 
        <p class="tour-price">Цена: <span class="info-value">{{ tour.price }} руб.</span></p>
      </div>

      <div class="tour-gallery">
        <h2>Фотографии из туров</h2>
        <swiper
          :slides-per-view="3"
          :space-between="30"
          pagination
          navigation
          loop
        >
          <swiper-slide v-for="(image, index) in galleryImages" :key="index">
            <img :src="image" alt="Tour Image" class="gallery-image" />
          </swiper-slide>
        </swiper>
      </div>

      <div class="tour-days">
        <h2>Программа тура по дням</h2>
        <div v-for="(day, index) in tour.days" :key="index" class="tour-day">
          <div class="day-number">День {{ day.dayNumber }}</div>
          <div class="day-details">{{ day.details }}</div>
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <p v-if="loading">Загрузка данных...</p>
    <p v-else-if="error">Ошибка загрузки данных: {{ error.message }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/swiper-bundle.css';
import { useRoute } from '#app';

const route = useRoute();
const tourId = route.params.id;

// Состояния загрузки и ошибок
const loading = ref(true);
const error = ref(null);
const tour = ref(null); 
const galleryImages = ref([]);

async function fetchTourData() {
  loading.value = true;
  error.value = null;

  try {
    // Запрос на получение данных о туре
    const { data: tourData } = await useFetch(`http://localhost:8080/api/tours/${tourId}`);
    if (!tourData.value) {
      throw new Error('Данные о туре не найдены');
    }
    tour.value = tourData.value;

    // Запрос на получение галереи
    const { data: galleryData } = await useFetch('http://localhost:8080/api/gallery');
    if (galleryData.value) {
      galleryImages.value = galleryData.value;
    }
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
}

// Вызов функции для получения данных
fetchTourData();
</script>




<style scoped>
/* Стили для всей страницы */
body {
  margin: 0;
  font-family: 'Roboto', sans-serif;
}

.page-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
}

.tour-title {
  margin: 0;
  font-size: 2.5em;
}

.tour-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.tour-description {
  text-align: justify;
  margin-bottom: 30px;
  font-size: 1.2em;
}

.tour-info {
  display: flex;
  justify-content: space-between;
  font-size: 1.2em;
  margin-bottom: 20px;
  background-color: #f9f9f9;
  padding: 15px;
  border-radius: 8px;
}

.tour-info .info-value {
  font-weight: bold;
}

/* Стили для слайдера */
.swiper {
  width: 100%;
  height: auto; /* Автоматическая высота */
  margin-bottom: 40px;
}

.gallery-image {
  width: 100%;
  height: 300px; /* Задайте высоту для всех изображений */
  object-fit: cover;
}

.tour-gallery h2 {
  margin: 0 0 20px;
  font-size: 1.5em;
}

.tour-days {
  margin-top: 30px;
}

.tour-day {
  margin-bottom: 20px;
  padding: 20px;
  background-color: #f7f7f7;
  border-left: 5px solid #333;
  border-radius: 8px;
}

.day-number {
  font-size: 1.4em;
  font-weight: bold;
  color: #555;
}

.day-details {
  margin-top: 10px;
  font-size: 1.1em;
  line-height: 1.4;
  color: #666;
}

@media (max-width: 768px) {
  .page-header {
    padding: 15px 0;
  }

  .tour-title {
    font-size: 2em;
  }

  .tour-description {
    font-size: 1.1em;
  }

  .tour-info {
    font-size: 1em;
    flex-direction: column;
    align-items: flex-start;
  }

  .tour-info p {
    margin-bottom: 10px;
  }

  .day-number {
    font-size: 1.2em;
  }

  .day-details {
    font-size: 1em;
  }
}
</style>
