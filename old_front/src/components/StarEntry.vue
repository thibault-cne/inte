<template>
  <div class="starCard">
    <div class="starLogo">
      <svg v-if="rank == 0" viewBox="0 0 512 512">
        <defs>
          <radialGradient
            id="B0"
            gradientTransform="translate(-0.94 -0.83) scale(1.88, 1.67)"
          >
            <stop offset="0%" stop-color="#cb9855" />
            <stop offset="6.42%" stop-color="#c3904d" />
            <stop offset="12.83%" stop-color="#bc8845" />
            <stop offset="19.25%" stop-color="#b4803e" />
            <stop offset="25.67%" stop-color="#ac7836" />
            <stop offset="32.08%" stop-color="#a4702f" />
            <stop offset="38.5%" stop-color="#9d6828" />
            <stop offset="44.92%" stop-color="#956121" />
            <stop offset="51.33%" stop-color="#8d591a" />
            <stop offset="57.75%" stop-color="#855214" />
            <stop offset="64.17%" stop-color="#7c4b0c" />
            <stop offset="77%" stop-color="#6c3d00" />
            <stop offset="78.92%" stop-color="#6b3d00" />
            <stop offset="80.83%" stop-color="#6b3d00" />
            <stop offset="82.75%" stop-color="#6a3c00" />
            <stop offset="84.67%" stop-color="#6a3c00" />
            <stop offset="86.58%" stop-color="#693c00" />
            <stop offset="88.5%" stop-color="#693c00" />
            <stop offset="90.42%" stop-color="#683b00" />
            <stop offset="92.33%" stop-color="#683b00" />
            <stop offset="94.25%" stop-color="#673b00" />
            <stop offset="96.17%" stop-color="#673b00" />
            <stop offset="100%" stop-color="#663a00" />
          </radialGradient>
          <radialGradient
            id="B1"
            gradientTransform="translate(0.12 0.47) scale(1.77, 1.07)"
          >
            <stop offset="0%" stop-color="rgba(202, 151, 83, 0.64)" />
            <stop offset="8.33%" stop-color="rgba(201, 147, 90, 0.59)" />
            <stop offset="16.67%" stop-color="rgba(199, 143, 96, 0.53)" />
            <stop offset="25%" stop-color="rgba(197, 140, 102, 0.48)" />
            <stop offset="33.33%" stop-color="rgba(195, 136, 106, 0.43)" />
            <stop offset="41.67%" stop-color="rgba(192, 133, 111, 0.37)" />
            <stop offset="50%" stop-color="rgba(188, 130, 114, 0.32)" />
            <stop offset="58.33%" stop-color="rgba(184, 127, 117, 0.27)" />
            <stop offset="66.67%" stop-color="rgba(179, 125, 120, 0.21)" />
            <stop offset="75%" stop-color="rgba(168, 125, 125, 0.16)" />
            <stop offset="83.33%" stop-color="rgba(156, 126, 128, 0.11)" />
            <stop offset="100%" stop-color="rgba(128, 128, 128, 0)" />
          </radialGradient>
        </defs>
        <polygon
          fill="url(#B0)"
          stroke="url(#B0)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
        <polygon
          fill="url(#B1)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
      </svg>
      <svg v-if="rank == 1" viewBox="0 0 512 512">
        <defs>
          <radialGradient
            id="S0"
            gradientTransform="translate(-1 -1) scale(2, 2)"
          >
            <stop offset="0%" stop-color="#ffffff" />
            <stop offset="5.58%" stop-color="#f7f7f7" />
            <stop offset="11.17%" stop-color="#efefef" />
            <stop offset="16.75%" stop-color="#e7e7e7" />
            <stop offset="22.33%" stop-color="#dfdfdf" />
            <stop offset="27.92%" stop-color="#d8d8d8" />
            <stop offset="33.5%" stop-color="#d0d0d0" />
            <stop offset="39.08%" stop-color="#c8c8c8" />
            <stop offset="44.67%" stop-color="#c1c1c1" />
            <stop offset="50.25%" stop-color="#b9b9b9" />
            <stop offset="55.83%" stop-color="#b2b2b2" />
            <stop offset="67%" stop-color="#a3a3a3" />
            <stop offset="69.75%" stop-color="#a2a2a2" />
            <stop offset="72.5%" stop-color="#a0a0a0" />
            <stop offset="75.25%" stop-color="#9f9f9f" />
            <stop offset="78%" stop-color="#9e9e9e" />
            <stop offset="80.75%" stop-color="#9d9d9d" />
            <stop offset="83.5%" stop-color="#9b9b9b" />
            <stop offset="86.25%" stop-color="#9a9a9a" />
            <stop offset="89%" stop-color="#999999" />
            <stop offset="91.75%" stop-color="#989898" />
            <stop offset="94.5%" stop-color="#969696" />
            <stop offset="100%" stop-color="#949494" />
          </radialGradient>
          <radialGradient
            id="S1"
            gradientTransform="translate(0.38 0.52) scale(1.25, 0.96)"
          >
            <stop offset="0%" stop-color="rgba(255, 255, 255, 0.53)" />
            <stop offset="8.33%" stop-color="rgba(244, 244, 244, 0.49)" />
            <stop offset="16.67%" stop-color="rgba(233, 233, 233, 0.44)" />
            <stop offset="25%" stop-color="rgba(222, 222, 222, 0.4)" />
            <stop offset="33.33%" stop-color="rgba(211, 211, 211, 0.35)" />
            <stop offset="41.67%" stop-color="rgba(200, 200, 200, 0.31)" />
            <stop offset="50%" stop-color="rgba(189, 189, 189, 0.27)" />
            <stop offset="58.33%" stop-color="rgba(179, 179, 179, 0.22)" />
            <stop offset="66.67%" stop-color="rgba(169, 169, 169, 0.18)" />
            <stop offset="75%" stop-color="rgba(158, 158, 158, 0.13)" />
            <stop offset="83.33%" stop-color="rgba(148, 148, 148, 0.09)" />
            <stop offset="100%" stop-color="rgba(128, 128, 128, 0)" />
          </radialGradient>
        </defs>
        <polygon
          fill="url(#S0)"
          stroke="url(#S0)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
        <polygon
          fill="url(#S1)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
      </svg>
      <svg v-if="rank == 2" viewBox="0 0 512 512">
        <defs>
          <radialGradient
            id="G0"
            gradientTransform="translate(-0.66 -0.61) scale(1.32, 1.22)"
          >
            <stop offset="0%" stop-color="#ffffff" />
            <stop offset="4.17%" stop-color="#fdfcfc" />
            <stop offset="8.33%" stop-color="#fbfafa" />
            <stop offset="12.5%" stop-color="#f9f7f7" />
            <stop offset="16.67%" stop-color="#f8f4f3" />
            <stop offset="20.83%" stop-color="#f8f1ef" />
            <stop offset="25%" stop-color="#f7eeea" />
            <stop offset="29.17%" stop-color="#f8ebe3" />
            <stop offset="33.33%" stop-color="#f9e7db" />
            <stop offset="37.5%" stop-color="#fae5cf" />
            <stop offset="41.67%" stop-color="#fbe2bd" />
            <stop offset="50%" stop-color="#ffdf53" />
            <stop offset="54.17%" stop-color="#ffdc50" />
            <stop offset="58.33%" stop-color="#ffd94e" />
            <stop offset="62.5%" stop-color="#ffd74b" />
            <stop offset="66.67%" stop-color="#ffd448" />
            <stop offset="70.83%" stop-color="#ffd145" />
            <stop offset="75%" stop-color="#ffce42" />
            <stop offset="79.17%" stop-color="#ffcb3f" />
            <stop offset="83.33%" stop-color="#ffc93b" />
            <stop offset="87.5%" stop-color="#ffc638" />
            <stop offset="91.67%" stop-color="#ffc335" />
            <stop offset="100%" stop-color="#ffbd2e" />
          </radialGradient>
          <radialGradient
            id="G1"
            gradientTransform="translate(0.27 0.38) scale(1.47, 1.25)"
          >
            <stop offset="0%" stop-color="#ffffff" />
            <stop offset="3.13%" stop-color="rgba(254, 254, 254, 0.93)" />
            <stop offset="6.25%" stop-color="rgba(254, 254, 254, 0.86)" />
            <stop offset="9.38%" stop-color="rgba(253, 253, 253, 0.79)" />
            <stop offset="12.5%" stop-color="rgba(253, 252, 252, 0.72)" />
            <stop offset="15.63%" stop-color="rgba(253, 251, 251, 0.65)" />
            <stop offset="18.75%" stop-color="rgba(253, 250, 249, 0.59)" />
            <stop offset="21.88%" stop-color="rgba(253, 250, 247, 0.52)" />
            <stop offset="25%" stop-color="rgba(253, 249, 244, 0.45)" />
            <stop offset="28.13%" stop-color="rgba(253, 248, 240, 0.38)" />
            <stop offset="31.25%" stop-color="rgba(254, 248, 232, 0.31)" />
            <stop offset="37.5%" stop-color="rgba(255, 255, 0, 0.17)" />
            <stop offset="40.63%" stop-color="rgba(254, 239, 119, 0.16)" />
            <stop offset="43.75%" stop-color="rgba(251, 225, 128, 0.14)" />
            <stop offset="46.88%" stop-color="rgba(247, 212, 131, 0.13)" />
            <stop offset="50%" stop-color="rgba(241, 199, 133, 0.11)" />
            <stop offset="53.13%" stop-color="rgba(235, 187, 135, 0.1)" />
            <stop offset="56.25%" stop-color="rgba(226, 176, 137, 0.09)" />
            <stop offset="59.38%" stop-color="rgba(216, 165, 138, 0.07)" />
            <stop offset="62.5%" stop-color="rgba(204, 155, 139, 0.06)" />
            <stop offset="65.63%" stop-color="rgba(190, 146, 139, 0.04)" />
            <stop offset="68.75%" stop-color="rgba(173, 139, 138, 0.03)" />
            <stop offset="75%" stop-color="rgba(128, 128, 128, 0)" />
          </radialGradient>
        </defs>
        <polygon
          fill="url(#G0)"
          stroke="url(#G0)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
        <polygon
          fill="url(#G1)"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
      </svg>
    </div>
    <div class="starText">
      <b>{{ receiver }}</b> <br />
      a reçu une étoile {{ rankText }} il y a {{ ilyaText }} :<br />
      {{ message }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    message: String,
    receiver: Number,
    rank: Number,
    date: Number,
  },
  data() {
    return {
      ilyaText: "tkt",
      rankText: "ça bug",
    };
  },
  created() {
    let ilya = Date.now() / 1000;
    ilya = ilya - this.date;
    ilya = Math.floor(ilya / 3600);
    if (ilya === 0) {
      this.ilyaText = "moins d'1h";
    } else if (ilya <= 23) {
      this.ilyaText = `${ilya}h`;
    } else if (ilya <= 167) {
      ilya = Math.floor(ilya / 24);
      this.ilyaText = `${ilya}j`;
    } else if (ilya <= 719) {
      ilya = Math.floor(ilya / 24 / 7);
      this.ilyaText = ilya === 1 ? `${ilya} semaine` : `${ilya} semaines`;
    } else {
      ilya = Math.floor(ilya / 24 / 7 / 4);
      this.ilyaText = `${ilya} mois`;
    }
    if (this.rank === 0) {
      this.rankText = "de bronze";
    }
    if (this.rank === 1) {
      this.rankText = "d'argent";
    }
    if (this.rank === 2) {
      this.rankText = "d'or";
    }
  },
};
</script>

<style scoped>
/* Les cards sont trop grandes dans Mozilla jusqu'à ce que j'ouvre l'inspecteur
ntm */
.starCard {
  display: flex;
  padding: 5px;
  margin-left: 0.5vw;
  margin-right: 0.5vw;
  border-radius: 0.5vh;
  /* border: solid 2px; */
  height: 75%;
  min-width: 25vw;
  max-width: 30vw;
  flex-grow: 0;
  flex-shrink: 0;
  color: black;
  background: white;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
}

.starText {
  /* border: 2px solid black; */
  border-radius: 0.5vh;
  width: 100%;
  padding: 0.5%;
  font-size: 0.9vw;
  background: rgb(247, 246, 246);
}

.starLogo {
  margin-right: 5px;
  height: 100%;
  aspect-ratio: 1/1;
  /* border: solid 2px; */
  /* background-color: rgb(186, 185, 185); */
}
</style>
