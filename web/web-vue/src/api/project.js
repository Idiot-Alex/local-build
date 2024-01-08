import axios from 'axios'

export const projectList = (params) => {
  return axios.get('/api/project/list', {params})
}

export const saveProject = (data) => {
  return axios.post('/api/project/save', {data})
}

export const delProject = (data) => {
  return axios.post('/api/project/del', {data})
}