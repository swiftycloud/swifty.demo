<template>
  <el-container id="app">
    <el-header height="64px" class="sw-header">
      <a href="/" class="sw-brand">swifty.demo</a>
      <a href="#" class="sw-logout" @click.prevent="logOut" v-if="this.$store.state.token">Log Out</a>
    </el-header>
    <el-main>
      <router-view/>
    </el-main>
  </el-container>
</template>

<script>
export default {
  methods: {
    logOut () {
      this.$store.commit('clearToken')
      this.$router.push({ name: 'signin' })
    }
  }
}
</script>