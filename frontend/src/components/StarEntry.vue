<template>
  <div class="starCard">
    <div class="starLogo"></div>
    <div class="starText">
      {{ receiver }} a reçu une étoile {{ rank }} il y a {{ ilyaText }} :<br />
      {{ message }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    message: String,
    receiver: Number,
    rank: String,
    date: Number,
  },
  data() {
    return {
      ilyaText: "2 semaines",
    };
  },
  created() {
    // let ilya = new Date().getTime();
    // ilya = ilya - this.date;
    // ilya = Math.floor(ilya/3600000)
    // switch (ilya) {
    //     case 0:
    //         this.ilyaText = "moins d'1h"
    //         break;
    //     case 1-23:
    //         this.ilyaText = `${ilya}h`
    //         break;
    //     case 24-167:
    //         ilya = Math.floor(ilya/24)
    //         this.ilyaText = `${ilya}j`
    //         break;
    //     case 168-719:
    //         ilya = Math.floor((ilya/24)/7)
    //         this.ilyaText = 1 ? `${ilya} semaine` : `${ilya} semaines`
    //         break;
    //     default:
    //         ilya = Math.floor(((ilya/24)/7)/4)
    //         this.ilyaText = `${ilya}mois`
    //         break;
    // }
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
