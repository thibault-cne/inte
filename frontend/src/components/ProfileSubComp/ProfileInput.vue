<template>
  <v-row>
    <v-container class="container">
      <v-text-field class="input" prepend-icon="mdi-vuetify">
        <template v-slot:append>
          <v-icon icon="mdi-chevron-right-circle-outline" />
        </template>
      </v-text-field>
    </v-container>
  </v-row>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-profileInput",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {};
  },
});
</script>
<style scoped lang="scss"></style>
