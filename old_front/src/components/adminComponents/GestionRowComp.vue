<template>
  <router-link :to="route" class="row">{{ texte }}</router-link>
</template>

<script>
export default {
  props: {
    route: String,
    texte: String,
  },
};
</script>

<style>
.row {
  width: 95%;
  height: 20%;
  margin-bottom: 1%;
  border-radius: 5px;
  padding-left: 5%;
  background: rgb(233, 233, 233);
  color: black;
  text-decoration: none;
  font-size: 2vh;
}
</style>
