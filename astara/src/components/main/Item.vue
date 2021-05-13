<template>
  <div class="item">
    <!--{{done}}-->
    <input v-if="done" type="checkbox" checked >
    <input v-else type="checkbox">
    <!--<input type="checkbox" >-->
    <span>{{name}}</span>
    <span>{{deadline}}</span>
  </div>
</template>

<script>
export default {
  name: 'Item',
  props: {
    name: String,
    deadline: String,
    done: Boolean,
  }
}

</script>

<style scoped lang="scss">

.item{
  background-color:#333533;
  padding:10px;
  width:90%;
  height:fit-content;
  margin:10px auto;

  display:grid;
  grid-template-columns:5% 75% 20%;

  transition: box-shadow .25s ease;

  cursor:pointer;

  & span:first-of-type{ text-transform:uppercase;}
  & span:last-child{ opacity:0.2; }
  & span{ text-align:left; }

  &:hover{ box-shadow:0 0 10px rgba(0,0,0,.4); }

}
</style>
