<template>
  <div>
    <p>Hello World</p>
    <!--<header class="app__header"> -->
      <!--<p v-on:click="goHome" class="app__header__logo">dapp</p> -->
    <!--</header> -->
  </div>
</template>
<script>
export default {
  name: 'hello-world',
  methods: {
    goHome: function () {
      this.$router.push({ name: 'dapp' })
    }
  }
}
</script>
<style scoped>
</style>
