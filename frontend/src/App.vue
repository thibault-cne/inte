<template>
  <v-app>
    <NavBar />
    <v-main>
      <router-view />
    </v-main>
    <FooterBar />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NavBar from "./components/NavBar.vue";
import FooterBar from "./components/FooterBar.vue";

export default defineComponent({
  name: "App",
  data() {
    return {
      //
    };
  },
  components: { NavBar, FooterBar },
});
</script>
