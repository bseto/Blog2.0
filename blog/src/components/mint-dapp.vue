<template>
  <div id="app">
    <metamask-util/>
  </div>
</template>

<script>
import MetamaskUtil from '@/components/metamask-util'
export default {
  name: 'mint-dapp',
  components: {
    'metamask-util': MetamaskUtil
  }
}
</script>
<style scoped>
</style>
