<template>
    <div class="starCard">
        {{receiver}} a reçu une étoile {{rank}} il y a {{ilyaText}} :<br>
        {{message}}
    </div>
</template>

<script>

export default {
    props: {
        message: String,
        receiver: Number,
        rank: String,
        date: Number,
    },
    data() {
        return {
            ilyaText: "2 semaines",
        }
    },
    created() {
        // let ilya = new Date().getTime();
        // ilya = ilya - this.date;
        // ilya = Math.floor(ilya/3600000)

        // switch (ilya) {
        //     case 0:
        //         this.ilyaText = "moins d'1h"
        //         break;

        //     case 1-23:
        //         this.ilyaText = `${ilya}h`
        //         break;

        //     case 24-167:
        //         ilya = Math.floor(ilya/24)
        //         this.ilyaText = `${ilya}j`
        //         break;

        //     case 168-719:
        //         ilya = Math.floor((ilya/24)/7)
        //         this.ilyaText = 1 ? `${ilya} semaine` : `${ilya} semaines`
        //         break;
        
        //     default:
        //         ilya = Math.floor(((ilya/24)/7)/4)
        //         this.ilyaText = `${ilya}mois`
        //         break;
        // }
    }
}
</script>

<style scoped>
.starCard {
    /* padding: 10px; */
    margin-top: 1vh;
    margin-bottom: 1vh;
    border-radius: 2vh;
    border: solid;
    height: 20vh;
    width: 50vw;
}
</style>