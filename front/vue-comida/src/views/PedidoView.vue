<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 via-amber-50 to-yellow-50 relative overflow-hidden">
    
    <!-- Elementos decorativos de fondo -->
    <div class="absolute inset-0 pointer-events-none">
      <!-- Overlay animado -->
      <div class="absolute inset-0 bg-gradient-to-br from-orange-200/30 via-amber-200/20 to-yellow-200/30 animate-bgMove"></div>
      <!-- Círculos animados -->
      <div class="absolute top-20 left-10 w-64 h-64 bg-orange-200/20 rounded-full blur-3xl animate-float1"></div>
      <div class="absolute bottom-20 right-10 w-80 h-80 bg-yellow-200/20 rounded-full blur-3xl animate-float2"></div>
      <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-amber-200/15 rounded-full blur-3xl animate-float3"></div>
      <div class="absolute top-10 right-1/3 w-40 h-40 bg-orange-300/20 rounded-full blur-2xl animate-float4"></div>
      <div class="absolute bottom-10 left-1/4 w-52 h-52 bg-yellow-300/20 rounded-full blur-2xl animate-float5"></div>
    </div>

    <!-- Contenedor principal -->
    <div class="relative z-10 p-8">
      
      <!-- Header mejorado -->
      <div class="text-center mb-16">
        <!-- Logo arriba del ícono -->
        <div class="flex justify-center mb-4">
          <img src="@/assets/galaLogo.png" alt="Gala Logo" class="h-60 drop-shadow-xl" />
        </div>
        <div class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-r from-orange-400 to-yellow-500 rounded-full shadow-xl mb-6">
          <span class="text-4xl">🍽️</span>
        </div>
        <h1 class="text-5xl font-bold bg-gradient-to-r from-orange-600 to-yellow-600 bg-clip-text text-transparent mb-4">
          ¡Escoge tu Menú Premium!
        </h1>
        <p class="text-xl text-gray-600 font-medium max-w-2xl mx-auto">
          Deliciosas opciones preparadas con ingredientes frescos y mucho amor
        </p>
      </div>

      <!-- Grid de tarjetas mejorado -->
      <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8 max-w-7xl mx-auto">
        
        <!-- Tarjeta de menú con diseño premium -->
        <div
          v-for="(menu, index) in menus"
          :key="menu.id"
          class="group bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl hover:shadow-2xl transition-all duration-500 overflow-hidden flex flex-col transform hover:-translate-y-2 border border-white/50"
          :class="[
            index === 1 ? 'md:scale-105 relative z-20' : '',
            'hover:scale-105'
          ]"
        >
          <!-- Badge especial para el menú del centro -->
          <div v-if="index === 1" class="absolute top-4 right-4 z-30">
            <span class="bg-gradient-to-r from-orange-500 to-yellow-500 text-white px-3 py-1 rounded-full text-sm font-bold shadow-lg">
              ⭐ Popular
            </span>
          </div>

          <!-- Imagen con overlay -->
          <div class="relative overflow-hidden">
            <img
              :src="menu.image"
              :alt="menu.title"
              class="w-full h-56 object-cover group-hover:scale-110 transition-transform duration-700"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/30 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          </div>

          <!-- Contenido con mejor espaciado -->
          <div class="p-8 flex flex-col flex-grow">
            <div class="flex items-center justify-between mb-4">
              <h2 class="text-2xl font-bold text-gray-800 group-hover:text-orange-600 transition-colors">
                {{ menu.title }}
              </h2>
              <div class="flex items-center space-x-1">
                <span class="text-yellow-400">⭐</span>
                <span class="text-sm font-medium text-gray-600">4.8</span>
              </div>
            </div>
            
            <p class="text-gray-600 flex-grow text-base leading-relaxed mb-6">
              {{ menu.description }}
            </p>

            <!-- Precio ficticio -->
            <div class="flex items-center justify-between mb-6">
              <span class="text-3xl font-bold text-orange-600">${{ menu.price }}</span>
              <span class="text-sm text-gray-500 line-through">${{ menu.originalPrice }}</span>
            </div>

            <!-- Botón con gradiente y animación -->
            <button
              @click="agregarAlPedido(menu)"
              class="relative overflow-hidden bg-gradient-to-r from-orange-500 to-yellow-500 hover:from-orange-600 hover:to-yellow-600 text-white py-4 px-6 rounded-2xl font-bold text-lg shadow-xl hover:shadow-2xl transition-all duration-300 transform active:scale-95 group"
            >
              <span class="relative z-10 flex items-center justify-center space-x-2">
                <span>🛒</span>
                <span>Agregar al pedido</span>
              </span>
              <div class="absolute inset-0 bg-gradient-to-r from-yellow-400 to-orange-400 translate-x-full group-hover:translate-x-0 transition-transform duration-300"></div>
            </button>
          </div>
        </div>

      </div>

      <!-- Información adicional -->
      <div class="mt-20 text-center">
        <div class="bg-white/60 backdrop-blur-sm rounded-2xl p-8 max-w-4xl mx-auto shadow-lg border border-white/50">
          <h3 class="text-2xl font-bold text-gray-800 mb-4">🚚 Entrega rápida y segura</h3>
          <p class="text-gray-600 text-lg">
            Todos nuestros menús se preparan el mismo día y se entregan frescos a tu domicilio entre las 21:00 y 22:00 hrs.
          </p>
        </div>
      </div>

    </div>

    <!-- Modal de dirección y día de entrega -->
    <transition name="fade">
      <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/30 backdrop-blur-sm">
        <div class="bg-white rounded-3xl shadow-2xl p-8 w-full max-w-md border-2 border-amber-100 relative">
          <button @click="showModal = false" class="absolute top-4 right-4 text-gray-400 hover:text-orange-600 text-2xl font-bold">×</button>
          <div class="flex flex-col items-center mb-6">
            <img src="@/assets/galaLogo.png" alt="Logo" class="h-40 mb-2 drop-shadow-lg" />
            <h2 class="text-2xl font-bold bg-gradient-to-r from-orange-600 to-yellow-600 bg-clip-text text-transparent mb-2">Confirmar pedido</h2>
            <p class="text-gray-500 text-center mb-2">Completa tus datos para recibir tu pedido</p>
            <span v-if="selectedMenu" class="text-lg font-semibold text-orange-600">{{ selectedMenu.title }}</span>
          </div>
          <form @submit.prevent="confirmarPedido" class="space-y-5">
            <div>
              <label class="block mb-2 text-sm font-bold text-gray-700">Dirección de entrega</label>
              <input v-model="address" type="text" placeholder="Ej: Av. Siempreviva 123" class="w-full px-4 py-3 border-2 border-amber-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all shadow-sm bg-gray-50 focus:bg-white" required />
            </div>
            <div>
              <label class="block mb-2 text-sm font-bold text-gray-700">Día de entrega</label>
              <input v-model="deliveryDate" type="date" class="w-full px-4 py-3 border-2 border-amber-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all shadow-sm bg-gray-50 focus:bg-white" required />
            </div>
            <button type="submit" class="w-full bg-gradient-to-r from-orange-500 via-amber-500 to-yellow-500 text-white py-3 rounded-xl font-bold text-lg shadow-xl hover:shadow-2xl hover:scale-[1.02] transition-all duration-300 focus:outline-none focus:ring-4 focus:ring-orange-500/50">Confirmar pedido</button>
          </form>
        </div>
      </div>
    </transition>
  </div>
  <!-- Footer con redes sociales -->
 <footer class="w-full bg-gradient-to-r from-gray-900 via-gray-800 to-gray-900 text-gray-300 py-8 mt-20">
    <div class="max-w-6xl mx-auto flex flex-col items-center space-y-6">
      
      <!-- 🔗 Redes sociales -->
      <div class="flex space-x-6">
        <!-- Instagram -->
        <a href="https://www.instagram.com/grupogala.catering/" target="_blank" rel="noopener" 
           class="p-3 rounded-full bg-gray-700 hover:bg-pink-500 transition-all duration-300 transform hover:scale-110 hover:shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" width="28" height="28">
            <path d="M12 2.2c3.2 0 3.6 0 4.9.1 1.2.1 1.9.2 2.4.4.6.2 1 .5 1.5.9.4.4.7.9.9 1.5.2.5.4 1.2.4 2.4.1 1.3.1 1.7.1 4.9s0 3.6-.1 4.9c-.1 1.2-.2 1.9-.4 2.4-.2.6-.5 1-0.9 1.5-.4.4-.9.7-1.5.9-.5.2-1.2.4-2.4.4-1.3.1-1.7.1-4.9.1s-3.6 0-4.9-.1c-1.2-.1-1.9-.2-2.4-.4-.6-.2-1-.5-1.5-.9-.4-.4-.7-.9-.9-1.5-.2-.5-.4-1.2-.4-2.4C2.2 15.6 2.2 15.2 2.2 12s0-3.6.1-4.9c.1-1.2.2-1.9.4-2.4.2-.6.5-1 0.9-1.5.4-.4.9-.7 1.5-.9.5-.2 1.2-.4 2.4-.4C8.4 2.2 8.8 2.2 12 2.2zm0 1.8c-3.1 0-3.4 0-4.6.1-1 .1-1.5.2-1.9.3-.5.2-.8.3-1.2.7-.4.4-.5.7-.7 1.2-.1.4-.3.9-.3 1.9-.1 1.2-.1 1.5-.1 4.6s0 3.4.1 4.6c.1 1 .2 1.5.3 1.9.2.5.3.8.7 1.2.4.4.7.5 1.2.7.4.1.9.3 1.9.3 1.2.1 1.5.1 4.6.1s3.4 0 4.6-.1c1-.1 1.5-.2 1.9-.3.5-.2.8-.3 1.2-.7.4-.4.5-.7.7-1.2.1-.4.3-.9.3-1.9.1-1.2.1-1.5.1-4.6s0-3.4-.1-4.6c-.1-1-.2-1.5-.3-1.9-.2-.5-.3-.8-.7-1.2-.4-.4-.7-.5-1.2-.7-.4-.1-.9-.3-1.9-.3-1.2-.1-1.5-.1-4.6-.1zm0 3.4a6.6 6.6 0 1 1 0 13.2 6.6 6.6 0 0 1 0-13.2zm0 10.9a4.3 4.3 0 1 0 0-8.6 4.3 4.3 0 0 0 0 8.6zm5.4-10.9a1.5 1.5 0 1 1 0-3.1 1.5 1.5 0 0 1 0 3.1z"/>
          </svg>
        </a>

        <!-- Facebook -->
        <a href="https://www.facebook.com/grupogalacatering/" target="_blank" rel="noopener" 
           class="p-3 rounded-full bg-gray-700 hover:bg-blue-600 transition-all duration-300 transform hover:scale-110 hover:shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" width="28" height="28">
            <path d="M22 12a10 10 0 1 0-11.5 9.9v-7h-2v-3h2v-2.3c0-2 1.2-3.1 3-3.1.9 0 1.8.1 1.8.1v2h-1c-1 0-1.3.6-1.3 1.2V12h2.2l-.4 3h-1.8v7A10 10 0 0 0 22 12z"/>
          </svg>
        </a>

        <!-- WhatsApp -->
        <a href="https://wa.me/" target="_blank" rel="noopener" 
           class="p-3 rounded-full bg-gray-700 hover:bg-green-500 transition-all duration-300 transform hover:scale-110 hover:shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" width="28" height="28">
            <path d="M20 3.5A10.5 10.5 0 0 0 3.4 17.7L2 22l4.4-1.4A10.5 10.5 0 1 0 20 3.5zm-8 .9a8.6 8.6 0 0 1 7.2 13.4l-.4.5.9 3.1-3.1-.9-.5.4A8.6 8.6 0 1 1 12 4.4zm4.5 10.2c-.2-.1-1.3-.6-1.4-.6-.2 0-.3-.1-.5.1-.1.1-.5.6-.6.7-.1.1-.2.1-.4 0-.2-.1-.9-.3-1.7-1-.6-.5-1-1.1-1.1-1.3-.1-.2 0-.3.1-.4l.3-.3c.1-.1.1-.2.2-.3.1-.1.1-.2 0-.4s-.5-1.2-.7-1.7c-.2-.4-.4-.3-.5-.3h-.4c-.1 0-.3 0-.5.2-.2.1-.7.6-.7 1.4s.7 1.6.8 1.7c.1.2 1.4 2.1 3.4 3 .5.2.9.3 1.2.4.5.1.9.1 1.2-.1.3-.2 1.1-.8 1.3-1.3.1-.4.1-.7 0-.8 0 0-.1-.1-.3-.2z"/>
          </svg>
        </a>
      </div>

      <!-- 🔤 Texto -->
      <span class="text-sm text-gray-400 hover:text-gray-200 transition-colors">
        &copy; 2025 <strong class="text-white">Grupo Galla</strong>. Todos los derechos reservados.
      </span>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import sushiImg from '@/assets/Sushi.jpg'
