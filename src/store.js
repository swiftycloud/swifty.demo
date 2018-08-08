import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)

export default new Vuex.Store({
  plugins: [createPersistedState()],

  state: {
    token: null,
    auth_endpoint: null,
    profile_endpoint: null,
    picture_endpoint: null
  },

  mutations: {
    updateToken (state, value) {
      state.token = value
    },

    updateAuthEndpoint (state, value) {
      state.auth_endpoint = value
    },

    updateProfileEndpoint (state, value) {
      state.profile_endpoint = value
    },

    updatePictureEndpoint (state, value) {
      state.picture_endpoint = value
    },

    clearToken (state) {
      state.token = null
    }
  },

  actions: {
    signUp ({ state }, { email, password }) {
      return axios.post(state.auth_endpoint + '/signup?userid=' + email + '&password=' + password)
    },

    signIn ({ state }, { email, password }) {
      return axios.post(state.auth_endpoint + '/signin?userid=' + email + '&password=' + password)
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
    }
  }
})
