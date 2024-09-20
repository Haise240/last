import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import ToursPage from '@/views/ToursPage.vue'
import AboutPage from '@/views/AboutPage.vue'
import ContactPage from '@/views/ContactPage.vue'
import AdminPage from '@/views/AdminPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import TourPage from '@/views/TourPage.vue'
import TermsPage from '@/views/TermsPage.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/tours', component: ToursPage },
  { path: '/about', component: AboutPage },
  { path: '/contacts', component: ContactPage },
  { path: '/login', component: LoginPage },
  { path: '/tours/:id', component: TourPage },
  {
    path: '/admin',
    component: AdminPage,
    meta: { requiresAuth: true } // Защищенный маршрут
  },
  { path: '/terms', component: TermsPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
