<template>
  <v-container>
    <h1 class="text-center">Debug les étoiles globales</h1>
    <v-container>
      <v-row>
        <v-col md="2" lg="3" sm="12">
          <!--  -->
        </v-col>
        <v-col md="8" lg="6" sm="12">
          <i-stars :status="status" height="550px"></i-stars>
        </v-col>
        <v-col md="2" lg="3" sm="12">
          <!--  -->
        </v-col>
      </v-row>
    </v-container>
    <h1 class="text-center">Les étoiles du profil</h1>
    <v-container>
      <v-row>
        <v-col md="2" lg="3" sm="12">
          <!--  -->
        </v-col>
        <v-col md="8" lg="6" sm="12">
          <i-stars :status="status" :preload="true" height="550px"></i-stars>
        </v-col>
        <v-col md="2" lg="3" sm="12">
          <!--  -->
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import IStars from "@/components/Stars.vue";
import { LoggedIn } from "@/models/LoggedIn";

export default defineComponent({
  name: "DebugStars",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { IStars },
});
</script>
