<template>
  <div class="sign-up">
    <el-form label-width="120px" label-position="left" ref="signUpForm" :model="form" :rules="rules" @submit.native.prevent="submitForm()">
      <el-form-item label="Auth URL" prop="auth_url">
        <el-input placeholder="Auth function URL" type="text" v-model="form.auth_url"></el-input>
      </el-form-item>
      <el-form-item label="Function URL" prop="func_url">
        <el-input placeholder="Profile management URL" type="text" v-model="form.func_url"></el-input>
      </el-form-item>
      <el-form-item label="Your email" prop="email">
        <el-input placeholder="Email" type="email" v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="City" prop="city">
        <el-input placeholder="City" type="text" v-model="form.city"></el-input>
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input placeholder="Password" type="password" v-model="form.password"></el-input>
      </el-form-item>
      
      <div class="sw-form-actions">
        <button type="submit" style="display: none"></button>
        <el-button type="primary" @click="submitForm">Sign Up</el-button>
      </div>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'sign-up-form',
  data () {
    return {
      form: {
        auth_url: null,
        func_url: null,
        email: null,
        city: null,
        password: null
      },
      rules: {
        auth_url: [
          { required: true, message: 'Please enter your auth url', trigger: 'blur' }
        ],
        func_url: [
          { required: true, message: 'Please enter your function url', trigger: 'blur' }
        ],
        email: [
          { required: true, message: 'Please enter your email', trigger: 'blur' },
          { type: 'email', message: 'Please input correct email address', trigger: 'blur' }
        ],
        password: [
          { required: true, message: 'Please enter your password', trigger: 'blur' }
        ]
      }
    }
  },

  created () {
    if (localStorage.getItem('auth_url')) {
      this.form.auth_url = localStorage.getItem('auth_url')
    }

    if (localStorage.getItem('func_url')) {
      this.form.func_url = localStorage.getItem('func_url')
    }
  },

  methods: {
    submitForm () {
      this.$refs['signUpForm'].validate(valid => {
        if (valid) {
          localStorage.setItem('auth_url', this.form.auth_url)
          localStorage.setItem('func_url', this.form.func_url)

          this.axios.post(this.form.auth_url + '?action=signup&userid=' + this.form.email + '&password=' + this.form.password + '&city=' + this.form.city, {
            // action: 'signup',
            // userid: this.form.email,
            // password: this.form.password,
            // city: this.form.city
          }).then(response => {
            this.$notify.success({
              title: 'Success',
              message: 'SignUp successful'
            })
          })
        }
      })
    }
  }
}
</script>