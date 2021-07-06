import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import './styles/reset.css'
import './styles/app.css'

createApp(App).use(router).use(store).mount('#app')
