<template>
  <div class="sidebar">
    <div id="sidebarWrap">
      <div class="logo noselect">Astara</div>
      <p class="noselect splitChar" data-splitting>AREAS</p>
      <div id="areas">
        <router-link :to="{name: 'Main'}">MAIN</router-link>
        <div class="area" v-for="area in areas" :key="area.id" >
          <router-link :to="{ name: 'Area', params: { name: area.slug} }" >{{area.name}}</router-link>
          <span class="deleteable" v-if="area.deleteable">X</span>
        </div>
      </div> 
      <div>
        <button @click="logOut">Log Out</button>
      </div>
      <div>
        <button @click="switchTheme">Change theme</button>
      </div>
    </div>
  </div>
</template>

<script>
/*const axios = require('axios');
const Axios = axios.create({
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 1000,
  withCredentials: true,
});*/

import Axios from '@/auth/auth';
import $ from 'jquery';

  export default{
    name: 'Sidebar',
    props: {areas: Object},
    emits: ['switchTheme'],
    data(){
      return {
      }
    },
    methods: {
      logOut(){ Axios.get('/auth/logout').then(()=>{ this.$router.push({name:'Login'}) }); },
      switchTheme(){
        $('#app').toggleClass('lightTheme');
      },
    }
  }
</script>

<style scoped lang="scss">

@import '@/assets/style/common';

.sidebar {
  color:var(--text);
  background-color:var(--primary);
  border-right:1px solid var(--secondary);
  display: flex;
  flex-direction: column;
  align-items:flex-start;
  text-align:left;

  padding:1vh 1.5vw;

  & #sidebarWrap{
    position:sticky;
    top:1vh;

    & > div:first-child {
      color:goldenrod;
      margin-bottom:3vh;
      font-size:1.5rem;
      //width:max-content;
    }
  }

    a{
      color:var(--text);
      text-decoration:none;
    }


    #areas{ 
      margin-top:15px;
      text-align:left; 
      font-size:.9rem;
    }

    .area{
      margin-bottom: 5px;
      cursor: pointer;

      text-transform:uppercase;

      .deleteable{
        margin-left:0;
        color:var(--primary);
        transition: color .25s ease, margin .25s ease;
      }

      &:hover .deleteable{
        margin-left:15px;
        color:var(--contrary); 
      }
    }
  }
</style>
