<template>
  <div id="profile">
    <h2>Perfil</h2>
    
    <span v-if="err" class="errorMsg" style="margin-bottom:15px;">{{errMsg}}</span>

    <form @submit.prevent >
      <label>Nombre de usuario
        <p class="info">El nombre de usuario es único y te identifica en la aplicación.</p>
        <input type="text" name="username" v-model="userCopy.username" @input="changeUserName" autocomplete="off" spellcheck="false" minlength="3" maxlength="30">
      </label>

      <label>Email
        <p class="info">Recuerda que también puedes usar tu correo electrónico para acceder a la aplicación.</p>
        <input type="text" name="email" v-model="userCopy.email" autocomplete="off" spellcheck="false" minlength="3" maxlength="50">
      </label>

      <hr>

      <label>Cambiar contraseña
        <p v-if="!same" class="info">Introduce tu contraseña antigua para cambiarla.</p>
        <p v-else class="info" style="margin-bottom:5px;">Introduce la nueva contraseña.</p>

        <span v-if="pwdErr" class="errorMsg" style="margin-bottom:15px;">{{pwdMsg}}</span>

        <input id="profilePwd" class="pwdTgg" type="password" v-model="userCopy.password" v-if="!same" autocomplete="off" minlength="1" maxlength="50">
        <div v-if="same">
          <label>Nueva contraseña
            <input class="pwdTgg" type="password" v-model="newPass" autocomplete="off" minlength="1" maxlength="50">
          </label>
          <label>Repetir contraseña
            <input class="pwdTgg" type="password" v-model="checkNP" autocomplete="off" minlength="1" maxlength="50">
          </label>
        </div>
      </label>

      <div id="pwdButtons">
        <label class="inputRow">
          <input @change="togglePwd" type="checkbox" v-model="show">
          <span v-if="show">Ocultar</span>
          <span v-else>Mostrar</span>
        </label>
        <button v-if="!same" @click="comparePass">Comprobar contraseña</button>
        <button v-else @click="comparePass">Cambiar contraseña</button>
      </div>

      <hr>

      <p>Theme</p>
      <label class="inputRow">
        <input @change="checkTheme" type="checkbox" v-model="userCopy.theme">
        <span v-if="!userCopy.theme">Oscuro</span>
        <span v-else>Claro</span>
      </label>

      <div id="modalPwdButtons">
        <button @click="submitData">Aceptar</button>
        <button @click="closeModal" >Cancelar</button>
      </div>
    </form>

    <button id="logOut" @click="logOut">Salir</button>
  </div>
</template>

<script>
import { GetErrMsg } from '@/js/error.js';
import Axios from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'Profile',
  props: {
    user: Object,
  },
  emits:[ 'changeUser'],
  data() {
    return {
      userCopy: {},
      originalName:"",
      newPass:"",
      checkNP:"",

      show:false,
      same: false,

      err:false,
      errMsg:"",

      pwdErr:false,
      pwdMsg: "",

    }
  },
  methods: {
    logOut(){ Axios.get('/auth/logout').then(()=>{ this.$router.push({name:'Login'}) }); },
    changeUserName(){ this.$emit('changeUser',this.userCopy.username); },
    changeTheme(iswhite){
      if(iswhite)
          $('#app').addClass('lightTheme');
      else
          $('#app').removeClass('lightTheme');
    },
    checkTheme(e){
        if(e.target.checked){ this.changeTheme(true);   this.userCopy.theme = true; }
        else{ this.changeTheme(false);  this.userCopy.theme = false; }
    },
    togglePwd(){
      let pwdInputs = $('.pwdTgg');
      if($(pwdInputs[0]).attr('type') == "password"){ $(pwdInputs).attr('type','text'); }
      else { $(pwdInputs).attr('type','password'); }
    },
    comparePass(){
      if(!this.same){ //the password is not the same
        Axios.post('user/profile/checkpass',{ password: this.userCopy.password }).then((res)=>{
          this.same = res.data.same;
          if(!res.data.same){
            this.pwdErr = true;
            this.pwdMsg = GetErrMsg('missPass');
          }else{
            this.pwdErr = false;
            this.pwdMsg = "";
          }
        });
      }else{ //the password is the same
        if(this.newPass != this.checkNP){ //passwords are differents
          this.pwdErr = true;
          this.pwdMsg = GetErrMsg('diffPass');
        }else
          this.pwdErr = false;
      }
    },
    checkEqual(){
      if($('#checkNP').val() == $('#newPass').val())
        this.setError('noError');
      else
        this.setError('diffPass');
    },
    checkInputs(){
      let userErr = "";
      let emailErr = "";

      if(this.userCopy.username.length < 3)
        userErr = GetErrMsg('shortUserName');

      if(this.userCopy.email.length < 5 || !this.userCopy.email.includes('@'))
        emailErr = GetErrMsg('wrongEmail');

      if(userErr.length > 0 && emailErr.length > 0){
        userErr = userErr.charAt(0).toUpperCase() + userErr.slice(1);
        this.errMsg = userErr + ' , ' + emailErr;
      }
      else
        this.errMsg = userErr + emailErr;

      if((userErr.length + emailErr.length) == 0)
        return true

      return false;
    },
    submitData(){
      if(!this.checkInputs()){
        this.err = true;
      }else{
        this.err = false;
      }
    },
    closeModal(){
      this.show = false;
      this.same =  false;

      this.pwdErr = false;
      this.message =  "";

      this.newPass = "";
      this.checkNP = "";

      this.$emit('changeUser',this.originalName);
      this.changeTheme(this.user.theme);
    },
    openModal(){
      this.userCopy = $.extend(true,{},this.user);
      this.originalName = this.userCopy.username;
    },
  },
  mounted(){
    this.openModal()
    console.log('mounted');
  },
  unmounted(){
    this.closeModal()
    console.log('unmounted');
  }

}

</script>

<style lang="scss">
#profile{
  display:flex;
  flex-direction:column;

  width:inherit;

  overflow-y:scroll;

  h2{ margin-bottom:15px; }

  input { margin:10px 0; }

  .info{
    color:var(--tertiary);
    font-size:.8rem;
    margin-top:5px;
  }

  label{
    display:flex;
    flex-direction:column;
  }
  .inputRow{ flex-direction:row; }

  #pwdButtons label{ margin-right:15px; }

  //bottom buttons
  #modalPwdButtons{
    display:flex;
    flex-direction:row;
    gap:.7rem;
  }
  #logOut{
    position:absolute;
    top:15px;
    right:15px;
  }
}
</style>
