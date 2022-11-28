import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/doc-overview",
    component: () => import("@/layout/Layout.vue"),
    children: [
      {
        path: "build",
        name: "build",
        component: () => import("@/views/get-started/Build.vue"),
      },
      {
        path: "file-structure",
        name: "file-structure",
        component: () => import("@/views/get-started/FileStructure.vue"),
      },
      {
        path: "setup-theme-skeleton",
        name: "setup-theme-skeleton",
        component: () => import("@/views/get-started/SetupThemeSkeleton.vue"),
      },
      {
        path: "vue-laravel-integration",
        name: "vue-laravel-integration",
        component: () =>
          import("@/views/get-started/VueLaravelIntegration.vue"),
      },
      {
        path: "internationalization",
        name: "internationalization",
        component: () => import("@/views/get-started/Internationalization.vue"),
      },
      {
        path: "rtl",
        name: "rtl",
        component: () => import("@/views/get-started/RTL.vue"),
      },
      {
        path: "api-setup",
        name: "api-setup",
        component: () => import("@/views/get-started/APISetup.vue"),
      },
      {
        path: "doc-overview",
        name: "doc-overview",
        component: () => import("@/views/get-started/Overview.vue"),
      },
      {
        path: "updates",
        name: "updates",
        component: () => import("@/views/get-started/Updates.vue"),
      },
      {
        path: "changelog",
        name: "changelog",
        component: () => import("@/views/get-started/Changelog.vue"),
      },
      {
        path: "utilities",
        name: "utilities",
        meta: {
          desc: "extended utility classes",
        },
        component: () => import("@/views/base/Utilities.vue"),
      },
      {
        path: "helpers/flex-layouts",
        name: "flex-layouts",
        meta: {
          desc: "extended flex layout classes",
        },
        component: () => import("@/views/base/helpers/FlexLayouts.vue"),
      },
      {
        path: "helpers/text",
        name: "text",
        meta: {
          desc: "extended text classes",
        },
        component: () => import("@/views/base/helpers/Text.vue"),
      },
      {
        path: "helpers/background",
        name: "backkground",
        meta: {
          desc: "extended background classes",
        },
        component: () => import("@/views/base/helpers/Background.vue"),
      },
      {
        path: "helpers/borders",
        name: "borders",
        meta: {
          desc: "extended borders classes",
        },
        component: () => import("@/views/base/helpers/Borders.vue"),
      },
      {
        path: "helpers/opacity",
        name: "opacity",
        component: () => import("@/views/base/helpers/Opacity.vue"),
      },
      {
        path: "forms",
        name: "forms",
        meta: {
          desc: "forms elements",
        },
        component: () => import("@/views/base/Forms.vue"),
      },
      {
        path: "buttons",
        name: "buttons",
        meta: {
          desc: "buttons elements",
        },
        component: () => import("@/views/base/Buttons.vue"),
      },
      {
        path: "indicator",
        name: "indicator",
        meta: {
          desc: "indicator element",
        },
        component: () => import("@/views/base/Indicator.vue"),
      },
      {
        path: "rotate",
        name: "rotate",
        meta: {
          desc: "Rotate element",
        },
        component: () => import("@/views/base/Rotate.vue"),
      },
      {
        path: "tables",
        name: "tables",
        meta: {
          desc: "extended bootstrap tables",
        },
        component: () => import("@/views/base/Tables.vue"),
      },
      {
        path: "cards",
        name: "cards",
        meta: {
          desc: "card elements",
        },
        component: () => import("@/views/base/Cards.vue"),
      },
      {
        path: "symbol",
        name: "symbol",
        meta: {
          desc: "symbol elements",
        },
        component: () => import("@/views/base/Symbol.vue"),
      },
      {
        path: "badges",
        name: "badges",
        meta: {
          desc: "badge elements",
        },
        component: () => import("@/views/base/Badges.vue"),
      },
      {
        path: "pulse",
        name: "pulse",
        meta: {
          desc: "pulse elements",
        },
        component: () => import("@/views/base/Pulse.vue"),
      },
      {
        path: "rating",
        name: "rating",
        component: () => import("@/views/base/Rating.vue"),
      },
      {
        path: "ribbon",
        name: "ribbon",
        component: () => import("@/views/base/Ribbon.vue"),
      },
      {
        path: "bullets",
        name: "bullets",
        meta: {
          desc: "bullets elements",
        },
        component: () => import("@/views/base/Bullets.vue"),
      },
      {
        path: "accordion",
        name: "accordion",
        meta: {
          desc: "accordion elements",
        },
        component: () => import("@/views/base/Accordion.vue"),
      },
      {
        path: "alerts",
        name: "alerts",
        component: () => import("@/views/base/Alerts.vue"),
      },
      {
        path: "carousel",
        name: "carousel",
        meta: {
          desc: "carousel elements",
        },
        component: () => import("@/views/base/Carousel.vue"),
      },
      {
        path: "overlay",
        name: "overlay",
        meta: {
          desc: "overlay elements",
        },
        component: () => import("@/views/base/Overlay.vue"),
      },
      {
        path: "separator",
        name: "separator",
        meta: {
          desc: "separator elements",
        },
        component: () => import("@/views/base/Separator.vue"),
      },
      {
        path: "tabs",
        name: "tabs",
        meta: {
          desc: "tabs elements",
        },
        component: () => import("@/views/base/Tabs.vue"),
      },
      {
        path: "breadcrumb",
        name: "breadcrumb",
        meta: {
          desc: "breadcrumb elements",
        },
        component: () => import("@/views/base/Breadcrumb.vue"),
      },
      {
        path: "modal",
        name: "modal",
        meta: {
          desc: "modal elements",
        },
        component: () => import("@/views/base/Modal.vue"),
      },
      {
        path: "pagination",
        name: "pagination",
        meta: {
          desc: "pagination elements",
        },
        component: () => import("@/views/base/Pagination.vue"),
      },
      {
        path: "vue-select",
        name: "vue-select",
        component: () => import("@/views/forms/VueSelect.vue"),
      },
      {
        path: "vee-validate",
        name: "vee-validate",
        component: () => import("@/views/forms/VeeValidate.vue"),
      },
      {
        path: "vue-currency-input",
        name: "vue-currency-input",
        component: () => import("@/views/forms/VueCurrencyInput.vue"),
      },
      {
        path: "tinymce",
        name: "tinymce",
        component: () => import("@/views/editors/TinyMCE.vue"),
      },
      {
        path: "fullcalendar",
        name: "fullcalendar",
        component: () => import("@/views/general/FullCalendar.vue"),
      },
      {
        path: "kt-datatables",
        name: "kt-datatables",
        component: () => import("@/views/general/KTDatatables.vue"),
      },
      {
        path: "element-ui",
        name: "element-ui",
        component: () => import("@/views/general/ElementUI.vue"),
      },
      {
        path: "icons/duotune",
        name: "duotune",
        meta: {
          desc: "duotune svg icons",
        },
        component: () => import("@/views/icons/Duotune.vue"),
      },
      {
        path: "icons/bootstrap-icons",
        name: "bootstrap-icons",
        meta: {
          desc: "free, high quality, open source icon library",
        },
        component: () => import("@/views/icons/BootstrapIcons.vue"),
      },
      {
        path: "icons/font-awesome",
        name: "font-awesome",
        meta: {
          desc: "awesome font icons",
        },
        component: () => import("@/views/icons/FontAwesome.vue"),
      },
      {
        path: "icons/line-awesome",
        name: "line-awesome",
        meta: {
          desc: "line font icons",
        },
        component: () => import("@/views/icons/LineAwesome.vue"),
      },
    ],
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/doc-overview",
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior(to) {
    if (to.hash) {
      return {
        el: to.hash,
        top: 10,
        behavior: "smooth",
      };
    }
  },
});

router.beforeEach((to) => {
  if (!to.hash) {
    // Scroll page to top on every route change
    setTimeout(() => {
      window.scrollTo(0, 0);
    }, 100);
  }
});

export default router;
