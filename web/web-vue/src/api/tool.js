import axios from '@/axios/axios.js'

export const toolList = (data) => {
  return axios.get('/api/tool/list', {
    params: data
  })
}

export const saveTool = (data) => {
  return axios.post('/api/tool/save', data)
}

export const delTool = (data) => {
  return axios.post('/api/tool/del', data)
}