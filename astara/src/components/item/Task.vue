<template>
  <div class="task">
    <Error ref="error" />
    <br />
    <!--<input type="checkbox" v-model=statusCopy disabled>-->
    <!--<span>{{task.id}}</span>-->
    <span>{{taskCopy.status}}</span>
    <span v-show="!onEdit" class="taskName">{{taskCopy.name}}</span> 
    <input v-show="onEdit" type="text" v-model="taskCopy.name" :placeholder="taskCopy.name">

    <span v-show="!onEdit" class="taskDeadline">{{taskCopy.deadline}}</span> 
    <input v-show="onEdit" type="date" v-model="taskCopy.deadline">

    <span v-show="!onEdit" class="taskDated">{{taskCopy.dated}}</span>
    <input v-show="onEdit" type="date" v-model="taskCopy.dated">

    <div>
      <button v-if="!onEdit" @click="edit">Editar</button>
      <button v-if="onEdit" @click="submitData">Aceptar</button>
      <button v-if="onEdit" @click="cancel">Cancelar</button>
      <button @click="remove">X</button>
    </div>
  </div>
</template>

<script>
import Axios from '@/auth/auth';
import { GetErrMsg } from '@/js/error.js';

import Error from '@/components/error/Error.vue';

import $ from 'jquery';

export default {
  name: 'Task',
  components: {
    Error,
  },
  emits: ['nodeleted','deleted','getTasks'],
  props: {
    task: {
      id: 0,
      name: "",
      deadline: "",
      dated: "",
      status: false,
    },
  },
  data(){
    return {
      onEdit: false,
      statusCopy:false,
      taskCopy: {},
      err: false,
      errMsg: "",
    }
  },
  methods: {
    submitData(e){
      if(!this.validateInputs(e))
        return;

      //this.taskCopy.name = $(e.path[2]).children('.taskName').val();
      //this.taskCopy.deadline = $(e.path[2]).children('.taskDeadline').val();
      //this.taskCopy.dated = $(e.path[2]).children('.taskDated').val();

      Axios.post('user/task/edit', {
        name: this.taskCopy.name, 
        deadline: this.taskCopy.deadline, 
        dated: this.taskCopy.dated, 
        task_id: this.task.id,
      }).then((res)=>{
        if(!res.data.updated){
          this.$refs.error.setErr(GetErrMsg());
          this.cancel();
          return;
        }
        this.$emit('getTasks',false);
        this.onEdit = false;
      });
    },
    validateInputs(){

      if(this.taskCopy.name.length < 4){
        this.$refs.error.setErr(GetErrMsg());
        return false;
      }

      let taskdeadline = new Date(this.task.deadline).getTime();
      let taskcopydeadline = new Date(this.taskCopy.deadline).getTime();
      ///console.log(taskdeadline == taskcopydeadline);
      let taskdated = new Date(this.task.dated).getTime();
      let taskcopydated = new Date(this.taskCopy.dated).getTime();

      //brute comparison
      if(
          this.taskCopy.name.localeCompare(this.task.name) == 0 &&
          taskdeadline == taskcopydeadline && 
          taskdated == taskcopydated
        )
      return false;

      return true;
    },
    cancel(){
      this.onEdit = false;
      this.taskCopy.name = this.task.name;
      this.taskCopy.deadline = this.task.deadline;
      this.taskCopy.dated = this.task.dated;
    },
    edit(e){
      if($(e.path[0]).text().localeCompare('Editar') === 0)
        this.onEdit = true;
      else
        this.$refs.error.setErr(GetErrMsg());

      //$(e.path[2]).children('.taskName')[0]; //--> the name of the task
      //$(e.path[2]).children('.taskDeadline')[0]; //--> the deadline of the task
      //$(e.path[2]).children('.taskDated')[0]; //--> the dated time of the task

      //if($(e.path[2]).children('.taskName')[0].nodeName.localeCompare("SPAN") == 0){
      //  $(e.path[2]).children('.taskName').replaceWith("<input class='taskName' type='text' value='" + this.taskCopy.name + "' minlength='4' maxlength='20'>");
      //  $(e.path[2]).children('.taskDeadline').replaceWith("<input class='taskDeadline' type='date' value='" + this.taskCopy.deadline + "'>");
      //  $(e.path[2]).children('.taskDated').replaceWith("<input class='taskDated' type='date' value='" + this.taskCopy.dated + "'>");
      //}
    },
    remove(){
      Axios.post('/area/remove-target',{id: this.task.id}).then((res)=>{
        if(res.data.error || res.status == 400){
          this.$emit('nodeleted');
        }else{
          if(!res.data.deleted)
            this.$emit('nodeleted');
          else
            this.$emit('deleted');
        }
      });
    },
  },
  created(){
    this.taskCopy = $.extend(true,{},this.task);
  }
}

</script>

<style lang="scss">
.task{
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
