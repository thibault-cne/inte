<template>
  <div class="header">
    <div class="gauche">
      <div class="nomSite">nomSite</div>
    </div>
    <div class="droite">
      <div class="rechercheUser">rechercheUser</div>
      <div class="notifs">notifs</div>
      <div class="accesProfil">accesProfil</div>
    </div>
  </div>
</template>

<script></script>

<style scoped>
.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  width: 100%;
  background: rgb(130, 193, 244);
}

.gauche {
  position: absolute;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  height: 100%;
  width: 100%;
  z-index: 1;
}

.droite {
  position: absolute;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
  height: 100%;
  width: 100%;
}

.nomSite {
  margin-left: 0.5vw;
  height: 80%;
  width: 30%;
  background: white;
  border-radius: 4px;
}

.rechercheUser {
  margin-right: 1vw;
  height: 50%;
  width: 20%;
  background: white;
  border-radius: 4px;
}

.notifs {
  margin-right: 1vw;
  height: 40%;
  width: 3%;
  background: white;
  border-radius: 4px;
}

.accesProfil {
  margin-right: 1vw;
  height: 70%;
  width: 15%;
  background: white;
  border-radius: 4px;
}
</style>
