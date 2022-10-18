<template>
  <div class="header">
    <div class="gauche">
      <div class="nomSite">Criminal'inté</div>
    </div>
    <div class="droite">
      <div class="rechercheUser">
        <input
          id="searchBar"
          name="bar"
          type="text"
          placeholder="Rechercher un élève"
        />
        <div class="rechercheLogo">
          <svg viewBox="0 0 1024 1024">
            <path
              d="M848.471 928l-263.059-263.059c-48.941 36.706-110.118 55.059-177.412 55.059-171.294 0-312-140.706-312-312s140.706-312 312-312c171.294 0 312 140.706 312 312 0 67.294-24.471 128.471-55.059 177.412l263.059 263.059-79.529 79.529zM189.623 408.078c0 121.364 97.091 218.455 218.455 218.455s218.455-97.091 218.455-218.455c0-121.364-103.159-218.455-218.455-218.455-121.364 0-218.455 97.091-218.455 218.455z"
            ></path>
          </svg>
        </div>
      </div>
      <div class="notifs">
        <svg viewBox="0 0 1024 1024">
          <path
            fill="white"
            d="M511.8 912c62.2 0 96.2-44 96.2-106h-192.6c0 62 34 106 96.4 106zM824 704.4c-30.8-40.6-91.4-64.4-91.4-246.2 0-186.6-82.4-261.6-159.2-279.6-7.2-1.8-12.4-4.2-12.4-11.8v-5.8c0-26.8-22-49.4-48.8-49.2-26.8-0.4-48.8 22.4-48.8 49.2v5.8c0 7.4-5.2 10-12.4 11.8-77 18.2-159.2 93-159.2 279.6 0 181.8-60.6 205.4-91.4 246.2-19.8 26.2-1 63.6 31.8 63.6h560.2c32.6 0 51.4-37.6 31.6-63.6z"
          />
        </svg>
      </div>
      <div class="accesProfil">accesProfil (ou pas en fait on verra)</div>
    </div>
  </div>
</template>

<script></script>

<style scoped>
.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  width: 100%;
  background: rgb(0, 0, 0);
}

.gauche {
  position: absolute;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  height: 100%;
  width: 100%;
}

.droite {
  position: absolute;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
  height: 100%;
  width: 100%;
}

@import url("https://fonts.googleapis.com/css2?family=Black+Ops+One&display=swap");
.nomSite {
  margin-left: 0.5vw;
  height: 80%;
  width: 30%;
  border-radius: 0.5vh;
  font-family: "Black Ops One", cursive;
  color: white;
  font-size: 7vh;
  user-select: none;
}

.rechercheUser {
  margin-right: 1vw;
  border-radius: 3px;
  height: 35%;
  width: 12%;
  background: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

input[type="text"] {
  all: unset;
  padding: 5px;
  height: 100%;
  width: 100%;
  font-family: var(--main-font);
  font-size: 1.8vh;
}

.rechercheLogo {
  height: 100%;
  aspect-ratio: 1;
}
.notifs {
  margin-right: 1vw;
  height: 60%;
  width: 3%;
  border-radius: 0.5vh;
}

.accesProfil {
  margin-right: 1vw;
  height: 70%;
  width: 15%;
  background: white;
  border-radius: 0.5vh;
}
</style>
