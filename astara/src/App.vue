<template>
  <router-view/>
</template>

<script>
</script>

<style lang="scss">

@import '@/assets/style/common';

/*Chrome, Safari, Edge*/
::-webkit-scrollbar { width:7px; }
::-webkit-scrollbar-track { background-color:var(--primary); }
::-webkit-scrollbar-thumb { background-color:var(--secondary); }
::-webkit-scrollbar-corner { background-color:var(--primary); }
::-webkit-scrollbar-resizer { background-color:var(--primary); }

/*Firefox*/
* {
  scrollbar-width:thin;
  scrollbar-color:var(--secondary) var(--primary);
}

html{
  font-size:16px;

  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  color: var(--text);
}

*,*::before,*::after{
  margin:0;
  padding:0;
  box-sizing:border-box;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}


.noselect{
  -webkit-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}

input, textarea{
  border:none;
  border-bottom:1px solid var(--tertiary);

  &:focus{
    outline:none;
    border-bottom:1px solid var(--gold);
  }
}

::placeholder{ color:var(--tertiary); }
::-webkit-placeholder{ color:var(--tertiary); }
::-moz-input-placeholder{ color:var(--tertiary); }
::-ms-input-placeholder{ color:var(--tertiary); }

::-webkit-calendar-picker-indicator {
    filter: invert(1);
    opacity:0.4;
}

input{
  height:25px;
}

input, button, textarea{
  padding:.5rem;
  background-color:transparent;
  color:var(--text);
}

button{
  padding:.5rem 1.5rem;
  border:none;
  letter-spacing:2px;
  transition:all .1s ease;
  cursor:pointer;

  &:hover{
    color:var(--gold);
  }
}

.errorMsg{
  color:var(--errorTxt);
  background-color:var(--errorBg);
  padding:1vh;
  border:2px solid var(--errorTxt);
  border-radius:5px;

  transition: all .25s ease;
  text-align:center;
}

.pwdReveal{
  display:grid;
  grid-template-rows:1fr;
  grid-template-columns:8fr 2fr;
}

.checkBox {
  position:absolute;
  height:0;
  width:0;
  opacity:0;
  cursor:pointer;
}

.custCheck {
  position: absolute;
  top:0;
  left:0;
  height:20px;
  width:20px;
  border:1px solid var(--tertiary);
  transition:all .2s ease;
}

.checkWrap {
  position:relative;
  cursor:pointer;

  & .custCheck:hover {
    border:1px solid var(--gold);
    background-color:var(--gold);
  }
}

.checkWrap input:checked ~ .custCheck{
  background-color:var(--gold) !important;
}

@import '@/assets/style/media';

</style>
