<template>
  <div class="lotteryContainer">
    <router-link to="/lottery" class="lotteryCard">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="-15 -15 670 542"
        class="logo"
      >
        <path
          d="M447.1 224c0-12.56-4.781-25.13-14.35-34.76l-174.9-174.9C249.1 4.786 236.5 0 223.1 0C211.4 0 198.9 4.786 189.2 14.35L14.35 189.2C4.783 198.9-.0011 211.4-.0011 223.1c0 12.56 4.785 25.17 14.35 34.8l174.9 174.9c9.625 9.562 22.19 14.35 34.75 14.35s25.13-4.783 34.75-14.35l174.9-174.9C443.2 249.1 447.1 236.6 447.1 224zM96 248c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1S120 210.8 120 224S109.3 248 96 248zM224 376c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1s23.1 10.75 23.1 23.1S237.3 376 224 376zM224 248c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1S248 210.8 248 224S237.3 248 224 248zM224 120c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1s23.1 10.75 23.1 23.1S237.3 120 224 120zM352 248c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1s23.1 10.75 23.1 23.1S365.3 248 352 248zM591.1 192l-118.7 0c4.418 10.27 6.604 21.25 6.604 32.23c0 20.7-7.865 41.38-23.63 57.14l-136.2 136.2v46.37C320 490.5 341.5 512 368 512h223.1c26.5 0 47.1-21.5 47.1-47.1V240C639.1 213.5 618.5 192 591.1 192zM479.1 376c-13.25 0-23.1-10.75-23.1-23.1s10.75-23.1 23.1-23.1s23.1 10.75 23.1 23.1S493.2 376 479.1 376z"
        />
      </svg>
      <div class="etat">
        <div v-if="lotteryDone" class="compteur">{{ this.toDemainChrono }}</div>
        <div v-else class="pret">Tente ta chance !</div>
      </div>
    </router-link>
  </div>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  data() {
    return {
      lotteryDone: true,
      toDemainChrono: "",
      chrono: setInterval(() => {
        let now = new Date().getTime();
        let jour = 86400000;
        let aujourdhui = Math.floor(now / jour) * jour;
        let toDemain = jour - (now - aujourdhui) - 7200000;
        let toDemainH = Math.floor(toDemain / 3600000);
        let toDemainM = Math.floor((toDemain % 3600000) / 60000);
        let toDemainS = Math.floor(((toDemain % 3600000) % 60000) / 1000);
        this.toDemainChrono = `${toDemainH}h ${toDemainM}min ${toDemainS}s`;
      }, 1000),
    };
  },
  async created() {
    await getRequest("/daily-game/check", "data").then((res) => {
      this.lotteryDone = res.data.status;
    });
  },
};
</script>

<style scoped>
.lotteryContainer {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.lotteryCard {
  /* j'ai dû bidouiller la width (40%) pour que ça déborde pas de la case,
  sinon à cause d'un espèce d'aspect ratio bizarre ça casse les proportions de la grid */
  width: 40%;
  height: 80%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-decoration: none;
  border-radius: 10px;
  background: lightgrey;
}

.logo {
  width: 100%;
  height: 85%;
  /* background: lightskyblue; */
  /* fill: rgb(243, 166, 243); */
  fill: rgb(255, 255, 255);
  stroke: black;
  stroke-width: 10px;
  transition: 200ms ease;
}

.logo:hover {
  fill: rgb(171, 115, 255);
  stroke: rgb(102, 0, 255);
}

.etat {
  width: 100%;
  height: 15%;
  /* background: rgb(139, 0, 42); */
  text-align: center;
  color: black;
}

.pret {
  width: 100%;
  height: 100%;
  /* background: lightgreen; */
  font-size: 1.4vw;
}

.compteur {
  width: 100%;
  height: 100%;
  /* background: lightcoral; */
  font-size: 1.6vw;
}
</style>
