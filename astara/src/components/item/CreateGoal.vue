<template>
  <div>
    <form @submit.capture="createGoal">
      <input type="text" placeholder="Nombre">
      <label> Descripción
        <input type="text" placeholder="Descripción">
      </label>
      <button type="submit">Añadir</button>
    </form>
  </div>
</template>

<script>
import Axios from '@/auth/auth';

export default {
  name: 'CreateGoal',
  emits: ['goalCreated'],
  methods: {
    async createGoal(e){
      e.preventDefault();
      //let formLenght = e.target.length -1;

      let name  = e.target[0].value;
      let description = e.target[1].value;
      
      if(name == ""){
        console.log("error");
        return;
      }

      Axios.post('/user/goal/create',{
        slug:this.$route.params.name,
        name:name,
        description:description,
      }).then(async (res)=>{
        console.log(res);
        if(!res.data.created)
          console.log("error");
        else
          this.$emit('taskCreated');
      });
    },
  }

}
</script>

<style lang="scss">

</style>
