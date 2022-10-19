<template>
  <v-app-bar class="navbar" density="compact" prominent>
    <v-app-bar-nav-icon
      variant="text"
      @click.stop="drawer = !drawer"
    ></v-app-bar-nav-icon>

    <v-app-bar-title>Site de tah l'inté</v-app-bar-title>

    <template v-slot:append>
      <!-- <v-btn icon="mdi-dots-vertical"></v-btn> -->
      <v-btn :href="logginUrl">Login</v-btn>
      <v-btn @click="checkStatus()">Test login</v-btn>
    </template>
  </v-app-bar>
  <v-navigation-drawer v-model="drawer" bottom temporary color="#424549">
    <v-list>
      <v-list-item v-for="item in items" :key="item.title" :to="item.path">
        <v-list-item class="nav-item">{{ item.title }}</v-list-item>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { isLogged } from "@/requests/logged";
import { base_backend_url } from "@/requests/axios";

export default {
  data: () => ({
    logginUrl: base_backend_url + "/auth/login",
    drawer: false,
    group: null,
    items: [
      {
        title: "Home",
        path: "/",
      },
      {
        title: "Debug Stars",
        path: "/debug-stars",
      },
    ],
  }),
  methods: {
    checkStatus() {
      isLogged().then((r: LoggedIn) => {
        console.log(r);
      });
    },
  },
};
</script>
<style lang="scss">
@use "@/assets/styles/scss/standars/colors";

.nav-item {
  color: colors.$text-primary !important;
  cursor: pointer;
}

.navbar {
  background-color: colors.$bg-footer !important;
  color: colors.$footer !important;
}
</style>
