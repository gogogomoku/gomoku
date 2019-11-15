import Vue from 'vue'
import App from './App.vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faRobot, faTimesCircle, faUser, faCrown } from '@fortawesome/free-solid-svg-icons'
import { faTimesCircle as farTimesCircle } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faTimesCircle, faRobot, faUser, farTimesCircle, faCrown)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
