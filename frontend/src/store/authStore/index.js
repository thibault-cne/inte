import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";
import { actions } from "./actions";
import { mutations } from "./mutations";
import { getters } from "@/store/authStore/getters";

const authStore = createStore({
  state: {
    accessToken: null,
    refreshToken: null,
  },
  getters,
  mutations,
  actions,
  plugins: [createPersistedState()],
});

export { authStore };
