<template>
  <el-row>
    <el-col :lg="10">
      <el-form :xs="24" :sm="18" :md="14" :lg="10" label-width="170px" label-position="left" ref="signUpForm" @submit.native.prevent="submitForm" :model="form">
        <el-form-item label="User email">
          <el-input placeholder="Email" type="text" v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="City">
          <el-input placeholder="City" type="text" v-model="form.city"></el-input>
        </el-form-item>
        <el-form-item label="Auth URL">
          <el-input placeholder="Auth function URL" type="text" v-model="$store.state.auth_endpoint" disabled></el-input>
        </el-form-item>
        <el-form-item label="Facebook Auth URL">
          <el-input placeholder="Facebook function URL" type="text" v-model="$store.state.facebook_endpoint" disabled></el-input>
        </el-form-item>
        <el-form-item label="Profile Function URL">
          <el-input placeholder="Profile management URL" type="text" v-model="$store.state.profile_endpoint" disabled></el-input>
        </el-form-item>
        <el-form-item label="Picture Function URL">
          <el-input placeholder="Picture management URL" type="text" v-model="$store.state.picture_endpoint" disabled></el-input>
        </el-form-item>
        <el-form-item label="Tasks Function URL">
          <el-input placeholder="Tasks management URL" type="text" v-model="$store.state.tasks_endpoint" disabled></el-input>
        </el-form-item>
        <el-form-item label="Picture">
          <el-upload
            class="picture-uploader"
            action=""
            accept="image/*"
            :http-request="handlerPictureUpload"
            :show-file-list="false"
            :on-success="handlePictureSuccess"
            :before-upload="beforePictureUpload">
            <img v-if="user_image" :src="user_image" class="picture">
            <i v-else class="el-icon-plus picture-uploader-icon"></i>
          </el-upload>
        </el-form-item>

        <!-- <el-form-item label="Facebook">
          <el-button type="success" plain @click="connectFacebook" v-if="!form.facebook">Connect</el-button>
          <el-button type="danger" plain @click="disconnectFacebook" v-if="form.facebook">Disconnect</el-button>
        </el-form-item> -->

        <div class="sw-form-actions">
          <button type="submit" style="display: none"></button>
          <el-button type="primary" @click="submitForm">Update profile</el-button>
        </div>
      </el-form>
    </el-col>
  </el-row>
</template>

<script>
import * as config from '@/config'

export default {
  data () {
    return {
      form: {
        email: null,
        city: null,
        cookie: null,
        facebook: null,
      },

      user_image: null
    }
  },

  created () {
    this.fetchProfile()
    this.fetchPicture()

    if ('code' in this.$route.query) {
      this.$store.dispatch('connectFacebook', {
        code: this.$route.query.code,
        redirect_uri: encodeURIComponent(window.location.origin + window.location.pathname)
      }).then(() => {
        return this.fetchProfile()
      }).then(() => {
        this.$router.push({ name: 'profile' })
      })
    }
  },

  methods: {
    fetchProfile () {
      return this.$store.dispatch('getProfile').then(response => {
        this.form = response.data
      })
    },

    fetchPicture () {
      return this.$store.dispatch('getPicture').then(response => {
        this.user_image = 'data:image/jpeg;base64,' + response.data.img
      })
    },

    submitForm () {
      return this.$store.dispatch('updateProfile', {
        email: this.form.email,
        city: this.form.city
      }).then(() => {
        this.$notify.success({
          title: 'Success',
          message: 'Profile updated'
        })
      })
    },

    handlerPictureUpload(option) {
      var file = option.file
      this.$store.dispatch('deletePicture').finally(() => {
        var reader = new FileReader()
        reader.readAsDataURL(file)

        reader.onload = () => {
          this.$store.dispatch('uploadPicture', reader.result.replace(/^data:image\/(.+);base64,/, '')).then(() => {
            this.user_image = reader.result
          })
        }
      })
    },

    handlePictureSuccess (res, file) {
      this.user_image = URL.createObjectURL(file.raw);
    },

    beforePictureUpload (file) {
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isLt2M) {
        this.$message.error('Picture size can not exceed 2MB!');
      }

      return isLt2M;
    },

    connectFacebook () {
      const redirect_uri = encodeURIComponent(window.location.origin + window.location.pathname)
      const url = 'https://www.facebook.com/v3.0/dialog/oauth?client_id=' + config.FACEBOOK_CLIENT_ID + '&redirect_uri=' + redirect_uri

      window.location.href = url
    },

    disconnectFacebook () {
      this.$store.dispatch('disconnectFacebook', this.$route.query).then(() => {
        this.fetchProfile()
      })
    }
  }
}
</script>
