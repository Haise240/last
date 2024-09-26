<template>
  <div>
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
</template>

<script>
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/swiper-bundle.css';

export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      tour: {
        name: '',
        description: '',
        duration: 0,
        price: 0,
        days: [],
      },
      galleryImages: [],
    };
  },
  async asyncData({ params }) {
    const tourId = params.id; // Получаем ID тура из параметров маршрута
    try {
      const tourResponse = await fetch(`http://localhost:8080/api/tours/${tourId}`);
      const tourData = await tourResponse.json();
      const galleryResponse = await fetch('http://localhost:8080/api/gallery');
      const galleryData = await galleryResponse.json();

      return {
        tour: tourData,
        galleryImages: galleryData.map(image => `${image}`),
      };
    } catch (error) {
      console.error('Ошибка при получении данных:', error);
      return {
        tour: {},
        galleryImages: [],
      };
    }
  },
};
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