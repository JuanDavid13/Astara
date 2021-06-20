<template>
  <div id="areaTasks">
    <Error ref="error" />
    <button id="addTask" @click="addTask">+ Tarea</button>
    <CreateTask v-if="creatingTask" @taskCreated="getTasks"  @cancelAddTask="cancelAddTask"/>

    <transition-group
      name="search-fade"
      @before-enter="beforeEnter"
      @enter="enter"
      @leave="leave"
      mode="out-in"
    >
    <div id="tasks" v-for="(task, index) in computedTasks" :key="task.id">
      <Task :task="task" :data-index="index"
        @deleted="getTasks"
        @nodeleted="nodeleted"
        @getTasks="getTasks"
      />
    </div>
    </transition-group>
    <span id="loadTasks" :class="{ loaderInv: allTasksLoaded }" >Cargar más</span>
  </div>
</template>

<script>
import { GetErrMsg } from '@/js/error.js';
import Axios from '@/auth/auth';

import Task from '@/components/item/Task.vue';
import CreateTask from '@/components/item/CreateTask.vue';
import Error from '@/components/error/Error.vue';

import $ from 'jquery';

export default {
  name: 'AreaTask',
  emits: ['onTasks'],
  components: {
    Task,
    CreateTask,
    Error,
  },
  props:['query'],
  data(){
    return {
      Tasks: [],
      task:{
        name:"",
        deadline:"",
        dated:"",
        status:false,
      },
      allTasksLoaded:false,
      creatingTask:false,
    }
  },
  computed: {
    /**
    * Método computado necesario para realizar la búsqueda de las tareas.
    * 
    * @function
    */
    computedTasks(){
      return this.Tasks.filter(task  => {
        return task.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    },
  },
  methods: {
    /**
    * Función que pide a la API las tareas.
    * 
    * @function
    * @param { bool } paginated - Si se necesitan nuevas tareas o no.
    */
    getTasks(paginated){
      if(paginated == null)
        paginated = false;

      let url = '/area/' + this.$route.params.name + '/paginated-tasks/' + this.Tasks.length + '/' + paginated;
      Axios.get(url).then((res)=>{
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
          return;
        }
        //this.Tasks = this.Tasks.concat(JSON.parse(res.data.tasks));
        if(this.Tasks.length == JSON.parse(res.data.tasks).length)
          this.allTasksLoaded = true;
        this.Tasks = JSON.parse(res.data.tasks);

      });
    },
    /**
    * Función que establece el estado de "creando tarea" como verdadero.
    * 
    * @function
    */
    addTask(){
      this.creatingTask = true;
    },
    /**
    * Función que re-establece el estado de "creando tarea" a su valor normal (false). 
    * 
    * @function
    */
    cancelAddTask(){
      this.creatingTask = false;
    },
    /**
    * Función que establece un error si la tarea no ha sido añadida.
    * 
    * @function
    */
    nodeleted(){
      this.$refs.error.setErr(GetErrMsg());
    },
    /**
    * Función auxilar para la animación de entrada las tareas.
    * 
    * @function
    * @param { [node]array } tasks - Array de DOM nodes.
    */
    beforeEnter(tasks){
      //console.log($(tasks).children(0)[0].dataset.index);
      let index = $(tasks).children(0)[0].dataset.index;
      $(tasks).css({
        "opacity":0,
        //"height":0,
        "transform":"translateX(-5vw)",
        "transition":"all .5s ease-in-out",
        "transition-delay":index*.5+"s",
      });
    },
    /**
    * Función auxilar para la animación de las tareas una vez entradas.
    * 
    * @function
    * @param { [node]array } tasks - Array de DOM nodes.
    */
    enter(tasks){
      $('.areaSection').ready(()=>{
        $(tasks).css({
          "opacity":1,
          //"height":"4rem",
          "transform":"translateX(0vw)",
          "transition":"all .5s ease-in-out",
        });
      });
    },
    /**
    * Función auxilar para la animación de las tareas antes de dejar la escena.
    * 
    * @function
    * @param { [node]array } tasks - Array de DOM nodes.
    */
    leave(tasks){
      $(tasks).css({
        "opacity":0,
        "height":0,
        "transition":"all .5s ease-in-out",
        "transform":"translateX(-5vw)",
      });
    },
  },
  mounted(){
    this.$emit('onTasks');

    const options = {
      root: null,
      threshold: 1,
      rootMargin: "0px"
    };

    $(document).ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          setInterval(()=>{
            if(!this.allTasksLoaded && entry.isIntersecting){
              this.getTasks(true);
            }
          },200);
        });
      },options);
      observer.observe($('#loadTasks')[0]);
    });
  }
}
</script>

<style scoped lang="scss">
  #loadTasks{
    transition:all .25s ease;
  }

  .loaderInv{
    opacity:0;
  }

  #areaTasks{
    margin-top:25px;
    #addTask{
      border-top:1px solid var(--tertiary);
      border-bottom:1px solid var(--tertiary);
    }
  }
</style>
