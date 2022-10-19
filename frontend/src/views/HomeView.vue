<template>
  <div>
    <hello-world />
    <v-btn :href="logginUrl">Login</v-btn>
    <v-btn @click="checkStatus()">Test login</v-btn>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

// Components
import HelloWorld from "../components/HelloWorld.vue";
import { LoggedIn } from "@/models/LoggedIn";
import { isLogged } from "@/requests/logged";
import { base_backend_url } from "@/requests/axios";

export default defineComponent({
  name: "HomeView",
  data() {
    return {
      logginUrl: base_backend_url + "/auth/login",
    };
  },
  methods: {
    checkStatus() {
      isLogged().then((r: LoggedIn) => {
        console.log(r);
      });
    },
  },
  components: {
    HelloWorld,
  },
});
</script>
