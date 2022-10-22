<template>
  <v-chip variant="text">
    <template v-slot:prepend>
      <v-icon :icon="icon" />
    </template>
    <v-tooltip activator="parent" location="start">{{ tooltip }}</v-tooltip>
    {{ text }}
  </v-chip>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-profileInfo",
  props: {
    text: { type: String, required: true },
    icon: { type: String, required: true },
    tooltip: { type: String, required: true },
  },
  data() {
    return {};
  },
});
</script>
<style scoped lang="scss"></style>
