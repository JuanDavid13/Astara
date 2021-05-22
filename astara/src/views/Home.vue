<template>
  <div class="home" @switchTheme="switchTheme">
    <Sidebar :areas=areas />
    <router-view :key="$route.fullPath"></router-view>
  </div>
</template>

<script>
// @ is an alias to /src
import Sidebar from '@/components/main/Sidebar.vue'

import "splitting/dist/splitting.css";
import Splitting from "splitting";

const axios = require('axios');
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  //headers:{'Accept':'application/json','Accept-Charset':'utf-8'},
  //headers:{'Authorization':'Bearer {Token}},
  timeout: 1000,
  withCredentials: true,
  //responseType: 'json',
  //responseEncoding: 'utf8',
});

export default {
  name: 'Home',
  components: {
    Sidebar,
 //   Main
  },
  data() {
    return{
      areas: {},
      goals: {}
    }
  },
  methods: {
    getAreas(){ Axios.get('/area').then((res)=>{ this.areas = JSON.parse(res.data); }); },
    switchTheme(){ console.log("llega"); },
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

/*Chrome, Safari, Edge*/
::-webkit-scrollbar { width:7px; }
::-webkit-scrollbar-track { background-color:var(--primary); }
::-webkit-scrollbar-thumb { background-color:var(--secondary); }

/*Firefox*/
* {
  scrollbar-width:thin;
  scrollbar-color:var(--secondary) var(--primary);
}

.home{
  color:var(--text);
  background-color:var(--primary);
  min-height:100vh;

  display:grid;
  grid-template-columns: 1fr 4.5fr;
  grid-template-rows:1fr;
}

@keyframes effect {
  from{ transform: translateY(1em); }
}

.splitChar.splitting .char{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.02s * var(--char-index));
}

.splitChar{overflow:hidden;}

</style>
