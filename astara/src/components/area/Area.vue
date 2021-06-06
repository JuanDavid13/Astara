<template>
  <div>
    <span id="areaName" style="text-transform: uppercase;">{{AreaName}}</span>
    <button id="editNameBt" v-if="deleteable" @click="editAreaName" >Editar</button>

    <button v-if="deleteable" @click="deleteArea">Eliminar</button>

    <br />
    <button @click="addTask">+ Tarea</button>
    <form @submit.capture="createTask">
      <input type="text" v-model="task.name" placeholder="Nombre">
      <label>Fecha límite
        <input type="date" v-model="task.deadline" placeholder="Fecha límite">
      </label>
      <label>Lo voy a hacer el día
        <input type="date" v-model="task.dated" placeholder="fechado para">
      </label>
      <button type="submit">Añadir</button>
    </form>
    <Task :task="task" id="newTaks" />
    <!--<p>se abre formulario y se crea una nueva "tarea", las tareas se mueven hacia abajo con margin y luego se hace una animación tope guapa</p>-->
    <input v-model="query" type="text">
    <transition-group
      name="search-fade"
      @before-enter="beforeEnter"
      @enter="enter"
      @leave="leave"
    >
      <div id="items" v-for="(item,index) in computedTasks" :key="item.name">
        <Item :data-index="index" :name="item.name" :deadline="item.deadline" :done="item.done"/>
      </div>
    </transition-group>
    <span id="load">Cargar más</span>
  </div>
</template>

<script>
import Item from '@/components/main/Item.vue';

import Task from '@/components/item/Task.vue';
import Axios, { AreaCorrespond }from '@/auth/auth';

import $ from 'jquery';

export default {
  name: 'Area',
  components: {
    Item,
    Task,
  },
  emits: ['updateSidebar'],
  data() {
    return {
      deleteable: false,
      AreaName: "",
      Items: [],
      task:{
        name:"",
        deadline:"",
        dated:"",
        status:false,
      },
      query:"",
    }
  },
  computed: {
    computedTasks(){
      return this.Items.filter(task  => {
        return task.name.toLowerCase().indexOf(this.query.toLowerCase()) !== -1;
      });
    }
  },
  methods: {
    beforeEnter(tasks){
      $(tasks).css({
        "opacity":0,
        "height":0,
        "transform":"translateX(5vw)",
      });
    },
    enter(tasks){
      let index = $(tasks).children(1)[0].dataset.index;
      $(tasks).css({
        "opacity":1,
        "height":"3rem",
        "transform":"translateX(0vw)",
        "transition":"all .25s ease",
        "transition-delay":(index*0.1)+"s",
      });
    },
    leave(tasks){
      $(tasks).css({
        "opacity":0,
        "height":0,
        "transform":"translateX(-5vw)",
      });
    },
    //delete from here
    getItemsFromAreas(slug) {
      console.log(slug);
      Axios.get('/area').then((res)=>{
        console.log(res);
      })
    },
    async editAreaName(e){
      if($('#areaName')[0].nodeName.localeCompare("SPAN") == 0){ 
        $('#areaName').replaceWith("<input id='areaName' type='text' value='" + this.AreaName + "'>");
      $(e.target).text('Aceptar');
      }else{
        if($('#areaName').val().localeCompare(this.AreaName) != 0){
          await Axios.post('/area/edit',{ name: $('#areaName').val() }).then((res)=>{
            if(!res.data.changed){
              this.AreaName = "Error";
            }else{
              this.AreaName = $('#areaName').val();
            }
          });
        }
        $('#areaName').replaceWith("<span id='areaName' style='text-transform: uppercase;'>"+this.AreaName +"</span>");
        $(e.target).text('Editar');
      }
    },
    async createTask(e){
      e.preventDefault();
      //let formLenght = e.target.length -1;

      const name  = e.target[0].value;
      const deadline = e.target[1].value;
      const dated =  e.target[2].value;
      
      Axios.post('/user/task/create',{
        slug:this.$route.params.name,
        name:name,
        deadline:deadline,
        dated:dated,
      }).then(async (res)=>{
        console.log(res);
        if(!res.data.created){
          console.log("error");
        }else{
          let data = await AreaCorrespond(this.$route.params.name);
          this.Items = JSON.parse(data);
        }
      });
    },
    deleteArea(){
      this.$emit('updateSidebar',this.$route.params.name);
      Axios.post("/area/delete",{slug:this.$route.params.name}).then((res)=>{
        if(res.data.deleted){
          this.$emit('updateSidebar',this.$route.params.name);
          this.$router.push({name:'Main'});
        }
      })
    }
  },
  async created() {
    const slug = this.$router.currentRoute._value.params.name;
    let data = await AreaCorrespond(slug)
    console.log(data);
    this.deleteable = data.deleteable;
    this.AreaName = data.areaName;
    //this.Items = JSON.parse(data);
  },
  mounted(){
    const options = {
      root: null,
      threshold: 0.2,
      rootMargin: "-10px"
    };

    $('#load').ready(()=>{
      let observer = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          if(!entry.isIntersecting || entry.intersectionRatio == 1)
            return;
          console.log(entry);
        });
      },options);

      observer.observe($('#load')[0]);
    });
  }
}
</script>
<style>
#newTask{
  border:2px solid red;
  border-radius:5px;
  overflow:hidden;
}
</style>
