import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify' // Import the config from above
import router from './router/index.js'

const app = createApp(App)

app.use(router)
app.use(vuetify)
app.mount('#app')
