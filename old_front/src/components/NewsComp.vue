<template>
  <div class="newsContainer">{{ newsText }}</div>
</template>

<script>
export default {
  data() {
    return {
      newsText: "ça bug",
    };
  },
  created() {
    this.newsText =
      "----------------- il faut récup ça dans la DB ----------------- J'arrive pas à revenir à la ligne WTF ---------------------- Bienvenue sur le site de l'inté, amusez-vous bien !";
  },
};
</script>

<style scoped>
.newsContainer {
  height: 70%;
  width: 86%;
  margin: 7%;
  overflow-y: auto;
  font-size: 2vh;
}
</style>
