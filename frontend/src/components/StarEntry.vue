<template>
  <div class="starCard">
    <div class="starLogo">
      <svg v-if="rank == 0" width="90" height="90" viewBox="-50 -50 562 562">
        <polygon
          fill="#967444"
          stroke="#967444"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
      </svg>
      <svg v-if="rank == 1" width="90" height="90" viewBox="-50 -50 562 562">
        <polygon
          fill="#c6c6c6"
          stroke="#c6c6c6"
          stroke-width="37.6152"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-miterlimit="10"
          points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
        />
      </svg>
      <svg v-if="rank == 2" width="90" height="90" viewBox="-50 -50 562 562">
        <polygon
          fill="#ffcf00"
          stroke="#ffcf00"
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
      {{ receiver }} <br />
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
.starCard {
  display: flex;
  padding: 5px;
  margin-left: 0.5vw;
  margin-right: 0.5vw;
  border-radius: 0.5vh;
  border: solid 2px;
  height: 75%;
  min-width: 24.5vw;
  max-width: 30vw;
  flex-shrink: 0;
  flex-grow: 0;
  color: black;
}

.starText {
  border: 2px solid black;

  border-radius: 0.5vh;
  width: 100%;
  padding: 0.5%;
  font-size: 1vw;
}

.starLogo {
  margin-right: 5px;
  height: 100%;
  aspect-ratio: 1/1;
  /* border: solid 2px; */
  /* background-color: rgb(186, 185, 185); */
}
</style>
