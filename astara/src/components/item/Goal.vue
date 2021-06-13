<template>
  <div class="goal">
    <Error ref="error"/>
    <div class="goalHeading">
      <!--<input type="checkbox" v-model="userCopy.status">-->
      <h4 v-if="!onEdit">{{goalCopy.name}}</h4> 
      <input v-if="onEdit" type="text" v-model="goalCopy.name">
      <!--<span>{{goalCopy.progress}}</span>-->
    </div>
    <p v-if="!onEdit" class="limit">{{goalCopy.deadline}}</p>
    <input v-if="onEdit" type="date" v-model="goalCopy.deadline" >

    <p v-if="!onEdit" class="desc">{{goalCopy.description}}</p>
    <input v-if="onEdit" type="text" v-model="goalCopy.description" >

    <div class="actionBtns" >
      <button v-if="!onEdit" @click="edit">Editar</button>
      <button v-if="onEdit" @click="submitData">Aceptar</button>
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
      if(!this.validateInputs(e))
        return;

      Axios.post('goal/edit', {
        name: this.goalCopy.name, 
        deadline: this.goalCopy.deadline, 
        description: this.goalCopy.description, 
        goal_id: this.goal.id,
      }).then((res)=>{
        if(!res.data.updated){
          this.$refs.error.setErr(GetErrMsg());
          return
        }
        this.$emit('getGoals',true);
        this.onEdit = false;
      });
    },
    validateInputs(){
      if(this.goalCopy.name.length < 4){
        this.$refs.error.setErr(GetErrMsg());
        return false;
      }

      let goalCopyDeadline = new Date(this.goalCopy.deadline).getTime();
      let goalDeadline = new Date(this.goal.deadline).getTime();

      if(goalCopyDeadline == 0 || goalCopyDeadline == null){ //check this
        this.$refs.error.setErr(GetErrMsg());
        return false;
      }


      if(
            this.goalCopy.name.localeCompare(this.goal.name) == 0 &&
            this.goalCopy.description.localeCompare(this.goal.description) == 0 &&
            goalCopyDeadline == goalDeadline

        )
      return false;

      return true;
    },
    cancel(){
      this.onEdit = false;
      this.goalCopy.name = this.goal.name;
      this.goalCopy.description = this.goal.description;
      this.goalCopy.deadline = this.goal.deadline;
    },
    edit(e){
      if($(e.path[0]).text().localeCompare('Editar') === 0)
        this.onEdit = true;
      else
        this.$refs.error.setErr(GetErrMsg());
    },
    taskCreated(){
      this.$emit('getGoals',true);
      this.createTask = false; 
    },
    taskRemoved(){
      this.$emit('getGoals',true);
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
      Axios.post('/area/remove-target',{ id: this.goal.id }).then((res)=>{
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
        }else{
          if(!res.data.deleted)
            this.$refs.error.setErr(GetErrMsg());
          else
            this.$emit('getGoals',true);
        }
      });
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
