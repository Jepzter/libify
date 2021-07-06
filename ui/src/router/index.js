import { createRouter, createWebHistory } from 'vue-router'
import Landing from "@/components/Landing.vue";
import Playlist from "@/components/Playlist.vue";
import Settings from "@/components/Settings.vue";

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
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router;