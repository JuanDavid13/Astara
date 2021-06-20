<template>
  <div id="sidebar" class="sidebar" @click="drop" >
    <div id="sidebarWrap">
      <div class="logo noselect splitChar" data-splitting>Astara</div>
      <div id="areas">
        <p class="noselect splitChar" data-splitting>AREAS</p>
        <router-link :to="{name: 'Main'}">MAIN</router-link>
        <transition-group name="list" mode="out-in">
          <div class="area" v-for="area in areas" :key="area.name" >
            <router-link :to="{ name: 'Tasks', params: { name: area.slug } }" >{{area.name}}</router-link>
          </div>
        </transition-group>
      </div> 
      <div>
        <form @submit.capture="addArea">
          <input id="addArea" v-model="newArea" type="text" minlegth="3" maxlength="20" spellcheck="false" autocomplete="off" placeholder="+ Area">
        </form>
      </div>
      <div>
        <a href="#" @click="$emit('openprofile')"><span>{{username}}</span></a>
      </div>
    </div>
  </div>
</template>

<script>
import $ from 'jquery';

import Axios from '@/auth/auth';

  export default{
    name: 'Sidebar',
    props: {username: String},
    emits: ['openprofile', 'deleteArea', 'updateAreaName'],
    data() {
      return {
        newArea:"",
        areas: [],
        dropped:false,
      }
    },
    methods: {
      /**
      *
      * Función auxiliar para hacer el sidebar responsive.
      *
      * @function
      */
      drop(){
        if($('#sidebar').hasClass('hiddable')){
          if(!this.dropped){
            $('#sidebar').css('height','90vh');
            $('#sidebarWrap').css('height','90vh');
            $('#sidebarWrap > div:not(:first-child)').css('height','auto');
            this.dropped = true;
          }else{
            $('#sidebar').css('height','75px');
            $('#sidebarWrap').css('height','75px');
            $('#sidebarWrap > div:not(:first-child)').css('height','0');
            this.dropped = false;
          }
        }
      },
      /**
      *
      * Función para añadir una nueva área.
      *
      * @function
      * @param { e } e - DOM event.
      */
      addArea(e){
        e.preventDefault();
        Axios.post("/area/create",{name:this.newArea}).then((res)=>{
          if(res.data.added){
            this.areas = JSON.parse(res.data.areas);
            $(document).ready(()=>{ 
              $('#addArea').val(""); 
              $('#addArea').blur();
            });
          }
        })
      },
      /**
      *
      * Función que elimina un área del sidebar.
      *
      * @function
      * @param { string } slug - Slug del area a eliminar.
      */
      deleteArea(slug){
        let pos = this.areas.findIndex(area => area.slug === slug);
          this.areas.splice(pos,1);
      },
      /**
      *
      * Función que actualiza el nombre del area en el sidebar.
      *
      * @function
      * @param { string } oldName - Nombre antiguo del área.
      * @param { string } name - Nombre nuevo del área.
      */
      updateAreaName(oldName,name){
        let pos = this.areas.findIndex(area => area.name === oldName);
        this.areas[pos].name = name;
      },
      /**
      *
      * Función que carga todas las areas en el sidebar.
      *
      * @function
      */
      getAreas(){
        Axios.get("/area").then((res)=>{
          this.areas = JSON.parse(res.data); 
        });
      }
    },
    created(){
      this.getAreas();
    },
  }
</script>

<style scoped lang="scss">

@import '@/assets/style/common';

.sidebar {
  color:var(--text);
  //background-color:transparent;
  border-right:1px solid var(--tertiary);
  text-align:left;

  max-height:100vh;
  position:sticky;
  top:0;

  padding:3vh;

  & * { max-width:100%; }
  & :is(p,a,span) {
    //equivalente a
    //p, a, span {

    overflow:hidden;
    text-overflow:ellipsis;
    white-space:nowrap;
  }

  & #sidebarWrap{
    //position:sticky;
    //top:1vh;

    height:100%;

    display:grid;
    grid-template-columns:1fr;
    grid-template-rows: minmax(0,5%) 85% minmax(0,10%);

    & > div:first-child {
      color:var(--gold);
      font-size:1.5rem;
      letter-spacing:2px;
    }

    *{ width:100%; }


    #areas{ 
      margin-top:25px;
      text-align:left; 
      font-size:.9rem;
      overflow-x:hidden;
      overflow-y:scroll;

      .list-enter-active,
      .list-leave-active {
        transition: all .5s ease;
      }
      .list-enter-from,
      .list-leave-to {
        opacity: 0;
        transform: translateX(-3rem);
        position:absolute;
      }
    }

  }

  #addArea{
    border:none;
    overflow:hidden;
    padding: calc(0.5rem + 1px);
    border-radius:5px;

    &:hover { cursor:pointer; }
    &:focus, &:focus-visible{
      padding:0.5rem;
      border:1px solid var(--gold);
      outline:none;
    }

  }

  a{
    color:var(--text);
    text-decoration:none;
  }

  .area{
    margin-bottom: 5px;
    cursor: pointer;
    text-transform:uppercase;
    letter-spacing:2px;
  }

}
#sidebar.hiddable{
  height:50px;
}
.hiddable {
  display:flex !important;
  flex-direction:column;

  background-color:var(--primary);
  #sidebarWrap {
      overflow:hidden;
    #areas{
      //overflow:hidden;
    }
    & > div:first-child{
      height:50px;
      cursor:pointer;
    }
    & > div:not(:first-child){
      height:0 !important;
      //overflow:hidden;
    }
  }
}
</style>
