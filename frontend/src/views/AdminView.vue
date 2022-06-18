<template lang="html">
  <div class="grid-container">
    <div class="red"><UsersGestionCardCompVue /></div>
    <div class="yellow"></div>
  </div>
</template>
<script>
import UsersGestionCardCompVue from "@/components/adminComponents/UsersGestionCardComp.vue";

export default {
  name: "AdminView",
  components: { UsersGestionCardCompVue },
};
</script>
<style lang="scss">
.grid-container {
  display: grid;
  height: 100%;
  width: 100%;
  grid-template-rows: 0.5fr repeat(2, 1fr 1fr 0.5fr);
  grid-template-columns: 0.5fr repeat(2, 1fr 1fr 0.5fr);
}

.red {
  grid-area: 2 / 2 / 6 / 4;
  background: rgb(235, 113, 113);
  display: flex;
  align-items: center;
  justify-content: center;
}

.yellow {
  grid-area: 2 / 5 / 6 / 7;
  background: rgb(235, 235, 113);
}
</style>
