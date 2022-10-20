<template>
  <v-card>
    <v-card-title
      >Profil de {{ status.user.personal_description }}</v-card-title
    >

    <i-typeText :message="status.user.personal_description"></i-typeText>
  </v-card>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { defineComponent } from "vue";
import ITypeText from "./TypeText.vue";

export default defineComponent({
  name: "i-profile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { ITypeText },
});
</script>
<style lang=""></style>
