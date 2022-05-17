<template>
  <div>
    <input type="checkbox" id="toggle" class="visuallyHidden" />
    <div class="navbar">
      <router-link to="/" class="navbarItem">
        <div class="navbarLogo">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512">
            <path
              d="M575.8 255.5C575.8 273.5 560.8 287.6 543.8 287.6H511.8L512.5 447.7C512.5 450.5 512.3 453.1 512 455.8V472C512 494.1 494.1 512 472 512H456C454.9 512 453.8 511.1 452.7 511.9C451.3 511.1 449.9 512 448.5 512H392C369.9 512 352 494.1 352 472V384C352 366.3 337.7 352 320 352H256C238.3 352 224 366.3 224 384V472C224 494.1 206.1 512 184 512H128.1C126.6 512 125.1 511.9 123.6 511.8C122.4 511.9 121.2 512 120 512H104C81.91 512 64 494.1 64 472V360C64 359.1 64.03 358.1 64.09 357.2V287.6H32.05C14.02 287.6 0 273.5 0 255.5C0 246.5 3.004 238.5 10.01 231.5L266.4 8.016C273.4 1.002 281.4 0 288.4 0C295.4 0 303.4 2.004 309.5 7.014L564.8 231.5C572.8 238.5 576.9 246.5 575.8 255.5L575.8 255.5z"
            />
          </svg>
        </div>
        <div class="navbarText">home</div>
      </router-link>
      <router-link to="/login" class="navbarItem">
        <div class="navbarLogo">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
            <path
              d="M272 304h-96C78.8 304 0 382.8 0 480c0 17.67 14.33 32 32 32h384c17.67 0 32-14.33 32-32C448 382.8 369.2 304 272 304zM48.99 464C56.89 400.9 110.8 352 176 352h96c65.16 0 119.1 48.95 127 112H48.99zM224 256c70.69 0 128-57.31 128-128c0-70.69-57.31-128-128-128S96 57.31 96 128C96 198.7 153.3 256 224 256zM224 48c44.11 0 80 35.89 80 80c0 44.11-35.89 80-80 80S144 172.1 144 128C144 83.89 179.9 48 224 48z"
            />
          </svg>
        </div>
        <div class="navbarText">profil</div>
      </router-link>
      <router-link to="/classement" class="navbarItem">
        <div class="navbarLogo">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512">
            <path
              d="M406.1 61.65C415.4 63.09 419.4 74.59 412.6 81.41L374.6 118.1L383.6 170.1C384.1 179.5 375.3 186.7 366.7 182.4L320.2 157.9L273.3 182.7C264.7 187 255 179.8 256.4 170.5L265.4 118.4L227.4 81.41C220.6 74.59 224.6 63.09 233.9 61.65L286.2 54.11L309.8 6.332C314.1-2.289 326.3-1.93 330.2 6.332L353.8 54.11L406.1 61.65zM384 256C401.7 256 416 270.3 416 288V480C416 497.7 401.7 512 384 512H256C238.3 512 224 497.7 224 480V288C224 270.3 238.3 256 256 256H384zM160 320C177.7 320 192 334.3 192 352V480C192 497.7 177.7 512 160 512H32C14.33 512 0 497.7 0 480V352C0 334.3 14.33 320 32 320H160zM448 416C448 398.3 462.3 384 480 384H608C625.7 384 640 398.3 640 416V480C640 497.7 625.7 512 608 512H480C462.3 512 448 497.7 448 480V416z"
            />
          </svg>
        </div>
        <div class="navbarText">classement</div>
      </router-link>
      <router-link to="/facebook-feed" class="navbarItem">
        <div class="navbarLogo"></div>
        <div class="navbarText">facebook-feed</div>
      </router-link>
      <router-link to="/cestjusteuntest" class="navbarItem">
        <div class="navbarLogo"></div>
        <div class="navbarText">cestjusteuntest</div>
      </router-link>
      <label for="toggle" class="toggleBtn">click</label>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
};
</script>

<style scoped>
.navbar {
  position: absolute;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-end;
  height: 100%;
  width: 5vw;
  background: rgb(244, 130, 244);
  transition: 200ms ease;
}

#toggle:checked ~ .navbar {
  width: 12vw;
}

.visuallyHidden {
  position: absolute;
  left: -100vw;
}

.navbarItem {
  display: flex;
  align-items: center;
  height: 5vw;
  width: 100%;
  text-decoration: none;
}

.navbarItem:hover {
  background: rgb(237, 102, 237);
}

.navbarLogo {
  margin-left: 0.5vw;
  min-height: 3vw;
  max-height: 3vw;
  min-width: 3vw;
  max-width: 3vw;
  background: white;
  border-radius: 1vw;
  /* box-shadow: none */
}

.navbarText {
  opacity: 0;
  margin-left: 0.5vw;
  color: black;
  font-size: 1.3vw;
}

#toggle:checked ~ .navbar .navbarItem .navbarText {
  opacity: 1;
  transition-delay: 200ms;
}

.toggleBtn {
  height: 2vw;
  width: 2vw;
  background: black;
  margin-top: auto;
  margin-bottom: 1vw;
  margin-right: 1.5vw;
  color: white;
}
</style>
