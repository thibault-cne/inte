<template>
  <div class="infoBar">
    <div class="recapProfil">recapProfil</div>
    <div class="lotterie"><LotteryComp /></div>
    <div class="calendrier">calendrier</div>
  </div>
</template>

<script>
import LotteryComp from "./LotteryComp.vue";
export default {
  components: { LotteryComp },
};
</script>

<style scoped>
.infoBar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  height: 100%;
  width: 100%;
}

.recapProfil {
  height: 80%;
  width: 40%;
  background: white;
  border-radius: 4px;
  /* box-shadow: none */
}

.lotterie {
  height: 60%;
  width: 10%;
  background: white;
  border-radius: 4px;
  flex-basis: 10%;
  flex-shrink: 0;
  /* box-shadow: none */
}

.calendrier {
  height: 80%;
  width: 30%;
  background: white;
  border-radius: 4px;
  /* box-shadow: none */
}
</style>