import pizzaImg from '@/assets/Pizza.jpg'
import burgerImg from '@/assets/Burger.jpg'
import shawarmaImg from '@/assets/shawarma.jpg'
import tacosImg from '@/assets/Tacos.jpg'
import casherImg from '@/assets/casher.jpg'

interface MenuItem {
  id: number
  title: string
  description: string
  image: string
  price: number
  originalPrice: number
}

const menus = ref<MenuItem[]>([
  {
    id: 1,
    title: 'Sushi Deluxe',
    description: '12 piezas de sushi fresco con salmón, atún y langostinos, acompañadas de jengibre, wasabi y salsa de soja.',
    image: sushiImg,
    price: 0.00,
    originalPrice: 15000.90
  },
  {
    id: 2,
    title: 'Burger House',
    description: 'Hamburguesa gourmet con carne 100% vacuna, queso cheddar, panceta crocante y pan brioche artesanal.',
    image: burgerImg,
    price: 0.00,
    originalPrice: 25000.90
  },
  {
    id: 3,
    title: 'Pizza Pepperoni',
    description: 'Pizza estilo americano con abundante pepperoni, queso mozzarella derretido y salsa de tomate casera en masa a la piedra.',
    image: pizzaImg,
    price: 0.00,
    originalPrice: 35000.90
  },
  {
    id: 4,
    title: 'Comida Árabe',
    description: 'Shawarma de cordero, hummus cremoso, falafel crocante y pan pita artesanal, acompañado de salsa tahini.',
    image: shawarmaImg,
    price: 0.00,
    originalPrice: 28000.90
  },
  {
    id: 5,
    title: 'Comida Judía',
    description: 'Kugel de papa dorado, pastrami casero, bagels frescos y ensalada israelí con aderezo de oliva.',
    image:casherImg,
    price: 0.00,
    originalPrice: 32000.90
  },
  {
    id: 6,
    title: 'Comida Mexicana',
    description: 'Tacos de carnitas, guacamole fresco, nachos con queso y salsa picante, todo acompañado de limonada artesanal.',
    image: tacosImg,
    price: 0.00,
    originalPrice: 29000.90
  }
])

