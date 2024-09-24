<template>
  <div class="login-page">
    <h1>Вход в панель администратора</h1>
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="username">Логин:</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div class="form-group">
        <label for="password">Пароль:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit" class="btn">Войти</button>
    </form>
    <p v-if="error" class="error-message">{{ error }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
      error: ''
    }
  },
  methods: {
    async login() {
      // Пример: отправка данных на сервер для проверки
      try {
        const response = await fetch('/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            username: this.username,
            password: this.password
          })
        })

        if (response.ok) {
          const data = await response.json()
          // Сохраняем токен в localStorage
          localStorage.setItem('authToken', data.token)
          // Перенаправляем на админ-панель
          this.$router.push('/admin')
        } else {
          this.error = 'Неправильный логин или пароль'
        }
      } catch (err) {
        this.error = 'Произошла ошибка, попробуйте позже'
      }
    }
  }
}
</script>

<style scoped>
/* Стили для страницы логина */
.login-page {
  max-width: 400px;
  margin: 100px auto;
  padding: 20px;
  background: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.btn {
  width: 100%;
  padding: 10px;
  background-color: #42b983;
  color: white;
  border-radius: 5px;
  text-align: center;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn:hover {
  background-color: #369972;
}

.error-message {
  color: red;
  margin-top: 10px;
  text-align: center;
}
</style>
