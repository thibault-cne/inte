import axios from "axios";

// Instagram feed api
export function instagramFeed() {
  axios.get("https://rss.app/feeds/9yteWA5mbhfzeosR.xml").then((res) => {
    console.log(res.data);
  });
}
