<template>
  <div class="sign">
    <el-card class="sign-card">
      <el-tabs>
        <el-tab-pane label="Sign Up">
          <sign-up-form></sign-up-form>
        </el-tab-pane>
        <el-tab-pane label="Sign In">
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
  components: { SignUpForm, SignInForm }
}
</script>