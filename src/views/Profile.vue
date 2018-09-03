<template>
  <div class="profile">
    <el-row>
      <el-col>
        <el-tabs v-model="activeTabName">
          <el-tab-pane label="TASKS" name="tasks"></el-tab-pane>
          <el-tab-pane label="PROFILE" name="profile"></el-tab-pane>
        </el-tabs>

        <profile-form></profile-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ProfileForm from '@/components/ProfileForm.vue'

export default {
  components: { ProfileForm },

  data () {
    return {
      activeTabName: 'profile'
    }
  },

  watch: {
    activeTabName (val) {
      this.$router.push({ name: val })
    }
  },

  created () {
    if (this.$store.state.token === null) {
      this.$router.push({ name: 'signin' })
    }
  }
}
</script>
