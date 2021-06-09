<template>
  <div class="goal">
    <span v-if="err" class="errorMsg" style="margin-bottom:15px;">{{errMsg}}</span>
    <br />
    <!--<input type="checkbox" v-model=statusCopy disabled>-->
    <span>{{goal.status}}</span>
    <span>{{goal.name}}</span> 
    <span>{{goal.deadline}}</span> 
    <div>
      <button @click="createNested">Crear</button>
      <button @click="edit">Editar</button>
      <button v-if="onEdit" @click="cancel">Cancelar</button>
      <button @click="remove">Eliminar</button>
    </div>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import { GetErrMsg } from '@/js/error.js';

//import nestedTask from '@/components/item/NestedTask.vue';

import $ from 'jquery';

export default {
  name: 'Task',
  //components:{
  //  nestedTask,
  //},
  emits: ['goalDeleted','getGoals'],
  props: {
    goal: {
      id: 0,
      name: "",
      deadline: "",
      status: false,
    },
  },
  data(){
    return {
      onEdit: false,
      statusCopy:false,
      goalCopy: {},

      err: false,
      errMsg: "",
    }
  },
  methods: {
    submitData(e){
      this.taskCopy.name = $(e.path[2]).children()[2].value;
      this.taskCopy.deadline = $(e.path[2]).children()[3].value;
      this.taskCopy.dated = $(e.path[2]).children()[4].value;

      if (this.validateInputs()){
        Axios.post('user/task/edit', {
          name: this.taskCopy.name, 
          deadline: this.taskCopy.deadline, 
          dated: this.taskCopy.dated, 
          task_id: this.task.id,
        }).then((res)=>{
          if(!res.data.updated){
            this.err = true;
            this.errMsg = GetErrMsg();
            return
          }
          this.$emit('getTasks');
          this.turnBackInputs(e);
        });
      }
    },
    turnBackInputs(e){
      let name = $(e.path[2]).children()[2];
      let deadline = $(e.path[2]).children()[3];

      console.log(name);
      console.log(deadline);

      if($(name)[0].nodeName.localeCompare("INPUT") == 0){
        $(name).replaceWith("<span>" + this.taskCopy.name + "</span>");
        $(deadline).replaceWith("<span>" + this.taskCopy.deadline + "</span>");
        $(dated).replaceWith("<span>" + this.taskCopy.dated + "</span>");

        let editBtn = $(e.path[1]).children()[1];
        $(editBtn).text('Editar')
        this.onEdit = false;
      }
    },
    validateInputs(){
      if(this.taskCopy.name.localeCompare("") == 0 || this.taskCopy.deadline.localeCompare("") == 0 || this.taskCopy.deadline.localeCompare("0000-00-00") == 0){
        this.err = true;
        this.errMsg = GetErrMsg('lackInputs');
        return false;
      }

      if(this.taskCopy.name.length < 4){
        this.err = true;
        this.errMsg = GetErrMsg('shotName');
        return false;
      }

      if(
            this.taskCopy.name.localeCompare(this.task.name) == 0 &&
            this.taskCopy.deadline.localeCompare(this.task.deadline) == 0 &&
            this.taskCopy.dated.localeCompare(this.task.dated) == 0
        )
      return false;

      return true;
    },
    createNested(){
      console.log(this.task.id);
    },
    cancel(e){
      let name = $(e.path[2]).children()[2];
      let deadline = $(e.path[2]).children()[3];
      let dated = $(e.path[2]).children()[4];

      if($(name)[0].nodeName.localeCompare("SPAN") != 0){
        $(name).replaceWith("<span>" + this.task.name + "</span>");
        $(deadline).replaceWith("<span>" + this.task.deadline + "</span>");
        $(dated).replaceWith("<span>" + this.task.dated + "</span>");

        let editBtn = $(e.path[1]).children()[1];
        $(editBtn).text('Editar')
        this.onEdit = false;
      }
    },
    edit(e){
      this.taskCopy = $.extend(false, {}, this.task);
      let name = $(e.path[2]).children()[2];
      let deadline = $(e.path[2]).children()[3];
      let dated = $(e.path[2]).children()[4];

      if($(name)[0].nodeName.localeCompare("SPAN") == 0){
        $(name).replaceWith("<input type='text' value='" + this.taskCopy.name + "' minlength='4' maxlength='20'>");
        $(deadline).replaceWith("<input type='date' value='" + this.taskCopy.deadline + "'>");
        $(dated).replaceWith("<input type='date' value='" + this.taskCopy.dated + "'>");

        $(e.path[0]).text('Aceptar');
        this.onEdit = true;
      }else{
        this.submitData(e);
      }
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
.goal{
  margin-top:15px;
  height:4rem;
  padding:15px;

  border:1px solid var(--tertiary);
  border-radius:5px;

  display: flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
