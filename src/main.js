import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import ElementUI from 'element-ui'
import './styles.scss'
import router from './router'

Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
