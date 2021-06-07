<template>
  <div class="task">
    <!--<input type="checkbox" v-model=statusCopy disabled>-->
    <span>{{task.status}}</span>
    <span>{{task.id}}</span>
    <span>{{task.name}}</span> 
    <span>{{task.deadline}}</span> 
    <span>{{task.dated}}</span>
    <div>
      <button @click="createNested">Crear</button>
      <button @click="edit">Editar</button>
      <button @click="remove">Eliminar</button>
    </div>
  </div>
</template>

<script>
import Axios from '@/auth/auth';

export default {
  name: 'Task',
  emits: ['taskDeleted'],
  props: {
    task: {
      id: 0,
      name: "",
      deadline: "",
      dated: "",
      status: false,
    },
  },
  data(){
    return {
      statusCopy:false,
    }
  },
  methods: {
    setNew(){
      this.statusCopy = this.task.status;
    },
    createNested(){
      console.log(this.task.id);
    },
    edit(){
      console.log(this.task.id);
    },
    remove(){
      Axios.post('/user/task/delete',{id:this.task.id}).then((res)=>{
        if(!res.data.deleted)
          console.log("error");
        else
          this.$emit('taskDeleted');
        //nota:
        //podría pasarle el id del eliminado y así poder hacer una transición
      });
    },
  }
}

</script>

<style lang="scss">
.task{
  height:4rem;
  padding:15px;

  border:1px solid var(--tertiary);
  border-radius:5px;

  display: flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
