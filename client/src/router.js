import VueRouter from 'vue-router'
import Vue from 'vue'

Vue.use(VueRouter)

// { RouteConfig }
const routes = [
  {
    path: '/',
    name: 'home',
    redirect: {
      name: 'goods-list'
    }
  },
  {
    path: '/goods/list',
    name: 'goods-list',
    components: {
      default: () => import('./views/Goods/goods-list'),
      sidebar: () => import('./components/sidebar')
    },
  },
  {
    path: '/goods/write-in',
    name: 'goods-write-in',
    components: {
      default: () => import('./views/Goods/goods-write-in'),
      sidebar: () => import('./components/sidebar')
    },
  }
];

const routesHelper = function (routes, cmptNames) {
  for (let r of routes) {
    if (r.components) {
      r.meta = r.meta || {};
      cmptNames.forEach(n => {
        r.meta[n] = !!r.components[n]
      });
    }
  }
}
routesHelper(routes, ['sidebar'])


export default new VueRouter({
  routes
})