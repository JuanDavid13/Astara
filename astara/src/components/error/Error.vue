<template>
  <div class="error" :class="{mistaken: error}">
    <span>{{message}}</span>
    <button @click="closeErr">X</button>
  </div>
</template>
<script>
export default{
  name: 'error',
  data(){
    return{
      message:"",
      error: false,
    }
  },
  methods:{
    setErr(msg){
      this.message = msg;
      this.error = true;
    },
    closeErr(){
      this.error = false;
      setTimeout(()=>{
        this.message = "Error inesperado";
      },500);
    }
  }
}
</script>
<style scoped lang="scss">
@import '@/assets/style/common.scss';
.error{
  position:relative;
  height:0;  
  max-height:5rem;
  padding:0;
  overflow:hidden;
  background-color:var(--errorBg);
  color:var(--errorTxt);
  transition:all .25s ease;
  border-radius:5px;
  
  & button{
    position:absolute;
    top:5px;
    right:5px;
    background-color:transparent;
    color:var(--errorTxt);
    border:none;
    cursor:pointer;
    font-size:1rem;
  }
} 
.mistaken{
  height:fit-content;  
  padding:10px;
  border:2px solid var(--errorTxt);
}
</style>
