import { createApp } from 'vue'
// import './style.css'
import '@/assets/styles.scss';
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config'

const app = createApp(App)
app.use(router)
app.use(PrimeVue)
app.mount('#app')
