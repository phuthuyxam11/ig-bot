<template>
  <!--begin::Aside-->
  <div
    id="kt_aside"
    class="aside aside-hoverable"
    :class="[
      asideTheme === 'light' && 'aside-light',
      asideTheme === 'dark' && 'aside-dark',
    ]"
    data-kt-drawer="true"
    data-kt-drawer-name="aside"
    data-kt-drawer-activate="{default: true, lg: false}"
    data-kt-drawer-overlay="true"
    data-kt-drawer-width="{default:'200px', '300px': '250px'}"
    data-kt-drawer-direction="start"
    data-kt-drawer-toggle="#kt_aside_mobile_toggle"
  >
    <!--begin::Brand-->
    <div class="aside-logo flex-column-auto" id="kt_aside_logo">
      <!--begin::Logo-->
      <a href="#" v-if="asideTheme === 'dark'">
        <img
          alt="Logo"
          src="media/logos/default-dark.svg"
          class="h-25px app-sidebar-logo-default"
        />
      </a>
      <a href="#" v-if="asideTheme === 'light'">
        <img
          v-if="themeMode === 'light'"
          alt="Logo"
          src="media/logos/default.svg"
          class="theme-light-show h-25px app-sidebar-logo-default"
        />
        <img
          v-else
          alt="Logo"
          src="media/logos/default-dark.svg"
          class="theme-dark-show h-25px app-sidebar-logo-default"
        />
      </a>
      <!--end::Logo-->

      <!--begin::Aside toggler-->
      <div
        id="kt_aside_toggle"
        class="btn btn-icon w-auto px-0 btn-active-color-primary aside-toggle"
        data-kt-toggle="true"
        data-kt-toggle-state="active"
        data-kt-toggle-target="body"
        data-kt-toggle-name="aside-minimize"
      >
        <span class="svg-icon svg-icon-1 rotate-180">
          <inline-svg src="media/icons/duotune/arrows/arr080.svg" />
        </span>
      </div>
      <!--end::Aside toggler-->
    </div>
    <!--end::Brand-->

    <!--begin::Aside menu-->
    <div class="aside-menu flex-column-fluid">
      <KTMenu></KTMenu>
    </div>
    <!--end::Aside menu-->

    <!--begin::Footer-->
    <div
      class="aside-footer flex-column-auto pt-5 pb-7 px-5"
      id="kt_aside_footer"
    >
      <a
        href="https://preview.keenthemes.com/metronic8/vue/docs/#/doc-overview"
        class="btn btn-custom btn-primary w-100"
        data-bs-toggle="tooltip"
        data-bs-trigger="hover"
        data-bs-delay-show="8000"
        title="Check out the complete documentation with over 100 components"
      >
        <span class="btn-label">
          {{ t("docsAndComponents") }}
        </span>
        <span class="svg-icon btn-icon svg-icon-2">
          <inline-svg src="media/icons/duotune/general/gen005.svg" />
        </span>
      </a>
    </div>
    <!--end::Footer-->
  </div>
  <!--end::Aside-->
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { useI18n } from "vue-i18n/index";
import KTMenu from "@/layout/aside/Menu.vue";
import { asideTheme } from "@/core/helpers/config";
import { useStore } from "vuex";

export default defineComponent({
  name: "KTAside",
  components: {
    KTMenu,
  },
  props: {
    lightLogo: String,
    darkLogo: String,
  },
  setup() {
    const { t } = useI18n();
    const store = useStore();

    const themeMode = computed(() => {
      return store.getters.getThemeMode;
    });

    return {
      asideTheme,
      themeMode,
      t,
    };
  },
});
</script>
