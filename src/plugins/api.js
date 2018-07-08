import axios from 'axios'

let api = {
  axios: axios,

  auth_func_url: null,
  profile_func_url: null,
  picture_func_url: null,

  signUp (email, password) {
    return this.axios.post(this.auth_func_url + '/signup?userid=' + email + '&password=' + password)
  },

  signIn (email, password) {
    return this.axios.post(this.auth_func_url + '/signin?userid=' + email + '&password=' + password)
  },

  createProfile (data) {
    return this.axios.post(this.profile_func_url, data)
  },

  getProfile () {
    return this.axios.get(this.profile_func_url)
  },

  updateProfile (data) {
    return this.axios.put(this.profile_func_url, data)
  },

  getPicture () {
    return this.axios.get(this.picture_func_url)
  },

  uploadPicture (data) {
    return this.axios.put(this.picture_func_url, data)
  },

  deletePicture () {
    return this.axios.delete(this.picture_func_url)
  }
}

if (localStorage.getItem('token')) {
  api.axios.defaults.headers.common['Authorization'] = 'Bearer: ' + localStorage.getItem('token')
}

if (localStorage.getItem('auth_func_url')) {
  api.auth_func_url = localStorage.getItem('auth_func_url')
}

if (localStorage.getItem('profile_func_url')) {
  api.profile_func_url = localStorage.getItem('profile_func_url')
}

if (localStorage.getItem('picture_func_url')) {
  api.picture_func_url = localStorage.getItem('picture_func_url')
}

export default api