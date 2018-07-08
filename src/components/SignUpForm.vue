<template>
  <div class="sign-up">
    <el-form label-width="170px" label-position="left" ref="signUpForm" :model="form" :rules="rules" @submit.native.prevent="submitForm()">
      <el-form-item label="Auth URL" prop="auth_endpoint">
        <el-input placeholder="Auth function URL" type="text" v-model="auth_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Profile Function URL" prop="profile_endpoint">
        <el-input placeholder="Profile management URL" type="text" v-model="profile_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Picture Function URL" prop="picture_endpoint">
        <el-input placeholder="Picture management URL" type="text" v-model="picture_endpoint"></el-input>
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
import api from '@/plugins/api'

export default {
  data () {
    return {
      form: {
        auth_endpoint: null,
        profile_endpoint: null,
        picture_endpoint: null,
        email: null,
        city: null,
        password: null
      },

      rules: {
        auth_endpoint: [
          { required: true, message: 'Please enter your auth url', trigger: 'blur' }
        ],
        profile_endpoint: [
          { required: true, message: 'Please enter your profile function url', trigger: 'blur' }
        ],
        picture_endpoint: [
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

  computed: {
    auth_endpoint: {
      get () {
        return this.$store.state.auth_endpoint
      },
      set (value) {
        this.form.auth_endpoint = value
        this.$store.commit('updateAuthEndpoint', value)
      }
    },

    profile_endpoint: {
      get () {
        return this.$store.state.profile_endpoint
      },
      set (value) {
        this.form.profile_endpoint = value
        this.$store.commit('updateProfileEndpoint', value)
      }
    },

    picture_endpoint: {
      get () {
        return this.$store.state.picture_endpoint
      },
      set (value) {
        this.form.picture_endpoint = value
        this.$store.commit('updatePictureEndpoint', value)
      }
    }
  },

  created () {
    this.form.auth_endpoint = this.auth_endpoint
    this.form.profile_endpoint = this.profile_endpoint
    this.form.picture_endpoint = this.picture_endpoint
  },

  methods: {
    submitForm () {
      this.$refs['signUpForm'].validate(valid => {
        if (valid) {
          this.$store.dispatch('signUp', this.form).then(() => {
            return this.$store.dispatch('signIn', this.form)
          }).then(response => {
            if ('error' in response.data) {
              this.$notify.error({
                title: 'SignIn Error',
                message: response.data.error
              });

              return Promise.reject()
            } else {
              this.$store.commit('updateToken', response.data.token)
              return this.$store.dispatch('createProfile', {
                email: this.form.email,
                city: this.form.city
              })
            }
          }).then(response => {
            this.$router.push({ name: 'profile' })
          })
        }
      })
    }
  }
}
</script>