import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// Plugins
import { registerPlugins } from '@/plugins'

const app = createApp(App)

app.use(createPinia())
app.use(router)

registerPlugins(app)

app.mount('#app')
