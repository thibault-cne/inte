<template>
  <div>Hello</div>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { getRequest } from "@/requests/getRequests";
import { Stars } from "@/models/Stars";
import { defineComponent } from "vue";

export default defineComponent({
  name: "StarsView",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {
      stars: {} as Stars,
    };
  },
  mounted() {
    console.log(this.status);
    if (this.status.logged) {
      getRequest("/stars/", "").then((r) => {
        console.log(r);
        //this.stars = <Stars>r.data;
      });
      console.log(this.stars);
    }
    console.log(this.stars);
  },
});
</script>
<style scoped lang="scss"></style>
