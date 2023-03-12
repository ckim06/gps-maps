/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import VueGoogleMaps from '@fawmi/vue-google-maps'
import { mapKey } from '../mapkey'
const app = createApp(App)

registerPlugins(app)
app.use(VueGoogleMaps, {
    load: {
        key: mapKey,
    },
}).mount('#app')
