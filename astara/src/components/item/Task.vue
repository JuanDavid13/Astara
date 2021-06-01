<template>
  <div class="task">
    <p>{{task.name}}</p> 
    <p>{{task.deadline}}</p> 
    <p>{{task.dated}}</p> 
    <p>{{task.status}}</p> 
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
