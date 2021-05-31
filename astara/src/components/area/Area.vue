<template>
  <div>
    <span>{{$route.params.name}}</span>
    <div id="items" v-for="item in Items" :key="item.name">
      <Item :name="item.name" :deadline="item.deadline" :done="item.done"/>
    </div>
  </div>
</template>

<script>
import Item from '@/components/main/Item.vue';
import Axios, { AreaCorrespond }from '@/auth/auth';

export default {
  name: 'Area',
  components: {
    Item
  },
  data() {
    return {
      Items: [],
    }
  },
  methods: {
    getItemsFromAreas(slug) {
      console.log(slug);
      Axios.get('/area').then((res)=>{
        console.log(res);
      })
    }
  },
  async created() {
    const slug = this.$router.currentRoute._value.params.name;
    let data = await AreaCorrespond(slug)
    console.log("data");
    console.log(data);
    this.Items = JSON.parse(data);
  }
}
</script>
<style>

</style>
