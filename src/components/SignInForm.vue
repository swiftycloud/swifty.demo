<template>
  <div class="sign-in">
    <el-form label-width="170px" label-position="left" ref="signInForm" :model="form" :rules="rules" @submit.native.prevent="submitForm()">
      <el-form-item label="Auth URL" prop="auth_endpoint">
        <el-input placeholder="Auth function URL" type="text" v-model="auth_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Facebook Auth URL" prop="facebook_endpoint">
        <el-input placeholder="Facebook function URL" type="text" v-model="facebook_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Profile Function URL" prop="profile_endpoint">
        <el-input placeholder="Profile management URL" type="text" v-model="profile_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Picture Function URL" prop="picture_endpoint">
        <el-input placeholder="Picture management URL" type="text" v-model="picture_endpoint"></el-input>
      </el-form-item>
      <el-form-item label="Tasks Function URL" prop="tasks_endpoint">
        <el-input placeholder="Tasks management URL" type="text" v-model="tasks_endpoint"></el-input>
      </el-form-item>

      <div class="sw-form-actions sw-form-actions--facebook">
        <el-button type="primary" plain @click.prevent="signInWithFacebook">Sign In With Facebook</el-button>
      </div>

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
import * as config from '@/config'

export default {
  data () {
    return {
      form: {
        auth_endpoint: null,
        facebook_endpoint: null,
        profile_endpoint: null,
        picture_endpoint: null,
        tasks_endpoint: null,
        email: null,
        password: null
      },

      rules: {
        auth_endpoint: [
          { required: true, message: 'Please enter your auth url', trigger: 'blur' }
        ],
        facebook_endpoint: [
          { required: true, message: 'Please enter your facebook function url', trigger: 'blur' }
        ],
        profile_endpoint: [
          { required: true, message: 'Please enter your profile function url', trigger: 'blur' }
        ],
        picture_endpoint: [
          { required: true, message: 'Please enter your picture function url', trigger: 'blur' }
        ],
        tasks_endpoint: [
          { required: true, message: 'Please enter your tasks function url', trigger: 'blur' }
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
    this.form.auth_endpoint = this.auth_endpoint
    this.form.facebook_endpoint = this.facebook_endpoint
    this.form.profile_endpoint = this.profile_endpoint
    this.form.picture_endpoint = this.picture_endpoint
    this.form.tasks_endpoint = this.tasks_endpoint

    if ('code' in this.$route.query) {
      this.$store.dispatch('signInWithFacebook', {
        code: this.$route.query.code,
        redirect_uri: encodeURIComponent(window.location.origin + window.location.pathname)
      }).then(response => {
        this.$store.commit('updateToken', response.data.token)
        this.$router.push({ name: 'profile' })
      })
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

    facebook_endpoint: {
      get () {
        return this.$store.state.facebook_endpoint
      },
      set (value) {
        this.form.facebook_endpoint = value
        this.$store.commit('updateFacebookEndpoint', value)
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
    },

    tasks_endpoint: {
      get () {
        return this.$store.state.tasks_endpoint
      },
      set (value) {
        this.form.tasks_endpoint = value
        this.$store.commit('updateTasksEndpoint', value)
      }
    }
  },

  methods: {
    submitForm () {
      this.$refs['signInForm'].validate(valid => {
        if (valid) {
          this.$store.dispatch('signIn', this.form).then(response => {
            if ('error' in response.data) {
              this.$notify.error({ title: 'Error', message: response.data.error })
            } else {
              this.$notify.success({ title: 'Success',message: 'SignIn successful' })
              this.$store.commit('updateToken', response.data.token)
              this.$router.push({ name: 'tasks' })
            }
          })
        }
      })
    },

    signInWithFacebook () {
      const redirect_uri = encodeURIComponent(window.location.origin + window.location.pathname)
      const url = 'https://www.facebook.com/v3.0/dialog/oauth?client_id=' + config.FACEBOOK_CLIENT_ID + '&redirect_uri=' + redirect_uri + '&scope=email'

      window.location.href = url
    }
  }
}
</script>
