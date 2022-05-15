<template>
  <div class="infoBar">
    <div class="recapProfil">recapProfil</div>
    <div class="lotterie">lotterie</div>
    <div class="calendrier">calendrier</div>
  </div>
</template>

<script></script>

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
