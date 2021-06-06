import axios from 'axios'
import router from '@/router';

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
export function Validate() {
  return Axios.get("/auth/validate").then((res)=>{
    if (/2.{2}/.test(res.status))
      return true;
    else
      return false;
  });
}

export function AreaCorrespond(Area) {
  return Axios.post("/area/correspond",{slug:Area}).then((res)=>{
    if(!res.data.correspond)
      router.push({name:'Main'});
    else
      return res.data;
  });
}

export default Axios;
