<template>
  <div class="terms-container">
    <!-- Header Section -->
    <header class="terms-header">
      <div class="container">
        <h1>Условия и правила</h1>
        <p>Пожалуйста, внимательно ознакомьтесь с нашими условиями перед бронированием.</p>
      </div>
    </header>

    <!-- Terms and Conditions Content Section -->
    <section class="terms-content">
      <div class="container">
        <div class="terms-section" v-for="(section, index) in termsSections" :key="index">
          <h2>{{ section.title }}</h2>
          <p>{{ section.content }}</p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'TermsPage',
  data() {
    return {
      termsSections: [
        {
          title: '1. Общие положения',
          content:
            'Настоящие условия регулируют использование нашего сайта и услуг. Используя наш сайт, вы соглашаетесь соблюдать данные условия.'
        },
        {
          title: '2. Бронирование',
          content:
            'Все бронирования считаются окончательными только после получения полной оплаты. Мы оставляем за собой право отменить бронирование в случае неполной оплаты.'
        },
        {
          title: '3. Отмена и возврат средств',
          content:
            'Вы можете отменить бронирование за 14 дней до начала тура для полного возврата средств. Отмена менее чем за 14 дней может повлечь за собой удержание части суммы.'
        },
        {
          title: '4. Ответственность',
          content:
            'Мы не несем ответственности за задержки, отмены или изменения в маршрутах, вызванные непредвиденными обстоятельствами, такими как погодные условия, природные катаклизмы или другие форс-мажорные обстоятельства.'
        },
        {
          title: '5. Защита данных',
          content:
            'Мы собираем и обрабатываем личные данные в соответствии с политикой конфиденциальности. Ваши данные будут использованы только для обработки бронирований и связи с вами.'
        },
        {
          title: '6. Изменения в условиях',
          content:
            'Мы оставляем за собой право изменять данные условия в любое время. Мы рекомендуем регулярно проверять их на наличие изменений.'
        }
      ]
    };
  }
};
</script>

<style scoped>
/* General Container Styling */
.terms-container {
  padding: 40px 0;
  background-color: #f4f4f4;
}

/* Header Styling */
.terms-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
}

.terms-header h1 {
  font-size: 2.5em;
  margin-bottom: 10px;
}

.terms-header p {
  font-size: 1.2em;
}

/* Terms Content Section */
.terms-content {
  padding: 40px 0;
  background-color: white;
}

/* Terms Section Styling */
.terms-section {
  margin-bottom: 30px;
}

.terms-section h2 {
  font-size: 1.8em;
  margin-top: 20px;
  color: #35495e;
  border-bottom: 2px solid #35495e;
  padding-bottom: 10px;
}

.terms-section p {
  font-size: 1em;
  margin-top: 10px;
  line-height: 1.6;
}

/* Container for Inner Content */
.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}

/* Responsive Styling */
@media (max-width: 768px) {
  .terms-header h1 {
    font-size: 2em;
  }

  .terms-header p {
    font-size: 1em;
  }

  .terms-section h2 {
    font-size: 1.5em;
  }

  .terms-section p {
    font-size: 0.9em;
  }
}
</style>
