<template>
  <div id="wrap">
    <logoName :size="2"/>
    <div id="contWrap">
      <div id="loginTxt">
        <h2 class="noselect">
          <p class="splitWords" data-splitting="words">Order is key</p>
          <p class="splitWords" data-splitting="words">to control chaos</p>
        </h2>
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
#wrap{
  min-height:100vh;
  display:grid;
  grid-template-rows:10% 90%;
  grid-template-columns:1fr;

  background-color:var(--primary);
  padding:2rem;

  #loginTxt{ font-size:2vw; }

  #contWrap{
    display:flex;
    flex-direction:row;
    flex-wrap: wrap;

    padding:10vh 0;
    
    & > div:first-child{ 
      flex: 0 0 60%;
    }
    & > div:last-child{ 
      flex: 0 0 40%;
    }
  }
}


@keyframes effect {
  from{ 
    transform: translateY(1em);
  }
}
.splitWords.splitting .word{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.05s * var(--word-index));
}
.splitWords{overflow:hidden;}

.splitChars.splitting .char{
  animation: effect .5s cubic-bezier(.5, 0, .5, 1) both;
  animation-delay: calc(.05s * var(--char-index));
}
.splitChars{overflow:hidden;}

@import '@/assets/style/media.scss';
</style>
