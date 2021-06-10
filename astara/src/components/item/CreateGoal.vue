<template>
  <div>
    <Error ref="error"/>
    <form @submit.capture="createGoal">
      <input type="text" placeholder="Nombre">
      <label> Descripción
        <input type="text" placeholder="Descripción">
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
  emits: ['goalCreated','cancelAddGoal'],
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

      if(name == ""){
        this.$refs.error.setErr(e,GetErrMsg('nameRequired'));
        return;
      }

      //Axios.post('/user/goal/create',{
      //  slug:this.$route.params.name,
      //  name:name,
      //  description:description,
      //}).then(async (res)=>{
      //  console.log(res);
      //  if(!res.data.created)
      //    console.log("error");
      //  else
      //    this.$emit('taskCreated');
      //});
    },
  }

}
</script>

<style lang="scss">

</style>
