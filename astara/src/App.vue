<template>
  <!--<div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </div>-->
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
  border:2px solid lightgrey;
  transition:background-color .25s ease;
}

.checkWrap {
  position:relative;
  cursor:pointer;
}

.checkWrap input:checked ~ .custCheck{
  background-color:lightgrey;
}

.noselect{
  -webkit-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}

input {
  background-color:var(--primary);
  color:var(--text);

  border:1px solid var(--contrary);
  border-radius:5px;

  padding:7px;
}
input:focus{
  outline:2px solid (--primary);
}

button{
  background-color:var(--contrary);
  color:var(--primary);

  border:1px solid var(--primary);
  border-radius:5px;

  padding:7px;
}

</style>
