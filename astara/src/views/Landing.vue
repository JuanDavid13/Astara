<template>
  <div id="landing">
    <nav>
      <div>
        <logoName :size="2"/>
      </div>
      <router-link :to="{ name: 'Login' }"><button class="simpleBtn">Entrar</button></router-link>
    </nav>
    <header>
      <div id="headerImg">
          <div class="img"></div>
      </div>
      <div>
      <svg id="logo" width="10%" height="10%" version="1.1" viewBox="0 0 46.029 46" xmlns="http://www.w3.org/2000/svg" xmlns:cc="http://creativecommons.org/ns#" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
       <metadata>
        <rdf:RDF>
         <cc:Work rdf:about="">
          <dc:format>image/svg+xml</dc:format>
          <dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage"/>
          <dc:title/>
         </cc:Work>
        </rdf:RDF>
       </metadata>
       <g id="headerLogoPath" class="logoPath" transform="translate(-67.199 -88.399)" fill-opacity="0" stroke="#F0A202" stroke-width="2">
        <circle transform="rotate(-90)" cx="-111.4" cy="90.229" r="22" opacity=".998"/>
        <ellipse cx="90.229" cy="111.4" rx="9.7289" ry="9.7289" opacity=".998"/>
        <path transform="rotate(-90)" d="m-104.93 82.991a22 22 0 0 1-6.4733-15.097" opacity=".998"/>
        <path transform="rotate(-90)" d="m-111.41 112.22a22 22 0 0 1 6.1568-14.466" opacity=".998"/>
        <path transform="rotate(-90)" d="m-111.37 67.199a22 22 0 0 1-6.5907 15.702 22 22 0 0 1-15.823 6.2941" opacity=".998"/>
        <path transform="rotate(-90)" d="m-133.97 91.135a22 22 0 0 1 22.561 20.993" opacity=".998"/>
       </g>
      </svg>
        <h1 id="bannerTxt" class="splitWords" data-splitting="words">Astara, una nueva app de tareas que te ayudará organizar tu vida.</h1>
      </div>
      <div id="centerBtn">
        <button @click="features">Más</button>
      </div>
    </header>
    <main>
      <hr>
      <div class="feature">
        <div class="featureImg" id="firstImg">
          <img class="img" src="/img/tareasAstara.png" alt="tareas">
        </div>
        <div class="featureCont">
          <h2 class="splitChars" data-splitting>Crea listas de tareas y objetivos</h2>
          <p>
            Ya no necesitara recordar tus tareas, mantén tus objetivos y tareas a un solo click,
            recuerda todas tus tareas de forma sencilla y rápida.
          </p>
        </div>
      </div>

      <hr>
      <div class="feature reverse">
        <div class="featureImg" id="secondImg">
          <img class="img" src="/img/tareasAstara.png" alt="tareas">
        </div>
        <div class="featureCont">
          <h2 class="splitChars" data-splitting>Añade nuevas areas</h2>
          <p>
            Creando areas podrás mantener organizas todos tus objetivos.<br />
            En cada Area podrás encontrar una sección para las tareas y otra para tus objetivos.
            Mantente organizado en todo momento, con las aras.
          </p>
        </div>
      </div>

      <hr>
      <div class="feature">
        <div class="featureImg" id="thirdImg">
          <img class="img" src="/img/tareasAstara.png" alt="tareas">
        </div>
        <div class="featureCont">
          <h2 class="splitChars" data-splitting>Encuentra tus tareas</h2>
          <p>
            Utiliza la búsqueda para encontrar tus tareas o objetivos.<br />
          </p>
        </div>
      </div>

      <hr>
      <div class="feature reverse" >
        <div class="featureImg" id="fourImg">
          <img class="img" src="/img/tareasAstara.png" alt="tareas">
        </div>
        <div class="featureCont">
          <h2 class="splitChars" data-splitting>Short-Cuts</h2>
          <p>
            Sé más productivo con los short-cuts.<br />
          </p>
        </div>
      </div>

      <hr>
      <div class="feature" id="fifthFeature">
        <div class="featureImg" id="fifthImg">
          <img class="img" src="/img/tareasAstara.png" alt="tareas">
        </div>
        <div class="featureCont">
          <h2 class="splitChars" data-splitting>White mode</h2>
          <p>
            Para los amantes del white mode.<br />
          </p>
        </div>
      </div>
    <div class="centerBtn">
      <button><router-link :to="{ name: 'Login' }">Entrar</router-link></button>
    </div>
    </main>
    <footer>
      <span>Desarrollado por: Juan David Ramírez Bernal</span>
    </footer>
  </div>
