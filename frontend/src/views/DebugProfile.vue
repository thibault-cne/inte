<template>
  <v-container>
    <h1 class="text-center">Debug le profile</h1>
    <v-container>
      <v-row>
        <v-col md="3" sm="12">
          <!--  -->
        </v-col>
        <v-col md="6" sm="12">
          <i-profile :status="status" height="550px"></i-profile>
        </v-col>
        <v-col md="3" sm="12">
          <!--  -->
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>
<script lang="ts">
import IProfile from "@/components/Profile.vue";
import { defineComponent } from "vue";
import { LoggedIn } from "@/models/LoggedIn";

export default defineComponent({
  name: "DebugProfile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { IProfile },
});
</script>
<style scoped lang="scss"></style>
