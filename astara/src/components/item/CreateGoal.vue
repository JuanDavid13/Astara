<template>
  <div id="createGoal">
    <Error ref="error"/>
    <form @submit.capture="createGoal">
      <input id="nombre" type="text" autocomplete="off" placeholder="Nombre" spellcheck="false" minlength="3" max-length="50">
      <label>
        <span class="info">Fecha límite </span>
        <input id="deadline" type="date" >
      </label>
      <!--<input type="text" placeholder="Descripción">-->
      <textarea id="descripcion" placeholder="Descripción"></textarea>
      <div>
        <button type="submit">Añadir</button>
        <button type="reset" @click="cancelAddGoal">Cancelar</button>
      </div>
    </form>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import Error from '@/components/error/Error.vue';
import { GetErrMsg } from '@/js/error.js';

import $ from 'jquery';

export default {
  name: 'CreateGoal',
  components:{
    Error,
  },
  emits: ['updateGoals','cancelAddGoal'],
  methods: {
    /**
    *
    * Función que emite el evento 'cancelAddGoal' para que
    * el padre re-establezca el estado de "creating goal".
    *
    * @function
    */
    cancelAddGoal(){ this.$emit('cancelAddGoal'); },
    /**
    *
    * Función que crea un nuevo objetivo.
    *
    * @async
    * @function
    * @param { DOM event } e - DOM event.
    */
    async createGoal(e){
      e.preventDefault();

      let name  = $('#nombre').val();
      let description = $('#descripcion').val();
      let deadline = $('#deadline').val();

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
        else{
          this.$emit('updateGoals',true);
          this.$emit('cancelAddGoal'); //close form
        }
      });
    },
  }

}
</script>

<style scoped lang="scss">
#createGoal {
  margin-top:25px;
  form{
    display:grid;
    grid-template-rows:1fr minmax(0,min-conent) 1fr;
    grid-template-columns:repeat(2,1fr);
    gap:1rem;
    //width:40%;

    textarea{ grid-column:1/3; }
  }
  .info{
    color:var(--tertiary);
  }
}

</style>
