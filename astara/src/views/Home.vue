<template>
  <div class="home">
    <Sidebar :areas=areas />
    <router-view :key="$route.fullPath"></router-view>
  </div>
</template>

<script>
// @ is an alias to /src
import Sidebar from '@/components/main/Sidebar.vue'

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
    getAreas(){
      Axios.get('/area').then((res)=>{
        this.areas = JSON.parse(res.data);
      });
    }
  },
  created(){
   this.getAreas();
  },
}
</script>

<style lang="scss">
$black:#242423;
$lightB:#333533;


/*Chrome, Safari, Edge*/
::-webkit-scrollbar { width:7px; }
::-webkit-scrollbar-track { background-color:$black; }
::-webkit-scrollbar-thumb { background-color:grey; }

/*Firefox*/
* {
  scrollbar-width:thin;
  scrollbar-color:grey $black;
}

.home{
  color:white;
  background-color:$black;
  min-height:100vh;

  display:grid;
  grid-template-columns: 1fr 4.5fr;
  grid-template-rows:1fr;
}
</style>
