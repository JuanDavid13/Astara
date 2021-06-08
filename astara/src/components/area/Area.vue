<template>
  <div class="areaSection">
    <span id="areaName" style="text-transform: uppercase;">{{AreaName}}</span>
    <button id="editNameBt" v-if="deleteable" @click="editAreaName" >Editar</button>
    <button v-if="deleteable" @click="deleteArea">Eliminar</button>

    <br />
    <!--    switch goes here    -->
    <button @click="addTask">+ Tarea</button>
    <CreateTask @taskCreated="getTasks"/>
    <input v-model="query" type="text">
    <div>
      <transition-group
        name="search-fade"
        @before-enter="beforeEnter"
        @enter="enter"
        @leave="leave"
        mode="out-in"
      >
      <div id="tasks" v-for="(task, index) in computedTasks" :key="task.id">
        <Task :task="task" :data-index="index" @taskDeleted="getTasks" @getTasks="getTasks"/>
      </div>

      </transition-group>
    </div>
    <span v-if="!allLoaded" id="load">Cargar más</span>
  </div>
</template>

<script>
import Task from '@/components/item/Task.vue';
import CreateTask from '@/components/item/CreateTask.vue';

import Axios, { AreaCorrespond }from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'Area',
  components: {
    Task,
    CreateTask,
  },
  emits: ['updateAreaName','deleteArea'],
  data() {
    return {
      deleteable: false,
      AreaName: "",
      Tasks: [],
      Goals: [],
      task:{
        name:"",
        deadline:"",
        dated:"",
        status:false,
      },
      goal: {
        name:"",
        deadline:"",
      },
      query:"",
      total:10,
      allLoaded: false,
    }
  },
  computed: {
    computedTasks(){
      return this.Tasks.filter(task  => {
        return task.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    }
  },
  methods: {
    beforeEnter(tasks){
      let index = $(tasks).children(1)[0].dataset.index;
      $(tasks).css({
        "opacity":0,
        //"height":0,
        "transform":"translateX(-5vw)",
        "transition":"all .5s ease-in-out",
        "transition-delay":(index*.1)+"s",
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
    getSlugfromName(name){
      return name.repplace(" ","-").trim();
    },
    async editAreaName(e){
      if($('#areaName')[0].nodeName.localeCompare("SPAN") == 0){ 
        $('#areaName').replaceWith("<input id='areaName' type='text' value='" + this.AreaName + "' minlength='3' maxlength='20'>");
      $(e.target).text('Aceptar');
      }else{
        if($('#areaName').val().localeCompare(this.AreaName) != 0){
          if($('#areaName').val().length < 3){
            $('#areaName').replaceWith("<span id='areaName' style='text-transform: uppercase;'>"+this.AreaName +"</span>");
            return;
          }
          await Axios.post('/area/edit',{ oldName: this.AreaName ,name: $('#areaName').val() }).then((res)=>{
            if(!res.data.changed){
              this.AreaName = "Error";
            }else{
              this.$emit('updateAreaName', this.AreaName, $('#areaName').val());
              this.AreaName = $('#areaName').val();
              let slug = $('#areaName').val();
              window.history.pushState($('#areaName').val(), $('#areaName').val(), '/area/'+slug);
            }
          });
        }
        $('#areaName').replaceWith("<span id='areaName' style='text-transform: uppercase;'>"+this.AreaName +"</span>");
        $(e.target).text('Editar');
      }
    },
    deleteArea(){
      Axios.post("/area/delete",{slug:this.$route.params.name}).then((res)=>{
        if(res.data.deleted){
          this.$emit('deleteArea',this.$route.params.name);
          this.$router.push({name:'Main'});
        }
      })
    },
    getTasks(){
      Axios.post('/area/tasks',{slug: this.$route.params.name }).then((res)=>{
        this.Tasks = JSON.parse(res.data.tasks);
      });
    },
  },
  async created() {
    const slug = this.$router.currentRoute._value.params.name;
    let data = await AreaCorrespond(slug)
    this.deleteable = data.deleteable;
    this.AreaName = data.areaName;
    
    //this.getTasks();
  },
  mounted(){
    const options = {
      root: null,
      //threshold: 0.3,
      threshold: 1,
      //rootMargin: "-50px"
      rootMargin: "0px"
    };

    $('#load').ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          //if(!entry.isIntersecting || entry.intersectionRatio == 1)
            //return;
          if(!this.allLoaded){
            console.log(entry);
            //setInterval(()=>{
              if(entry.isIntersecting){
                console.log(entry);
                console.log('asked form more');

                Axios.post('/area/paginated-tasks',{
                  offset: this.Tasks.length,
                  slug:this.$route.params.name
                }).then((res)=>{
                  if(res.data.error){
                    console.log('error');
                  }else{
                    this.Tasks = this.Tasks.concat(JSON.parse(res.data.tasks));

                    if(this.Tasks.length >= this.total)
                      this.allLoaded = true;
                  }
                });
              }
            //},3000);
          }
        });
      },options);

      observer.observe($('#load')[0]);
    });
  }
}
</script>
<style lang="scss">
.areaSection{
  padding:1.5rem;
}
#newTask{
  border:2px solid red;
  border-radius:5px;
  overflow:hidden;
}
</style>
