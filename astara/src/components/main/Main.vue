<template>
  <div id="main">
  <button>+ Tarea</button>
  <button>+ Objetivo</button>
  <!--<p class="splitChar" data-splitting>Todos los objetivos</p>
  <div id="items" v-for="item in Items" :key="item.name">
    <Item :name="item.name" :deadline="item.deadline" />
  </div>-->

  <!--<p>Objetivos del usuario</p>
  <div id="goals" v-for="goal in DashGoals" :key="goal.id">
    <DashGoal :id="goal.id" :name="goal.name" :deadline="goal.deadline" :progress="goal.progress" />
  </div>-->

  <!--<div @click="hideModal" id="modal">
    <div id="modalCont">
      <span>Hello</span>
    </div>
  </div>-->
</div>
</template>

<script>
//import Item from '@/components/main/Item.vue';
//import DashGoal from '@/components/item/Goal.vue';

import Axios from '@/auth/auth';


export default {
  name: 'Main',
  components: {
    //Item,
    //DashGoal
  },
  data() {
    return{
      Items: [],
      DashGoals: [],
      //Task: [],
      //Events: []
    }
  },
  methods: {
    getTargets(){
      Axios.get('/user/targets').then((res)=>{ this.Items = JSON.parse(res.data); });
    },
    getGoals(){
      Axios.get('/user/goal').then((res)=>{ this.DashGoals = JSON.parse(res.data); });
    },
  },
  created() {
    this.getTargets();
    this.getGoals();
  }
}

</script>

<style lang="scss">

@import '@/assets/style/common';

#items div{ height:fit-content; }

#main{ position:relative; }

</style>
