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
          <defs>
            <radialGradient
              id="G0"
              gradientTransform="translate(-0.66 -0.61) scale(1.32, 1.22)"
            >
              <stop offset="0%" stop-color="#ffffff" />
              <stop offset="4.17%" stop-color="#fdfcfc" />
              <stop offset="8.33%" stop-color="#fbfafa" />
              <stop offset="12.5%" stop-color="#f9f7f7" />
              <stop offset="16.67%" stop-color="#f8f4f3" />
              <stop offset="20.83%" stop-color="#f8f1ef" />
              <stop offset="25%" stop-color="#f7eeea" />
              <stop offset="29.17%" stop-color="#f8ebe3" />
              <stop offset="33.33%" stop-color="#f9e7db" />
              <stop offset="37.5%" stop-color="#fae5cf" />
              <stop offset="41.67%" stop-color="#fbe2bd" />
              <stop offset="50%" stop-color="#ffdf53" />
              <stop offset="54.17%" stop-color="#ffdc50" />
              <stop offset="58.33%" stop-color="#ffd94e" />
              <stop offset="62.5%" stop-color="#ffd74b" />
              <stop offset="66.67%" stop-color="#ffd448" />
              <stop offset="70.83%" stop-color="#ffd145" />
              <stop offset="75%" stop-color="#ffce42" />
              <stop offset="79.17%" stop-color="#ffcb3f" />
              <stop offset="83.33%" stop-color="#ffc93b" />
              <stop offset="87.5%" stop-color="#ffc638" />
              <stop offset="91.67%" stop-color="#ffc335" />
              <stop offset="100%" stop-color="#ffbd2e" />
            </radialGradient>
            <radialGradient
              id="G1"
              gradientTransform="translate(0.27 0.38) scale(1.47, 1.25)"
            >
              <stop offset="0%" stop-color="#ffffff" />
              <stop offset="3.13%" stop-color="rgba(254, 254, 254, 0.93)" />
              <stop offset="6.25%" stop-color="rgba(254, 254, 254, 0.86)" />
              <stop offset="9.38%" stop-color="rgba(253, 253, 253, 0.79)" />
              <stop offset="12.5%" stop-color="rgba(253, 252, 252, 0.72)" />
              <stop offset="15.63%" stop-color="rgba(253, 251, 251, 0.65)" />
              <stop offset="18.75%" stop-color="rgba(253, 250, 249, 0.59)" />
              <stop offset="21.88%" stop-color="rgba(253, 250, 247, 0.52)" />
              <stop offset="25%" stop-color="rgba(253, 249, 244, 0.45)" />
              <stop offset="28.13%" stop-color="rgba(253, 248, 240, 0.38)" />
              <stop offset="31.25%" stop-color="rgba(254, 248, 232, 0.31)" />
              <stop offset="37.5%" stop-color="rgba(255, 255, 0, 0.17)" />
              <stop offset="40.63%" stop-color="rgba(254, 239, 119, 0.16)" />
              <stop offset="43.75%" stop-color="rgba(251, 225, 128, 0.14)" />
              <stop offset="46.88%" stop-color="rgba(247, 212, 131, 0.13)" />
              <stop offset="50%" stop-color="rgba(241, 199, 133, 0.11)" />
              <stop offset="53.13%" stop-color="rgba(235, 187, 135, 0.1)" />
              <stop offset="56.25%" stop-color="rgba(226, 176, 137, 0.09)" />
              <stop offset="59.38%" stop-color="rgba(216, 165, 138, 0.07)" />
              <stop offset="62.5%" stop-color="rgba(204, 155, 139, 0.06)" />
              <stop offset="65.63%" stop-color="rgba(190, 146, 139, 0.04)" />
              <stop offset="68.75%" stop-color="rgba(173, 139, 138, 0.03)" />
              <stop offset="75%" stop-color="rgba(128, 128, 128, 0)" />
            </radialGradient>
          </defs>
          <polygon
            fill="url(#G0)"
            stroke="url(#G0)"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
          <polygon
            fill="url(#G1)"
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
          <defs>
            <radialGradient
              id="S0"
              gradientTransform="translate(-1 -1) scale(2, 2)"
            >
              <stop offset="0%" stop-color="#ffffff" />
              <stop offset="5.58%" stop-color="#f7f7f7" />
              <stop offset="11.17%" stop-color="#efefef" />
              <stop offset="16.75%" stop-color="#e7e7e7" />
              <stop offset="22.33%" stop-color="#dfdfdf" />
              <stop offset="27.92%" stop-color="#d8d8d8" />
              <stop offset="33.5%" stop-color="#d0d0d0" />
              <stop offset="39.08%" stop-color="#c8c8c8" />
              <stop offset="44.67%" stop-color="#c1c1c1" />
              <stop offset="50.25%" stop-color="#b9b9b9" />
              <stop offset="55.83%" stop-color="#b2b2b2" />
              <stop offset="67%" stop-color="#a3a3a3" />
              <stop offset="69.75%" stop-color="#a2a2a2" />
              <stop offset="72.5%" stop-color="#a0a0a0" />
              <stop offset="75.25%" stop-color="#9f9f9f" />
              <stop offset="78%" stop-color="#9e9e9e" />
              <stop offset="80.75%" stop-color="#9d9d9d" />
              <stop offset="83.5%" stop-color="#9b9b9b" />
              <stop offset="86.25%" stop-color="#9a9a9a" />
              <stop offset="89%" stop-color="#999999" />
              <stop offset="91.75%" stop-color="#989898" />
              <stop offset="94.5%" stop-color="#969696" />
              <stop offset="100%" stop-color="#949494" />
            </radialGradient>
            <radialGradient
              id="S1"
              gradientTransform="translate(0.38 0.52) scale(1.25, 0.96)"
            >
              <stop offset="0%" stop-color="rgba(255, 255, 255, 0.53)" />
              <stop offset="8.33%" stop-color="rgba(244, 244, 244, 0.49)" />
              <stop offset="16.67%" stop-color="rgba(233, 233, 233, 0.44)" />
              <stop offset="25%" stop-color="rgba(222, 222, 222, 0.4)" />
              <stop offset="33.33%" stop-color="rgba(211, 211, 211, 0.35)" />
              <stop offset="41.67%" stop-color="rgba(200, 200, 200, 0.31)" />
              <stop offset="50%" stop-color="rgba(189, 189, 189, 0.27)" />
              <stop offset="58.33%" stop-color="rgba(179, 179, 179, 0.22)" />
              <stop offset="66.67%" stop-color="rgba(169, 169, 169, 0.18)" />
              <stop offset="75%" stop-color="rgba(158, 158, 158, 0.13)" />
              <stop offset="83.33%" stop-color="rgba(148, 148, 148, 0.09)" />
              <stop offset="100%" stop-color="rgba(128, 128, 128, 0)" />
            </radialGradient>
          </defs>
          <polygon
            fill="url(#S0)"
            stroke="url(#S0)"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
          <polygon
            fill="url(#S1)"
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
          <defs>
            <radialGradient
              id="B0"
              gradientTransform="translate(-0.94 -0.83) scale(1.88, 1.67)"
            >
              <stop offset="0%" stop-color="#cb9855" />
              <stop offset="6.42%" stop-color="#c3904d" />
              <stop offset="12.83%" stop-color="#bc8845" />
              <stop offset="19.25%" stop-color="#b4803e" />
              <stop offset="25.67%" stop-color="#ac7836" />
              <stop offset="32.08%" stop-color="#a4702f" />
              <stop offset="38.5%" stop-color="#9d6828" />
              <stop offset="44.92%" stop-color="#956121" />
              <stop offset="51.33%" stop-color="#8d591a" />
              <stop offset="57.75%" stop-color="#855214" />
              <stop offset="64.17%" stop-color="#7c4b0c" />
              <stop offset="77%" stop-color="#6c3d00" />
              <stop offset="78.92%" stop-color="#6b3d00" />
              <stop offset="80.83%" stop-color="#6b3d00" />
              <stop offset="82.75%" stop-color="#6a3c00" />
              <stop offset="84.67%" stop-color="#6a3c00" />
              <stop offset="86.58%" stop-color="#693c00" />
              <stop offset="88.5%" stop-color="#693c00" />
              <stop offset="90.42%" stop-color="#683b00" />
              <stop offset="92.33%" stop-color="#683b00" />
              <stop offset="94.25%" stop-color="#673b00" />
              <stop offset="96.17%" stop-color="#673b00" />
              <stop offset="100%" stop-color="#663a00" />
            </radialGradient>
            <radialGradient
              id="B1"
              gradientTransform="translate(0.12 0.47) scale(1.77, 1.07)"
            >
              <stop offset="0%" stop-color="rgba(202, 151, 83, 0.64)" />
              <stop offset="8.33%" stop-color="rgba(201, 147, 90, 0.59)" />
              <stop offset="16.67%" stop-color="rgba(199, 143, 96, 0.53)" />
              <stop offset="25%" stop-color="rgba(197, 140, 102, 0.48)" />
              <stop offset="33.33%" stop-color="rgba(195, 136, 106, 0.43)" />
              <stop offset="41.67%" stop-color="rgba(192, 133, 111, 0.37)" />
              <stop offset="50%" stop-color="rgba(188, 130, 114, 0.32)" />
              <stop offset="58.33%" stop-color="rgba(184, 127, 117, 0.27)" />
              <stop offset="66.67%" stop-color="rgba(179, 125, 120, 0.21)" />
              <stop offset="75%" stop-color="rgba(168, 125, 125, 0.16)" />
              <stop offset="83.33%" stop-color="rgba(156, 126, 128, 0.11)" />
              <stop offset="100%" stop-color="rgba(128, 128, 128, 0)" />
            </radialGradient>
          </defs>
          <polygon
            fill="url(#B0)"
            stroke="url(#B0)"
            stroke-width="37.6152"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-miterlimit="10"
            points="
	259.216,29.942 330.27,173.919 489.16,197.007 374.185,309.08 401.33,467.31 259.216,392.612 117.104,467.31 144.25,309.08 
	29.274,197.007 188.165,173.919 "
          />
          <polygon
            fill="url(#B1)"
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
  height: 90%;
  width: 90%;
  margin-left: 5%;
  border-radius: 0 0 10px 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 2px solid #1d1d1b;
  border-top: none;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
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

/* .recapCard:nth-of-type(2) > .recapLogo > svg > polygon {
  fill: var(--gold-star-color);
  stroke: var(--gold-star-color);
} */

/* .recapCard:nth-of-type(3) > .recapLogo > svg > polygon {
  fill: var(--silver-star-color);
  stroke: var(--silver-star-color);
} */

/* .recapCard:nth-of-type(4) > .recapLogo > svg > polygon {
  fill: var(--bronze-star-color);
  stroke: var(--bronze-star-color);
} */

.recapNumber {
  /* background-color: #dadada; */
  border-radius: 5px;
  font-size: 150%;
  user-select: none;
}
</style>
