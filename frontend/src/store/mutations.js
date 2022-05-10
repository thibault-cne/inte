const mutations = {
  updateStorage(state, { access, refresh }) {
    state.accessToken = access;
    state.refreshToken = refresh;
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
