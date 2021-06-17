<template>
  <div id="loginWrap">
    <nav>
      <div>
        <router-link :to="{name: 'Landing'}"><logoName :size="2"/></router-link>
      </div>
      <router-link :to="{ name: 'Landing' }"><button class="simpleBtn">Inicio</button></router-link>
    </nav>
    <div id="contWrap">
      <div id="loginTxt">
        <h2 class="noselect splitWords" data-splitting="words">Order is key to control chaos </h2>
        <p>
          Entra y empieza organiza tus tareas y objetivos para ser más productivo.
        </p>
      </div>
      <loginForm />
    </div>
  </div>
</template>

<script>
import LoginForm from '@/components/loginForm.vue'
import logoName from '@/components/commons/logoName.vue'

import "splitting/dist/splitting.css";
import Splitting from "splitting";

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
  },
  mounted(){
    Splitting();
  }
} 
</script>

<style lang="scss">
  nav{
    position:sticky;
    top:0;
    padding:1rem 3rem;
    z-index:99;
    display:flex;
    flex-direction:row;
    justify-content:space-between;
    align-items:center;

    & > div:first-child{
      display:flex;
      flex-direction:row;
      //justify-content:space-between;
      align-items:center;
      gap:2rem;
      font-size:2rem;
      color:var(--gold);
    }
  }

#loginWrap{

  background-color:var(--primary);
  min-height:100vh;

  #contWrap{
    //margin-top:4rem;
    display:grid;
    grid-template-columns:3fr 2fr;
    grid-template-rows:1fr;

    min-height:80vh;
    padding:0 2rem;
  }

  #loginTxt{
    height:inherit;
    display:flex;
    flex-direction:column;
    h2{
      font-size:6vw;
      margin-bottom:100px;
      max-width:15ch;
    }
    p{
      font-family: 'Lora', serif;
    }

  }

}


@keyframes effect {
  from{ 
    transform: translateY(1em);
  }
}

@keyframes effectReverse {
  from{ 
    transform: translateY(1em);
    opacity:0;
  }
  to{
    //transform: translateY(0em);
    opacity:1;
  }
}
.splitWords.splitting .word{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.05s * var(--word-index));
}
.splitWords{overflow:hidden;}

.splitChars.splitting .char{
  animation: effectReverse .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.05s * var(--char-index));
}
.splitChars{overflow:hidden;}

.splitCharsReverse.splitting .char{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.05s * var(--char-index));
}
.splitCharsReverse{overflow:hidden;}


</style>
