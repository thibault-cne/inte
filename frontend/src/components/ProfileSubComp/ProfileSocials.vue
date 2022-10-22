<template>
  <v-btn :icon="icon" :color="color"> </v-btn>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-profileSocials",
  props: {
    color: { type: String, required: true },
    icon: { type: String, required: true },
  },
});
</script>
<style scoped lang="scss"></style>
