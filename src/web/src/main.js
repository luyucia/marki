import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import './assets/default.css'
Vue.config.productionTip = false

Vue.use(VueResource)
Vue.use(ElementUI);

new Vue({
  render: h => h(App),
}).$mount('#app')
