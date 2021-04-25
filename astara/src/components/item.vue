<template>
  <div class="item">
    <span>{{name}}</span>
    <span>{{priority}}</span>
  </div>
</template>

<script>
export default{
  name: 'item',
  props:{
    name: String,
    priority: Number 
  }
}
</script>

<style scoped lang="scss">
  .item{
    display:flex;
    flex-direction:row;
  }
</style>
