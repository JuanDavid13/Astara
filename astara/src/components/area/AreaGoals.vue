<template>
  <button @click="addGoal">+ Goal</button>
  <CreateGoal v-if="creatingGoal" @updateGoals="getGoals" @cancelAddGoal="cancelAddGoal"/>
  <div>
    <transition-group
      name="search-fade"
      @before-enter="beforeEnter"
      @enter="enter"
      @leave="leave"
      mode="out-in"
    >
    <div id="goal" v-for="(goal, index) in computedGoals" :key="goal.id">
      <Goal :goal="goal" :data-index="index" @remove="remove"/>
    </div>

    </transition-group>
  </div>
  <span v-if="!allGoalsLoaded" id="loadGoals">Cargar más</span>
</template>

<script>
import Goal from '@/components/item/Goal.vue';
//import Goal from '@/components/item/goal.vue';
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
      let route = '/area/'+ this.$route.params.name +'/paginated-tasks/' + this.Goals.length;
      Axios.get(route).then((res)=>{
        if(res.data.error){
          console.log('error');
          return;
        }
        this.Goals = JSON.parse(res.data.goals);
      });
    },
    remove(id){
      Axios.post('/area/remove-target',{id: id}).then((res)=>{
        console.log(res);
        if(res.data.error){
          console.log('error');
        }else{
          if(!res.data.deleted)
            console.log('error')
          else
            this.getGoals();
        }
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
          if(!this.allLoaded){
            if(entry.isIntersecting){
              this.getGoals();
              console.log(entry);
            }
          }
        });
      },options);
      observer.observe($('#loadGoals')[0]);
    });
  }

}
</script>

<style lang="scss">

</style>
