<template>
  <div>
    <Error ref="error"/>
    <form @submit.capture="createGoal">
      <input type="text" placeholder="Nombre">
      <input type="text" placeholder="Descripción">
      <label> Fecha límite 
        <input type="date">
      </label>
      <button type="submit">Añadir</button>
      <button @click="cancelAddGoal">Cancelar</button>
    </form>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import Error from '@/components/error/Error.vue';
import { GetErrMsg } from '@/js/error.js';

export default {
  name: 'CreateGoal',
  components:{
    Error,
  },
  emits: ['updateGoals','cancelAddGoal'],
  data() {
    return{

    }
  },
  methods: {
    cancelAddGoal(){ this.$emit('cancelAddGoal');
    },
    async createGoal(e){
      e.preventDefault();
      //let formLenght = e.target.length -1;

      let name  = e.target[0].value;
      let description = e.target[1].value;
      let deadline = e.target[2].value;

      if(name == ""){
        this.$refs.error.setErr(GetErrMsg('nameRequired'));
        return;
      }
      if(deadline == "" || deadline == "0000-00-00"){
        this.$refs.error.setErr(GetErrMsg('nameRequired'));
        return;
      }

      Axios.post('/goal/create',{
        slug:this.$route.params.name,
        name:name,
        description:description,
        deadline: deadline,
      }).then(async (res)=>{
        if(res.data.error)
          this.$refs.error.setErr(GetErrMsg());
        else
          this.$emit('updateGoals',true);
      });
    },
  }

}
</script>

<style lang="scss">

</style>
