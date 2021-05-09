<template>
  <div class="home">
    <Sidebar :areas=areas />
    <form @submit.prevent class="form">
      <input type="text" ref="input" v-model="name" placeholder="Nombre de usuario">
      <input v-if="found" type="text" v-model="pass" placeholder="contraseña">
      <button type="submit" @click="clicked">Aceptar</button>
    </form>
    <div v-for="item in items" :key="item.id">
      <item :name="item.Name" :priority="item.Priority" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import item from '@/components/item.vue'
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
    item,
    Sidebar
  },
  data() {
    return{
      name: "",
      found: false,
      items: [],
      areas: []
    }
  },
  methods: {
    clicked() {
      //axios.get('http://localhost:3000/api/v1/',{withCredentials: true}).then((response)=>{
      //  console.log(response);
      //  console.log(response.headers['set-cookie']);
      //  this.$data.items = response.data;
      //  //console.log(this.$data);
      //  //document.cookie="name=nose";
      //  console.log(document.cookie);
      //});
      ////console.log(this.$refs.input);
      ////this.$refs.input.value = "has clickado";
      Axios.post('/login',{name:this.name}).then((res)=>{
        this.found = res.data.found.value;
      })
    },
    getAreas(){
      Axios.get('/user/areas').then((res)=>{
        console.log(res);
        this.areas = res.data;
        console.log(this.areas);
      });
    }

  },
  created(){
   this.getAreas();
  },
}
</script>

<style scoped lang="scss">
$black:#242423;
$lightB:#333533;

.home{
  color:white;
  background-color:$black;
  min-height:100vh;

  display:grid;
  grid-template-columns: 1fr 4fr;
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
