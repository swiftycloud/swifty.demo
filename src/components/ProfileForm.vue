<template>
  <el-form label-width="170px" label-position="left" ref="signUpForm" @submit.native.prevent="submitForm" :model="form">
    <el-form-item label="User email">
      <el-input placeholder="Email" type="text" v-model="form.email"></el-input>
    </el-form-item>
    <el-form-item label="City">
      <el-input placeholder="City" type="text" v-model="form.city"></el-input>
    </el-form-item>
    <el-form-item label="Auth URL">
      <el-input placeholder="Auth function URL" type="text" v-model="$store.state.auth_endpoint" disabled></el-input>
    </el-form-item>
    <el-form-item label="Profile Function URL">
      <el-input placeholder="Profile management URL" type="text" v-model="$store.state.profile_endpoint" disabled></el-input>
    </el-form-item>
    <el-form-item label="Picture Function URL">
      <el-input placeholder="Picture management URL" type="text" v-model="$store.state.picture_endpoint" disabled></el-input>
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
    
    <div class="sw-form-actions">
      <button type="submit" style="display: none"></button>
      <el-button type="primary" @click="submitForm">Update profile</el-button>
    </div>
  </el-form>
</template>

<script>
export default {
  data () {
    return {
      form: {
        email: null,
        city: null
      },

      user_image: null
    }
  },

  created () {
    this.fetchProfile()
    this.fetchPicture()
  },

  methods: {
    fetchProfile () {
      return this.$store.dispatch('getProfile').then(response => {
        this.form.email = response.data.email
        this.form.city = response.data.city
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
      }).then(response => {
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
          this.$store.dispatch('uploadPicture', reader.result.replace(/^data:image\/(.+);base64,/, '')).then(response => {
            this.user_image = reader.result
          })
        }

        reader.onerror = function (error) {
          console.log('Error: ', error);
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
    }
  }
}
</script>