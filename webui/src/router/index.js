import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SearchView from '../views/SearchView.vue'
import ProfileView from '../views/ProfileView.vue'
import PersonalProfileView from '../views/PersonalProfileView.vue'
import SettingsView from '../views/SettingsView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/home', component: HomeView},
		{path: '/search', component: SearchView},
		{path: '/personalProfile', component: PersonalProfileView},
		{path: '/profile/:username', component: ProfileView},
		{path: '/settings', component: SettingsView},
	]
})

export default router
