<template>
  <div id="profile">
      <h2>Perfil</h2>
      <p>Nombre de usuario</p>
      <p class="info">El nombre de usuario es único y te identifica en la aplicación.</p>
      <input type="text" :value="username">
      <p>Email</p>
      <p class="info">Recuerda que también puedes usar tu correo electrónico para acceder a la aplicación.</p>
      <input type="text" :value="email">
      <hr>
      <p>Cambiar contraseña</p>
      <p class="info">Introduce tu contraseña antigua para cambiar la.</p>
      <div>
        <input id="profilePwd" type="password">
        <button>Cambiar contraseña</button>
      </div>
      <label>
          <input @change="togglePwd" type="checkbox" v-model="show">
          <span v-if="show">Ocultar</span>
          <span v-else>Mostrar</span>
      </label>
      <hr>
      <p>Theme</p>
      <label>
        <input @change="checkTheme" type="checkbox" v-model="theme">
        <span v-if="!theme">Oscuro</span>
        <span v-else>Claro</span>
      </label>
  </div>
</template>

<script>
import $ from 'jquery';

export default {
  name: 'Profile',
  data() {
    return {
        username:"algo",
        email:"algo@gmail.com",
        theme:false,
        show:false,
    }
  },
  methods: {
    checkTheme(e){
        if(e.target.checked){ $('#app').addClass('lightTheme');
        }else{ $('#app').removeClass('lightTheme'); }
    },
    togglePwd(){
        console.log($('#profilePwd').type);
        if($('#profilePwd')[0].type == "password") {
               $('#profilePwd')[0].type = 'text';
        }else{ $('#profilePwd')[0].type = 'password'; }
    }
  },
}

</script>

<style lang="scss">
#profile{
    display:flex;
    flex-direction:column;

    width:50vw;

    h2{ margin-bottom:15px; }

    input { margin-bottom:10px; }

    .info{
        color:var(--tertiary);
        font-size:.8rem;
        margin:5px 0 10px 0;
    }

    label span{
        margin-left:15px;
    }
}
</style>
