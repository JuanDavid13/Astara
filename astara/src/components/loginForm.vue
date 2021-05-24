<template>
  <div id="formWrap">
    <div id="formAction">
      <transition name="swipe" mode="out-in">
        <span class="loginAction" v-if="signUp">Registrate</span>
        <span class="loginAction" v-else>Entrar</span>
      </transition>
    </div>
    <span v-if="error" class="errorMsg">{{message}}</span>
    <form @submit.prevent class="form">
      <label v-if="signUp"><span>Nombre de usuario <span class="mandatory">*</span></span><input type="text" v-model="user"></label>
      <input v-else type="text" v-model="user" placeholder="Nombre de usuario o email">

      <input id="pass" v-if="found && !signUp" type="password" v-model="pass" placeholder="contraseña">

      <div id="signUpForm" v-if="signUp">
        <label><span>Correo electrónico</span>
          <input type="text" v-model="email">
        </label>
        <label><span>Contraseña <span class="mandatory">*</span></span>
          <input id="SUpass" type="password" @input="checkEqual" v-model="pass">
        </label>
        <label id="lastLabel"><span>Repite la contraseña <span class="mandatory">*</span></span>
          <input id="SUcheckpass" @input="checkEqual" type="password" v-model="checkpass">
        </label> 
        <span><span class="mandatory">*</span> Significa que es obligatorio</span>
      </div>

      <button type="submit" @click="clicked">Aceptar</button>
    </form>
    <p v-if="signUp" style="text-align:center;">¿Ya eres usuario de Astara? <a href="#" @click.prevent="SignIn" class="formLink" >Entra</a></p>
    <p v-else style="text-align:center;">¿Nuevo en Astara? <a href="#" @click.prevent="SignUp" class="formLink">Hazte una cuenta</a></p>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import $ from 'jquery';

export default{
    name: 'LoginForm',
    data() {
      return {
        user: "",
        pass: "",
        checkpass: "",
        email: "",
        signUp:false,
        found: false,
        error: false,
        message: "",
      }
    },
    methods:{
      setError(errorName){
        switch(errorName) {
          case 'noError': 
            this.error = false;
            break;
          case 'noUser': {
            this.error = true;
            this.message = "El usuario introducido no existe";
          }break;
          case 'missPass': {
            this.error = true;
            this.message = "La contraseña introducida es incorrecta";
          }break;
          case 'diffPass':{
            this.error = true;
            this.message = "Las contraseñas no coinciden";
          }break;
          case 'lackInput': {
            this.error = true;
            this.message = "Debe de completar todos los campos obligatorios";
          }break;
          case 'alreadyCreated': {
            this.error = true;
            this.message = "Ya existe un usuario con ese nombre o email";
          }break;
          default:{
            this.error = true;
            this.message = "Ha ocurrido un error";
          }
        }
      },
      clicked(){
        if(this.signUp){
          console.log("nose");
          Axios.post('login/create',{ 
            user: this.user,
            pass: this.pass,
            email: this.email,
          }).then((res)=>{
            if(res.data.error == true){
              if(res.data.alreadyCreated == true){
                this.setError('alreadyCreated');
                this.clearVal();
              }else{ this.setError(); this.clearVal(); }
            }else{
              if(res.data.created == true)
                this.$router.push({name: 'Home'});
            }
            console.log(res);
          })
        }else{
          if(this.pass.localeCompare("") == 0){
            Axios.post('login/',{ user: this.user }).then((res)=>{
              $('#pass').focus();
              if(res.data.found == "true"){
                this.found = true;
                this.setError('noError');
              }else{ 
                this.found = false;
                this.setError('noUser'); } });
          }else{
            Axios.post('login/check',{ user: this.user, pass: this.pass }).then((res)=>{
              if(res.data.logged == "true"){
                this.$router.push({name:'Home'});
              }else{ this.setError('missPass'); } });
          }
        }
      },
      SignUp(){ this.error = false; this.signUp = true; },
      SignIn(){ this.signUp = false; this.clearVal(); },
      checkEqual(){
        if($('#SUcheckpass').val() == $('#SUpass').val()){
          this.setError('noError');
        }else{
          this.setError('diffPass');
        }
      },
      clearVal(){
        this.pass="";
        this.checkpass="";
        this.email="";
        this.found = false;
      }
    }
  }

</script>

<style scoped lang="scss">
#formWrap{
  border:1px solid var(--contrary);
  border-radius:7px;
  padding: 15px;
  height:fit-content;

  box-shadow: 0 0 20px rgba(0,0,0,.2);

  label{
    display:flex;
    flex-direction:column;
    margin-bottom:15px;
  }

  //no me sirve hacer un lable:last-child
  #lastLabel{ margin-bottom:15px; }
  #pass{margin-top:15px;}

  p:last-child{ margin-top:15px; }
  button{ margin-top:15px; }
}

form{
  display: flex;
  flex-direction: column;
}

.formLink{
  color:var(--gold);
  cursor:pointer;
  text-decoration:none;

  &:focus, &:hover{ text-decoration:underline var(--gold); }
}

.loginAction{
  display:inline-block;
  font-size:1.5rem;
  margin-bottom:15px;
}

.swipe-enter-active, .swipe-leave-active {
  transition: all .25s ease;
}

.swipe-enter-from, .swipe-leave-to{
  transform: translateX(1rem);
  opacity:0;
}

.mandatory{ color:var(--errorTxt); }

.errorMsg{
  display:block;
  width:70%;
  margin:0 auto 15px auto;
}
</style>
