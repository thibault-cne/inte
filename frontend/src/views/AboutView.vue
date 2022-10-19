<template>
  <div class="about">
    <h1>This is an about page</h1>
    <StarsView :status="status" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { LoggedIn } from "@/models/LoggedIn";
import StarsView from "@/components/StarsView.vue";

export default defineComponent({
  name: "AboutView",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  components: { StarsView },
});
</script>
