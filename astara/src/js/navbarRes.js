import $ from 'jquery';

/**
 * Función auxiliar para hacher el sidebar responsive.
 *
 * @function
 */
export function resizeNav(){
  setInterval(()=>{
    if($(window).width() <= 650)
      $('#sidebar').addClass('hiddable');
    else
      $('#sidebar').removeClass('hiddable');

  },1000);
}
