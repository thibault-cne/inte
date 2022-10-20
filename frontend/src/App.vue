<template>
  <v-app>
    <NavBar :status="status" />
    <v-main>
      <router-view :status="status" />
    </v-main>
    <FooterBar />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NavBar from "./components/NavBar.vue";
import FooterBar from "./components/FooterBar.vue";
import { LoggedIn } from "@/models/LoggedIn";
import { isLogged } from "@/requests/logged";

export default defineComponent({
  name: "App",
  created() {
    isLogged().then((r: LoggedIn) => {
      this.status = r;
      console.log("User status:", r);
    });
  },
  data() {
    return {
      status: {
        logged: false,
        user: {},
      } as LoggedIn,
    };
  },
  components: { NavBar, FooterBar },
});
</script>
