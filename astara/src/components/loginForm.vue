<template>
  <div>
    <form @submit.prevent class="form">
      <span style="color:red">{{message}}</span>
      <input type="text" ref="input" v-model="user" placeholder="Nombre de usuario">
      <input v-if="found" type="text" v-model="pass" placeholder="contraseña">
      <button type="submit" @click="clicked">Aceptar</button>
    </form>
    <div>
      <p>¿Nuevo en Astara? <span>Hazte una cuenta</span></p>
    </div>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
  export default{
    name: 'LoginForm',
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
        console.log(process.env);
        if(this.pass.localeCompare("") == 0){
          Axios.post('login/',{
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
          Axios.post('login/check',{
            user: this.user,
            pass: this.pass
          }).then((res)=>{
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

  input{margin-bottom: 15px;}

  span{color: red;}

}

</style>
