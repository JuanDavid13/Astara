<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Atara"/>
    <input type="text" ref="input">
    <button @click="clicked">Aceptar</button>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

const axios = require('axios');

export default {
  name: 'Home',
  components: {
    HelloWorld
  },
  methods: {
    clicked() {
      axios.get('http://localhost:3000/api/v1/').then((response)=>{
        console.log(response);
      });
      //console.log(this.$refs.input);
      //this.$refs.input.value = "has clickado";

    }
  }
}
</script>
