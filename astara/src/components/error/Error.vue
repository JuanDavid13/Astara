<template>
  <div class="error errClosed" :class="{err: hasError}">
    {{message}}
    <button @click="closeErr">X</button>
  </div>
</template>

<script>
export default{
  name: 'Error',
  data() {
    return{
      hasError: false,
      message:"Error inesperado",
    }
  },
  methods:{
    closeErr(){
      this.hasError = false;
      setTimeout(()=>{ this.message = "Error inesperado"; },500);
    },
    setErr(e,msg){
      this.hasError = true;
      this.message = msg; 
    },
  }
}
</script>

<style lang="scss">
.error{
  color:red;
  height:0;
  overflow:hidden;
  transition:height .25s ease-in-out;
}
.err{
  height:3rem;
}
</style>
