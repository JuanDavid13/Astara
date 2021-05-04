<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Orus"/>
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
import HelloWorld from '@/components/HelloWorld.vue'
import item from '@/components/item.vue'

const axios = require('axios');

export default {
  name: 'Home',
  components: {
    HelloWorld,
    item
  },
  data() {
    return{
      name:"",
      pass:"",
      found: false,
      items: []
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
      axios.post('http://localhost:3000/api/v1/login',{
        name:this.name
      }).then((res)=>{
        if(res.data.Found == "true")
          this.found = true;
        else
          this.found = false;
      })
      //this.found = res.data.found.value;
    }
  }
}
</script>

<style scoped lang="scss">
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
