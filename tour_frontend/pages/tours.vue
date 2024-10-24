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
          <button
            v-for="(label, order) in sortOptions"
            :key="order"
            @click="setSortOrder(order)"
            :class="{ active: sortOrder === order }"
          >
            {{ label }}
          </button>
        </div>
      </div>
    </section>

    <!-- Tours Catalog Section -->
    <section class="tours-catalog">
      <div class="container">
        <div v-if="loading" class="loading-text">Загрузка туров...</div>
        <div v-else-if="filteredAndSortedTours.length === 0" class="no-tours-text">
          Нет доступных туров по выбранным фильтрам.
        </div>
        <div v-else class="tour-cards">
          <div v-for="tour in filteredAndSortedTours" :key="tour.id" class="tour-card">
            <img :src="formatImageUrl(tour.image_url)" :alt="tour.name || 'Изображение тура'" />
            <h2 class="tour-title">{{ tour.name || 'Название не указано' }}</h2>
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
      sortOrder: 'display-order',
      sortOptions: {
        'price-asc': 'Цена ↑',
        'price-desc': 'Цена ↓',
        'duration-asc': 'Длительность ↑',
        'duration-desc': 'Длительность ↓'
      },
      tours: [],
      loading: true,
    };
  },

  computed: {
    filteredAndSortedTours() {
      let filteredTours = this.tours;

      // Filter by duration
      if (this.filters.duration) {
        filteredTours = filteredTours.filter((tour) => {
          const duration = tour.duration;
          if (this.filters.duration === '1-3') return duration <= 3;
          if (this.filters.duration === '4-7') return duration >= 4 && duration <= 7;
          if (this.filters.duration === '8+') return duration >= 8;
          return true;
        });
      }

      // Filter by price range
      if (this.filters.priceRange) {
        filteredTours = filteredTours.filter((tour) => {
          const price = tour.price;
          if (this.filters.priceRange === '0-20000') return price <= 20000;
          if (this.filters.priceRange === '20000-50000') return price > 20000 && price <= 50000;
          if (this.filters.priceRange === '50000+') return price > 50000;
          return true;
        });
      }

      // Sort based on the selected order
      const sortFunctions = {
        'display-order': (a, b) => a.DisplayOrder - b.DisplayOrder,
        'price-asc': (a, b) => a.price - b.price,
        'price-desc': (a, b) => b.price - a.price,
        'duration-asc': (a, b) => a.duration - b.duration,
        'duration-desc': (a, b) => b.duration - a.duration,
      };

      return filteredTours.sort(sortFunctions[this.sortOrder]);
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
      return imageURL?.Valid ? imageURL.String : '1.jpg'; // Fallback image
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
/* General Styles */
body {
  margin: 0;
  font-family: Arial, sans-serif;
}

.loading-text, .no-tours-text {
  min-height: 50px; /* Задайте минимальную высоту для избежания сдвигов */
  text-align: center;
  font-size: 1.2em;
  color: #555;
}

/* Header Section */
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

/* Container */
.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}

/* Sorting Buttons */
.sorting-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-bottom: 10px;
}

.sorting-buttons button {
  padding: 10px 20px;
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.3s, box-shadow 0.3s;
}

.sorting-buttons button:hover {
  background-color: #e0e0e0;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.sorting-buttons button.active {
  background-color: #007bff;
  color: white;
  border-color: #007bff;
}

/* Tours Catalog Section */
.tours-catalog {
  padding: 40px 0;
  min-height: 200px; /* Добавьте минимальную высоту для предотвращения скачков */
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
  text-align: center;
  transition: transform 0.3s ease;
}

.tour-card:hover {
  transform: translateY(-5px);
}

.tour-card img {
  width: 100%;
  height: 180px; /* Фиксированная высота для предотвращения сдвигов */
  object-fit: cover;
  background-color: #f0f0f0;
  border-bottom: 1px solid #ddd;
  display: block;
}

/* Skeleton Placeholder для плавной загрузки */
.tour-card-placeholder {
  width: 100%;
  height: 180px;
  background-color: #f0f0f0;
  animation: shimmer 1.5s infinite;
}

@keyframes shimmer {
  0% {
    background-position: -500px 0;
  }
  100% {
    background-position: 500px 0;
  }
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

/* Responsive Design */
@media (max-width: 1024px) {
  .tour-card {
    flex-basis: calc(50% - 20px);
  }
}

@media (max-width: 768px) {
  .tour-card {
    flex-basis: 100%;
  }

  .sorting-buttons {
    flex-direction: column;
  }

  .sorting-buttons button {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .tour-card img {
    height: 130px;
  }

  .btn {
    width: 70%;
  }
}
</style>
