<template>
  <div class="sidebar">
    <div class="logo">Astara</div>
    <div id="areas">
      <div class="area" v-for="area in areas" :key="area.name" >
        {{area.name}}
        <span class="deleteable" v-if="area.deleteable">X</span>
      </div>
    </div> 
  </div>
</template>

<script>
  export default{
    name: 'Sidebar',
    //props:{
    //  areas: Array
    //},
    props: ['areas'],
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
    align-items:flex-start;

    padding:2vw;

    & > div:first-child {
      color:goldenrod;
      margin-bottom:3vh;
      font-size:1.5rem;
    }
    #areas{
      display:flex;
      flex-direction:column;
      align-items:flex-start;
    }

    .area{
      margin-bottom: 5px;
      cursor: pointer;

      .deleteable{
        margin-left:10px;
        color:$lightB;
        transition: color .25s ease;
      }
      
      &:hover .deleteable{ color:gray; }

    }

  }
</style>
