import $ from 'jquery';

export function parallax(){
  $(window).scroll(()=>{
    let windowsoff = $(window).scrollTop();

    let firstImgoff = $('#firstImg .img').offset().top;
    //let  percentage = (windowsoff-firstImgoff);
    $('#firstImg .img').css('top',((windowsoff-firstImgoff)+500)*.15 +'px');

    let secondImgoff = $('#secondImg .img').offset().top;
    $('#secondImg .img').css('top',((windowsoff-secondImgoff)+500)*.15 + 'px');

    let thirdImgoff = $('#thirdImg .img').offset().top;
    $('#thirdImg .img').css('top',((windowsoff-thirdImgoff)+500)*.15 + 'px');

    let  fourImgoff = $('#fourImg .img').offset().top;
    $('#fourImg .img').css('top',((windowsoff-fourImgoff)+500)*.15 + 'px');

    let  fifthImgoff = $('#fifthImg .img').offset().top;
    $('#fifthImg .img').css('top',((windowsoff-fifthImgoff)+500)*.15 + 'px');
  });
}
