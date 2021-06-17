<template>
  <div class="areaSection">
    <Error ref="error" />
    <div id="areaHeader">
      <div>
        <span v-if="!onEdit" id="areaName">{{AreaNameCopy}}</span>
        <input v-if="onEdit" type="text" v-model="AreaNameCopy" placeholder="Nombre">
      </div>

      <div>
        <span>Buscar: </span>
        <input v-model="query" type="text">
      </div>

      <div id="actionBtns">
        <button id="editNameBt" v-if="deleteable" @click="editAreaName">Editar</button>
        <button v-if="onEdit"  @click="cancelEdit">Cancel</button>
        <button v-if="deleteable" @click="deleteArea">X</button>
      </div>
    </div>

    <div id="switches">
      <router-link class="switch" :class="{active: onTasks}" :to="{ name: 'Tasks'}" >Tareas</router-link>
      <router-link class="switch" :class="{active: onGoals}" :to="{ name: 'Goals'}" >Objetivos</router-link>
    </div>
    
    <router-view :query="query" @onTasks="gotoTasks" @onGoals="gotoGoals"></router-view>
  </div>
</template>

<script>
import Axios, { AreaCorrespond }from '@/auth/auth';
import $ from 'jquery';

import Error from '@/components/error/Error.vue';
import { GetErrMsg } from '@/js/error.js';

export default {
  name: 'Area',
  components:{ Error },
  emits: ['updateAreaName','deleteArea'],
  data() {
    return {
      onEdit:false,
      onTasks:true,
      onGoals:false,
      deleteable: false,
      Slug: "",
      AreaName: "",
      AreaNameCopy: "",

      query:"",
    }
  },
  methods: {
    gotoTasks(){ this.onGoals = false; this.onTasks = true; },
    gotoGoals(){ this.onGoals = true; this.onTasks = false; },
    getSlugfromName(name){ 
      return name.replace(" ","-"); 
    },
    deleteArea(){
      Axios.post("/area/delete",{slug:this.$route.params.name}).then((res)=>{
        if(res.data.deleted){
          this.$emit('deleteArea',this.$route.params.name);
          this.$router.push({name:'Main'});
        }
      });
    },
    editAreaName(e){
      if(this.onEdit){
        this.onEdit = false;
        $(e.target).text('Editar');

        if(this.AreaName.length < 3){
          this.$refs.error.setErr(GetErrMsg());
          return;
        }

        Axios.post('/area/edit',{ 
          area: this.$route.params.name,
          name: this.AreaNameCopy
        }).then((res)=>{
          console.log(res);
          if(!res.data.changed)
            this.$refs.error.setErr(GetErrMsg());
          else{
            this.$emit('updateAreaName', this.AreaName, this.AreaNameCopy);
            //this.AreaName = $('#areaName').val();
            let slug = this.getSlugfromName(this.AreaNameCopy);
            window.history.pushState(this.AreaNameCopy, this.AreaNameCopy, '/area/'+slug);
          }
        });
      }else{
        this.onEdit = true;
        $(e.target).text('Aceptar');
        
      }
    },
    cancelEdit(){
      this.onEdit = false;
      this.AreaNameCopy = this.AreaName;
    }
  },
  async created() {
    let data = await AreaCorrespond(this.$router.currentRoute._value.params.name);
    this.deleteable = data.deleteable;
    this.Slug = this.$route.params.name;
    this.AreaName = data.areaName;
    this.AreaNameCopy = this.AreaName;
  },
}
</script>
<style lang="scss">

.areaSection{
  padding:1.5rem;
  #areaHeader{
    display:flex;
    flex-direction:row;
    justify-content:space-between;
    align-items:flex-start;
    flex-wrap:wrap;
  }
  #areaName{
    font-size:2rem;
    text-transform:uppercase;
    font-family: 'Lora', serif;
  }

  #lupa{
    width:25px;
    height:25px;
    stroke:var(--gold);
    path{
      stroke:var(--gold);
    }    
  }

  #switches{
    margin-top:25px;
  }

  #actionBtns{
    button{
      color:var(--tertiary);
      &:hover{
        color:var(--gold);
      }
    }
  }
}
#newTask{
  border:2px solid red;
  border-radius:5px;
  overflow:hidden;
}

.switch{
  text-decoration:none;
  margin-right:15px;
  color:var(--tertiary);
}

.active{
  color:var(--text);
}
</style>
