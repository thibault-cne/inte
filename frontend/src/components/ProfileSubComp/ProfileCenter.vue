<template>
  <v-container>
    <span>
      <h4>{{ status.user.name }}</h4>
      <h5>Year : {{ status.user.current_year }}A</h5>
      <h5>Promotion {{ status.user.promotion_year }}</h5>
    </span>
    <v-container class="container">
      <i-profileInfo
        :text="status.user.hometown"
        icon="mdi-map-marker"
        tooltip="Location"
      ></i-profileInfo>
      <i-profileInfo
        :text="status.user.studies"
        icon="mdi-school"
        tooltip="Studies"
      ></i-profileInfo>
    </v-container>
    <v-divider></v-divider>
    <v-container class="pa-6">
      <v-row justify="space-around">
        <i-profileSocials
          v-if="status.user.facebook_id"
          icon="mdi-facebook"
          color="blue"
        ></i-profileSocials>
        <i-profileSocials
          v-if="status.user.snapchat_id"
          icon="mdi-snapchat"
          color="yellow"
        ></i-profileSocials>
        <i-profileSocials
          v-if="status.user.google_id"
          icon="mdi-google"
          color="red"
        ></i-profileSocials>
        <i-profileSocials
          v-if="status.user.instagram_id"
          icon="mdi-instagram"
          color="pink"
        ></i-profileSocials>
      </v-row>
    </v-container>
  </v-container>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { LoggedIn } from "@/models/LoggedIn";
import IProfileInfo from "./ProfileInfo.vue";
import IProfileSocials from "./ProfileSocials.vue";

export default defineComponent({
  name: "i-profileCenter",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { IProfileInfo, IProfileSocials },
});
</script>
<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  margin-top: 2vh;
}
</style>
