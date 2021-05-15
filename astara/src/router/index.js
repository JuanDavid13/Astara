import { createRouter, createWebHistory } from 'vue-router'

import Home from '../views/Home.vue'
import Login from '../views/Login.vue'

//import Main from '../views/Main.vue'
import Main from '../components/main/Main.vue'

//import Area from '../views/Area.vue'
import Area from '../components/area/Area.vue'

const Auth = require('../auth/auth');

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    redirect: '/', //little trick here
    meta: { auth: true },
    beforeEnter: async (to, from, next) => {
      let isAuthenticated = await Auth.Validate();
      if(/*to.meta.auth &&*/ !isAuthenticated)
        next({name: 'Login'});
      else
        next();
    },
    children: [
      {
        path: '/area/:name',
        name: 'Area',
        component: Area,
      },
      {
        path: '',
        name: 'Main',
        component: Main
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: { auth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { auth: false}
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
