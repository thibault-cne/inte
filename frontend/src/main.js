import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import gAuthPlugin from "vue3-google-oauth2";

createApp(App)
  .use(store)
  .use(router)
  .use(gAuthPlugin, {
    clientId:
      "938768924694-4b483tmfbeogob3pmn4s5u5ujngqm43b.apps.googleusercontent.com",
    scope: "email profile",
    prompt: "consent",
  })
  .mount("#app");
