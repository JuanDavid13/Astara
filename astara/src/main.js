import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'



createApp(App).use(router).mount('#app')

//move from here
import $ from 'jquery';
let oldkey = 0;
$(document).keydown((e)=>{
  if(oldkey == 18 && e.keyCode == 65){ $('#modal').toggleClass('modalActive');  } //toggle modal
  oldkey = e.keyCode;
});
