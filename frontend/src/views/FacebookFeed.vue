<template>
  <div class="facebook-feed"></div>
</template>

<script>
import axios from "axios";
import { parseXML } from "@/javascript/xmlToJson";

export default {
  name: "FacebookFeed",
  mounted() {
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
  },
};
</script>

<style scoped></style>
