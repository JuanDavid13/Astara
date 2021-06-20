import axios from 'axios'
import router from '@/router';

/**
 * Axios es una instancia de axios, que contiene una configuración global
 * para realizar llamadas a la API de una mejor forma.
 *
 * @type {axios}
 *
 */
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});

//interceptor
Axios.interceptors.response.use((res)=>{
  return res;
},(error)=>{ 
  if(error.response.status == 401)
    router.push({name:'Login'});
 return Promise.reject(error);
});

//this function complements the interceptor
/**
 * Comprueba si el estatus de la respuesta está
 * dentro del rango de los 200.
 *
 * @function
 * @name Validate
 * @returns { bool } true si está en el rango de los 200, false, si no.
*/
export function Validate() {
  return Axios.get("/auth/validate").then((res)=>{
    if (/2.{2}/.test(res.status))
      return true;
    else
      return false;
  });
}

/**
 * Comprueba que el área corresponde al usuario.
 *
 * @function
 * @param { string } Area - Slug del area
 * @returns { string } data - Datos realtivos al área.
*/
export function AreaCorrespond(Area) {
  return Axios.post("/area/correspond",{slug:Area}).then((res)=>{
    if(!res.data.correspond)
      router.push({name:'Main'});
    else
      return res.data;
  });
}

export default Axios;
