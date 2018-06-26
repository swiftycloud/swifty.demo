<template>
  <el-form label-width="120px" label-position="left" ref="signUpForm" :model="form">
    <el-form-item label="User email">
      <el-input placeholder="Email" type="text" v-model="form.email"></el-input>
    </el-form-item>
    <el-form-item label="City">
      <el-input placeholder="City" type="text" v-model="form.city"></el-input>
    </el-form-item>
    <el-form-item label="Auth URL">
      <el-input placeholder="Auth function URL" type="text" v-model="form.auth_url" disabled></el-input>
    </el-form-item>
    <el-form-item label="Function URL">
      <el-input placeholder="Profile management URL" type="text" v-model="form.func_url" disabled></el-input>
    </el-form-item>
    <el-form-item label="Picture">
      <el-upload
        class="avatar-uploader"
        action="https://jsonplaceholder.typicode.com/posts/"
        :show-file-list="false"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload">
        <img v-if="form.user_image" :src="form.user_image" class="avatar">
        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
      </el-upload>
    </el-form-item>
    
    <div class="sw-form-actions">
      <button type="submit" style="display: none"></button>
      <el-button type="primary">Update profile</el-button>
    </div>
  </el-form>
</template>

<script>
export default {
  data () {
    return {
      form: {
        email: null,
        city: null,
        auth_url: null,
        func_url: null,
        user_image: null
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
    handleAvatarSuccess(res, file) {
      this.form.user_image = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('Avatar picture must be JPG format!');
      }
      if (!isLt2M) {
        this.$message.error('Avatar picture size can not exceed 2MB!');
      }
      return isJPG && isLt2M;
    }
  }
}
</script>

<style>
.header {
  font-size: 16px;
  font-weight: normal;
  text-transform: uppercase;
  color: #1989fa;
  border-bottom: 1px solid #e4e7ed;
  padding-bottom: 15px;
  margin-top: 30px;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 136px;
  height: 136px;
  line-height: 136px;
  text-align: center;
}
.avatar {
  width: 136px;
  height: 136px;
  display: block;
}
</style>