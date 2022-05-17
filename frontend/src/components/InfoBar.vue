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
  margin-left: 3%;
  margin-right: 5%;
  /* box-shadow: none */
}

.lotterie {
  background: white;
  border-radius: 4px;
  flex-basis: 12%;
  /* box-shadow: none */
}

.calendrier {
  height: 80%;
  width: 30%;
  background: white;
  border-radius: 4px;
  margin-left: 5%;
  margin-right: 3%;
  /* box-shadow: none */
}
</style>
