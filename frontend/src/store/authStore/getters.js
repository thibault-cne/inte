const getters = {
  loggedIn: (state) => state.accessToken != null,
  accessToken: (state) => state.accessToken,
  refreshToken: (state) => state.refreshToken,
};

export { getters };
