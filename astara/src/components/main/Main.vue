<template>
  <div id="main">
    <div>
      <span>Objetivos principales</span>
      <!--<div v-for="(goal, index) in mainGoals" :key="goal.id">
        <MainGoal :goal="goal" />
      </div>-->
    </div>
    <div>
      <span>Tareas</span>
      <!--<div v-for="(task, index) in mainTasks" :key="task.id">
        <Task :task="task" />
      </div>-->
    </div>
  </div>
</template>

<script>
//import Task from '@/components/item/Task.vue';
//import MainGoal from '@/components/item/Goal.vue';

//import Axios from '@/auth/auth';

export default {
  name: 'Main',
  components: {
    //Task,
    //MainGoal,
  },
  data() {
    return{
      mainGoals: [],
      mainTasks: [],
    }
  },
  methods: {
    getMainTargets(){
      //Axios.get('/task').then((res)=>{

      //});
    },
    getMainGoals(){
      //Axios.get('/goal').then((res)=>{

      //});
    },
  },
  created() {
    this.getMainTargets();
    this.getMainGoals();
  }
}

</script>

<style lang="scss">

@import '@/assets/style/common';

#items div{ height:fit-content; }

#main{ position:relative; }

</style>
