<template>
  <div id="profile">
      <h2>Perfil</h2>

      <p>Nombre de usuario</p>
      <p class="info">El nombre de usuario es único y te identifica en la aplicación.</p>
      <input type="text" v-model="userCopy.username" @input="changeUserName" autocomplete="off" autocorrect="off" minlength="3" maxlength="20">

      <p>Email</p>
      <p class="info">Recuerda que también puedes usar tu correo electrónico para acceder a la aplicación.</p>
      <input type="text" :value="userCopy.email" autocomplete="off" autocorrect="off" minlength="3" maxlength="20">

      <hr>

      <p>Cambiar contraseña</p>
      <p v-if="!same" class="info">Introduce tu contraseña antigua para cambiarla.</p>
      <p v-else class="info">Introduce la nueva contraseña.</p>

      <span v-if="pwdErr" class="errorMsg" style="margin-bottom:15px;">{{message}}</span>

      <input id="profilePwd" class="pwdTgg" type="password" v-model="userCopy.password" v-if="!same" autocomplete="off" minlength="1" maxlength="50">
      <div v-if="same">
        <label>Nueva contraseña
          <input class="pwdTgg" type="password" v-model="newPass" autocomplete="off" minlength="1" maxlength="50">
        </label>
        <label>Repetir contraseña
          <input class="pwdTgg" type="password" v-model="checkNP" autocomplete="off" minlength="1" maxlength="50">
        </label>
      </div>

      <div id="pwdButtons">
        <label>
          <input @change="togglePwd" type="checkbox" v-model="show">
          <span v-if="show">Ocultar</span>
          <span v-else>Mostrar</span>
        </label>
        <button @click="comparePass">Cambiar contraseña</button>
      </div>

      <hr>

      <p>Theme</p>
      <label>
        <input @change="checkTheme" type="checkbox" v-model="userCopy.theme">
        <span v-if="!userCopy.theme">Oscuro</span>
        <span v-else>Claro</span>
      </label>

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

      pwdErr:false,
      message: "",

    }
  },
  methods: {
    changeUserName(){ this.$emit('changeUser',this.userCopy.username); },
    changeTheme(iswhite){
      if(iswhite)
          $('#app').addClass('lightTheme');
      else
          $('#app').removeClass('lightTheme');
    },
    checkTheme(e){
        if(e.target.checked){ this.changeTheme(true);   this.userCopy.theme = true; }
        else{                 this.changeTheme(false);  this.userCopy.theme = false; }
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
            this.message = GetErrMsg('missPass');
          }else{
            this.pwdErr = false;
            this.message = "";
          }
        });
      }else{ //the password is the same
        if(this.newPass != this.checkNP){ //passwords are differents
          this.pwdErr = true;
          this.message = GetErrMsg('diffPass');
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
    }
  },
}

</script>

<style lang="scss">
#profile{
    display:flex;
    flex-direction:column;

    width:inherit;

    overflow-y:scroll;

    h2{ margin-bottom:15px; }

    input { margin-bottom:10px; }

    .info{
        color:var(--tertiary);
        font-size:.8rem;
        margin:5px 0 10px 0;
    }

    label span{ margin-left:15px; }
    #pwdButtons{
        label{ margin-right:15px; }
    }
}
</style>
