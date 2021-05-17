<template>
  <div>
  <p>Todos los objetivos</p>
  <div id="items" v-for="item in Items" :key="item.name">
    <Item :name="item.name" :deadline="item.deadline" />
  </div>

</div>
</template>

<script>
import Item from '@/components/main/Item.vue';
const axios = require('axios');
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});

export default {
  name: 'Main',
  components: {
    Item
  },
  data() {
    return{
      Items: []
    }
  },
  methods: {
    getTargets(){
      Axios.get('/user/targets').then((res)=>{ this.Items = JSON.parse(res.data); });
    }

  },
  created() {
    this.getTargets()
  }
}

</script>

<style scoped lang="scss">
 #items div{
   height:fit-content;
}

</style>
