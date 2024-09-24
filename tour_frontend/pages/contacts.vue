<template>
  <div class="contacts-page">
    <!-- Header Section -->
    <header class="page-header">
      <div class="container">
        <h1>Контакты</h1>
        <p>Свяжитесь с нашим туристическим агентством. Мы всегда рады помочь вам!</p>
      </div>
    </header>

    <!-- Contact Information Section -->
    <section class="contact-info">
      <div class="container">
        <div class="info-block">
          <h2>Наши контакты</h2>
          <ul>
            <li>
              <i class="fas fa-phone-alt"></i>
              <a href="tel:+79188225170">+7 (918) 822 51-70</a>
            </li>
            <li>
              <i class="fas fa-envelope"></i>
              <a href="mailto:info@example.com">info@example.com</a>
            </li>
            <li>
              <i class="fas fa-map-marker-alt"></i>
              г. Владикавказ, ул. Бр. Темировых, 69
            </li>
          </ul>
        </div>

        <div class="info-block">
          <h2>Часы работы</h2>
          <ul>
            <li>Понедельник - Пятница: 09:00 - 18:00</li>
            <li>Суббота: 10:00 - 15:00</li>
            <li>Воскресенье: выходной</li>
          </ul>
        </div>
      </div>
    </section>

    <!-- Contact Form Section -->
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

    <!-- Map Section -->
    <section class="map-section">
      <div class="container">
        <h2>Наше местоположение</h2>
        <div class="map">
          <!-- Вставьте здесь ссылку на вашу карту, например, Google Maps iframe -->
          <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2917.56641868201!2d44.66386717663307!3d43.008463593890774!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x405aa0f155f7f09f%3A0x500a8c2d7d60fad1!2z0YPQuy4g0JHRgNCw0YLRjNC10LIg0KLQtdC80LjRgNC-0LLRi9GFLCA2OSwg0JLQu9Cw0LTQuNC60LDQstC60LDQtywg0KDQtdGB0L8uINCh0LXQstC10YDQvdCw0Y8g0J7RgdC10YLQuNGPIOKAlCDQkNC70LDQvdC40Y8sIDM2MjAxNQ!5e0!3m2!1sru!2sru!4v1726590855083!5m2!1sru!2sru" 
          width="800" 
          height="600" 
          style="border:0;" 
          allowfullscreen="" 
          loading="lazy" 
          referrerpolicy="no-referrer-when-downgrade"></iframe>
        </div>
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
        // Отправляем данные на сервер
        const response = await axios.post('http://localhost:8080/api/messages', this.formData);
        
        if (response.status === 200) {
          alert('Сообщение успешно отправлено!');
          // Сброс формы после отправки
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
/* Основные стили для страницы ContactsPage.vue */

.page-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.page-header h1 {
  font-size: 36px;
  margin-bottom: 10px;
}

.page-header p {
  font-size: 18px;
}

.contact-info {
  padding: 50px 0;
}

.info-block {
  margin-bottom: 30px;
}

.info-block h2 {
  font-size: 28px;
  margin-bottom: 15px;
}

.info-block ul {
  list-style: none;
  padding: 0;
}

.info-block ul li {
  font-size: 18px;
  margin-bottom: 10px;
  display: flex;
  align-items: center;
}

.info-block ul li i {
  margin-right: 10px;
}

.contact-form-section {
  padding: 50px 0;
}

.contact-form-section h2 {
  font-size: 28px;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  font-size: 16px;
  margin-bottom: 5px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.form-group textarea {
  resize: vertical;
}

.btn {
  background: #42b983;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
}

.btn:hover {
  background: #369972;
}

.map-section {
  padding: 50px 0;
  background: #f4f4f4;
}

.map-section h2 {
  font-size: 28px;
  margin-bottom: 20px;
}

.map {
  width: 100%;
  max-width: 100%;
  height: 450px;
  overflow: hidden;
  border-radius: 10px;
}

/* Основные стили для формы */
form {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background: #35495e; /* Белый фон для формы */
  border-radius: 8px; /* Скругленные углы */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Тень для формы */
}

/* Стиль для группы полей формы */
.form-group {
  margin-bottom: 20px;
}

/* Стилизация меток */
label {
  display: block;
  margin-bottom: 8px; /* Увеличен отступ снизу для меток */
  font-weight: 500; /* Полужирное начертание для меток */
  color: #ffffff; /* Тёмный цвет текста для лучшего контраста */
}

/* Стили для полей ввода и текстовой области */
input, textarea {
  width: 100%;
  padding: 12px; /* Увеличен отступ внутри полей для удобства ввода */
  border: 1px solid #ddd; /* Светлый цвет границы для полей ввода */
  border-radius: 5px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1); /* Внутренняя тень для полей ввода */
  font-size: 16px; /* Увеличен размер шрифта для лучшей читаемости */
}

/* Увеличенные отступы для полей ввода при фокусе */
input:focus, textarea:focus {
  border-color: #348f66; /* Цвет границы при фокусе */
  outline: none; /* Убираем стандартную обводку браузера */
  box-shadow: 0 0 0 2px rgba(150, 100, 80, 0.2); /* Добавляем тень при фокусе */
}

/* Высота для текстовой области */
textarea {
  height: 150px;
  resize: vertical; /* Позволяет изменять размер области только по вертикали */
}
</style>
