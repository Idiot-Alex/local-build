import axios from 'axios'

export const initEnv = () => {
  return axios.get('/api/env/init')
}