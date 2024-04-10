import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import Photo from './components/Photo.vue'
import UserModal from './components/UserModal.vue'
import CommentModal from './components/CommentModal.vue'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Photo", Photo);
app.component("UserModal", UserModal);
app.component("CommentModal", CommentModal);
app.use(router)
app.mount('#app')
