import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "Home",
    },
    component: HomeView,
  },
  {
    path: "/login",
    name: "login",
    meta: {
      title: "Login",
    },
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/LoginView.vue"),
  },
  {
    path: "/facebook-feed",
    name: "facebook-feed",
    meta: {
      title: "Facebook Feed",
    },
    component: () =>
      import(
        /* webpackChunkName: "facebook-feed" */ "../views/FacebookFeed.vue"
      ),
  },
  {
    path: "/admin",
    name: "admin",
    meta: {
      title: "Admin",
    },
    component: () =>
      import(/* webpackChunkName: "admin" */ "../views/AdminView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

// This callback runs before every route change, including on page load.
router.beforeEach((to, from, next) => {
  // This goes through the matched routes from last to first, finding the closest route with a title.
  // e.g., if we have `/some/deep/nested/route` and `/some`, `/deep`, and `/nested` have titles,
  // `/nested`'s will be chosen.
  const nearestWithTitle = to.matched
    .slice()
    .reverse()
    .find((r) => r.meta && r.meta.title);

  // If a route with a title was found, set the document (page) title to that value.
  if (nearestWithTitle) {
    document.title = nearestWithTitle.meta.title;
  }
  next();
});

export default router;
