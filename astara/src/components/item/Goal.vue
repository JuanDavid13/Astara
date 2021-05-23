<template>
  <div @click="clicked(id)" class="dashGoal">
    <span>{{name}}</span>
    <span>{{deadline}}</span>
    <span>{{progress}}%</span>
  </div>
</template>

<script>
import Axios from '@/auth/auth';

export default {
  name: 'DashGoal',
  props: ['id','name', 'deadline', 'progress'],
  data() {
    return {

    }
  },
  methods: {
    clicked(id){
      console.log(id);
        Axios.post('/user/goal/delete').then((res)=>{ console.log(res); }); 
    }
  },
}
</script>

<style scoped lang="scss">
.dashGoal {
  display:flex;
  flex-direction:row;
  justify-content:space-around;
  width:100%;
}
</style>
