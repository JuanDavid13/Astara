<template>
  <div id="areaGoals">
    <Error ref="error" />
    <button id="addGoal" @click="addGoal">+ Goal</button>
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
  </div>
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
  emits: ['onGoals'],
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
      tasks:[],
      allGoalsLoaded:false,
      creatingGoal:false,
    }
  },
  computed: {
    /**
    * Método computado que permite la búsqueda de nuevos elementos.
    *
    * @function
    */
    computedGoals(){
      return this.Goals.filter(goal => {
        return goal.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    }
  },
  methods: {
    /**
    * Función que establece el estado de "creando objetivo".
    *
    * @function
    */
    addGoal(){
      this.creatingGoal = true;
    },
    /**
    * Función que re-establece el estado de "creando objetivo" a su valor normal.
    *
    * @function
    */
    cancelAddGoal(){ this.creatingGoal = false; },
    /**
    * Función que pide a la API los objetivos.
    *
    * @function
    * @param { bool } paginated - Si la llamada necesita o no paginación
    */
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
    this.$emit('onGoals');

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
#areaGoals{
  margin-top:25px;
  #addGoal{
    border-top:1px solid var(--tertiary);
    border-bottom:1px solid var(--tertiary);
  }
}
</style>
