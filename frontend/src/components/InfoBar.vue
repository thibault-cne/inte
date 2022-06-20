<template>
  <div class="infoBar">
    <div class="lotterie"><LotteryComp /></div>
    <div class="recapProfil"><RecapPointsStars /></div>
    <div class="calendrier">calendrier</div>
  </div>
</template>

<script>
import LotteryComp from "./LotteryComp.vue";
import RecapPointsStars from "./RecapPointsStars.vue";
export default {
  components: { LotteryComp, RecapPointsStars },
};
</script>

<style scoped>
.infoBar {
  display: grid;
  grid-template-areas:
    "lotterie recapProfil calendrier"
    "lotterie infos calendrier";
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 2fr;
  height: 100%;
  width: 100%;
}

.recapProfil {
  grid-area: recapProfil;
  height: 100%;
  width: 100%;
  align-self: flex-start;
  /* box-shadow: none */
}

.lotterie {
  grid-area: lotterie;
  height: 100%;
  width: 100%;
  /* box-shadow: none */
}

.calendrier {
  grid-area: calendrier;
  height: 100%;
  width: 100%;
  /* box-shadow: none */
}
</style>
