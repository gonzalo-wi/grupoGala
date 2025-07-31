/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",   // 👈 importante incluir .vue y .ts
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}