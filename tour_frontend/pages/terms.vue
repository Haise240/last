<template>
  <div class="terms-container">
    <header class="terms-header">
      <div class="container">
        <h1>Условия и правила</h1>
        <p>Пожалуйста, внимательно ознакомьтесь с нашими условиями перед бронированием.</p>
      </div>
    </header>

    <section class="terms-content">
      <div class="container">
        <div class="terms-section">
          <h2>1. Общие положения</h2>
          <p>
            Настоящие условия регулируют использование нашего сайта и услуг. Используя наш сайт, вы
            соглашаетесь соблюдать данные условия.
          </p>
        </div>

        <div class="terms-section">
          <h2>2. Бронирование</h2>
          <p>
            Все бронирования считаются окончательными только после получения полной оплаты. Мы
            оставляем за собой право отменить бронирование в случае неполной оплаты.
          </p>
        </div>

        <div class="terms-section">
          <h2>3. Отмена и возврат средств</h2>
          <p>
            Вы можете отменить бронирование за 14 дней до начала тура для полного возврата средств.
            Отмена менее чем за 14 дней может повлечь за собой удержание части суммы.
          </p>
        </div>

        <div class="terms-section">
          <h2>4. Ответственность</h2>
          <p>
            Мы не несем ответственности за задержки, отмены или изменения в маршрутах, вызванные
            непредвиденными обстоятельствами, такими как погодные условия, природные катаклизмы или
            другие форс-мажорные обстоятельства.
          </p>
        </div>

        <div class="terms-section">
          <h2>5. Защита данных</h2>
          <p>
            Мы собираем и обрабатываем личные данные в соответствии с политикой конфиденциальности.
            Ваши данные будут использованы только для обработки бронирований и связи с вами.
          </p>
        </div>

        <div class="terms-section">
          <h2>6. Изменения в условиях</h2>
          <p>
            Мы оставляем за собой право изменять данные условия в любое время. Мы рекомендуем
            регулярно проверять их на наличие изменений.
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'TermsPage'
}
</script>

<style scoped>
.terms-container {
  padding: 40px 0;
  background-color: #f4f4f4;
}

.terms-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.terms-header h1 {
  font-size: 36px;
  margin-bottom: 10px;
}

.terms-header p {
  font-size: 18px;
}

.terms-content {
  padding: 40px 0;
  background: white;
}

.terms-section {
  margin-bottom: 30px;
}

.terms-section h2 {
  font-size: 26px;
  margin-top: 20px;
  color: #35495e;
  border-bottom: 2px solid #35495e;
  padding-bottom: 10px;
}

.terms-section p {
  font-size: 16px;
  margin-top: 10px;
  line-height: 1.6;
}

.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
}
</style>
