import $ from 'jquery';

export function parallax(){
  $(window).scroll(()=>{
    let windowsoff = $(window).scrollTop();

    let firstImgoff = $('#firstImg .img').offset().top;
    $('#firstImg .img').css('top',-(firstImgoff - windowsoff)*.2);

    let secondImgoff = $('#secondImg .img').offset().top;
    $('#secondImg .img').css('top',-(secondImgoff - windowsoff)*.2);

    let thirdImgoff = $('#thirdImg .img').offset().top;
    $('#thirdImg .img').css('top',-(thirdImgoff- windowsoff)*.2);

    let  fourImgoff = $('#fourImg .img').offset().top;
    $('#fourImg .img').css('top',-(fourImgoff - windowsoff)*.2);

    let  fifthImgoff = $('#fifthImg .img').offset().top;
    $('#fifthImg .img').css('top',-(fifthImgoff - windowsoff)*.2);
  });
}
