<template>
  <div class="starCard">
    <div class="starLogo"></div>
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
  border-radius: 2vh;
  border: solid 2px;
  height: 75%;
  min-width: 24.5vw;
  max-width: 30vw;
  flex-shrink: 0;
  flex-grow: 0;
  color: black;
}

.starText {
  font-size: 2.25vh;
}

.starLogo {
  margin-right: 5px;
  height: 100%;
  aspect-ratio: 1/1;
  background-color: rgb(186, 185, 185);
}
</style>
