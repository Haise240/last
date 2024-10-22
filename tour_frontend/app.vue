<template>
  <div id="app">
    <!-- Верхнее меню для больших экранов -->
    <header class="desktop-nav">
      <nav>
        <div class="logo">
          <NuxtLink to="/">
            <h1>Explore Ossetia</h1>
          </NuxtLink>
        </div>
        <ul class="nav-links">
          <li><NuxtLink to="/">Главная</NuxtLink></li>
          <li><NuxtLink to="/tours">Туры</NuxtLink></li>
          <li><NuxtLink to="/about">О нас</NuxtLink></li>
          <li><NuxtLink to="/contacts">Контакты</NuxtLink></li>
        </ul>
      </nav>
    </header>

    <!-- Точка монтирования для компонентов -->
    <NuxtPage />

    <div class="mobile-nav">
      <NuxtLink to="/">
        <font-awesome-icon :icon="['fas', 'home']" />
      </NuxtLink>
      <NuxtLink to="/tours">
        <font-awesome-icon :icon="['fas', 'route']" />
      </NuxtLink>
      <NuxtLink to="/about">
        <font-awesome-icon :icon="['fas', 'info-circle']" />
      </NuxtLink>
      <NuxtLink to="/contacts">
        <font-awesome-icon :icon="['fas', 'phone']" />
      </NuxtLink>
    </div>

    <!-- Футер -->
    <footer>
      <div class="container">
        <ul class="footer-links">
          <li><NuxtLink to="/terms">Условия использования</NuxtLink></li>
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
  background-color: #fafafa;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

a {
  text-decoration: none;
  color: inherit;
}

#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

main {
  flex: 1;
}

footer {
  background: #35495e;
  color: #fff;
  text-align: center;
  padding: 15px 0;
  margin-top: auto;
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

.container {
  width: 90%;
  margin: 0 auto;
  max-width: 1200px;
}

header {
  background: #f4f4f4;
  color: #333;
  width: 100%;
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
  margin-left: 10%;
  font-size: 28px;
  font-weight: 600;
  color: #35495e;
}

.desktop-nav .nav-links {
  list-style: none;
  display: flex;
  align-items: normal;
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

  footer {
    display: none;
  }

  .mobile-nav {
    height: 8%;
    z-index: 1;
    display: flex;
  }

  body {
    padding-bottom: 60px;
  }
}
</style>
