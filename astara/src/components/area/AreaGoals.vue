<template>
  <button @click="addGoal">+ Goal</button>
  <CreateGoal v-if="creatingGoal" @updateGoals="getGoals" @cancelAddGoal="cancelAddGoal"/>
  <div>
    {{computedGoals}}
    <transition-group
      name="search-fade"
      @before-enter="beforeEnter"
      @enter="enter"
      @leave="leave"
      mode="out-in"
    >
    <div id="goal" v-for="(goal, index) in computedGoals" :key="goal.id">
      <Goal :goal="goal" :data-index="index" @goalDeleted="getGoals" @getGoals="getGoals"/>
    </div>

    </transition-group>
  </div>
  <span v-if="!allGoalsLoaded" id="loadGoals">Cargar más</span>
</template>

<script>
import Goal from '@/components/item/Goal.vue';
import CreateGoal from '@/components/item/CreateGoal.vue';

import Axios, { AreaCorrespond }from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'AreaGoals',
  components: {
    Goal,
    CreateGoal,
  },
  props:['query'],
  data() {
    return {
      Goals: [],
      goal: {
        name:"",
        description:"",
        deadline:"",
        status:false,
      },
      allGoalsLoaded:false,
      creatingGoal:false,
    }
  },
  computed: {
    computedGoals(){
      return this.Goals.filter(goal => {
        return goal.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    }
  },
  methods: {
    addGoal(){ this.creatingGoal = true; },
    cancelAddGoal(){ this.creatingGoal = false; },
    getGoals(){
      console.log('asked');
      //let slug = this.getSlug();
      let route = '/area/'+ this.getSlug() +'/paginated-tasks/' + this.Goals.length;
      Axios.get(route).then((res)=>{
        console.log(res);
        if(res.data.error){
          console.log('error');
          return;
        }
        this.Goals = res.goals;
      });
    },
  },
  mounted(){
    const options = {
      root: null,
      threshold: 1,
      rootMargin: "0px"
    };

    $('#load').ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          if(!this.allLoaded){
            console.log(entry);
            if(entry.isIntersecting){
              if(entry.target.id.localeCompare('loadTasks') == 0)
                this.getGoals();
            }
          }
        });
      },options);
      console.log('observing');
      observer.observe($('#loadGoals')[0]);
    });
  }

}
</script>

<style lang="scss">

</style>
