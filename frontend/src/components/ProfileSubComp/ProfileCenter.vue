<template>
  <v-container>
    <span>
      <h4>{{ status.user.name }}</h4>
      <h5>Year : {{ status.user.current_year }}A</h5>
      <h5>Promotion {{ status.user.promotion_year }}</h5>
    </span>
    <span>
      <i-profile-input class="mt-2" :status="status"></i-profile-input>
    </span>
  </v-container>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { LoggedIn } from "@/models/LoggedIn";
import IProfileInput from "./ProfileInput.vue";

export default defineComponent({
  name: "i-profileCenter",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { IProfileInput },
});
</script>
<style scoped lang="scss">
.mt-2 {
  margin-top: 2vh;
}
</style>
