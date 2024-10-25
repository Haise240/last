<template>
  <div>
    <!-- Header -->

    <!-- Hero Section -->
    <section class="hero">
      <div class="overlay"></div>

      <div class="hero-content">
        <h2>Откройте для себя красоту Северной Осетии</h2>
        <p>
          Присоединяйтесь к нам в незабываемом путешествии по живописным ландшафтам и богатой
          культуре Северной Осетии.
        </p>
        <NuxtLink to="/tours" class="btn cta-button">Выбрать тур</NuxtLink>
      </div>
    </section>

    <!-- Popular Tours -->
    <section class="tours" id="tours">
      <div class="container">
        <h2>Популярные туры</h2>
        <div class="tour-cards">
          <div class="tour-card">
            <h3>Путешествие в Сердце Кавказа</h3>
            <p>Испытайте азарт восхождения по потрясающим горам Северной Осетии.</p>
            <NuxtLink to="/tour/77" class="btn">Подробнее</NuxtLink>
          </div>
          <div class="tour-card">
            <h3>Горные Ущелья и Культурные Традиции</h3>
            <p>Откройте для себя богатую историю и традиции осетинского народа.</p>
            <NuxtLink to="/tour/79" class="btn">Подробнее</NuxtLink>
          </div>
          <div class="tour-card">
            <h3>Джип-Приключение по Горам Осетии</h3>
            <p>Откройте для себя захватывающее джип-приключение по горным маршрутам Северной Осетии!</p>
            <NuxtLink to="/tour/64" class="btn">Подробнее</NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Advantages -->
    <section class="advantages">
      <div class="container">
        <h2>Наши преимущества</h2>
        <div class="advantages-cards">
          <div class="advantage-card">
            <h3>Опытные гиды</h3>
            <p>
              Наши гиды знают местность как свои пять пальцев и гарантируют безопасность в каждом туре.
            </p>
          </div>
          <div class="advantage-card">
            <h3>Уникальные маршруты</h3>
            <p>Мы предлагаем маршруты, которые невозможно найти в других компаниях.</p>
          </div>
          <div class="advantage-card">
            <h3>Поддержка 24/7</h3>
            <p>Наши специалисты всегда на связи, чтобы помочь вам с любыми вопросами.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Telegram Section -->
    <section class="contact-info">
      <div class="container">
        <h2>Наш Телеграм канал</h2>
        <p>Присоединяйтесь к нашему Телеграм-каналу, чтобы быть в курсе последних новостей и предложений!</p>
        <a href="https://t.me/" class="btn telegram-button">Подписаться на канал</a>
      </div>
    </section>

    <!-- Contact Section -->
    <section class="contact" id="contacts">
      <div class="container">
        <h2>Свяжитесь с нами</h2>
        <p>У вас есть вопросы? Хотите создать индивидуальный тур? Ответим на все вопросы!</p>
        <div class="phone">WhatsApp: +7 (928) 494 59-04</div>
        <div class="phone">Телефон для бронирования: +7 (918) 822 51-70</div>
        <a href="tel:+79188225170" class="btn contact-button">Позвонить</a>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="name">Имя</label>
            <input type="text" id="name" v-model="formData.name" required />
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" v-model="formData.email" required />
          </div>
          <div class="form-group">
            <label for="phone">Телефон</label>
            <input type="tel" id="phone" v-model="formData.phone" required />
          </div>
          <div class="form-group">
            <label for="message">Сообщение</label>
            <textarea id="message" v-model="formData.message" required></textarea>
          </div>
          <button type="submit" class="btn contact-button">Отправить</button>
        </form>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      formData: {
        name: '',
        email: '',
        phone: '',
        message: ''
      }
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await axios.post('http://localhost:8080/api/messages', this.formData);
        
        if (response.status === 200) {
          alert('Сообщение успешно отправлено!');
          this.formData = {
            name: '',
            email: '',
            phone: '',
            message: ''
          };
        } else {
          alert('Произошла ошибка при отправке. Попробуйте еще раз.');
        }
      } catch (error) {
        alert('Произошла ошибка при отправке: ' + error.message);
      }
    }
  }
};
</script>

