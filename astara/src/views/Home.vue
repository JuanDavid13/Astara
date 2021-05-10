<template>
  <div class="home">
    <Sidebar :areas=areas />
    <Main :Items=goals />
  </div>
</template>

<script>
// @ is an alias to /src
import Sidebar from '@/components/main/Sidebar.vue'
import Main from '@/components/main/Main.vue'

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
    Main
  },
  data() {
    return{
      areas: {},
      goals: {}
    }
  },
  methods: {
    getAreas(){
      Axios.get('/user/areas').then((res)=>{
        console.log(res);
        this.areas = JSON.parse(res.data);
        console.log(this.areas);
      });
    },
    getTargets(){
      Axios.get('/user/targets').then((res)=>{
        console.log(res);
        this.goals = JSON.parse(res.data);
        console.log(this.goals);
      });
    }

  },
  created(){
   this.getAreas();
   this.getTargets();
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

.form{
    display:flex;
    flex-direction:column;
    width:60%;
    margin:0 auto;

     input{
        border:1px solid grey;
        margin-bottom:10px;
        padding:5px;

        &:focus{
          outline:2px solid #42b983;
        }
     }

     button{
       background-color:#42b983;
       padding:5px;
       border:none;

       &:focus{
          outline:2px solid #2c3e50;
       }
     }
}
</style>
