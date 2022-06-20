<template lang="">
  <div class="card">
    <div class="card-header">
      <h3>GESTION DES POINTS</h3>
      <svg
        width="77"
        height="92"
        viewBox="0 0 77 92"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M28.0979 1.8541C28.6966 0.0114746 31.3034 0.0114799 31.9021 1.8541L37.409 18.8024C37.6767 19.6265 38.4446 20.1844 39.3111 20.1844L57.1316 20.1844C59.0691 20.1844 59.8746 22.6636 58.3072 23.8024L43.8901 34.2771C43.1891 34.7864 42.8958 35.6891 43.1635 36.5132L48.6704 53.4615C49.2691 55.3041 47.1601 56.8364 45.5927 55.6976L31.1756 45.2229C30.4746 44.7136 29.5254 44.7136 28.8244 45.2229L14.4073 55.6976C12.8399 56.8364 10.7309 55.3041 11.3296 53.4615L16.8365 36.5132C17.1042 35.6891 16.8109 34.7864 16.1099 34.2771L1.69281 23.8024C0.125385 22.6636 0.93094 20.1844 2.86839 20.1844L20.6889 20.1844C21.5554 20.1844 22.3233 19.6265 22.591 18.8024L28.0979 1.8541Z"
        />
      </svg>
    </div>
    <div class="card-body">
      <GestionRowComp
        route="faudra mettre la bonne route quoi"
        texte="Modération des étoiles"
      />
    </div>
  </div>
</template>
<script>
import GestionRowComp from "./GestionRowComp.vue";
export default {
  components: {
    GestionRowComp,
  },
};
</script>
<style scoped lang="scss">
.card {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 80%;
  height: 80%;
  background: rgb(250, 250, 250);
  border-radius: 10px;
}

.card-header {
  width: 100%;
  height: 30%;
  background-color: rgb(0, 0, 0);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  border-radius: 10px 10px 0 0;
}

.card-header h3 {
  font-size: 1em;
  font-weight: bold;
  font-family: "Gill Sans", "Gill Sans MT", Calibri, "Trebuchet MS", sans-serif;
  color: rgb(211, 211, 211);
}

.card-header svg {
  fill: rgb(211, 211, 211);
}

.card-body {
  width: 90%;
  height: 30%;
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}
</style>
