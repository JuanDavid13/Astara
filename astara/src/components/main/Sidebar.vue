<template>
  <div class="sidebar">
    <div class="logo">Astara</div>
    <p>{{ name }}</p> 
      
  </div>
</template>

<script>
  export default{
    name: 'Sidebar',
    props:['name'],
    data(){
      return {

      }
    }
  }
</script>

<style scoped lang="scss">

$black:#242423;
$lightB:#333533;

  .sidebar {
    color:white;
    background-color:$lightB;
    display: flex;
    flex-direction: column;
  }
</style>
