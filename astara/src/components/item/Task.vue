<template>
  <div class="task">
    <Error ref="error" />
    <div class="taskCont">
      <div>
        <lable v-show="!onEdit" class="checkWrap">
          <input class="checkBox" type="checkbox" v-model="checked">
          <span class="custCheck" @click="check"></span>
        </lable>

        <h4 v-show="!onEdit" class="taskName splitWords" data-splitting="words">{{taskCopy.name}}</h4> 
        <input v-show="onEdit" class="inputEdit" type="text" v-model="taskCopy.name" placeholder="NOMBRE">

        <div v-show="!onEdit" class="taskDeadline">
          <span class="info">Fecha límite: </span>
          <span>{{taskCopy.deadline}}</span>
        </div> 
        <input v-show="onEdit" class="inputEdit" type="date" v-model="taskCopy.deadline">

        <div v-show="!onEdit" class="taskDated">
          <span class="info">Planeado para: </span>
          <span>{{taskCopy.dated}}</span>
        </div>
        <input v-show="onEdit" class="inputEdit" type="date" v-model="taskCopy.dated">
      </div>
      <div class="taskBtns">
        <button v-if="!onEdit" @click="edit">Editar</button>
        <button v-if="onEdit" @click="submitData">Aceptar</button>
        <button v-if="onEdit" @click="cancel">Cancelar</button>
        <button @click="remove">X</button>
      </div>
    </div>
  </div>
</template>

<script>
import "splitting/dist/splitting.css";
import Splitting from "splitting";

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
      checked:false,
      onEdit: false,
      taskCopy: {},
      err: false,
      errMsg: "",
    }
  },
  methods: {
    /**
    *
    * Función que envía a la API los datos de la tarea para editarlos.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
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
    /**
    *
    * Función que falida los inputs del formulario de edición de la tarea.
    *
    * @function
    */
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
    /**
    *
    * Función que cancela la edición de la tarea.
    *
    * @function
    */
    cancel(){
      this.onEdit = false;
      this.taskCopy.name = this.task.name;
      this.taskCopy.deadline = this.task.deadline;
      this.taskCopy.dated = this.task.dated;
    },
    /**
    *
    * Función que permite la editar la tarea.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
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
    /**
    *
    * Función que elimina una tarea.
    *
    * @function
    */
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
    /**
    *
    * Función que checkea una tarea como hecha.
    *
    * @function
    */
    check(){
      Axios.post('user/task/check',{ id: this.task.id }).then((res)=>{   
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
        }else{
          this.$emit('getTasks');
        }
      });
    }
  },
  created(){
    this.taskCopy = $.extend(true,{},this.task);
  },
  mounted(){
    Splitting();
  }
}

</script>

<style lang="scss">
.task{
  position:relative;
  margin-top:15px;
  height:fit-content;
  padding:15px;

  border-top:1px solid var(--tertiary);
  //border-radius:5px;

  display: flex;
  flex-direction:column;
  justify-content:flex-start;
  align-items:flex-start;

  .taskCont{
    width: 75%;
    & > div:first-child{
      display:flex;
      flex-direction:row;
      align-items:flex-start;
      //justify-content:space-between;
      justify-content:flex-start;
      gap:1rem;
      flex-wrap:wrap;
    }

    .taskName{
      margin-left:30px;
      font-weight:bold;
      text-transform:uppercase;
      letter-spacing:3px;
    }

    .info{
      color:var(--tertiary);
    }
  }

  .taskBtns{
    position:absolute;
    right:15px;
    top:10px;

    button{
      border:none;
      background-color:transparent;
      cursor:pointer;
      color:var(--tertiary);

      &:hover {
        color:var(--gold);
      }
    }
  }
}
</style>
