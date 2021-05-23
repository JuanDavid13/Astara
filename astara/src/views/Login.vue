<template>
  <div class="wrap">
    <div id="loginTxt">
      <logoName :size="1.5" />
      <h2>Order is key <br />to control chaos</h2>
    </div>
    <loginForm />
  </div>
</template>

<script>
import LoginForm from '@/components/loginForm.vue'
import logoName from '@/components/commons/logoName.vue'

export default {
  name: 'Login',
  components: {
    LoginForm,
    logoName,
  },
  data() {
    return {
      publicPath: 'http://localhost/',
    }
  }
} 
</script>

<style scoped lang="scss">
.wrap{
  min-height:100vh;
  display:flex;
  flex-direction:row;

  background-color:var(--primary);
  padding:2rem;

}
#loginTxt{
  display:flex;
  flex-direction:column;
  justify-content:center;
  gap:2rem;
  max-height:80vh;
  font-size:3vw;
}

@media only screen and(max-width:670px){
  .wrap{
    flex-direction:column;
  }
}
</style>
