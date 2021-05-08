<template>
  <div class="wrap">
    <h2>Únete a Astara y empieza a organizar tu vida</h2>
    <loginForm />
  </div>
</template>

<script>
import loginForm from '@/components/loginForm.vue'

export default {
  name: 'Login',
  components: {
    loginForm
  }
} 
</script>

<style scoped lang="scss">
.wrap{
  min-height:100vh;
  display:grid;
  grid-template-rows:1fr;
  grid-template-columns:repeat(2,1fr);
}

@media only screen and(max-width:670px){
  .wrap{
    grid-template-rows:20% 80%;
    grid-template-columns:1fr;
  }
}
</style>
