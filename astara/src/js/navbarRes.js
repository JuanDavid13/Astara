import $ from 'jquery';

export function resizeNav(){
  setInterval(()=>{
    if($(window).width() <= 650)
      $('#sidebar').addClass('hiddable');
    else
      $('#sidebar').removeClass('hiddable');

  },1000);
}
