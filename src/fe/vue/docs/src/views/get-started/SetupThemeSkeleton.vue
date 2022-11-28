<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="setup_skeleton">
          <router-link to="#setup_skeleton"></router-link>
          Setup skeleton
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          To get started with theme skeleton you should change router path to
          new path for clean router file, which is already added inside
          <code>src/router</code> folder. Change router path inside file main.js
          to the following path:
        </div>
        <!--end::Block-->

        <div class="pb-5">
          <CodeHighlighter lang="javascript">{{
            `import router from "@/router/clean.ts";`
          }}</CodeHighlighter>
        </div>

        <!--begin::Block-->
        <div class="py-5">
          Second thing what you need to change is our menus, in {{ appName }} we
          have two types of menu in header and aside bar.
        </div>
        <!--end::Block-->

        <div class="py-5">
          <!--begin::Information-->
          <div
            class="d-flex align-items-center rounded py-5 px-5 bg-light-info"
          >
            <span class="svg-icon svg-icon-3x svg-icon-info me-5">
              <inline-svg src="media/icons/duotune/general/gen044.svg" />
            </span>
            <div class="text-gray-600 fw-bold fs-5">
              Main menu is generated from the file
              <code>src/core/config/MainMenuConfig.ts</code>. Header menu is
              generated inside
              <code>src/layout/header/Menu.vue</code>
              and the aside menu is generated inside
              <code>src/layout/aside/Menu.vue</code>.
            </div>
            <!--end::Description-->
          </div>
          <!--end::Information-->
        </div>

        <div class="py-5">
          To get started with clean menu, you can update menu config path in
          both <code>Menu.vue</code> files from
          <code>src/core/config/MainMenuConfig.ts</code>
          to the following code:
        </div>

        <div class="pb-5">
          <CodeHighlighter lang="javascript">{{
            `import MainMenuConfig from "@/core/config/CleanMainMenuConfig";`
          }}</CodeHighlighter>
        </div>

        <!--begin::Block-->
        <div class="py-5">
          Now you can delete unneeded page view files from folder
          <code>src/views</code>
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->
    </div>
    <!--end::Card Body-->
  </div>
  <!--begin::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "setup-theme-skeleton",
  components: {
    CodeHighlighter,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Setup theme skeleton", ["Getting Started"]);
    });

    return {
      appName,
    };
  },
});
</script>