<style scoped>
/* Main Styles */
body {
  font-family: 'Roboto', sans-serif;
  color: #333;
  line-height: 1.6;
  background-color: #f9f9f9;
}

/* Hero Section */
.hero {
  position: relative;
  background-image: url('/public/HomePage/MainPageBack.jpg');
  background-size: cover;
  background-position: center;
  color: white;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.hero .overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
}

.hero-content {
  z-index: 1;
  max-width: 700px;
}

.hero-content h2 {
  font-size: 3.5rem;
  margin-bottom: 20px;
  font-weight: 700;
}

.hero-content p {
  font-size: 1.25rem;
  margin-bottom: 30px;
}

.cta-button {
  background: #ff6347;
  color: white;
  padding: 12px 30px;
  text-transform: uppercase;
  border-radius: 8px;
  font-weight: bold;
  transition: background 0.3s ease;
}

.cta-button:hover {
  background: #e55341;
}

/* Popular Tours Section */
.tours {
  padding: 60px 0;
  background: #f4f4f4;
}

.tours h2{
  font-size: 2.5rem;
  text-align: center;
  margin-bottom: 30px;
}

.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}

.tour-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: space-between;
}

.tour-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  padding: 25px;
  text-align: center;
  display: flex;
  flex-direction: column;
  flex-basis: 30%;
}

.tour-card h3 {
  font-size: 1.5rem;
  margin-bottom: 15px;
  font-weight: 600;
}

.tour-card p {
  font-size: 1rem;
  margin-bottom: 20px;
  flex-grow: 1;
}

.btn {
  background: #42b983;
  color: white;
  padding: 12px 20px;
  border-radius: 5px;
  font-weight: bold;
  transition: background 0.3s ease;
}

.btn:hover {
  background: #348f66;
}

/* Advantages Section */
.advantages {
  padding: 60px 0;
  background: #35495e;
  color: white;
}

.advantages h2 {
  font-size: 2.5rem;
  text-align: center;
  margin-bottom: 40px;
}

.advantages-cards {
  display: flex;
  gap: 20px;
  justify-content: space-between;
}

.advantage-card {
  background: #4e5d6c;
  border-radius: 12px;
  padding: 30px;
  text-align: center;
  flex-basis: 30%;
}

.advantage-card h3 {
  font-size: 1.75rem;
  margin-bottom: 20px;
}

.advantage-card p {
  font-size: 1rem;
}

/* Contact Info Section */
.contact-info {
  background-color: #f4f4f4;
  padding: 60px 0;
  text-align: center;
  backdrop-filter: blur(11.5rem);
}

.contact-info h2{
  font-size: 2.5rem;
}

.contact-info p{
  font-size: 1.5rem;
  margin: 5%;
}

.telegram-button {  
  background: #42b983;
  color: white;
  border-radius: 8px;
  padding: 12px 30px;
  text-transform: uppercase;
  font-weight: bold;
  transition: background 0.3s ease;
}

.telegram-button:hover {
  background: #348f66;
}

/* Contact Section */
.contact {
  position: relative;
  background-color: #35495e;
  background-size: cover;
  background-position: center;
  color: white;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.contact h2{
  font-size: 2.5rem;
}

.phone {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.contact-button {
  background: #42b983;
  color: white;
  padding: 12px 30px;
  text-transform: uppercase;
  font-weight: bold;
  transition: background 0.3s ease;
}

.contact-button:hover {
  background: #348f66;
}

form {
  max-width: 500px;
  margin: 3% auto;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
}


input,
textarea {
  width: 100%;
  padding: 10px;
  border: 2px solid #525252;
  border-radius: 5px;
  font-size: 1rem;
}
  
input:focus,
textarea:focus {
  outline: none;
}

/* Responsive Styles */
@media (max-width: 768px) {
  .tour-cards,
  .advantages-cards {
    flex-direction: column;
  }

  .hero-content h2 {
    font-size: 2.5rem;
  }

  .hero-content p {
    font-size: 1rem;
  }
}
</style>
