import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { authStore } from "./store/authStore";

createApp(App).use(authStore).use(router).mount("#app");
