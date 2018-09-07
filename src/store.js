import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import md5 from 'blueimp-md5'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)

export default new Vuex.Store({
  plugins: [createPersistedState()],

  state: {
    token: null,
    auth_endpoint: null,
    facebook_endpoint: null,
    profile_endpoint: null,
    picture_endpoint: null,
    tasks_endpoint: null
  },

  mutations: {
    updateToken (state, value) {
      state.token = value
    },

    updateAuthEndpoint (state, value) {
      state.auth_endpoint = value
    },

    updateFacebookEndpoint (state, value) {
      state.facebook_endpoint = value
    },

    updateProfileEndpoint (state, value) {
      state.profile_endpoint = value
    },

    updatePictureEndpoint (state, value) {
      state.picture_endpoint = value
    },

    updateTasksEndpoint (state, value) {
      state.tasks_endpoint = value
    },

    clearToken (state) {
      state.token = null
    }
  },

  actions: {
    signUp ({ state }, { email, password }) {
      return axios.post(state.auth_endpoint + '/signup?userid=' + email + '&password=' + md5(password))
    },

    signIn ({ state }, { email, password }) {
      return axios.post(state.auth_endpoint + '/signin?userid=' + email + '&password=' + md5(password))
    },

    createProfile ({ state }, payload) {
      return axios.post(state.profile_endpoint, payload, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    getProfile ({ state }) {
      return axios.get(state.profile_endpoint, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    updateProfile ({ state }, payload) {
      return axios.put(state.profile_endpoint, payload, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    connectFacebook ({ state }, { code }) {
      return axios.put(state.profile_endpoint, {
        facebook: 1,
        code: code
      }, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    disconnectFacebook ({ state }) {
      return axios.put(state.profile_endpoint, {
        facebook: 0
      }, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    signUpWithFacebook ({ state }, { code, redirect_uri }) {
      return axios.post(state.facebook_endpoint + '/signin', {
        code: code,
        redirect_uri: redirect_uri
      })
    },

    signInWithFacebook ({ state }, { code, redirect_uri }) {
      return axios.post(state.facebook_endpoint + '/signin', {
        code: code,
        redirect_uri: redirect_uri
      })
    },

    getPicture ({ state }) {
      return axios.get(state.picture_endpoint, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    uploadPicture ({ state }, payload) {
      return axios.post(state.picture_endpoint, payload, {
        headers: {
          'Authorization': 'Bearer ' + state.token,
          'Content-Type': 'application/json'
        }
      })
    },

    deletePicture ({ state }) {
      return axios.delete(state.picture_endpoint, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    getTasks ({ state }, filter) {
      let url = state.tasks_endpoint + '/tasks'
      url = filter !== 'all' ? url + '?status=' + filter : url

      return axios.get(url, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    addTask ({ state }, task) {
      return axios.post(state.tasks_endpoint + '/tasks', task, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    deleteTask ({ state }, task_id) {
      return axios.delete(state.tasks_endpoint + '/tasks/' + task_id, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    },

    doneTask ({ state }, task_id) {
      return axios.post(state.tasks_endpoint + '/tasks/' + task_id + '/done', {}, {
        headers: { 'Authorization': 'Bearer ' + state.token }
      })
    }
  }
})
