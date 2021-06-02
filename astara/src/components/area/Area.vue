<template>
  <div>
    <span>{{$route.params.name}}</span>
    <button @click="addTask">+ Tarea</button>
    <button @click="deleteArea">Eliminar</button>
    <form @submit.capture="createTask">
      <input type="text" v-model="task.name" placeholder="Nombre">
      <label> Fecha límite
        <input type="date" v-model="task.deadline" placeholder="Fecha límite">
      </label>
      <label> Lo voy a hacer el día
        <input type="date" v-model="task.dated" placeholder="fechado para">
      </label>
      <button type="submit">Añadir</button>
    </form>
    <Task :task="task" id="newTaks" />
    <!--<p>se abre formulario y se crea una nueva "tarea", las tareas se mueven hacia abajo con margin y luego se hace una animación tope guapa</p>-->
    <div id="items" v-for="item in Items" :key="item.name">
      <Item :name="item.name" :deadline="item.deadline" :done="item.done"/>
    </div>
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
      Items: [],
      task:{
        name:"",
        deadline:"",
        dated:"",
        status:false,
      }
    }
  },
  methods: {
    //delete from here
    getItemsFromAreas(slug) {
      console.log(slug);
      Axios.get('/area').then((res)=>{
        console.log(res);
      })
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
    this.Items = JSON.parse(data);
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
