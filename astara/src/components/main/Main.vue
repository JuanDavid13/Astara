<template>
  <div id="main">
  <p>Todos los objetivos</p>
  <div id="items" v-for="item in Items" :key="item.name">
    <Item :name="item.name" :deadline="item.deadline" />
  </div>

  <p>Objetivos del usuario</p>
  <div id="goals" v-for="goal in DashGoals" :key="goal.id">
    <DashGoal :id="goal.id" :name="goal.name" :deadline="goal.deadline" :progress="goal.progress" />
  </div>

  <div id="modal">
    <div id="modalCont">
      <span>Hello</span>
    </div>
  </div>

</div>
</template>

<script>
import Item from '@/components/main/Item.vue';
import DashGoal from '@/components/item/Goal.vue';
import $ from 'jquery';

console.log($('#main'));
const axios = require('axios');
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});

export default {
  name: 'Main',
  components: {
    Item,
    DashGoal
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
    getTargets(){ Axios.get('/user/targets').then((res)=>{ this.Items = JSON.parse(res.data); }); },
    getGoals(){ Axios.get('/user/goal').then((res)=>{ 
      console.log(res);
      this.DashGoals = JSON.parse(res.data); 
      console.log(this.DashGoals);
      }); 
    }
    //getTask(){ Axios.get('/user/task').then((res)=>{ this.Tasks = JSON.parse(res.data); }); }
    //getEvents(){ Axios.get('/user/event').then((res)=>{ this.Events = JSON.parse(res.data); }); }

  },
  created() {
    this.getTargets();
    this.getGoals();
  }
}

</script>

<style scoped lang="scss">
 #items div{
   height:fit-content;
}

#main{
  position:relative;
}

#modal{
  background-color:rgba(0,0,0,.5);
  position:absolute;
  height:100%;
  width:100%;
  top:0;
  left:0;

  display:none;

  #modalCont{

    position:absolute;
    top:20%;
    left:50%;

    transform:translate(-50%,-50%);

    padding:15px;
    background-color:lightblue;
    border-radius:5px;


    width:60%;
    height:30vh;

  }
}

.modalActive {
  display:block !important;
}

</style>
