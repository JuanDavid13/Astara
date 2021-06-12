<template>
  <div id="landing">
    <nav>
      <div>
        <svg id="navLogo" width="50px" height="50px" version="1.1" viewBox="0 0 46.029 46" xmlns="http://www.w3.org/2000/svg" xmlns:cc="http://creativecommons.org/ns#" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
         <metadata>
          <rdf:RDF>
           <cc:Work rdf:about="">
            <dc:format>image/svg+xml</dc:format>
            <dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage"/>
            <dc:title/>
           </cc:Work>
          </rdf:RDF>
         </metadata>
         <g class="logoPath" transform="translate(-67.199 -88.399)" fill-opacity="0" stroke="#F0A202" stroke-width="2">
          <circle transform="rotate(-90)" cx="-111.4" cy="90.229" r="22" opacity=".998"/>
          <ellipse cx="90.229" cy="111.4" rx="9.7289" ry="9.7289" opacity=".998"/>
          <path transform="rotate(-90)" d="m-104.93 82.991a22 22 0 0 1-6.4733-15.097" opacity=".998"/>
          <path transform="rotate(-90)" d="m-111.41 112.22a22 22 0 0 1 6.1568-14.466" opacity=".998"/>
          <path transform="rotate(-90)" d="m-111.37 67.199a22 22 0 0 1-6.5907 15.702 22 22 0 0 1-15.823 6.2941" opacity=".998"/>
          <path transform="rotate(-90)" d="m-133.97 91.135a22 22 0 0 1 22.561 20.993" opacity=".998"/>
         </g>
        </svg>
      </div>
      <button><router-link :to="{ name: 'Login' }">Entrar</router-link></button>
    </nav>
    <header>
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
        <h1>Astara, una nueva app de tareas que te ayudará organizar tu vida.</h1>
      </div>
      <div id="centerBtn">
        <button @click="features">Más</button>
      </div>
    </header>
    <main>
      <div class="feature">
        <div>
        </div>
        <div class="featureCont">
          <h2>Crea listas de tareas y objetivos</h2>
          <p>
            Ya no necesitara recordar tus tareas, mantén tus objetivos y tareas a un solo click, <br />
            recuerda todas tus tareas de forma sencilla y rápida.
          </p>
        </div>
      </div>

      <div class="feature reverse">
        <div>
        </div>
        <div class="featureCont">
          <h2>Añade nuevas areas</h2>
          <p>
            Creando areas podrás mantener organizas todos tus objetivos.<br />
            En cada Area podrás encontrar una sección para las tareas y otra para tus objetivos.
            Mantente organizado en todo momento, con las aras.
          </p>
        </div>
      </div>

      <div class="feature">
        <div>
        </div>
        <div class="featureCont">
          <h2>Encuentra tus tareas</h2>
          <p>
            Utiliza la búsqueda para encontrar tus tareas o objetivos.<br />
          </p>
        </div>
      </div>

      <div class="feature reverse">
        <div>
        </div>
        <div class="featureCont">
          <h2>Short-Cuts</h2>
          <p>
            Sé más productivo con los short-cuts.<br />
          </p>
        </div>
      </div>

      <div class="feature">
        <div>
        </div>
        <div class="featureCont">
          <h2>White mode</h2>
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

    </footer>
  </div>
</template>

<script>
export default {
  name: 'Landing',
}
</script>

<style scoped lang="scss">

@import '@/assets/style/common';

#landing{
  min-height:100vh;
  background-color:var(--primary);
  color:var(--text);

  transition: background-color .25s ease, color .25s ease;

  header, main, footer{
    padding:2rem  3rem;
  }

  .centerBtn{
    display:grid;
    place-content:center;
    
    button{
      padding: 10px 40px;
      font-size:1rem;
      background-color:var(--gold);
      color:var(--primary);
      border:none;
      box-shadow: 0 0 10px rgba(0,0,0,.5);
      cursor:pointer;
      transition:all .25s ease;
      &:hover{
        transform:scale(1.1);
        box-shadow: 0 0 20px var(--goldShadow);
      }
    }
  }

  nav{
    position:sticky;
    top:0;
    background-color:var(--primary);
    padding:1rem 3rem;
    z-index:99;
    display:flex;
    flex-direction:row;
    justify-content:space-between;
    align-items:center;
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
    min-height:100vh;
    overflow:hidden;
  }

  main{

    .feature{
      display:flex;
      flex-flow:row wrap;
      justify-content:space-between;
      align-items:center;
      gap:2rem;

      & > div{
        width:50%;
      }
    }

    .reverse{
      flex-flow:row-reverse wrap;
    }
  }

  footer{
    position:relative;
    min-height:20vh;
    width:100%;
    background-color:rgb(30,30,30);
  }
}
</style>
