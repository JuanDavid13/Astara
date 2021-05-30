<template>
  <div class="home">
    <Sidebar @openprofile="openProfile" :username="user.username"/>
    <router-view :key="$route.fullPath"></router-view>
    <div @click="hideModal" id="modal">
      <div id="modalCont">
        <Profile 
           v-if="modalOption == 1"
           ref="profile"
           :user="user"
           @changeUser="changeUser"
        />
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import { setShortCodes } from '@/js/shortcodes';

import Profile from '@/components/modals/profile.vue';
import Sidebar from '@/components/main/Sidebar.vue';
import Axios from '@/auth/auth';

import "splitting/dist/splitting.css";
import Splitting from "splitting";

import $ from 'jquery';


export default {
  name: 'Home',
  components: {
    Sidebar,
    Profile,
  },
  data() {
    return{
      user:{
        username:"",
        email:"",
        theme:true,
      },
      modalOption:0,
    }
  },
  methods: {
    changeUser(username){
      this.user.username = username;
    },
    openProfile(){
      $('#modal').addClass('modalActive');
      this.modalOption = 1;
      $(document).ready(()=>{ this.$refs.profile.openModal(); });
    },
    hideModal(e){
      if($('#modal').get(0) == (e.srcElement)){
        $('#modal').removeClass('modalActive');
        this.$refs.profile.closeModal();
      } 
    },
    closeModal(){
      $('#modal').removeClass('modalActive');
      this.$refs.profile.closeModal();
    }
  },
  created(){
    Axios.get("/user/profile/info").then((res)=>{ 
      this.user.username = res.data.name;
      this.user.email = res.data.email;
      this.user.theme = res.data.theme;

      if(res.data.theme)
        $('#app').addClass("lightTheme");
    });

    setShortCodes(this);
  },
  mounted(){
    Splitting();
  }
}
</script>

<style lang="scss">

@import '@/assets/style/common';

.home{
  color:var(--text);
  background-color:var(--primary);
  min-height:100vh;

  display:grid;
  grid-template-columns: minmax(0,2fr) 9fr;
  grid-template-rows:1fr;
}

@keyframes effect {
  from{ transform: translateY(2rem); }
}

.splitChar{overflow:hidden;}

.splitChar.splitting .char{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.02s * var(--char-index));
}

#modal{
  //make shadow a variable in common.scss
  background-color:rgba(0,0,0,.5);
  position:absolute;
  height:100%;
  width:100%;
  top:0;
  left:0;

  display:none;

  #modalCont{
    position:absolute;
    top:10vh;
    left:50%;
    transform:translateX(-50%);

    background-color:var(--primary);
    box-shadow:0 0 10px rgba(0,0,0,.2);
    border:1px solid var(--tertiary);
    border-radius:5px;

    height:fit-content;
    max-height:80vh;
    width:60vw;
    max-width:90vw;

    overflow-x:hidden;

    & div:first-child{
      padding:15px;
    }
  }
}

.modalActive { display:block !important; }

hr{
  border:none;
  border-top:1px solid var(--tertiary);
  opacity:.3;
  margin: 10px 0;
}
</style>
