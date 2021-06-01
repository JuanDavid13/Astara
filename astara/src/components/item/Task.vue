<template>
  <div class="task">
    <input type="checkbox" v-model=statusCopy disabled>
    <p>{{task.status}}</p> 
    <span>{{task.name}}</span> 
    <span>{{task.deadline}}</span> 
    <span>{{task.dated}}</span>
  </div>
</template>

<script>
export default {
  name: 'Task',
  props: {
    task: {
      name: "",
      deadline: "",
      dated: "",
      status: false,
    },
  },
  data(){
    return {
      statusCopy:false,
    }
  },
  methods: {
    setNew(){
      this.statusCopy = this.task.status;
    }
  }
}

</script>

<style lang="scss">
.task{
  padding:15px;

  border:1px solid var(--tertiary);
  border-radius:5px;

  display: flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
