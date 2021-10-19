import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/hello-world'
import MintDapp from '@/components/mint-dapp'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/dapp',
      name: 'dapp',
      component: MintDapp
    },
    {
      path: '/',
      name: 'hello-world',
      component: HelloWorld
    }
  ]
})
