<template>
    <!-- <div v-for="star in allStars" :key="star.id">
      <StarEntry
        :message="star.message"
        :date="star.date"
        :rank="star.star_rank"
        :receiver="star.receiver_user_id"
      /> -->
      <StarEntry
        message="ceci est un message"
        date=12345
        rank="d'or"
        receiver=1
      />
      <StarEntry
        message="ceci est un messageaz aerfhbzefgv erfdgarefdhgva  azbtza v  yzt yza vay vaze azet atztaztazt  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
        date=12124234
        rank="en or"
        receiver=1
      />
      <StarEntry
        message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
        date=1423
        rank="en or"
        receiver=1
      />
      <StarEntry
        message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
        date=1423
        rank="en or"
        receiver=1
      />
      <StarEntry
        message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
        date=1423
        rank="en or"
        receiver=1
      />
      <StarEntry
        message="ceci est unzbtza v  yzt yza vay v  azh vy v vazv v azvetv aiyzvet vazty vazy vazt"
        date=1423
        rank="en or"
        receiver=1
      />
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

</style>
