<template>
  <v-card>
    <v-card-title>Profil de {{ status.user.name }}</v-card-title>
    <v-row align="center">
      <v-col md="3"><v-card>1</v-card></v-col>
      <v-col md="5"><i-profileCenter :status="status"></i-profileCenter></v-col>
      <v-col md="4"><v-card>3</v-card></v-col>
    </v-row>
  </v-card>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { defineComponent } from "vue";
import IProfileCenter from "./ProfileSubComp/ProfileCenter.vue";

export default defineComponent({
  name: "i-profile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { IProfileCenter },
});
</script>
<style lang=""></style>
