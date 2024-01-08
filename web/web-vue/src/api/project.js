import axios from 'axios'

export const projectList = (data) => {
  return axios.get('/api/project/list', data)
}

export const saveProject = (data) => {
  return axios.post('/api/project/save', data)
}

export const delProject = (data) => {
  return axios.post('/api/project/del', data)
}