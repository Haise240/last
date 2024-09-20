<template>
  <div id="app">
    <!-- Верхнее меню для больших экранов -->
    <header class="desktop-nav">
      <nav>
        <div class="logo">
          <router-link to="/">
            <h1>Explore Ossetia</h1>
          </router-link>
        </div>
        <ul class="nav-links">
          <li><router-link to="/">Главная</router-link></li>
          <li><router-link to="/tours">Туры</router-link></li>
          <li><router-link to="/about">О нас</router-link></li>
          <li><router-link to="/contacts">Контакты</router-link></li>
        </ul>
      </nav>
    </header>

    <!-- Точка монтирования для компонентов -->
    <router-view />

    <div class="mobile-nav">
  <router-link to="/">
    <font-awesome-icon :icon="['fas', 'home']" />
  </router-link>
  <router-link to="/tours">
    <font-awesome-icon :icon="['fas', 'route']" />
  </router-link>
  <router-link to="/about">
    <font-awesome-icon :icon="['fas', 'info-circle']" />
  </router-link>
  <router-link to="/contacts">
    <font-awesome-icon :icon="['fas', 'phone']" />
  </router-link>
</div>


    <!-- Футер -->
    <footer>
      <div class="container">
        <ul class="footer-links">
          <li><router-link to="/terms">Условия использования</router-link></li>
          <li><a href="#">Социальные сети</a></li>
        </ul>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
/* Reset CSS */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Open Sans', sans-serif;
  color: #333;
  line-height: 1.6;
  background-color: #fafafa; /* Легкий фон */
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* Минимальная высота на весь экран */
  margin: 0;
}

a {
  text-decoration: none;
  color: inherit;
}

html, body {
  height: 100%;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
}

#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* Контейнер занимает всю высоту окна */
}

main {
  flex: 1; /* Этот блок занимает оставшееся пространство между header и footer */
}

footer {
  background: #35495e;
  color: #fff;
  text-align: center;
  padding: 15px 0;
  margin-top: auto; /* Прижимаем футер к низу */
}

footer .footer-links {
  list-style: none;
  display: flex;
  justify-content: center;
}

footer .footer-links li {
  margin: 0 15px;
}

footer .footer-links a {
  color: #fffff0;
  transition: color 0.3s;
}

footer .footer-links a:hover {
  color: #9b9b9b;
}

h1, h2, h3, h4, h5, h6 {
  margin-bottom: 1rem;
  font-weight: 700;
}

.btn {
  display: inline-block;
  padding: 10px 20px;
  background-color: #509496; /* Цвет кнопки */
  color: white;
  border-radius: 5px;
  text-align: center;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn:hover {
  background-color: #509496; /* Цвет кнопки при наведении */
}

.container {
  width: 90%;
  margin: 0 auto;
  max-width: 1200px;
}

header {
  background: #f4f4f4;
  color: #333;
  width: 100%; /* Ширина на весь экран */
  padding: 20px 0;
}

.desktop-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  margin: 0 auto;
}

.desktop-nav .logo h1 {
  font-size: 28px;
  font-weight: 600;
  color: #35495e;
}

.desktop-nav .nav-links {
  list-style: none;
  display: flex;
  align-items: center;
}

.desktop-nav .nav-links li {
  margin-left: 20px;
}

.desktop-nav .nav-links a {
  color: #35495e;
  font-size: 16px;
}

section {
  padding: 60px 0;
}

/* Нижняя панель навигации для мобильных устройств */
.mobile-nav {
  display: flex;
  justify-content: space-around;
  align-items: center;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  background-color: #f4f4f4;
  padding: 10px 0;
  box-shadow: 0 -2px 5px rgba(0, 0, 0, 0.1);
}

.mobile-nav a {
  color: #333;
  font-size: 24px;
  text-align: center;
  flex-grow: 1;
}

.mobile-nav a:hover {
  color: #509496;
}

/* Скрываем мобильное меню на больших экранах */
@media (min-width: 768px) {
  .mobile-nav {
    display: none;
  }

}

@media (max-width: 768px) {
  .desktop-nav {
    display: none;
  }

  .footer {
    display: none;
  }

  .mobile-nav {
    height: 7%;
    z-index: 1;
    display: flex;
  }

  body {
    padding-bottom: 60px;
  }
}
</style>