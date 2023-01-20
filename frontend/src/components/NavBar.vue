<template>
  <v-app-bar class="navbar" density="compact">
    <v-app-bar-nav-icon
      variant="text"
      @click.stop="drawer = !drawer"
    ></v-app-bar-nav-icon>

    <v-app-bar-title>Site de tah l'inté</v-app-bar-title>

    <template v-slot:append>
      <!-- <v-btn icon="mdi-dots-vertical"></v-btn> -->
      <v-btn v-if="!status.logged" :href="logginUrl">Login</v-btn>
      <v-btn v-else :href="disconnectUrl">Disconnect</v-btn>
    </template>
  </v-app-bar>
  <v-navigation-drawer v-model="drawer" bottom temporary color="#1e293b">
    <v-list>
      <v-list-item v-for="item in items" :key="item.title" :to="item.path">
        <v-list-item class="nav-item">{{ item.title }}</v-list-item>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { base_backend_url } from "@/requests/axios";

export default {
  name: "i-navbar",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data: () => ({
    logginUrl: base_backend_url + "/auth/login",
    disconnectUrl: base_backend_url + "/auth/logout",
    drawer: false,
    items: [
      {
        title: "Home",
        path: "/",
      },
      {
        title: "Debug Stars",
        path: "/debug-stars",
      },
      {
        title: "Debug Profile",
        path: "/debug-profile",
      },
      {
        title: "Debug Instagram",
        path: "/debug-instagram",
      },
    ],
  }),
  methods: {},
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
