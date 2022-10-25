<template>
    <v-container>
        <v-row justify="center">
            <v-col cols="12" md="6">
                <h1 class="text-center">Il faut mettre une image ici</h1>
            </v-col>
        </v-row>
        <v-row justify="center">
            <v-col cols="12" md="12" lg="6">
                <v-card>
                    <v-card-text class="mt-5 mb-5">
                        <v-row justify="center" class="mt-5 mb-5">
                            <v-col cols="8">
                                <!-- 
                                    Beautiful title
                                -->
                                <h1 class="text-center">Bienvenue sur le site de l'intégration</h1>
                            </v-col>
                        </v-row>
                        <v-row justify="center" class="mt-5 mb-5">
                            <v-col cols="6">
                                <!-- 
                                    Beautiful button
                                -->
                                <v-btn color="primary" class="white--text" block large :href="loginURL"> Se connecter avec Google </v-btn>
                                <!--  -->
                            </v-col>
                        </v-row>
                        <v-row justify="center" class="mt-5 mb-5">
                            <v-col cols="6">
                                <v-divider></v-divider>
                                <!-- 
                                    Help text
                                -->
                                <p class="text-center mt-5 mb-5" @click="triggerAlert()"><b>Problèmes de connexion ?</b></p>
                                <v-alert class="mb-5 text-center alert-popin" v-model="alert" type="warning" ref="alertComponent">
                                    <p>Pour te connecter à la plateforme, tu dois utiliser une adresse Google <b>@telecomnancy.net</b>.</p>
                                    <p>Si tu n'en as pas, contacte un <b>admin TN.net</b> pour qu'il t'en crée une.</p>
                                    <p>Si le problème persiste, essaie de désactiver Ad Block ou de te connecter en navigation privé.</p>
                                    <p>Si rien n'y fait, tu peux envoyer un mail à : </p>
                                    <!-- email -->
                                    <v-btn class="mt-5" href="mailto:admin@telecomnancy.net">admin@telecomnancy.net</v-btn>
                                </v-alert>
                                <v-divider></v-divider>
                            </v-col>
                        </v-row>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import { base_backend_url } from "@/requests/axios";

// Components

export default defineComponent({
    name: "LoginModal",
    data() {
        return {
            loginURL: base_backend_url + "/auth/login",
            alert: false,
            alertComponent: ref(null),
        };
    },
    methods: {
        triggerAlert() {
            if (this.alert) {
                let e = (this.$refs.alertComponent as HTMLFormElement)
                e.$el.classList.value = e.$el.classList.value.replace("popin", "popout");
                setTimeout(() => {
                    this.alert = false;
                    e.$el.classList.value = e.$el.classList.value.replace("popout", "popin");
                }, 500);
            } else {
                this.alert = true;
            }
        },
    },
});
</script>
  
<style scoped>
@keyframes popIn {
	0% {
		opacity: 0;
		transform: rotateX(-100deg);
		transform-origin: top;
	}

	100% {
		opacity: 1;
		transform: rotateX(0deg);
		transform-origin: top;
	}
}
@keyframes popOut {
	0% {
		opacity: 1;
		transform: rotateX(0deg);
		transform-origin: top;
	}

	100% {
		opacity: 0;
		transform: rotateX(70deg);
		transform-origin: top;
	}
}
.alert-popin {
    animation: popIn 0.5s ease 0s 1 normal forwards;
}
.alert-popout {
    animation: popOut 0.5s ease 0s 1 normal forwards; 
}
</style>