import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'

import $ from 'jquery';

createApp(App).use(router).mount('#app')


let oldkey = 0;
$(document).keydown((e)=>{
  console.log(oldkey);
  console.log(e);
  if(oldkey == 18 && e.keyCode == 65){
    console.log("combination");
    $('#modal').toggleClass('modalActive');
  }

  //console.log(e.KeyCode);
  oldkey = e.keyCode;
});
