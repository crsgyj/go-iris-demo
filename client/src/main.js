import Vue from 'vue'

import App from './App'
import router from './router'
import './styles/main.styl'
import Container from './components/container'
import { Button, Input, Table, Pagination, TableColumn, Dialog, Form, FormItem, Autocomplete, Popover, Loading, Message } from 'element-ui';
import axios from 'axios';

Vue.use(Button)
Vue.use(Input)
Vue.use(Table)
Vue.use(Pagination)
Vue.use(TableColumn)
Vue.use(Dialog)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Autocomplete)
Vue.use(Popover)
Vue.use(Loading)
// Vue.use(Message)
Vue.prototype.$message = Message;

Vue.prototype.$config = {
  Api: "http://127.0.0.1:12350/admin",
  baiduFY: "http://api.fanyi.baidu.com/api/trans/vip/translate"
}
Vue.prototype.$api = axios.create({
  baseURL: Vue.prototype.$config.Api,
  withCredentials: true
})
Vue.prototype.$http = axios.create()

const cmptHelper = (name, cpmt) => ({
  install: Vue => Vue.component(name, cpmt)
})

Vue.use(cmptHelper('Container', Container))

new Vue({
  el: '#app',
  render: h => h(App),
  router
})