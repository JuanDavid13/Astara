<template>
  <div id="home" class="home">
    <Sidebar @openprofile="openProfile" :username="user.username" ref="Sidebar"/>
    <router-view :key="$route.fullPath" @deleteArea="deleteArea" @updateAreaName="updateAreaName"></router-view>
    <div @click="hideModal" id="modal">
      <div id="modalCont">
        <Profile 
           v-if="modalOption == 1"
           ref="profile"
           :user="user"
           @changeUser="changeUser"
           @changeTheme="updateTheme"
        />
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import { setShortCodes } from '@/js/shortcodes';
import { resizeNav } from '@/js/navbarRes';

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
    /**
    * Función auxiliar para la eliminación de un area.
    *
    * @function
    * @param { string } slug - Slug del área a eliminar.
    */
    deleteArea(slug){ this.$refs.Sidebar.deleteArea(slug); },
    /**
    * Función auxiliar para la actualización del nombre del area.
    *
    * @function
    * @param { string } oldName - Nombre antiguo del área.
    * @param { string } name - Nuevo nombre del area.
    */
    updateAreaName(oldName, name){ this.$refs.Sidebar.updateAreaName(oldName, name); },
    /**
    * Función auxiliar para la actualización del nombre del usuario.
    *
    * @function
    * @param { string } username - Nuevo nombre de usuario.
    */
    changeUser(username){ this.user.username = username; },
    /**
    * Función auxiliar para el cambio de theme.
    *
    * @function
    * @param { bool } theme - 'true' si es claro, 'false' si es oscuro.
    */
    updateTheme(theme) {this.user.theme = theme},
    /**
    * Función para ablir el modal con el perfil de usuario;
    *
    * @function
    */
    openProfile(){
      $('#modal').addClass('modalActive');
      this.modalOption = 1;
      $(document).ready(()=>{ this.$refs.profile.openModal(); });
    },
    /**
    * Función para esconder el modal.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
    hideModal(e){
      if($('#modal').get(0) == (e.srcElement)){
        $('#modal').removeClass('modalActive');
        this.$refs.profile.closeModal();
      } 
    },
    /**
    * Función para cerrar el modal.
    *
    * @function
    */
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
    resizeNav();
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

a{
  text-decoration:none;
  *{
    text-decoration:none;
  }
}

.simpleBtn{
  background:transparent;
  color:var(--text);
  border:none;
  font-size:1rem;
  letter-spacing:2px;

  transition:color .25s ease;
  cursor:pointer;
  position:relative;

  &:before{
    content:"";
    height:2px;
    width:0%;
    background-color: var(--gold);
    position:absolute;
    bottom:0;
    left:0;
    transition:all .25s ease;
  }

  &:hover{
    border:none;
    &:before{
      width:100%;
    }
  }
  
}

</style>
