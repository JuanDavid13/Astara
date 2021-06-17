<template>
  <div id="main">
    <Error ref="error" />
    <div>
      <h3 class="splitChars" data-splitting>Objetivos principales</h3>
      <div v-for="goal in Goals" :key="goal.id">
        <Goal :goal="goal" />
      </div>
    </div>
    <div>
      <h3 class="splitChars" data-splitting>Tareas</h3>
      <div v-for="task in Tasks" :key="task.id">
        <Task :task="task" />
      </div>
    </div>
  </div>
</template>

<script>
import "splitting/dist/splitting.css";
import Splitting from "splitting";

import { GetErrMsg } from '@/js/error.js';
import Error from '@/components/error/Error.vue';

import Task from '@/components/item/Task.vue';
import Goal from '@/components/item/Goal.vue';

import Axios from '@/auth/auth';

export default {
  name: 'Main',
  components: {
    Task,
    Goal,
    Error,
  },
  data() {
    return{
      Goals: [],
      Tasks: [],
    }
  },
  methods: {
    getMainTargets(){
      //Axios.get('/task').then((res)=>{

      //});
    },
    getGoals(){
      let route = '/area/main/goals';
      Axios.get(route).then((res)=>{
        console.log(res);
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
          return;
        }
        this.Goals = JSON.parse(res.data.goals);

      });
    },
    getTasks(){
      let route = '/area/main/tasks';
      Axios.get(route).then((res)=>{
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
          return;
        }
        this.Tasks = JSON.parse(res.data.tasks);

      });
    },
  },
  created() {
    this.getTasks();
    this.getGoals();
  },
  mounted(){
    Splitting();
  }
}

</script>

<style lang="scss">

@import '@/assets/style/common';

#items div{ height:fit-content; }

#main{ 
  position:relative; 
  padding:1rem 2rem;

  h3{
    letter-spacing: 2px;
    text-transform:uppercase;
    font-size:2rem;
    font-weight:400;
    font-family: 'Lora', serif;
  }

}

</style>
