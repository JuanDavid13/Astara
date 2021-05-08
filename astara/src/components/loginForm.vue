<template>
    <form @submit.prevent class="form">
      <span style="color:red">{{message}}</span>
      <input type="text" ref="input" v-model="user" placeholder="Nombre de usuario">
      <input v-if="found" type="text" v-model="pass" placeholder="contraseña">
      <button type="submit" @click="clicked">Aceptar</button>
    </form>
</template>

<script>
const axios = require('axios');
  export default{
    name: 'loginForm',
    data() {
      return {
        message: "",
        found: false,
        user: "",
        pass: ""
      }
    },
    methods:{
      clicked(){
        if(this.pass.localeCompare("") == 0){
          axios.post('http://localhost:3000/api/v1/login/',{
            user: this.user
          }).then((res)=>{
            if(res.data.found == "true"){
              this.message = "";
              this.found = true;
            }else{
              this.message = "El usuario introducino no existe";
              this.found = false;
            }
          });
        }else{
          axios.post('http://localhost:3000/api/v1/login/check',{
            user: this.user,
            pass: this.pass
          },{ withCredentials: true }).then((res)=>{
            if(res.data.logged == "true"){
              this.$router.push({name:'Home'});
            }else{
              this.message = "Contraseña incorrecta";
            }
          });
        }
      }
    }
  }

</script>

<style scoped lang="scss">
form{
  display: flex;
  flex-direction: column;
  width: 70%;
  margin:0 auto;

  input{margin-bottom: 15px;}

  span{color: red;}

}

</style>
