<template>
  <div id="createTask">
    <Error ref="error" />
    <form @submit.capture="createTask">
      <input id="name" type="text" placeholder="Nombre" autocomplete="off" spellcheck="false" minlenght="3" maxlenght="50">
      <label>
        <span class="info">Fecha límite</span>
        <input id="deadline" type="date" placeholder="Fecha límite">
      </label>
      <label>
        <span class="info">Planeo hacerlo:</span>
        <input id="dated" type="date" placeholder="fechado para">
      </label>
      <div>
        <button type="submit">Añadir</button>
        <button @click="cancel" type="reset">Cancelar</button>
      </div>
    </form>
  </div>
</template>

<script>
import Axios from '@/auth/auth';

import { GetErrMsg } from '@/js/error.js';
import Error from '@/components/error/Error.vue';

import $ from 'jquery';

export default {
  name: 'CreateTask',
  components: {
    Error,
  },
  emits: ['taskCreated','cancelAddTask'],
  props:['id'],
  data(){
    return{
      parentId: -1,
    }
  },
  methods: {
    /**
    *
    * Función que crea una nueva tarea.
    *
    * @async
    * @function
    * @param { DOM event } e - DOM event.
    */
    async createTask(e){
      e.preventDefault();
      //let formLenght = e.target.length -1;

      let name  = $('#name').val();
      let deadline = $('#deadline').val();
      let dated =  $('#dated').val();
      
      if(name == "" || deadline == "" || this.id == 0){
        this.$refs.error.setErr(GetErrMsg('lackInput'));
        return;
      }
      if(this.id == null)
        this.parentId = -1;

      Axios.post('/user/task/create',{
        slug:this.$route.params.name,
        name:name,
        deadline:deadline,
        dated:dated,
        id:this.parentId,
      }).then((res)=>{
        if(!res.data.created){
          this.$refs.error.setError(GetErrMsg());
          return;
        }
        this.$emit('taskCreated', true);
        this.cancel();
      });
    },
    /**
    *
    * Función que emite el evento 'cancelAddTask' para que
    * el padre re-establezca el estado de "creating task".
    *
    * @function
    */
    cancel(){
      this.$emit('cancelAddTask');
    }
  },
  created(){
    this.parentId = this.id;
  },

}
</script>

<style scoped lang="scss">
#createTask{
  display:flex;
  flex-direction:column;
  gap:1rem;
  margin-bottom:25px;

  width:40%;
  form{
    display:flex;
    flex-direction:column;
    gap:1rem;

    input {
      width:100%;
    }

    .info{
      color:var(--tertiary);
    }
  }

} 
</style>
