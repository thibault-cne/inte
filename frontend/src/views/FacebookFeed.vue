<template>
  <div class="facebook-feed"></div>
</template>

<script>
// import axios from "axios";
// import { parseXML } from "@/javascript/xmlToJson";
import { getRequest } from "@/requests/getRequest";
import { authStore } from "@/store/authStore";

export default {
  name: "FacebookFeed",
  mounted() {
    /*
    axios
      .get("https://rss.app/feeds/9yteWA5mbhfzeosR.xml")
      .then((response) => {
        // parsing xml data
        parseXML(response.data).then((data) => {
          console.log(data);
        });
      })
      .catch((error) => {
        console.log(error);
      });
    */
    let params = {
      user_id: 1,
      points: 100,
    };
    getRequest("/points-api/add", "json", params).then((response) => {
      console.log(response);
    });
    authStore.commit("destroyToken");
  },
};
</script>

<style scoped></style>
