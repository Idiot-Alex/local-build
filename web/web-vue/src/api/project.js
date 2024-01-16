import axios from '@/axios/axios.js'

export const projectList = (data) => {
  return axios.post('/api/project/list', data)
}

export const saveProject = (data) => {
  return axios.post('/api/project/save', data)
}

export const delProject = (data) => {
  return axios.post('/api/project/del', data)
}