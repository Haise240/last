<template>
  <div class="about-page">
    <!-- Header Section -->
    <header class="page-header">
      <div class="container">
        <h1>О нас</h1>
        <p>Познакомьтесь с нашим туристическим агентством, узнайте о наших ценностях и миссии.</p>
      </div>
    </header>

    <!-- About Us Content Section -->
    <section class="about-content">
      <div class="container">
        <div class="about-text">
          <h2>Кто мы</h2>
          <p>
            Мы — туристическое агентство, посвященное предоставлению уникальных возможностей для знакомства с 
            культурой и природой Северной Осетии. Наша команда стремится создать незабываемые путешествия, 
            которые позволят вам полностью погрузиться в атмосферу этого удивительного региона.
          </p>
        </div>

        <div class="about-text">
          <h2>Наша миссия</h2>
          <p>
            Мы верим, что путешествия обогащают жизнь и открывают новые горизонты. Наша цель — предложить вам 
            уникальные туры, которые сочетают в себе приключения, культурные открытия и отдых на природе. 
            Мы стремимся предоставить каждому клиенту незабываемый опыт.
          </p>
        </div>

        <div class="about-text">
          <h2>Почему выбирают нас</h2>
          <ul>
            <li>Опытные гиды и экскурсоводы</li>
            <li>Индивидуальный подход к каждому клиенту</li>
            <li>Безопасность и комфорт на протяжении всего тура</li>
            <li>Широкий выбор маршрутов и программ</li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>



<style scoped>
/* Header Styling */
.page-header {
  background: linear-gradient(135deg, #355e5e, #35495e);
  color: white;
  padding: 60px 0;
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.page-header h1 {
  font-size: 2.5em;
  margin-bottom: 10px;
}

.page-header p {
  font-size: 1.2em;
}

/* About Content Styling */
.about-content {
  padding: 50px 0;
}

.about-text {
  margin-bottom: 30px;
}

.about-text h2 {
  font-size: 1.8em;
  margin-bottom: 15px;
}

.about-text p {
  font-size: 1em;
  line-height: 1.6;
}

.about-text ul {
  list-style: disc;
  padding-left: 20px;
}

.about-text ul li {
  margin-bottom: 8px;
}

/* Team Section Styling */
.team-section {
  background: #f4f4f4;
  padding: 50px 0;
}

.team-members {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.team-member {
  flex-basis: 30%;
  text-align: center;
  margin-bottom: 30px;
}

.team-member img {
  border-radius: 50%;
  width: 150px;
  height: 150px;
  object-fit: cover;
  margin-bottom: 15px;
}

.team-member h3 {
  font-size: 22px;
  margin-bottom: 5px;
}

.team-member p {
  font-size: 16px;
}

/* Testimonials Section Styling */
.testimonials-section {
  padding: 50px 0;
}

.testimonials {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.testimonial {
  flex-basis: 48%;
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

.testimonial p {
  font-style: italic;
  margin-bottom: 10px;
}

.testimonial h4 {
  font-size: 16px;
  color: #35495e;
  text-align: right;
}

/* Responsive Styling */
@media (max-width: 768px) {
  .team-member {
    flex-basis: 100%;
    margin-bottom: 20px;
  }

  .testimonial {
    flex-basis: 100%;
    margin-bottom: 20px;
  }

  .about-text h2 {
    font-size: 1.5em;
  }
}
</style>
