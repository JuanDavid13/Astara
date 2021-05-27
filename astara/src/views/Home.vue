<template>
  <div class="home">
    <Sidebar @openprofile="showModal" />
    <router-view :key="$route.fullPath"></router-view>

    <div @click="hideModal" id="modal">
      <div id="modalCont">
        <Profile 
           :username="user.username"
           :email="user.email"
           :theme="user.theme"
        />
        <div id="modalButtons">
          <button>Aceptar</button>
          <button>Cancelar</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
// @ is an alias to /src
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
        theme:false
      },
    }
  },
  methods: {
    hideModal(e){
      if($('#modal').get(0) == (e.srcElement)){ $('#modal').removeClass('modalActive'); } 
    },
    showModal(){
      $('#modal').addClass('modalActive');
    },
  },
  created(){
    Axios.get("/user/info").then((res)=>{ 
      this.user.username = res.data.name;
      this.user.email = res.data.email;
      this.user.theme = res.data.theme;
    });
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

    display:grid;
    grid-template-columns:1fr;
    grid-template-rows:minmax(0,9fr) 1fr;

    height:fit-content;
    max-height:80vh;
    width:60vw;
    max-width:90vw;

    overflow-x:hidden;

    & div:first-child{
      padding:15px;
    }

    #modalButtons{
      border-top:1px solid var(--tertiary);
      padding:15px;

      button{
        margin-right:15px;
      }
    }

  }
}

.modalActive {
  display:block !important;
}

hr{
  border:none;
  border-top:1px solid var(--tertiary);
  opacity:.3;
  margin: 10px 0;
}

</style>
