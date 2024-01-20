import axios from '@/axios/axios.js'

export const buildPlanList = (data) => {
  return axios.post('/api/build-plan/list', data)
}

export const saveBuildPlan = (data) => {
  return axios.post('/api/build-plan/save', data)
}

export const delBuildPlan = (data) => {
  return axios.post('/api/build-plan/del', data)
}