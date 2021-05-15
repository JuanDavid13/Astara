import axios from 'axios'
import router from '@/router';

const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});

//interceptor
Axios.interceptors.response.use((res)=>{
  console.log(res);
},(error)=>{
  console.log("next");
  console.log(router);
  router.push({name:'Login'});
  console.log(error);
});

//this function complements the interceptor
function Validate() {
  return Axios.get("/auth/validate").then((res)=>{
    console.log("validate");
    console.log(res);
    //return res.data;
    return true;
  });
}

function AreaCorrespond(Area) {
  return Axios.post("/area/correspond",{name:Area}).then((res)=>{
    console.log(res);
    return res.data;
  });
}

export { Validate, AreaCorrespond };
