<template>
  <div>
    <Error ref="error" />
    <form @submit.capture="createTask">
      <input type="text" placeholder="Nombre">
      <label>Fecha límite
        <input type="date" placeholder="Fecha límite">
      </label>
      <label>Planeo hacerlo:
        <input type="date" placeholder="fechado para">
      </label>
      <button type="submit">Añadir</button>
    </form>
  </div>
</template>

<script>
import Axios from '@/auth/auth';

import { GetErrMsg } from '@/js/error.js';
import Error from '@/components/error/Error.vue';

export default {
  name: 'CreateTask',
  components: {
    Error,
  },
  emits: ['taskCreated'],
  props:['id'],
  data(){
    return{
      parentId: -1,
    }
  },
  methods: {
    async createTask(e){
      e.preventDefault();
      //let formLenght = e.target.length -1;

      let name  = e.target[0].value;
      let deadline = e.target[1].value;
      let dated =  e.target[2].value;
      
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
      }).then(async (res)=>{
        if(!res.data.created){
          this.$refs.error.setError(GetErrMsg());
          return;
        }
        this.$emit('taskCreated', true);
      });
    },
  },
  created(){
    this.parentId = this.id;
  },

}
</script>

<style scoped lang="scss">

</style>
