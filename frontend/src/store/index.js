import { createStore } from "vuex";
import state from "./state";
import mutations from "./mutations";
import actions from "./actions";
import getters from "./getters";
import modules from "./modules";

export default createStore({
  state,
  getters,
  mutations,
  actions,
  modules,
});
