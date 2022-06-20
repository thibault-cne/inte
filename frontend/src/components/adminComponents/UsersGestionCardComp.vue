<template lang="">
  <div class="card">
    <div class="card-header">
      <h3>GESTION DES UTILISATEURS</h3>
      <svg
        width="77"
        height="92"
        viewBox="0 0 77 92"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <circle cx="54.449" cy="12.2041" r="12.2041" />
        <path
          d="M66.0408 50.2245C71.5637 50.2245 76.2315 45.5284 73.8129 40.5633C73.3199 39.5512 72.748 38.574 72.1002 37.6411C69.5346 33.9465 65.9009 31.1244 61.6862 29.5531C57.4714 27.9818 52.8771 27.7364 48.519 28.8498C47.4186 29.1309 46.3466 29.4952 45.3113 29.9376C40.2327 32.1076 40.687 38.7133 44.862 42.3288L51.1613 47.7839C52.9791 49.358 55.3031 50.2245 57.7077 50.2245H66.0408Z"
        />
        <circle cx="26.2857" cy="21.5918" r="14.0816" />
        <path
          d="M41.6327 66.1837C47.1555 66.1837 51.8062 61.5455 49.72 56.4318C49.7026 56.3892 49.6851 56.3467 49.6675 56.3042C48.3701 53.172 46.4685 50.326 44.0712 47.9288C41.674 45.5315 38.828 43.6299 35.6958 42.3325C32.5636 41.0351 29.2066 40.3673 25.8163 40.3673C22.4261 40.3673 19.069 41.0351 15.9368 42.3325C12.8047 43.6299 9.95869 45.5315 7.56143 47.9288C5.16416 50.326 3.26254 53.172 1.96515 56.3042C1.94754 56.3467 1.93005 56.3893 1.91268 56.4318C-0.17352 61.5455 4.47715 66.1837 10 66.1837L25.8163 66.1837H41.6327Z"
        />
      </svg>
    </div>
    <div class="card-body">
      <GestionRowComp route="admin/user-list" texte="Liste des utilisateurs" />
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
