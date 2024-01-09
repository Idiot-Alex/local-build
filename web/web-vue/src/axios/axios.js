import axios from 'axios'
import { app } from '@/main.js'

const instance = axios.create({
  baseURL: 'http://localhost:8000/',
  timeout: 1000 * 10
})

// 添加请求拦截器
instance.interceptors.request.use((config) => {
  // 在发送请求之前做些什么
  return config
}, function (error) {
  console.log(error)
  // 对请求错误做些什么
  return Promise.reject(error)
});

// 添加响应拦截器
instance.interceptors.response.use((response) => {
  console.log('response: ', response)
  // 2xx 范围内的状态码都会触发该函数。
  // 对响应数据做点什么
  if (response.status === 200) {
    return response.data
  }
  return response
}, (error) => {
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  app.config.globalProperties.$toast.add({
    severity: 'error',
    summary: error.message,
    detail: `${error.config.url}`,
    method: error.config.method,
    life: 3000,
    group: 'network'
  })
  return Promise.reject(error)
});

export default instance