</template>

<script>
import "splitting/dist/splitting.css";
import Splitting from "splitting";

import logoName from '@/components/commons/logoName.vue'
import { parallax } from '@/js/parallax.js';

import $ from 'jquery';

export default {
  name: 'Landing',
  components:{
    logoName,
  },
  mounted(){
    parallax();
    Splitting();

    const options1 = {
      root: null,
      threshold: 0,
      rootMargin: "0px"
    };

    const options2 = {
      root: null,
      threshold: 0.5,
      rootMargin: "0px"
    };

    $('document').ready(()=>{
      let observer1 = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          let h2 = $(entry.target)[0].children[1].children[0];
          if(entry.isIntersecting){
            $(h2).addClass('splitCharsReverse');
          }
        });
      },options1);

      let features = $('.feature');
      features.each((index)=> {
        observer1.observe(features[index]);
      });

      let observer2 = new IntersectionObserver((entries)=>{
        entries.forEach(entry =>{
          console.log(entry);
          if(entry.isIntersecting){
            $('#app').addClass('lightTheme');
          }else{
            $('#app').removeClass('lightTheme');
          }
        });
      },options2);

      console.log($('#fifthFeature'));
      observer2.observe($('#fifthFeature')[0]);

    });

  }
}
</script>

<style scoped lang="scss">

@import'@/assets/style/common';

#landing{
  min-height:100vh;
  background-color:var(--primary);
  color:var(--text);

  transition: background-color .25s ease, color .25s ease;

  header, main{
    padding:2rem  3rem;
    color:var(--text);
    h2{
      text-transform:uppercase;
    }
  }

  .centerBtn{
    display:grid;
    place-content:center;
    
    button{
      padding: 10px 40px;
      font-size:1rem;
      background-color:transparent;
      color:var(--gold);
      border:none;
      cursor:pointer;
      transition:all .25s ease;
    }
  }

  nav{
    position:sticky;
    top:0;
    padding:1rem 3rem;
    z-index:99;
    display:flex;
    flex-direction:row;
    justify-content:space-between;
    align-items:center;

    & > div:first-child{
      display:flex;
      flex-direction:row;
      //justify-content:space-between;
      align-items:center;
      gap:2rem;
      font-size:2rem;
      color:var(--gold);
    }
  }

  #logo{
    position:absolute;
    top:30%;
    left:0%;
    transform: rotate(-15deg);
    transform-origin:50%;
    width:50%;
    height:auto;

    #headerLogoPath{
      stroke:var(--contrary);
      opacity:0.05;
    }
  }

  .logoPath{
   stroke:var(--gold);
  }

  header{
    min-height:50vh;
    overflow:hidden;

    #bannerTxt{
      max-width:35ch;
      text-align:center;
    }

  }

  main{
    overflow:hidden;

    .feature{
      position:relative;
      display:grid;
      grid-template-columns:repeat(2, 1fr);
      grid-template-rows:1fr;
      grid-gap:2rem;
      overflow:visible;
      min-height:50vh;

      .featureCont{
        height:100%;
        position:relative;
        display:flex;

        h2 {
          font-size:2.5rem;
          margin-bottom:25px;
          position:absolute;
          top:5vh;
          width:100%;
          right:10vw;
          //opacity:0;
        }
        p{
          align-self:flex-end;
          margin-bottom:25px;
          font-family: 'Lora', serif;
        }
      }
    }

    .reverse{
      .featureImg{
        order:2;
      }
      .featureCont{
        order:1;
        h2{
          left:10vw;
        }
      }
    }
  }

  .featureImg{
    width:800px;
    height:450px;
    //background-color:var(--tertiary);
    position:relative;
    //overflow:hidden;

    .img{
      position:absolute;
      width:100%; 
      height:auto;
      background-color:orange;

      top:0;
      left:50%;
      transform:translateX(-50%);
      transition:all .2s ease;

      box-shadow:0 0 15px rgba(0,0,0,.2);
    }
  }

  footer{
    padding:1rem;
    position:relative;
    width:100%;
    background-color:rgba(0,0,0,.3);
    color:var(--contrary);
    transition:all .25s ease;
    text-align:center;
  }
}

</style>
