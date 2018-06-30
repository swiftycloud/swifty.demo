<template>
  <div class="sign-in">
    <el-form label-width="170px" label-position="left" ref="signInForm" :model="form" :rules="rules" @submit.native.prevent="submitForm()">
      <el-form-item label="Auth URL" prop="auth_url">
        <el-input placeholder="Auth function URL" type="text" v-model="form.auth_url"></el-input>
      </el-form-item>
      <el-form-item label="Profile Function URL" prop="profile_func_url">
        <el-input placeholder="Profile management URL" type="text" v-model="form.profile_func_url"></el-input>
      </el-form-item>
      <el-form-item label="Picture Function URL" prop="picture_func_url">
        <el-input placeholder="Picture management URL" type="text" v-model="form.picture_func_url"></el-input>
      </el-form-item>
      <el-form-item label="Your email" prop="email">
        <el-input placeholder="Email" type="email" v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input placeholder="Password" type="password" v-model="form.password"></el-input>
      </el-form-item>
      
      <div class="sw-form-actions">
        <button type="submit" style="display: none"></button>
        <el-button type="primary" @click="submitForm">Sign In</el-button>
      </div>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'sign-in-form',
  data () {
    return {
      form: {
        auth_url: null,
        profile_func_url: null,
        picture_func_url: null,
        email: null,
        city: null,
        password: null
      },
      rules: {
        auth_url: [
          { required: true, message: 'Please enter your auth url', trigger: 'blur' }
        ],
        profile_func_url: [
          { required: true, message: 'Please enter your profile function url', trigger: 'blur' }
        ],
        picture_func_url: [
          { required: true, message: 'Please enter your picture function url', trigger: 'blur' }
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

    if (localStorage.getItem('profile_func_url')) {
      this.form.profile_func_url = localStorage.getItem('profile_func_url')
    }

    if (localStorage.getItem('picture_func_url')) {
      this.form.picture_func_url = localStorage.getItem('picture_func_url')
    }
  },

  methods: {
    submitForm () {
      this.$refs['signInForm'].validate(valid => {
        if (valid) {
          localStorage.setItem('auth_url', this.form.auth_url)
          localStorage.setItem('func_url', this.form.func_url)

          this.axios.post(this.form.auth_url + '?action=signin&userid=' + this.form.email + '&password=' + this.form.password, {
            // action: 'signup',
            // userid: this.form.email,
            // password: this.form.password,
            // city: this.form.city
          }).then(response => {
            if ('error' in response.data) {
              this.$notify.error({
                title: 'Error',
                message: response.data.error
              });
            } else {
              this.$notify.success({
                title: 'Success',
                message: 'SignIn successful'
              })

              localStorage.setItem('token', response.data.token)
              this.$router.push({ name: 'profile' })
            }
          })
        }
      })
    }
  }
}
</script>