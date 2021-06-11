<template>
  <div class="areaSection">
    <span id="areaName" style="text-transform: uppercase;">{{AreaName}}</span>
    <button id="editNameBt" v-if="deleteable" @click="editAreaName" >Editar</button>
    <button v-if="deleteable" @click="deleteArea">Eliminar</button>

    <input v-model="query" type="text">

    <br />
    <router-link :to="{ name: 'Tasks', params: { name: 'TFG' }}" >tareas</router-link>
    <router-link :to="{ name: 'Goals', params: { name: 'TFG' }}" >objetivos</router-link>

    <router-view :query="query" ></router-view>
  </div>
</template>

<script>
import Axios, { AreaCorrespond }from '@/auth/auth';

export default {
  name: 'Area',
  emits: ['updateAreaName','deleteArea'],
  data() {
    return {
      deleteable: false,
      AreaName: "",

      query:"",
    }
  },
  methods: {
    getSlugfromName(name){ return name.repplace(" ","-").trim(); },
    getSlug(){ return this.$route.params.name; },
    deleteArea(){
      Axios.post("/area/delete",{slug:this.$route.params.name}).then((res)=>{
        if(res.data.deleted){
          this.$emit('deleteArea',this.$route.params.name);
          this.$router.push({name:'Main'});
        }
      });
    },
    async editAreaName(e){
      if($('#areaName')[0].nodeName.localeCompare("SPAN") == 0){ 
        $('#areaName').replaceWith("<input id='areaName' type='text' value='" + this.AreaName + "' minlength='3' maxlength='20'>");
        $(e.target).text('Aceptar');
      }else{
        if($('#areaName').val().localeCompare(this.AreaName) != 0){
          if($('#areaName').val().length < 3){
            $('#areaName').replaceWith("<span id='areaName' style='text-transform: uppercase;'>"+this.AreaName +"</span>");
            return;
          }
          await Axios.post('/area/edit',{ oldName: this.AreaName ,name: $('#areaName').val() }).then((res)=>{
            if(!res.data.changed)
              this.AreaName = "Error";
            else{
              this.$emit('updateAreaName', this.AreaName, $('#areaName').val());
              this.AreaName = $('#areaName').val();
              let slug = $('#areaName').val();
              window.history.pushState($('#areaName').val(), $('#areaName').val(), '/area/'+slug);
            }
          });
        }
        $('#areaName').replaceWith("<span id='areaName' style='text-transform: uppercase;'>"+this.AreaName +"</span>");
        $(e.target).text('Editar');
      }
    },
  },
  async created() {
    let data = await AreaCorrespond(this.$router.currentRoute._value.params.name);
    this.deleteable = data.deleteable;
    this.AreaName = data.areaName;
  },
}
</script>
<style lang="scss">
.areaSection{
  padding:1.5rem;
}
#newTask{
  border:2px solid red;
  border-radius:5px;
  overflow:hidden;
}
</style>
