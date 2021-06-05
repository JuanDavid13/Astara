import $ from 'jquery';

export function setShortCodes(context){
  //let target;
  //$('#modal').ready(()=>{ target = $('#modalCont')[0]; });
  //const observer = new MutationObserver((mutations)=>{
  //  console.log("do something");
  //  console.log(mutations);
  //  //mutations.forEach((mutation)=>{
  //  //  console.log("algo");
  //  //  console.log(mutation);
  //  //});
  //});
  //const config = { attributes: true, childList: true, characterData: true };

  $(document).keydown((e)=>{
    //observer.observe(target,config);

    if(e.keyCode == 27){ //esq key
      $('#modal').removeClass('modalActive'); 
      $(document).ready(()=>{
        switch( context.$data.modalOption ){
          case 1: { context.$refs.profile.closeModal(); }break;
          case 2: { context.$refs.profile.closeModal(); }break;
          case 3: { context.$refs.profile.closeModal(); }break;
        }
      });
    }

      if(/*e.ctrlKey && */e.altKey){
        switch(e.keyCode){
          case 80: { //p --> alt + p
            context.$data.modalOption = 1;
            $(document).ready(()=>{
              if($('#modal').hasClass('modalActive')){
                context.$refs.profile.closeModal();
                $('#modal').removeClass('modalActive');
              }else{
                context.$refs.profile.openModal();
                $('#modal').addClass('modalActive');
              }
            });
          }break; //repeat for the rest
          case 71: { //g --> alt + g
            $('#modal').toggleClass('modalActive'); 
            context.$data.modalOption = 2;
          }break;
          case 65: {
            $('#modal').toggleClass('modalActive'); 
            context.$data.modalOption = 3;
          }break;
        }
      }

      //observer.disconnect();
  });

}
