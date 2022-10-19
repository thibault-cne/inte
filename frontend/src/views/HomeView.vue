<template>
  <hello-world />
</template>

<script lang="ts">
import { defineComponent } from "vue";

// Components
import HelloWorld from "../components/HelloWorld.vue";
import { LoggedIn } from "@/models/LoggedIn";
import { isLogged } from "@/requests/logged";

export default defineComponent({
  name: "HomeView",
  created() {
    isLogged().then((r: LoggedIn) => {
      console.log(r);
    });
  },
  components: {
    HelloWorld,
  },
});
</script>
