<template>
  <div class="grid-container">
    <div class="info"><InfoBar /></div>
    <div class="etoiles">etoiles</div>
    <div class="feed">feed</div>
  </div>
</template>

<script>
import InfoBar from "../components/InfoBar.vue";

export default {
  name: "HomeView",
  components: { InfoBar },
};
</script>

<style scoped>
.grid-container {
  display: grid;
  height: 100%;
  width: 100%;
  grid-template-rows: 2fr 3fr;
  grid-template-columns: 3fr 2fr;
}

.info {
  grid-area: 1 / 1 / 2 / -1;
  background: rgb(237, 200, 106);
}

.etoiles {
  grid-area: 2 / 1 / -1 / 2;
  background: rgb(237, 106, 106);
}

.feed {
  grid-area: -2 / -2 / -1 / -1;
  background: rgb(134, 106, 237);
}
</style>
