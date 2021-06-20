import { createRouter, createWebHistory } from 'vue-router';

import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Landing from '../views/Landing.vue';

import Main from '../components/main/Main.vue';
import Area from '../components/area/Area.vue';
import AreaTasks from '../components/area/AreaTasks.vue';
import AreaGoals from '../components/area/AreaGoals.vue';

const Auth = require('../auth/auth');

const routes = [
  {
    path: '/dashboard',
    name: 'Home',
    component: Home,
    redirect: '/dashboard', //little trick here
    meta: { auth: true },
    beforeEnter: async (to, from, next) => {
      let isAuthenticated = await Auth.Validate();
      if(/*to.meta.auth &&*/ isAuthenticated)
        next();
      else
        next({name: 'Login'});
    },
    children: [
      {
        path: '',
        name: 'Main',
        component: Main
      },
      {
        path: '/area/:name',
        name: 'Area',
        component: Area,
        children: [
          {
            path: 'tasks',
            name: 'Tasks',
            component: AreaTasks 
          },
          {
            path: 'goals',
            name: 'Goals',
            component: AreaGoals 
          }
        ]
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { auth: false}
  },
  {
    path: '/',
    name: 'Landing' ,
    component: Landing,
    meta: { auth: false }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

/*
router.beforeEach( async (to,from,next)=>{
  const Auth = require('../auth/auth');
  let isAuthenticated = await Auth.Validate();
  if(to.meta.auth && !isAuthenticated)
    next({name: 'Login'});
  else
    next();
});
*/


export default router
