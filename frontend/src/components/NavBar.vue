<template>
    <div class="navbar">
        <NavBarItem />
        <NavBarItem />
        <NavBarItem />
        <NavBarItem />
        <NavBarItem />
        <NavBarItem />
        <NavBarItem />
    </div>
</template>

<script>
import NavBarItem from "./NavBarItem.vue"

export default {
    components: {NavBarItem}
}
</script>

<style scoped>
.navbar {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    height: 100%;
    width: 100%;
    background: rgb(244, 130, 244);
}
</style>