import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'

// Registrar IP al cargar la app
fetch('http://beneficios.grupogala.com.ar:8095/api/register-ip', {
  method: 'POST',
  credentials: 'include',
});

createApp(App)
  .use(router)
  .mount('#app')
