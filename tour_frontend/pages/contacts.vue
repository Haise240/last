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
    <section class="contact-form-section">
      <div class="container">
        <h2>Свяжитесь с нами</h2>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="name" class="form-head">Имя</label>
            <input type="text" id="name" v-model="formData.name" required />
          </div>
          <div class="form-group">
            <label for="email" class="form-head">Email</label>
            <input type="email" id="email" v-model="formData.email" required />
          </div>
          <div class="form-group">
            <label for="phone" class="form-head">Телефон</label>
            <input type="tel" id="phone" v-model="formData.phone" required />
          </div>
          <div class="form-group">
            <label for="message" class="form-head">Сообщение</label>
            <textarea id="message" v-model="formData.message" required></textarea>
          </div>
          <button type="submit" class="btn contact-button">Отправить</button>
        </form>
      </div>
    </section>

    <!-- Map Section -->
    <section class="map-section">
      <div class="container">
        <h2>Наше местоположение</h2>
        <div class="map">
          <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2917.56641868201!2d44.66386717663307!3d43.008463593890774!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x405aa0f155f7f09f%3A0x500a8c2d7d60fad1!2z0YPQuy4g0JHRgNCw0YLRjNC10LIg0KLQtdC80LjRgNC-0LLRi9GFLCA2OSwg0JLQu9Cw0LTQuNC60LDQstC60LDQtywg0KDQtdGB0L8uINCh0LXQstC10YDQvdCw0Y8g0J7RgdC10YLQuNGPIOKAlCDQkNC70LDQvdC40Y8sIDM2MjAxNQ!5e0!3m2!1sru!2sru!4v1726590855083!5m2!1sru!2sru" 
            width="100%" 
            height="450" 
            style="border:0;" 
            allowfullscreen="" 
            loading="lazy" 
            referrerpolicy="no-referrer-when-downgrade">
          </iframe>
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
/* Contacts Page Styles */
.contacts-page {
  background-color: #f9f9f9;
}

/* Header Section */
.page-header {
  background-color: #35495e;
  color: white;
  padding: 60px 0;
  text-align: center;
}

.page-header h1 {
  font-size: 2.5rem;
  margin-bottom: 15px;
  font-weight: 700;
}

.page-header p {
  font-size: 1.2rem;
  color: #ddd;
}

/* Contact Information Section */
.contact-info {
  padding: 60px 0;
  background-color: #f4f4f4;
  text-align: center;
}

.contact-info h2 {
  font-size: 2.5rem;
  margin-bottom: 30px;
}

.form-head{
  color:#333;
}
.info-block {
  margin-bottom: 40px;
}

.info-block ul {
  list-style: none;
  padding: 0;
  font-size: 1.2rem;
  color: #333;
}

.info-block ul li {
  margin: 10px 0;
}

.info-block i {
  margin-right: 10px;
  color: #42b983;
}

/* Contact Form Section */
.contact-form-section {
  background-color: #35495e;
  padding: 60px 0;
  color: #f4f4f4;
}

.contact-form-section h2 {
  font-size: 2.5rem;
  margin-bottom: 20px;
  text-align: center;
}

form {
  max-width: 500px;
  margin: 0 auto;
  background: #f4f4f4;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 20px;
  text-align: left;
}

label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
}

input[type="text"],
input[type="email"],
input[type="tel"],
textarea {
  width: 100%;
  padding: 10px;
  border: 2px solid #ddd;
  border-radius: 5px;
  font-size: 1rem;
  background: #f9f9f9;
}

input:focus,
textarea:focus {
  outline: none;
  border-color: #42b983;
}

textarea {
  min-height: 100px;
  resize: vertical;
}

.contact-button {
  background: #42b983;
  color: white;
  padding: 12px 30px;
  text-transform: uppercase;
  font-weight: bold;
  border-radius: 5px;
  transition: background 0.3s ease;
  display: inline-block;
  margin-top: 20px;
}

.contact-button:hover {
  background: #348f66;
}

/* Map Section */
.map-section {
  padding: 60px 0;
  text-align: center;
}

.map-section h2 {
  font-size: 2.5rem;
  margin-bottom: 30px;
}

/* Responsive Styles */
@media (max-width: 768px) {
  .page-header {
    padding: 40px 20px;
  }

  .contact-info h2,
  .contact-form-section h2,
  .map-section h2 {
    font-size: 2rem;
  }

  form {
    padding: 15px;
  }

  .info-block ul {
    font-size: 1rem;
  }
}

</style>
