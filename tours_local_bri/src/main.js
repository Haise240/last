import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// Импорт библиотек FontAwesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

// Импортируем иконки, которые будем использовать
import { faHome, faRoute, faInfoCircle, faPhone } from '@fortawesome/free-solid-svg-icons'

// Добавляем иконки в библиотеку
library.add(faHome, faRoute, faInfoCircle, faPhone)

const app = createApp(App)

// Регистрируем компонент FontAwesome globally
app.component('font-awesome-icon', FontAwesomeIcon)

app.use(router)

app.mount('#app')
