<template>
  <div id="createGoal">
    <Error ref="error"/>
    <form @submit.capture="createGoal">
      <input type="text" placeholder="Nombre">
      <label>
        <span class="info">Fecha límite </span>
        <input type="date">
      </label>
      <!--<input type="text" placeholder="Descripción">-->
      <textarea placeholder="Descripción"></textarea>
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

export default {
  name: 'CreateGoal',
  components:{
    Error,
  },
  emits: ['updateGoals','cancelAddGoal'],
  methods: {
    cancelAddGoal(){ this.$emit('cancelAddGoal'); },
    openCreate(){ this.opened = true; },
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
