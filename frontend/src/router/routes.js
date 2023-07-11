import { createWebHistory, createRouter } from "vue-router";
import ListUsers from "@/views/ListUsers.vue";
import MainLayout from "@/layouts/MainLayout";
import Error404 from "@/views/errors/Error404";
import Home from "@/views/Home";
import Login from "@/views/Login";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
        meta: {
            layout: MainLayout
        }
    },
    {
        path: "/admin",
        name: "Login",
        component: Login,
        meta: {
            layout: MainLayout
        }
    },
    {
        path: "/users",
        name: "ListUsers",
        component: ListUsers,
        props: true,
        meta: {
            layout: MainLayout
        }
    },
    {
        path: "/:pathMatch(.*)*",
        name: "404",
        component: Error404,
        meta: {
            layout: MainLayout
        }
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;