const showModal = ref(false)
const selectedMenu = ref<MenuItem | null>(null)
const address = ref('')
const deliveryDate = ref('')

const agregarAlPedido = (menu: MenuItem) => {
  selectedMenu.value = menu
  showModal.value = true
}

const confirmarPedido = async () => {
  showModal.value = false
  // Validar campos requeridos
  if (!selectedMenu.value?.title || !address.value || !deliveryDate.value) {
    alert('❌ Faltan datos requeridos para enviar el pedido')
    return
  }
  // Formatear fecha a ISO 8601 (UTC)
  const dateObj = new Date(deliveryDate.value)
  const fechaISO = dateObj.toISOString()
  try {
    await axios.post('http://192.168.0.55:8095/api/pedido', {
      combo: selectedMenu.value.title,
      direccion: address.value,
      fecha: fechaISO || ''
    })
    alert('✅ Pedido enviado correctamente')
  } catch (error) {
    alert('❌ Error al enviar el pedido')
  }
  selectedMenu.value = null
  address.value = ''
  deliveryDate.value = ''
}
</script>

<style>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active en versiones anteriores de Vue */ {
  opacity: 0;
}

/* Animaciones para los círculos y el overlay */
@keyframes float1 {
  0%, 100% { transform: translateY(0) scale(1); }
  50% { transform: translateY(-20px) scale(1.05); }
}
@keyframes float2 {
  0%, 100% { transform: translateY(0) scale(1); }
  50% { transform: translateY(30px) scale(1.08); }
}
@keyframes float3 {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.12); }
}
@keyframes float4 {
  0%, 100% { transform: translateX(0) scale(1); }
  50% { transform: translateX(20px) scale(1.07); }
}
@keyframes float5 {
  0%, 100% { transform: translateX(0) scale(1); }
  50% { transform: translateX(-25px) scale(1.09); }
}
@keyframes bgMove {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.85; }
}
.animate-float1 { animation: float1 7s ease-in-out infinite; }
.animate-float2 { animation: float2 9s ease-in-out infinite; }
.animate-float3 { animation: float3 11s ease-in-out infinite; }
.animate-float4 { animation: float4 8s ease-in-out infinite; }
.animate-float5 { animation: float5 10s ease-in-out infinite; }
.animate-bgMove { animation: bgMove 12s ease-in-out infinite; }
</style>
