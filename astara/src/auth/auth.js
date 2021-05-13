import axios from 'axios'
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});

//return if the token is valid or not
//valid = true
//not valid = false
function validate() {
  return Axios.get("/login/validate").then((res)=>{
    if(res.data.valid == "true")
      return true;
    return false;
  });
}

function AreaCorrespond(Area) {
  return Axios.post("/area/correspond",{name:Area}).then((res)=>{
    console.log(res);
    return res.data;
  });
}

export {validate as validate, AreaCorrespond as AreaCorrespond};
