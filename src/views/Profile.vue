<template>
  <div class="profile">
    <el-row>
      <el-col :span="7">
        <h2 class="header">Profile Management</h2>

        <profile-form></profile-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ProfileForm from '@/components/ProfileForm.vue'
export default {
  components: { ProfileForm },
}
</script>