
<template>
  <div>
    <!-- Header Section -->
    <header class="page-header">
      <div class="container">
        <h1>Наши туры</h1>
        <p>
          Выберите тур, который подходит именно вам, и отправляйтесь в незабываемое путешествие.
        </p>
      </div>
    </header>

    <!-- Filters and Sorting Section -->
    <section class="filters">
      <div class="container">
        <div class="filter-group">
          <label for="duration">Длительность:</label>
          <select id="duration" v-model="filters.duration">
            <option value="">Любая</option>
            <option value="1-3">1-3 дня</option>
            <option value="4-7">4-7 дней</option>
            <option value="8+">8+ дней</option>
          </select>
        </div>

        <div class="filter-group">
          <label for="price">Цена:</label>
          <select id="price" v-model="filters.priceRange">
            <option value="">Любая</option>
            <option value="0-20000">до 20,000 руб.</option>
            <option value="20000-50000">20,000 - 50,000 руб.</option>
            <option value="50000+">более 50,000 руб.</option>
          </select>
        </div>

        <div class="filter-group">
          <label for="sort">Сортировка:</label>
          <select id="sort" v-model="sortOrder">
            <option value="price-asc">Цена: по возрастанию</option>
            <option value="price-desc">Цена: по убыванию</option>
            <option value="duration-asc">Длительность: по возрастанию</option>
            <option value="duration-desc">Длительность: по убыванию</option>
          </select>
        </div>
      </div>
    </section>

    <!-- Tours Catalog Section -->
    <section class="tours-catalog">
      <div class="container">
        <div v-if="loading">Загрузка туров...</div>
        <div v-else-if="filteredAndSortedTours.length === 0">Нет доступных туров по выбранным фильтрам.</div>
        <div v-else class="tour-cards">
          <div v-for="tour in filteredAndSortedTours" :key="tour.id" class="tour-card">
            <img :src="formatImageUrl(tour.image_url)" :alt="tour.name" />
            <h1 class="tour-title">{{ tour.name }}</h1>
            <p class="tour-info">
              <span>Длительность: {{ tour.duration }} дней </span>
              <span class="price">Цена: {{ tour.price }} руб.</span>
            </p>
            <nuxt-link :to="`/tour/${tour.id}`" class="btn">Подробнее</nuxt-link>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ToursPage',
  data() {
    return {
      filters: {
        duration: '',
        priceRange: ''
      },
      sortOrder: 'price-asc',
      tours: [], // Массив туров
      loading: true // Индикатор загрузки
    };
  },

  computed: {
    filteredAndSortedTours() {
      let filteredTours = this.tours;

      // Фильтрация по длительности
      if (this.filters.duration) {
        filteredTours = filteredTours.filter((tour) => {
          const duration = tour.duration;
          if (this.filters.duration === '1-3') return duration <= 3;
          if (this.filters.duration === '4-7') return duration >= 4 && duration <= 7;
          if (this.filters.duration === '8+') return duration >= 8;
          return true;
        });
      }

      // Фильтрация по ценовому диапазону
      if (this.filters.priceRange) {
        filteredTours = filteredTours.filter((tour) => {
          const price = tour.price;
          if (this.filters.priceRange === '0-20000') return price <= 20000;
          if (this.filters.priceRange === '20000-50000') return price > 20000 && price <= 50000;
          if (this.filters.priceRange === '50000+') return price > 50000;
          return true;
        });
      }

      // Теперь просто возвращаем отсортированный массив без дополнительной сортировки
      return filteredTours;
    }

  },

  methods: {
    fetchTours() {
      axios
        .get('http://localhost:8080/api/tours')
        .then((response) => {
          console.log('Fetched Tours:', response.data); // Выводим полученные данные
          // Используем непосредственно строку image_url
          this.tours = response.data;
          this.loading = false;
        })
        .catch((error) => {
          console.error('Ошибка при получении туров:', error);
          this.loading = false;
        });
    },

    formatImageUrl(imageURL) {
      // Если URL пустой, используем изображение-заглушку
      if (imageURL && !imageURL.startsWith('http')) {
        return `http://localhost:8080/static/uploads/${imageURL}`;
      }
      return imageURL || '1.jpg'; // Фоллбэк, если imageURL пустой
    }

  },

  mounted() {
    this.fetchTours();
  }
}
</script>


<style scoped>
/* Общие стили */
body {
  margin: 0;
  font-family: Arial, sans-serif;
}

/* Стили для заголовка страницы */
.page-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
}

.page-header h1 {
  font-size: 2.5em;
  margin-bottom: 10px;
}

.page-header p {
  font-size: 1.2em;
  margin: 0;
}

/* Стили для контейнера */
.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}

/* Стили для секции фильтров */
.filters {
  background: #f9f9f9;
  padding: 20px 0;
  border-bottom: 1px solid #ddd;
}

.filter-group {
  display: inline-block;
  margin-right: 15px;
}

.filter-group label {
  margin-right: 10px;
}

.filter-group select {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 5px;
}


/* Стили для секции каталога туров */
.tours-catalog {
  padding: 40px 0;
}

.tour-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: space-between;
}

.tour-card {
  background: #fafff5;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  flex-basis: calc(33.333% - 20px);
  display: flex;
  flex-direction: column; /* Вертикальная структура */
  text-align: center;
}

.tour-card img {
  width: 100%;
  height: 180px;
  object-fit: cover;
  background-color: #f0f0f0;
  border-bottom: 1px solid #ddd;
}

.tour-card-content {
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  padding: 15px;
}

.tour-title {
  font-size: 1.2em;
  font-weight: bold;
  color: #333;
  margin: 15px 0;
}

.tour-info {
  font-size: 1em;
  color: #666;
  margin-bottom: 15px;
}

.btn {
  max-width: 70%;
  margin: auto ; /* Выровнять по центру с верхним отступом */
  margin-bottom: 10px;
  padding: auto ;
  background: #42b983;
  color: white;
  text-decoration: none;
  border-radius: 10px;
  font-size: 1em;
  display: inline-block;
}

.btn:hover {
  background: #369972;
}

/* Адаптивные стили */
@media (max-width: 1024px) {
  .tour-card {
    flex-basis: calc(50% - 20px);
  }
}

@media (max-width: 768px) {
  .tour-card {
    flex-basis: calc(100% - 20px);
  }

  .page-header h1 {
    font-size: 2em;
  }

  .page-header p {
    font-size: 1em;
  }

  .filter-group {
    display: block;
    margin-bottom: 10px;
  }

  .filter-group label {
    display: block;
    margin-bottom: 5px;
  }

  .filter-group select {
    width: 100%;
  }
}
</style>
