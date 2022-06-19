<template>
  <div class="recapContainer">
    <div class="recapCard">
      <div class="recapLogo">
        <svg id="Layer_1" viewBox="0 -10 120.03 132.88">
          <path
            style="fill-rule: evenodd"
            d="M28.89,4.45A17.42,17.42,0,0,1,36.46.67a19.64,19.64,0,0,1,8.15-.1,17.41,17.41,0,0,1,6.17,2.52A17.38,17.38,0,0,1,56.28.6,18.75,18.75,0,0,1,63.84.22,20.52,20.52,0,0,1,71,2.77a21.48,21.48,0,0,1,7.87,8,14.91,14.91,0,0,1,9,2.92,16,16,0,0,1,4.34,4.76,14.94,14.94,0,0,1,2.06,6.45c.41,5-1.65,10.45-7.68,14.84V50.44l5-.44C101.71,49.11,110,58.29,110,68.41V90.05c0,9.93-8.16,18.91-18,18l-5.36-.46V113a9.93,9.93,0,0,1-2.92,7h0a9.92,9.92,0,0,1-7,2.92H18.24A10,10,0,0,1,8.31,113c0-7.69.1-62,.13-73.69C1.49,33.57-.68,27,.18,21.08A17.13,17.13,0,0,1,4.56,12a18.88,18.88,0,0,1,8.38-5.26,19.15,19.15,0,0,1,13.26.73,15.69,15.69,0,0,1,2.69-3ZM86.62,62.18V96.3h7a4.22,4.22,0,0,0,4.19-4.19V66.78a4.62,4.62,0,0,0-4.61-4.6ZM61.06,67.29a4.1,4.1,0,0,1,8.14,0V96.64a4.1,4.1,0,0,1-8.14,0V67.29Zm-35.32,0a3.85,3.85,0,0,1,4.06-3.58,3.85,3.85,0,0,1,4.07,3.58V96.64a3.85,3.85,0,0,1-4.07,3.58,3.85,3.85,0,0,1-4.06-3.58V67.29ZM43.4,64.53a4.1,4.1,0,0,1,8.13,0V99.39a4.1,4.1,0,0,1-8.13,0V64.53ZM79.49,41.92l-1.68.1c-1.38,0-2.79,0-4.23,0l0,.36A13.75,13.75,0,0,1,63.17,56a15.62,15.62,0,0,1-7.83,0A14.39,14.39,0,0,1,47,50.18a17.47,17.47,0,0,1-2.29-5.66C44.35,43.26,44,42,40.94,42H15.43v71a2.83,2.83,0,0,0,2.81,2.81H76.7a2.79,2.79,0,0,0,2-.82h0a2.79,2.79,0,0,0,.82-2v-71ZM14.06,35.87C-5.81,22,16.13,2.22,29.72,17.48,28.63,4.77,46.69,2.76,50.66,12,56.51,1.93,72,5.28,75.25,17.88c8.81-4.95,20.4,9,6.68,17.65a49.47,49.47,0,0,1-8.68.43c-4.54-.08-5.68,2.24-5.73,6.38C67.41,51.45,56,52.72,52,46.77c-2.67-4-.35-10.9-11.09-10.9Z"
          />
        </svg>
      </div>
      <div class="recapNumber">{{ points }}</div>
    </div>
    <div class="recapCard">
      <div class="recapLogo">
        <svg viewBox="0 -50 512 562">
          <polygon
            fill="#1D1D1B"
            stroke="#1D1D1B"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
        </svg>
      </div>
      <div class="recapNumber">{{ stars2 }}</div>
    </div>
    <div class="recapCard">
      <div class="recapLogo">
        <svg viewBox="0 -50 512 562">
          <polygon
            fill="#1D1D1B"
            stroke="#1D1D1B"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
        </svg>
      </div>
      <div class="recapNumber">{{ stars1 }}</div>
    </div>
    <div class="recapCard">
      <div class="recapLogo">
        <svg viewBox="0 -50 512 562">
          <polygon
            fill="#1D1D1B"
            stroke="#1D1D1B"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
        </svg>
      </div>
      <div class="recapNumber">{{ stars0 }}</div>
    </div>
  </div>
</template>

<script>
import { getRequest } from "../requests/getRequest";
export default {
  name: "RecapPointsStars",
  data() {
    return {
      points: 69,
      stars0: 69,
      stars1: 69,
      stars2: 69,
    };
  },
  async created() {
    await getRequest("/profile/user/get/stats", "data").then((res) => {
      this.points = res.data.points;
      this.stars0 = res.data.bronze_stars;
      this.stars1 = res.data.silver_stars;
      this.stars2 = res.data.gold_stars;
    });
  },
};
</script>

<style scoped>
.recapContainer {
  height: 30%;
  width: 90%;
  border-radius: 0 0 10px 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 2px solid #1d1d1b;
  border-top: none;
  /* background: lightgrey; */
}

.recapCard {
  height: 100%;
  width: 100%;
  margin: 2%;
  display: flex;
  align-items: center;
  justify-content: space-around;
  /* background-color: #949494; */
  border-radius: 5px;
}

.recapLogo {
  width: 60%;
  aspect-ratio: 1;
  /* background-color: #ffffff; */
  border-radius: 5px;
}

.recapCard:nth-of-type(2) > .recapLogo > svg > polygon {
  fill: var(--gold-star-color);
  stroke: var(--gold-star-color);
}

.recapCard:nth-of-type(3) > .recapLogo > svg > polygon {
  fill: var(--silver-star-color);
  stroke: var(--silver-star-color);
}

.recapCard:nth-of-type(4) > .recapLogo > svg > polygon {
  fill: var(--bronze-star-color);
  stroke: var(--bronze-star-color);
}

.recapNumber {
  /* background-color: #dadada; */
  border-radius: 5px;
  font-size: 150%;
  user-select: none;
}
</style>
