<template>
  <div class="facebook-feed"></div>
</template>

<script>
import axios from "axios";
import * as convert from "xml-js";

export default {
  name: "FacebookFeed",
  mounted() {
    axios
      .get("https://rss.app/feeds/9yteWA5mbhfzeosR.xml")
      .then((response) => {
        let result2 = convert.xml2json(response.data, {
          compact: false,
          spaces: 4,
        });
        console.log(result2);
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>

<style scoped></style>
