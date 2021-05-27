<template>
  <div class="home">
    <Sidebar @openprofile="showModal" :areas=areas />
    <router-view :key="$route.fullPath"></router-view>

    <div @click="hideModal" id="modal">
      <div id="modalCont">
        <Profile />
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
import Sidebar from '@/components/main/Sidebar.vue'
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
      areas: {},
    }
  },
  methods: {
    hideModal(e){
      if($('#modal').get(0) == (e.srcElement)){ $('#modal').removeClass('modalActive'); } 
    },
    showModal(){
      $('#modal').addClass('modalActive');
    },
    getAreas(){
      Axios.get('/area').then((res)=>{ this.areas = JSON.parse(res.data); });
    },
    openProfile() {
      console.log("open from parent");
    },
  },
  created(){
   this.getAreas();
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
    top:15vh;
    left:50%;
    transform:translateX(-50%);

    background-color:var(--primary);
    box-shadow:0 0 10px rgba(0,0,0,.2);
    border:1px solid var(--tertiary);
    border-radius:5px;

    //width:60%;
    //height:40vh;
    height:fit-content;
    width:auto;

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
  margin: 10px 0;
}

</style>
