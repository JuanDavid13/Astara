<template>
  <button @click="addTask">+ Tarea</button>
  <CreateTask @taskCreated="getTasks"  />
  <transition-group
    name="search-fade"
    @before-enter="beforeEnter"
    @enter="enter"
    @leave="leave"
    mode="out-in"
  >
  <div id="tasks" v-for="(task, index) in computedTasks" :key="task.id">
    <Task :task="task" :data-index="index" @remove="remove" @getTasks="getTasks"/>
  </div>
  </transition-group>
  <span v-if="!allTasksLoaded" id="loadTasks">Cargar más</span>
</template>

<script>
import Task from '@/components/item/Task.vue';
import CreateTask from '@/components/item/CreateTask.vue';

import Axios from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'AreaTask',
  components: {
    Task,
    CreateTask,
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
    computedTasks(){
      return this.Tasks.filter(task  => {
        return task.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    },
  },
  methods: {
    getTasks(){
      Axios.post('/area/paginated-tasks',{
        offset: this.Tasks.length,
        slug:this.$route.params.name
      }).then((res)=>{
        if(res.data.error){
          console.log('error');
          return;
        }
        this.Tasks = this.Tasks.concat(JSON.parse(res.data.tasks));
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
            this.getTasks();
        }
      });
    },
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
    const options = {
      root: null,
      threshold: 1,
      rootMargin: "0px"
    };

    $('#load').ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          if(!this.allLoaded){
            if(entry.isIntersecting)
              this.getTasks();
          }
        });
      },options);
      observer.observe($('#loadTasks')[0]);
    });
  }
}
</script>

<style lang="scss">

</style>
