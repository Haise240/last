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

        <!-- Sorting Buttons Section -->
        <div class="sorting-buttons">
          <button @click="setSortOrder('price-asc')" :class="{ active: sortOrder === 'price-asc' }">Цена ↑</button>
          <button @click="setSortOrder('price-desc')" :class="{ active: sortOrder === 'price-desc' }">Цена ↓</button>
          <button @click="setSortOrder('duration-asc')" :class="{ active: sortOrder === 'duration-asc' }">Длительность ↑</button>
          <button @click="setSortOrder('duration-desc')" :class="{ active: sortOrder === 'duration-desc' }">Длительность ↓</button>
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
            <img :src="formatImageUrl(tour.image_url)" :alt="tour.name || 'Изображение тура'" />

            <h1 class="tour-title">{{ tour.name || 'Название не указано' }}</h1>

            <p class="tour-info">
              <span>Длительность: {{ tour.duration || 'не указано' }} дней</span>
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
      sortOrder: 'display-order', // Сортировка по умолчанию
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

      // Сортировка по DisplayOrder
      if (this.sortOrder === 'display-order') {
        filteredTours.sort((a, b) => a.DisplayOrder - b.DisplayOrder);
      } else if (this.sortOrder === 'price-asc') {
        filteredTours.sort((a, b) => a.price - b.price);
      } else if (this.sortOrder === 'price-desc') {
        filteredTours.sort((a, b) => b.price - a.price);
      } else if (this.sortOrder === 'duration-asc') {
        filteredTours.sort((a, b) => a.duration - b.duration);
      } else if (this.sortOrder === 'duration-desc') {
        filteredTours.sort((a, b) => b.duration - a.duration);
      }

      return filteredTours;
    }
  },

  methods: {
    fetchTours() {
      axios
        .get('http://localhost:8080/api/tours')
        .then((response) => {
          this.tours = response.data;
          this.loading = false;
        })
        .catch((error) => {
          console.error('Ошибка при получении туров:', error);
          this.loading = false;
        });
    },

    formatImageUrl(imageURL) {
      if (imageURL && imageURL.Valid) {
        return imageURL.String;
      }
      return '1.jpg'; // Фоллбэк изображение
    },

    setSortOrder(order) {
      this.sortOrder = order;
    }
  },


  mounted() {
    this.fetchTours();
  }
};
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

/* Стили для кнопок фильтрации и сортировки */
.filter-buttons,
.sorting-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-bottom: 10px;
}

.filter-buttons button,
.sorting-buttons button {
  padding: 10px 20px;
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.3s, box-shadow 0.3s;
}

.filter-buttons button:hover,
.sorting-buttons button:hover {
  background-color: #e0e0e0;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.filter-buttons button.active,
.sorting-buttons button.active {
  background-color: #007bff;
  color: white;
  border-color: #007bff;
}

.filter-buttons button.active:hover,
.sorting-buttons button.active:hover {
  background-color: #0056b3;
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
  flex-direction: column;
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
  margin: auto;
  margin-bottom: 10px;
  padding: 10px;
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

/* Адаптивные стили для планшетов */
@media (max-width: 1024px) {
  .tour-card {
    flex-basis: calc(50% - 20px);
  }

  .page-header h1 {
    font-size: 2em;
  }

  .page-header p {
    font-size: 1.1em;
  }
}

/* Адаптивные стили для мобильных устройств */
@media (max-width: 768px) {
  .tour-card {
    flex-basis: 100%;
  }

  .tour-card img {
    height: 150px;
  }

  .page-header h1 {
    font-size: 1.8em;
  }

  .page-header p {
    font-size: 1em;
  }

  .toggle-filters-btn {
    width: 100%;
  }

  .sorting-buttons {
    flex-direction: column;
  }

  .sorting-buttons button {
    width: 100%;
  }

  .tour-title {
    font-size: 1.1em;
  }

  .tour-info {
    font-size: 0.9em;
  }

  .btn {
    padding: 8px 0;
    font-size: 0.9em;
  }
}

/* Дополнительные стили для очень маленьких экранов (телефоны) */
@media (max-width: 480px) {
  .filters{
    margin: 0;
    padding: 20px;
  }
  .page-header h1 {
    font-size: 1.5em;
  }

  .page-header p {
    font-size: 0.9em;
  }

  .tour-card img {
    height: 130px;
  }

  .btn {
    width:30% ;
    font-size: 1em;
  }
}
</style>
