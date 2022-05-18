<template>
  <div class="stars">
    <table>
      <thead>
        <td>id</td>
        <td>receveur</td>
        <td>rank</td>
        <td>date</td>
        <td>message</td>
      </thead>
      <tr v-for="star in allStars" :key="star.id">
        <td>{{ star.id }}</td>
        <td>{{ star.receiver_user_id }}</td>
        <td>{{ star.star_rank }}</td>
        <td>{{ star.date }}</td>
        <td>{{ star.message }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  data() {
    return {
      allStars: ["21"],
    };
  },
  async created() {
    await getRequest("stars-info/all", "json").then((res) => {
      this.allStars = res.data;
    });
  },
};
</script>

<style>
.stars {
  max-height: 54vh;
  min-height: 54vh;
  overflow: auto;
  background: rgb(240, 160, 160);
}
</style>
