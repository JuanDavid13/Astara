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
      <label v-if="signUp"><span>Nombre de usuario <span class="mandatory">*</span></span><input type="text" v-model="user.username"></label>
      <input v-else type="text" v-model="user.username" placeholder="Nombre de usuario o email">

      <input id="pass" class="pwdTgg" v-if="found && !signUp" type="password" v-model="user.pass" placeholder="contraseña">
      <label v-if="found && !signUp">
        <input type="checkbox" @change="togglePwd">
        <span>Mostar contraseña</span>
      </label>

      <div id="signUpForm" v-if="signUp">
        <label><span>Correo electrónico</span>
          <input type="text" v-model="user.email">
        </label>
        <label><span>Contraseña <span class="mandatory">*</span></span>
          <input class="pwdTgg" type="password" v-model="user.pass">
        </label>
        <label id="lastLabel"><span>Repite la contraseña <span class="mandatory">*</span></span>
          <input class="pwdTgg" type="password" v-model="user.checkPass">
        </label> 
        <label>
          <input type="checkbox" @change="togglePwd">
          <span>Mostar contraseña</span>
        </label>
        <span><span class="mandatory">*</span> Significa que es obligatorio</span>
      </div>

      <button type="submit" @click="clicked">Aceptar</button>
    </form>
    <p v-if="signUp" style="text-align:center;">¿Ya eres usuario de Astara? <a href="#" @click.prevent="Sign(1)" class="formLink" >Entra</a></p>
    <p v-else style="text-align:center;">¿Nuevo en Astara? <a href="#" @click.prevent="Sign(2)" class="formLink">Hazte una cuenta</a></p>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import { GetErrMsg  } from '@/js/error.js';
import $ from 'jquery';

export default{
    name: 'LoginForm',
    data() {
      return {
        user:{
          username: "",
          email: "",
          pass: "",
          checkPass: "",
        },
        signUp:false,

        found: false,

        error: false,
        message: "",
      }
    },
    methods:{
      /**
      * Función que valida los inputs del formulario de login.
      *
      * @function
      * @returns { bool } Devuelve 'true' si los valores son validos y 'false' si no.
      */
      validateInputs(){
        if(this.signUp){
          if(this.user.username == "" || this.user.pass == "" || this.user.checkpass == ""){
            this.error = true;
            this.message = GetErrMsg('lackInput');
            return false;
          } 

          if(this.user.pass !== this.user.checkPass){
            this.error = true;
            this.message = GetErrMsg('diffPass');
            return false;
          }

          if(this.user.email != ""){
            if(this.user.email.indexOf('@') === -1){
              this.error = true;
              this.message = GetErrMsg('wrongEmail');
              return false;
            }
          }
        }else{
          if(this.user.username == ""){
            this.error = true;
            this.message = GetErrMsg('lackInput');
            return false;
          } 
        }
        return true;
      },
      /**
      * Función que logea al usuario o crea uno nuevo.
      *
      * @function
      */
      clicked(){
        if(!this.validateInputs())
          return;

        if(this.signUp){
          Axios.post('login/create',{
            user: this.user.username,
            pass: this.user.pass,
            email: this.user.email,
          }).then((res)=>{
            if(res.data.error == true){
              if(res.data.alreadyCreated == true){
                this.error = true;
                this.message = GetErrMsg('alreadyCreated');
                this.clearVal();
              }else{ 
                this.error = true;
                this.message = GetErrMsg();
                this.clearVal();
              }
            }else{
              if(res.data.created == true)
                this.$router.push({name: 'Home'});
            }
          })
        }else{
          if(this.user.pass.localeCompare("") == 0){ //otra forma de comparar dos strings
            Axios.post('login/',{ user: this.user.username }).then((res)=>{
              $(document).ready(()=>{
                $('#pass').focus();
              });
              if(!res.data.found){
                this.found = false;
                this.error = true;
                this.message = GetErrMsg('noUser');
              }else{
                this.found = true;
                this.error = false
                this.message = "";
              }
            });
          }else{
            Axios.post('login/check',{ user: this.user.username, pass: this.user.pass }).then((res)=>{
              if(res.data.logged == "true")
                this.$router.push({name:'Home'});
              else{
                this.error = true;
                this.message = GetErrMsg('missPass');
              } 
            });
          }
        }
      },
      /**
      * Función auxliar para cambiar de opción.
      *
      * @function
      * @param { int } option - Opción del formulario.
      */
      Sign(option){
        if(option==2){
          this.error = false;
          this.signUp = true;
        }else{
          this.signUp = false;
          this.clearVal(); }
      },
      /**
      * Función que vacía los inputs al cambiar de opción
      *
      * @function
      */
      clearVal(){
        this.user.pass="";
        this.user.checkPass="";
        this.user.email="";
        this.found = false;
      },
      /**
      * Función auxiliar para mostrar y ocultar la contraseña.
      *
      * @function
      */
      togglePwd(){
        let pwdInputs = $('.pwdTgg');
        if($(pwdInputs[0]).attr('type') == "password"){ $(pwdInputs).attr('type','text'); }
        else { $(pwdInputs).attr('type','password'); }
      },
    }
  }

</script>

<style scoped lang="scss">
#formWrap{
  //border-top:1px solid var(--contrary);
  border-left:1px solid var(--tertiary);

  //border-radius:7px;
  padding: 15px;
  height:fit-content;
  max-height:1000px;

  transition: all .25s ease;

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
