<template>
  <v-app>
    <NavBar :status="status" />
    <v-main>
      <router-view :status="status" @edit="(u: User) => edit(u)" />
    </v-main>
    <FooterBar />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NavBar from "./components/navBar.vue";
import FooterBar from "./components/footerBar.vue";
import { LoggedIn } from "@/models/LoggedIn";
import { isLogged } from "@/requests/logged";
import { User } from "@/models/User";

export default defineComponent({
  name: "App",
  created() {
    isLogged().then((r: LoggedIn) => {
      this.status = r;
      console.debug("User status:", r);
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
  methods: {
    edit(u: User) {
      this.status.user = u;
    },
  },
  components: { NavBar, FooterBar },
});
</script>
