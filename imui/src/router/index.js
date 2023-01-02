import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginView.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/RegisterView.vue')
  },
  {
    path: '/art',
    name: 'art',
    component: () => import('../views/ArtView.vue')
  },
  {
    path: '/upload',
    name: 'upload',
    component: () => import('../views/UploadView.vue')
  },
  {
    path: '/search',
    name: 'search',
    component: () => import('../views/SearchView.vue')
  },
  {
    path: '/favorite',
    name: 'favorite',
    component: () => import('../views/CollectView.vue')
  },
  {
    path: '/arts',
    name: 'artsmanage',
    component: () => import('../views/ArtsManageView.vue')
  },
  {
    path: '/user',
    name: 'user',
    component: () => import('../views/UserView.vue')
  },
  {
    path: '/me',
    name: 'me',
    component: () => import('../views/UserMeView.vue')
  },
  {
    path: '/more',
    name: 'more',
    component: () => import('../views/MoreView.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
