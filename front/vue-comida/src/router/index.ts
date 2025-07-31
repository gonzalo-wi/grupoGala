import { createRouter, createWebHistory } from 'vue-router'


import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PedidoView from '../views/PedidoView.vue'


const routes = [
  { path: '/', name: 'home', component: HomeView },
  { path: '/login', name: 'login', component: LoginView },
  { path: '/pedido', name: 'pedido', component: PedidoView },
  
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router