<template>
  <div class="goal">
    <Error ref="error"/>
    <div class="goalHeading">
      <!--<input type="checkbox" v-model="userCopy.status">-->
      <h4>{{goalCopy.name}}</h4> 
      <!--<span>{{goalCopy.progress}}</span>-->
    </div>
    <p class="limit">{{goalCopy.deadline}}</p>
    <p class="desc">{{goalCopy.description}}</p>
    <div class="actionBtns" >
      <button @click="edit">Editar</button>
      <button v-if="onEdit" @click="cancel">Cancelar</button>
      <button @click="remove">X</button>
    </div>
    <div class="nest">
      <div> <!--buttons part-->
        <button @click="createNested">Crear tarea</button>
        <!--<button v-if="goalCopy.tasks.length > 0" @click="viewNested">Ver tareas</button>-->
      </div>
      <CreateTask v-if="createTask" :id="goalCopy.id" @taskCreated="taskCreated"/>
      <div v-if="viewTasks"> <!--nested part-->
        <div class="nestedTasks" v-for="(task, index) in goalCopy.tasks" :key="task.id">
          <Task :task="task" :data-index="index"
            @nodeleted="nodeleted"
            @deleted="taskRemoved"
            @getTasks="getGoals"
          />
        </div>
      </div>
    </div>
 </div>
</template>

<script>
import Axios from '@/auth/auth';
import { GetErrMsg } from '@/js/error.js';

import Task from '@/components/item/Task.vue';
import CreateTask from '@/components/item/CreateTask.vue';
import Error from '@/components/error/Error.vue';

import $ from 'jquery';

export default {
  name: 'Goal',
  components:{
    Error,
    Task,
    CreateTask,
  },
  emits: ['remove','getGoals'],
  props: {
    goal: Object,
  },
  data(){
    return {
      showNested:false,
      onEdit: false,
      goalCopy: {},

      createTask:false,
      viewTasks:false,
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
            this.$refs.error.setErr(GetErrMsg());
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

      if($(name)[0].nodeName.localeCompare("INPUT") == 0){
        $(name).replaceWith("<span>" + this.taskCopy.name + "</span>");
        $(deadline).replaceWith("<span>" + this.taskCopy.deadline + "</span>");
        //$(dated).replaceWith("<span>" + this.taskCopy.dated + "</span>");

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
      console.log(e);
      return;
      //this.taskCopy = $.extend(false, {}, this.task);
      //let name = $(e.path[2]).children(2)[0];
      //let deadline = $(e.path[2]).children()[3];
      //let dated = $(e.path[2]).children()[4];

      //if($(name)[0].nodeName.localeCompare("SPAN") == 0){
      //  $(name).replaceWith("<input type='text' value='" + this.taskCopy.name + "' minlength='4' maxlength='20'>");
      //  $(deadline).replaceWith("<input type='date' value='" + this.taskCopy.deadline + "'>");
      //  $(dated).replaceWith("<input type='date' value='" + this.taskCopy.dated + "'>");

      //  $(e.path[0]).text('Aceptar');
      //  this.onEdit = true;
      //}else{
      //  this.submitData(e);
      //}
    },
    taskCreated(){
      this.$emit('getGoals');
      this.createTask = false; 
    },
    taskRemoved(){
      this.$emit('getGoals');
    },
    createNested(e){
      if(!this.createTask){
        this.createTask = true; 
        $(e.path[0]).text('Cancelar');
        return;
      }
      this.createTask = false; 
      $(e.path[0]).text('Crear tarea');
    },
    nodeleted(){
      this.$refs.error.setErr(GetErrMsg('taskNoDeleted'));
    },
    viewNested(e){
      if(!this.viewTasks){
        this.viewTasks = true; 
        $(e.path[0]).text('Ocultar tareas');
        return;
      }
      this.viewTasks = false; 
      $(e.path[0]).text('Ver tareas');

    },
    remove(){
      this.$emit('remove',this.goal.id);
    },
  },
  created(){
    this.goalCopy = $.extend(true, {}, this.goal);
  },
}

</script>

<style lang="scss">
@import '@/assets/style/common.scss';

.goal{
  position:relative;
  margin-top:15px;
  height:fit-content;
  max-height:2000px; //trick
  padding:15px;

  border:1px solid var(--tertiary);
  border-radius:5px;

  display: flex;
  flex-direction:column;
  justify-content:space-between;

  .goalHeading{
    display:flex;
    flex-flow:row nowrap;

  }
  .actionBtns{
    position:absolute;
    top:15px;
    right:15px;
  }

  .nest{
    display:flex;
    flex-direction:column;
    justify-content:flex-start;
    align-items:flex-end;

    & div:first-child button:first-child{
      margin-right:15px;
    }
  }
}
</style>
