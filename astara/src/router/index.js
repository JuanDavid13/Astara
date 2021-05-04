import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { auth: true }
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

router.beforeEach( async (to,from,next)=>{
  const Auth = require('../auth/auth');
  //let response = await Auth.validate();
  //if(response){
  //  console.log("algo");
  //  console.log(response);
  //}
  let isAuthenticated = await Auth.validate();
  if(to.meta.auth && !isAuthenticated)
    next({name: 'Login'});
  else
    next();
})


export default router
