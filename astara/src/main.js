import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'

createApp(App).use(router).mount('#app')

import $ from 'jquery';
$(document).keydown((e)=>{
  if(/*e.ctrlKey && */e.altKey){
    switch(e.keyCode){
      case 65: { $('#modal').toggleClass('modalActive'); }break;
    }
  }
});
