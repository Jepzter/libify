import { createRouter, createWebHistory } from 'vue-router'
import Landing from "@/components/Landing.vue";
import Playlist from "@/components/Playlist.vue";
import Settings from "@/components/Settings.vue";
import SearchSongs from "@/components/SearchSongs.vue";

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: Landing
  },
  {
    path: '/playlist/:id',
    name: 'Playlist',
    component: Playlist,
    props: true,
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings
  },
  {
    path: '/search/:search',
    name: 'SearchSongs',
    component: SearchSongs,
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router;