<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Atara"/>
    <input type="text" ref="input">
    <button @click="clicked">Aceptar</button>
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
      items: []
    }
  },
  methods: {
    clicked() {
      axios.get('http://localhost:3000/api/v1/',{withCredentials: true}).then((response)=>{
        console.log(response);
        console.log(response.headers['set-cookie']);
        this.$data.items = response.data;
        //console.log(this.$data);
        //document.cookie="name=nose";
        console.log(document.cookie);
      });
      //console.log(this.$refs.input);
      //this.$refs.input.value = "has clickado";

    }
  }
}
</script>
