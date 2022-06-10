<template>
  <div class="starContainer">
      <StarEntry v-for="star in allStars" :key="star.id"
        :message="star.message"
        :date="star.created_at"
        :rank="star.type"
        :receiver="star.receiver_name"
      />
  </div>
  <!-- <StarEntry
    message="ceci est un message"
    date="12345000000000000000"
    rank="d'or"
    receiver="Louis-Vincent CAPELLI"
  />
  <StarEntry
    message="ceci est un messageaz aerfhbzefgv  vaze azet atztaztazt  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
    date="121296684234"
    rank="en or"
    receiver="Louis-Vincent CAPELLI"
  />
  <StarEntry
    message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
    date="1623000000"
    rank="d'or"
    receiver="Louis-Vincent CAPELLI"
  />
  <StarEntry
    message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
    date="1423"
    rank="d'or"
    receiver="Louis-Vincent CAPELLI"
  />
  <StarEntry
    message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
    date="1423"
    rank="d'or"
    receiver="Louis-Vincent CAPELLI"
  />
  <StarEntry
    message="eci est un messageaz aerfhbzefgv  vaze azet atztaztazt  azh vy v vazv v aaztazâ az azetrae taztataz a ta tragazrha e azet"
    date="1423"
    rank="d'or"
    receiver="Louis-Vincent CAPELLI"
  /> -->
  <!-- </div> -->
</template>

<script>
import { getRequest } from "@/requests/getRequest";
import StarEntry from "./StarEntry.vue";
export default {
  components: { StarEntry },
  data() {
    return {
      allStars: [],
    };
  },
  async created() {
    await getRequest("/stars/all", "json").then((res) => {
      this.allStars = res.data;
    });
  },
};
</script>

<style scoped>
.starContainer {
  display: flex;
  height: 100%;
  align-items: center;
}
</style>
