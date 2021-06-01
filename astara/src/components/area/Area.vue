<template>
  <div>
    <span>{{$route.params.name}}</span>
    <button @click="addTask">+ Tarea</button>
    <button @click="deleteArea">Eliminar</button>
    <form @submit.capture="createTask">
      <input type="text" v-model="task.name">
      <input type="text" v-model="task.deadline">
      <input type="text" v-model="task.dated">
      <button type="submit">Añadir</button>
    </form>
    <Task :task="task" id="newTaks" />
    <!--<p>se abre formulario y se crea una nueva "tarea", las tareas se mueven hacia abajo con margin y luego se hace una animación tope guapa</p>-->
    <div id="items" v-for="item in Items" :key="item.name">
      <Item :name="item.name" :deadline="item.deadline" :done="item.done"/>
    </div>
  </div>
</template>

<script>
import Item from '@/components/main/Item.vue';

import Task from '@/components/item/Task.vue';
import Axios, { AreaCorrespond }from '@/auth/auth';

export default {
  name: 'Area',
  components: {
    Item,
    Task,
  },
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
    getItemsFromAreas(slug) {
      console.log(slug);
      Axios.get('/area').then((res)=>{
        console.log(res);
      })
    },
    createTask(e){
      e.preventDefault();
      console.log(this.$route.params.name);
    },
    deleteArea(){
      Axios.post("/area/delete",{slug:this.$route.params.name}).then((res)=>{
        if(res.data.deleted){
          this.$router.push({name:'Main'});
        }
      })
    }
  },
  async created() {
    const slug = this.$router.currentRoute._value.params.name;
    let data = await AreaCorrespond(slug)
    console.log("data");
    console.log(data);
    this.Items = JSON.parse(data);
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
