<template>
  <div class="sidebar">
    <div id="sidebarWrap">
      <div class="logo noselect">Astara</div>
      <div id="areas">
        <p class="noselect splitChar" data-splitting>AREAS</p>
        <router-link :to="{name: 'Main'}">MAIN</router-link>
        <div class="area" v-for="area in areas" :key="area.id" >
          <router-link :to="{ name: 'Area', params: { name: area.slug } }" >{{area.name}}</router-link>
        </div>
        <div>
          <form @submit.capture="addArea">
            <input id="addArea" v-model="newArea" type="text" minlegth="3" maxlength="20" spellcheck="false" autocomplete="off" placeholder="+ Area">
          </form>

        </div>
      </div> 
      <div>
        <div>
          <div>OOO</div>
          <a href="#" @click="$emit('openprofile')"><span>{{username}}</span></a>
          <div>...</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import $ from 'jquery';

import Axios from '@/auth/auth';

  export default{
    name: 'Sidebar',
    props: {username: String},
    emits: ['openprofile', 'updateAreas'],
    data() {
      return {
        newArea:"",
        areas: {},
      }
    },
    methods: {
      addArea(e){
        e.preventDefault();
        Axios.post("/area/create",{name:this.newArea}).then((res)=>{
          if(res.data.added){
            this.areas = JSON.parse(res.data.areas);
            $(document).ready(()=>{ 
              $('#addArea').val(""); 
              $('#addArea').blur();
            });
          }
        })
      },
      deleteArea(slug){
        let pos = this.areas.findIndex(area => area.slug === slug);
        let selector = "#areas div.area:nth-of-type(" + (pos +1) + ")";
        $(selector).fadeOut('fast',()=>{
          this.areas.splice(pos,1);
        });
      },
      getAreas(){
        Axios.get("/area").then((res)=>{
          this.areas = JSON.parse(res.data); 
          console.log(res.data);
        });
      }
    },
    created(){
      this.getAreas();
    },
  }
</script>

<style scoped lang="scss">

@import '@/assets/style/common';

.sidebar {
  color:var(--text);
  background-color:var(--primary);
  border-right:1px solid var(--secondary);
  text-align:left;

  padding:3vh;

  & * { max-width:100%; }
  & :is(p,a,span) {
    //equivalente a
    //p, a, span {

    overflow:hidden;
    text-overflow:ellipsis;
    white-space:nowrap;
  }

  & #sidebarWrap{
    position:sticky;
    top:1vh;

    height:100%;

    display:grid;
    grid-template-columns:1fr;
    grid-template-rows: minmax(0,5%) 85% minmax(0,10%);

    & > div:first-child {
      color:var(--gold);
      font-size:1.7rem;
    }


    #areas{ 
      text-align:left; 
      font-size:.9rem;
      overflow-x:hidden;
      overflow-y:scroll;
    }

    & > div:last-child {
      display:flex;
      align-items:flex-end;
    }

    & > div:last-child > div{
      height:fit-content;
      display:grid;
      grid-template-columns:2fr 4fr 1fr;
      grid-template-rows:1fr;

      span { overflow:hidden; }
    }

    #addArea{
      border:none;
      overflow:hidden;
      padding: calc(0.5rem + 1px);

      &:hover { cursor:pointer; }
      &:focus, &:focus-visible{
        padding:0.5rem;
        border:1px solid var(--gold);
        outline:none;
      }

    }

  }

  a{
    color:var(--text);
    text-decoration:none;
  }

  .area{
    margin-bottom: 5px;
    cursor: pointer;

    text-transform:uppercase;
  }
}
</style>
