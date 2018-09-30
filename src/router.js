import Vue from 'vue'
import Router from 'vue-router'
import Sign from './views/Sign.vue'
import Profile from './views/Profile.vue'
import Tasks from './views/Tasks.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Sign
    },
    {
      path: '/signin',
      name: 'signin',
      component: Sign
    },
    {
      path: '/signup',
      name: 'signup',
      component: Sign
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
    },
    {
      path: '/tasks',
      name: 'tasks',
      component: Tasks
    }
  ]
})
