<template>
  <div class="sign">
    <el-card class="sign-card">
      <el-tabs v-model="activeTabName">
        <el-tab-pane label="Sign Up" name="signup">
          <sign-up-form></sign-up-form>
        </el-tab-pane>
        <el-tab-pane label="Sign In" name="signin">
          <sign-in-form></sign-in-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import SignUpForm from '@/components/SignUpForm.vue'
import SignInForm from '@/components/SignInForm.vue'

export default {
  components: { SignUpForm, SignInForm },

  computed: {
    activeTabName: {
      get: function () {
        return this.$route.name === 'home' ? 'signup' : this.$route.name
      },
      set: function (newValue) {
        this.$router.push({ name: newValue })
      }
    }
  }
}
</script>