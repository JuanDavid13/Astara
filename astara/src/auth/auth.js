import axios from 'axios'

//return if the token is valid or not
//valid = true
//not valid = false
function validate() {
  return axios.get("http://localhost:3000/api/v1/login/validate", {withCredentials:true}).then((res)=>{
    if(res.data.valid == "true")
      return true;
    else
      return false;
  });
}

export {validate as validate};
