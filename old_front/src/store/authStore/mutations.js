const mutations = {
  updateStorage(state, credentials) {
    state.accessToken = credentials.access_token;
    state.refreshToken = credentials.refresh_token;
  },
  destroyToken(state) {
    state.accessToken = null;
    state.refreshToken = null;
  },
  setAccessToken(state, token) {
    state.accessToken = token;
  },
};

export { mutations };
