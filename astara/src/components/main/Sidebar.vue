<template>
  <div class="sidebar">
    <div class="logo">Astara</div>
    <p>AREAS</p>
    <div id="areas">
      <div class="area" v-for="area in areas" :key="area.id" >
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
    props: {areas: Object},
    data(){
      return {

      }
    },
    create() {
    }
  }
</script>

<style scoped lang="scss">

$black:#242423;
$lightB:#333533;

  .sidebar {
    color:white;
    background-color:$black;
    border-right:1px solid grey;
    display: flex;
    flex-direction: column;
    align-items:flex-start;

    padding:1vh 1.5vw;


    position:sticky;
    top:0;

    & > div:first-child {
      color:goldenrod;
      margin-bottom:3vh;
      font-size:1.5rem;
    }
    #areas{ 
      margin-top:15px;
      text-align:left; 
      font-size:.9rem;
    }

    .area{
      margin-bottom: 5px;
      cursor: pointer;

      text-transform:uppercase;

      .deleteable{
        margin-left:15px;
        color:$black;
        transition: color .25s ease;
      }
      
      &:hover .deleteable{ color:gray; }

    }

  }
</style>
