<template>
  <!--<div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </div>-->
  <router-view/>
</template>

<style lang="scss">

*,*::before,*::after{
  margin:0;
  padding:0;
  box-sizing:border-box;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
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
  border:1px solid grey;
}

.checkWrap {
  position:relative;
  cursor:pointer;
  -webkit-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}

.checkWrap input:checked ~ .custCheck{
  background-color:grey;
}

</style>
