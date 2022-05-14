<template>
    <div>
        <input type="checkbox" id="toggle" class="visuallyHidden">
        <div class="navbar">
            <div class="navbarItem">
                <div class="navbarLogo"></div>
                <div class="navbarText">TEST TEXTE</div>
            </div>
            <div class="navbarItem">
                <div class="navbarLogo"></div>
                <div class="navbarText">PAREIL MDR</div>
            </div>
            <div class="navbarItem">
                <div class="navbarLogo"></div>
                <div class="navbarText">PAREIL MDR</div>
            </div>
            <div class="navbarItem">
                <div class="navbarLogo"></div>
                <div class="navbarText">PAREIL MDR</div>
            </div>
            <label for="toggle" class="toggleBtn">click</label>
        </div>
    </div>
</template>

<script>

export default {
    components: {}
}
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
}

.navbarItem:hover {
    background: rgb(237, 102, 237);
}

.navbarLogo {
    margin-left: .5vw;
    height: 4vw;
    min-width: 4vw;
    background: white;
    border-radius: 1vw;
    /* box-shadow: none */
}

.navbarText {
    opacity: 0;
    margin-left: .5vw;
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