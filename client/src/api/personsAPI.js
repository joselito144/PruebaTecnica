import axios from 'axios'


export const PERSONAPI = (_token) => {
  let headers = {}


  const instance = axios.create({
    baseURL: 'http://localhost:8080/api/',
    headers,
  })


  return instance
}

export default PERSONAPI
