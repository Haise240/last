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
        <router-link to="/tours" class="btn cta-button">Выбрать тур</router-link>
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
            <router-link to="/tours/77" class="btn">Подробнее</router-link>
          </div>
          <div class="tour-card">
            <h3>Горные Ущелья и Культурные Традиции</h3>
            <p>Откройте для себя богатую историю и традиции осетинского народа.</p>
            <router-link to="/tours/79" class="btn">Подробнее</router-link>
          </div>
          <div class="tour-card">
            <h3>Джип-Приключение по Горам Осетии</h3>
            <p>Откройте для себя захватывающее джип-приключение по горным маршрутам Северной Осетии! </p>
            <router-link to="/tours/64" class="btn">Подробнее</router-link>
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
              Наши гиды знают местность как свои пять пальцев и гарантируют безопасность в каждом
              туре.
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

<!-- Contact Information -->
<section class="contact-info">
  <div class="container">
    <div>Блок под Телеграм канал</div>
    <p> фотографии </p>
  </div>
</section>


    <!-- Contact Section -->
    <section class="contact" id="contacts">
  <div class="container">
    
    <h2>Свяжитесь с нами</h2>
    <p>У вас есть вопросы? Хотите создать индивидуальный тур? Ответим на все вопросы!</p>
    <div class="phone">WhatsApp  +7 (928) 494 59-04</div>
    <div class="phone">Телефон для бронирования:  +7 (918) 822 51-70</div>
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
/* Основные стили для страницы HomePage.vue */
body {
  font-family: 'Roboto', sans-serif;
  color: #333;
  line-height: 1.6;
  background-color: #f9f9f9;
}

.hero {
  position: relative;
  background-image: url('@/assets/HomePage/MainPageBack.jpg');
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
  background: rgba(0, 0, 0, 0.5);
}

.hero-content {
  z-index: 1;
  max-width: 700px;
}

.hero-content h2 {
  font-size: 48px;
  margin-bottom: 20px;
  font-weight: 700;
}

.hero-content p {
  font-size: 20px;
  margin-bottom: 30px;
}

.cta-button {
  background: #ff6347;
  color: white;
  padding: 10px 20px;
  text-transform: uppercase;
  border-radius: 5px;
  text-decoration: none;
}

.cta-button:hover {
  background: #e55341;
}

.tours {
  padding: 50px 0;
  background: #f4f4f4;
}

.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}

.tour-cards {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.tour-card {
  background: white;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-bottom: 20px;
  flex-basis: 30%;
  padding: 20px;
  text-align: center;
  display: flex;
  flex-direction: column; /* Задаем вертикальный флекс-лейаут для карточки */
}

.tour-card h3 {
  font-size: 22px;
  margin-bottom: 15px;
  font-weight: 600;
}

.tour-card p {
  font-size: 16px;
  margin-bottom: 15px;
  flex-grow: 1; /* Растягиваем текстовое описание, чтобы кнопка всегда была внизу */
}

.btn {
  margin: auto;
  max-width: 40%;
  display: inline-block;
  margin-top: auto; /* Автоматически размещает кнопку внизу */
  padding: 10px 20px;
  background: #42b983;
  color: white;
  text-decoration: none;
  border-radius: 5px;
}

.btn:hover {
  background: #348f66;
}

.advantages {
  padding: 50px 0;
  background: #35495e;
  color: white;
}

.advantages h2 {
  text-align: center;
  margin-bottom: 30px;
}

.advantages-cards {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.advantage-card {
  background: #2c3e50;
  padding: 20px;
  border-radius: 10px;
  flex-basis: 30%;
  margin-bottom: 20px;
  text-align: center;
}

.advantage-card h3 {
  font-size: 22px;
  margin-bottom: 10px;
}

.reviews {
  padding: 50px 0;
  background: #f4f4f4;
}

.review-cards {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.review-card {
  background: white;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  padding: 20px;
  flex-basis: 48%;
  margin-bottom: 20px;
}

.review-card h4 {
  text-align: right;
  font-style: italic;
  margin-top: 10px;
}

.contact {
  position: relative; /* Для позиционирования псевдоэлемента */
  background-image: url('@/assets/HomePage/MainPageBack.jpg');
  background-size: cover; /* Растягиваем изображение на весь контейнер */
  background-position: center; /* Центрируем изображение */
  color: white;
  text-align: center;
  overflow: hidden; /* Чтобы скрыть любые элементы, выходящие за пределы контейнера */
  padding: 50px 0;
}

.contact::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('@/assets/HomePage/MainPageBack.jpg');
  background-size: cover;
  background-position: center;
  filter: blur(8px); /* Размытие изображения */
  z-index: 0; /* Псевдоэлемент находится за текстом */
}

.contact > * {
  position: relative;
  z-index: 1; /* Все внутренние элементы располагаются над псевдоэлементом */
}

/* Затемнение фона */
.contact::after {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5); /* Затемнение */
  z-index: 0; /* Наложение затемнения поверх размытого изображения */
}

.contact p {
  font-size: 20px;
  margin-bottom: 20px;
}

.contact-button {
  background: #42b983;
  color: white;
  padding: 10px 20px;
  text-transform: uppercase;
  border-radius: 5px;
  text-decoration: none;
}

.contact-button:hover {
  background: #348f66;
}

/* Основные стили для формы */
form {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background: #ffffff; /* Белый фон для формы */
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
  color: #333; /* Тёмный цвет текста для лучшего контраста */
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

.phone{
  font-size: 20px;
  margin-bottom: 20px;
}

/* Стиль для кнопки отправки */
button[type="submit"] {
  background-color: #42b983 ;
  color: white;
  border: none;
  padding: 12px 25px; /* Увеличены отступы для кнопки */
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px; /* Увеличен размер шрифта для кнопки */
  transition: background-color 0.3s, transform 0.2s; /* Плавный переход при наведении и нажатии */
}

/* Эффект при наведении на кнопку */
button[type="submit"]:hover {
  background-color: #348f66;
}

/* Эффект при нажатии на кнопку */
button[type="submit"]:active {
  transform: scale(0.98); /* Немного уменьшаем размер кнопки при нажатии */
}



@media screen and (max-width: 768px) {

  .btn,
  .cta-button{
    background: #42b983;
  color: white;
  border-radius: 5px;
  text-decoration: none;
  max-width: 50%;

  }
  
  .tour-cards,
  .advantages-cards,
  .review-cards {
    flex-direction: column;
  }

  .tour-card,
  .advantage-card,
  .review-card {
    flex-basis: 100%;
  }

  .hero-content h2 {
    font-size: 36px;
  }

  .hero-content p {
    font-size: 18px;
  }
}

form {
    padding: 15px;
  }

  input, textarea {
    font-size: 14px; /* Уменьшаем размер шрифта внутри полей на маленьких экранах */
  }

  button[type="submit"] {
    font-size: 18px; /* Уменьшаем размер шрифта для кнопки */
  }
  .contact-button {
  background: #42b983;
  color: white;
  border-radius: 5px;
  text-decoration: none;
  max-width: 50%;
}

.contact-button:hover {
  background: #348f66;
}
</style>
