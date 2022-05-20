<template>
  <div class="stars">
    <div v-for="star in allStars" :key="star.id">
      <StarEntry :message="star.message" :date="star.date" :rank="star.star_rank" :receiver="star.receiver_user_id"/>
    </div>
      <!-- <table>
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
      </table> -->
  </div>
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
.stars {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  height: 100%;
  width: 100%;
  overflow: auto;
  background: rgb(240, 160, 160);
  scrollbar-width: thin;
  scrollbar-color: rgb(173, 38, 38) rgb(227, 124, 124);
}

</style>
