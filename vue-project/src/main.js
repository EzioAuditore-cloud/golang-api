import {createApp} from 'vue'
import 'ant-design-vue/dist/antd.css';
import { Button, Form, Input, Table } from 'ant-design-vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './assets/css/style.css'
import config from './config/axios'
const app = createApp(App)
axios.defaults.baseURL = config.WEB_API_PATH
app.config.globalProperties.$http = axios
app.use(router).mount('#app')
app.use(Button)
app.use(Form)
app.use(Input)
app.use(Table)