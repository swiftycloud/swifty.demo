import Vue from 'vue'
import Router from 'vue-router'
import Sign from './views/Sign.vue'
import Profile from './views/Profile.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Sign
    },
    {
      path: '/sign',
      name: 'sign',
      component: Sign
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
    }
  ]
})
