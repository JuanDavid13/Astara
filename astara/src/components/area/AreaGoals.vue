<template>
  <Error ref="error" />
  <button @click="addGoal">+ Goal</button>
  <CreateGoal v-if="creatingGoal" @updateGoals="getGoals" @cancelAddGoal="cancelAddGoal"/>
  <div>
    <transition-group
      name="search-fade"
      mode="out-in"
    >
    <div id="goal" v-for="(goal, index) in computedGoals" :key="goal.id">
      <Goal :goal="goal" :data-index="index" @getGoals="getGoals"/>
    </div>

    </transition-group>
  </div>
  <span v-if="!allGoalsLoaded" id="loadGoals">Cargar más</span>
</template>

<script>
import { GetErrMsg } from '@/js/error.js';
import Goal from '@/components/item/Goal.vue';
//import Goal from '@/components/item/goal.vue';
import CreateGoal from '@/components/item/CreateGoal.vue';
import Error from '@/components/error/Error.vue';

import Axios from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'AreaGoals',
  components: {
    Goal,
    CreateGoal,
    Error,
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
    getGoals(paginated){
      if(paginated == null)
        paginated = false;
      let route = '/area/'+ this.$route.params.name +'/paginated-goals/' + this.Goals.length + '/' + paginated;
      Axios.get(route).then((res)=>{
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
          return;
        }

        if(this.Goals.length == JSON.parse(res.data.goals).length)
          this.allGoalsLoaded = true;
        this.Goals = JSON.parse(res.data.goals);
      });
    },
  },
  mounted(){
    const options = {
      root: null,
      threshold: 1,
      rootMargin: "0px"
    };

    $('document').ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          setInterval(()=>{
            if(!this.allGoalsLoaded && entry.isIntersecting){
              this.getGoals(true);
            }
          },200)
        });
      },options);
      observer.observe($('#loadGoals')[0]);
    });
  }

}
</script>

<style lang="scss">

</style>
