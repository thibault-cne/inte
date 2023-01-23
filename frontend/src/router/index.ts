import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import { isLogged } from "@/requests/logged";
import HomeView from "../views/HomeView.vue";
import { NavigationGuardNext, RouteLocationNormalized } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "Home",
    },
    component: HomeView,
    beforeEnter: checkAuth,
  },
  {
    path: "/login",
    name: "login",
    meta: {
      title: "Login",
    },
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/LoginModal.vue"),
  },
  {
    path: "/debug-stars",
    name: "debug-stars",
    meta: {
      title: "Debug Stars",
    },
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/DebugStars.vue"),
  },
  {
    path: "/debug-profile",
    name: "debug-profile",
    meta: {
      title: "Debug Profile",
    },
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/DebugProfile.vue"),
  },
  {
    path: "/debug-calendar",
    name: "debug-calendar",
    meta: {
      title: "Debug calendar",
    },
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/DebugCalendar.vue"),
  },
  {
    path: "/debug-admin",
    name: "debug-admin-page",
    meta: {
      title: "Debug admin page",
    },
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/DebugAdminPage.vue"),
  },
  {
    path: "/debug-user-management",
    name: "debug-user-management",
    meta: {
      title: "Debug user management",
    },
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(
        /* webpackChunkName: "about" */ "../views/DebugUserManagement.vue"
      ),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

async function checkAuth(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  const status = await isLogged();
  if (!status.logged && to.name !== "login") {
    // redirect the user to the login page
    next({ name: "login" });
    return;
  }
  next();
}

router.afterEach(async (to) => {
  // check meta to put title
  if (to.meta.title) {
    document.title = to.meta.title as string;
  }
});

export default router;
