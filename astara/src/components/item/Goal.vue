<template>
  <div class="goal">
    <Error ref="error"/>
    <div>
      <div class="goalCont">
        <div class="goalHeading">
          <!--<input type="checkbox" v-model="userCopy.status">-->
          <h4 id="goalName" class="splitWords" v-if="!onEdit" data-splitting="words">{{goalCopy.name}}</h4> 
          <input v-if="onEdit" type="text" v-model="goalCopy.name">
        </div>

        <div v-if="!onEdit" class="limit">
          <span class="info">Fecha límite: </span>
          <span>{{goalCopy.deadline}}</span>
        </div>
        <input v-if="onEdit" type="date" v-model="goalCopy.deadline" >

        <p v-if="!onEdit" class="desc info">{{goalCopy.description}}</p>
        <input v-if="onEdit" type="text" v-model="goalCopy.description" >
      </div>

      <div class="actionBtns" >
        <button v-if="!onEdit" @click="edit">Editar</button>
        <button v-if="onEdit" @click="submitData">Aceptar</button>
        <button v-if="onEdit" @click="cancel">Cancelar</button>
        <button @click="remove">X</button>
      </div>

      <div class="nest">
        <div> <!--buttons part-->
          <button id="addTask" @click="createNested">Crear tarea</button>
          <button v-if="goalCopy.tasks.length > 0" @click="viewNested">Ver tareas</button>
        </div>
        <CreateTask v-if="createTask" :id="goalCopy.id" @taskCreated="taskCreated" @cancelAddTask="cancelAddTask"/>
        <div id="nested" v-if="viewTasks"> <!--nested part-->
          <div class="nestedTasks" v-for="(task, index) in goalCopy.tasks" :key="task.id">
            <Task :task="task" :data-index="index"
              @nodeleted="nodeleted"
              @deleted="taskRemoved"
              @getTasks="updateTasks"
            />
          </div>
        </div>
      </div>
    </div>
 </div>
</template>

<script>
import "splitting/dist/splitting.css";
import Splitting from "splitting";

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
    /**
    *
    * Función que edita la información de un objetivo.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
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
    /**
    *
    * Función auxiliar que valida los inputs del formulario, al editar un objetivo.
    *
    * @function
    * @returns { bool } Devuelve true si los inputs son correctos o false si no.
    */
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
    /**
    *
    * Función que cancela la edición del objetivo.
    *
    * @function
    */
    cancel(){
      this.onEdit = false;
      this.goalCopy.name = this.goal.name;
      this.goalCopy.description = this.goal.description;
      this.goalCopy.deadline = this.goal.deadline;
    },
    /**
    *
    * Función que establece la edición de la información del objetivo.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
    edit(e){
      if($(e.path[0]).text().localeCompare('Editar') === 0)
        this.onEdit = true;
      else
        this.$refs.error.setErr(GetErrMsg());
    },
    /**
    *
    * Función que actualiza la información del objetivo si se ha creado una tarea anidada.
    *
    * @function
    */
    taskCreated(){
      this.$emit('getGoals',false);
      this.updateTasks();
      this.createTask = false; 
    },
    /**
    *
    * Función que actualiza la información del objetivo si se ha eliminado una tarea anidada.
    *
    * @function
    */
    taskRemoved(){
      this.$emit('getGoals',false);
      this.updateTasks();
    },
    /**
    *
    * Función que permite la creación de una tarea anidada.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
    createNested(e){
      if(!this.createTask){
        this.createTask = true; 
        $(e.path[0]).text('Cancelar');
        return;
      }
      this.createTask = false; 
      $(e.path[0]).text('Crear tarea');
    },
    /**
    *
    * Función que genera un error si no se ha podido eliminar la tarea.
    *
    * @function
    */
    nodeleted(){
      this.$refs.error.setErr(GetErrMsg('taskNoDeleted'));
    },
    /**
    *
    * Función que permite que se muestren las tareas anidadas del objetivo.
    *
    * @function
    * @param { DOM event } e - DOM event.
    */
    viewNested(e){
      if(!this.viewTasks){
        this.viewTasks = true; 
        $(e.path[0]).text('Ocultar tareas');
        return;
      }
      this.viewTasks = false; 
      $(e.path[0]).text('Ver tareas');
    },
    /**
    *
    * Función para la eliminación del objetivo.
    *
    * @function
    */
    remove(){
      Axios.post('/area/remove-target',{ id: this.goal.id }).then((res)=>{
        if(res.data.error){
          this.$refs.error.setErr(GetErrMsg());
        }else{
          if(!res.data.deleted)
            this.$refs.error.setErr(GetErrMsg());
          else{
            this.$emit('getGoals',true);
          }
        }
      });
    },
    /**
    *
    * Función que actualiza las tareas anidadas de un objetivo.
    *
    * @function
    */
    updateTasks(){
      Axios.post('/user/task/goal-tasks',{ id: this.goal.id }).then((res)=>{
        if(res.data.error)
          this.$refs.error.setErr(GetErrMsg());
        else
          this.goalCopy.tasks = JSON.parse(res.data.tasks);
      });
    }
  },
  created(){
    this.goalCopy = $.extend(true, {}, this.goal);
    if(this.goalCopy.tasks == null)
      this.goalCopy.tasks = [];
  },
  mounted(){
    Splitting();
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

  transition:all .25s ease;

  border-top:1px solid var(--tertiary);

  display: flex;
  flex-direction:column;
  justify-content:flex-start;

  .goalCont{
    display:flex;
    flex-direction:column;
    gap:1rem;
    max-width:80%;

    .desc {
      font-family: 'Lora', serif;
    }

  }

  .goalHeading{
    display:flex;
    flex-flow:row nowrap;

    #goalName{
      font-weight:bold;
      text-transform:uppercase;
      letter-spacing:3px;
    }

  }

  .info{
    color:var(--tertiary);
  }

  .actionBtns{
    position:absolute;
    top:15px;
    right:15px;

    button{
      color:var(--tertiary);
    }

  }

  .nest{
    display:flex;
    flex-direction:column;
    justify-content:flex-start;
    align-items:flex-end;

    & div:first-child button:first-child{
      margin-right:15px;
    }

    #nested{
      width:95%;
    }
  }
}
</style>
