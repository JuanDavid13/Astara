<template>
  <div id="profile">
    <h2>Perfil</h2>
    
    <span v-if="err" class="errorMsg" style="margin-bottom:15px;">{{errMsg}}</span>

    <form @submit.prevent >
      <label>Nombre de usuario
        <p class="info">El nombre de usuario es único y te identifica en la aplicación.</p>
        <input type="text" name="username" v-model="userCopy.username" @input="changeUserName" autocomplete="off" spellcheck="false" minlength="4" maxlength="30">
      </label>

      <label>Email
        <p class="info">Recuerda que también puedes usar tu correo electrónico para acceder a la aplicación.</p>
        <input type="text" name="email" v-model="userCopy.email" autocomplete="off" spellcheck="false" minlength="5" maxlength="50">
      </label>

      <hr>

      <label>Cambiar contraseña
        <p v-if="!same" class="info">Introduce tu contraseña antigua para cambiarla.</p>
        <p v-else class="info" style="margin-bottom:5px;">Introduce la nueva contraseña.</p>

        <span v-if="pwdErr" class="errorMsg" style="margin-bottom:15px;">{{pwdMsg}}</span>

        <input id="profilePwd" class="pwdTgg" type="password" v-model="userCopy.password" v-if="!same" autocomplete="off" minlength="7" maxlength="50">
        <div v-if="same">
          <label>Nueva contraseña
            <input class="pwdTgg" type="password" v-model="newPass" autocomplete="off" minlength="7" maxlength="50">
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
      userCopy: {
        username:"",
        email:"",
      },
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
      if(iswhite) $('#app').addClass('lightTheme');
      else $('#app').removeClass('lightTheme');
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
      if($('#checkNP').val() == $('#newPass').val()) this.setError('noError');
      else this.setError('diffPass');
    },
    checkUser(){
      if(this.originalName == this.userCopy.username) return null;

      if(this.userCopy.username == "") return -1

      if(this.userCopy.username.length < 6) return -2

      return 0;
    },
    checkEmail(){
      if(this.user.email == this.userCopy.email) return null; //equal
      
      if(this.userCopy.email == "") return -1; //deleted

      if(this.userCopy.email.length < 5) return -2;

      if(this.userCopy.email.indexOf('@') === -1) return -3;

      return 0; //added or updated
    },
    checkInputs(){
      let changes = {
        username:null,
        email:null,
        theme:null
      };

      let userErr = "";
      let emailErr = "";

      switch(this.checkUser()){
        case null:  { changes.username = null; }break;
        case -1:    { userErr = GetErrMsg('emptyuser'); }break;
        case -2:    { userErr = GetErrMsg('shortUserName'); }break;
        case 0:     { changes.username = this.userCopy.username; }break;
        default:    { userErr = GetErrMsg(); }break;
      }

      switch(this.checkEmail()){
        case null:  { changes.email = null; }break;
        case -1:    { changes.email = this.userCopy.email; }break;
        case -2:    { emailErr = GetErrMsg('shortEmail'); }break;
        case -3:    { emailErr = GetErrMsg('wrongEmail'); }break;
        case 0:     { changes.email = this.userCopy.email; }break;
        default:    { emailErr = GetErrMsg(); }break;
      }

      if(this.user.theme == this.userCopy.theme) changes.theme = null;

      if(this.user.theme != this.userCopy.theme) changes.theme = this.userCopy.theme;

      //genrate error message
      if(userErr.length > 0 && emailErr.length > 0){
        userErr = userErr.charAt(0).toUpperCase() + userErr.slice(1);
        this.errMsg = userErr + ' , ' + emailErr;
      }else{
        userErr = userErr.charAt(0).toUpperCase() + userErr.slice(1);
        emailErr = emailErr.charAt(0).toUpperCase() + emailErr.slice(1);
        this.errMsg = userErr + emailErr;
      }

      if((userErr.length + emailErr.length) != 0)
        return null; 

      return changes;
    },
    submitData(){

      let changes = this.checkInputs();

      if(changes == null){
        this.err = true;
        return;
      }
      this.err = false;

      if(changes.username == null && changes.email == null && changes.theme == null)
        return;

      Axios.post("user/profile/update",{
        //username: changes.username,
        //email: changes.email,
        //theme: changes.theme,
        changes: changes
      }).then((res)=>{
        if(!res.data.updated){
          this.err = true;
          this.errMsg = GetErrMsg('updateErr');
        }else{
          this.originalName = this.userCopy.username;
          $('#modal').removeClass('modalActive');
          this.closeModal();
        }
      });

    },
    closeModal(){
      //default values;
      this.newPass = "",
      this.checkNP = "",

      this.show = false,
      this.same = false,

      this.err = false,
      this.errMsg = "",

      this.pwdErr = false,
      this.pwdMsg = "",

      this.$emit('changeUser',this.originalName);
      this.changeTheme(this.user.theme);

      $('#modal').removeClass('activeModal');
    },
    openModal(){
      this.userCopy = $.extend(true,{},this.user);
      this.originalName = this.userCopy.username;
    },
  },
  mounted(){ this.openModal() },
  unmounted(){ this.closeModal(); }

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
