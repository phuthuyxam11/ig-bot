<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="overview">
          <router-link to="#overview"></router-link>
          Updating &nbsp;{{ appName }}
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          Like all HTML templates, you will need to manually handle all updates
          that we provide after purchase for free. The best practice is to
          always keep
          {{ appName }}'s source folder on an entire separate folder within your
          project and then use {{ appName }} as a development companion or
          support.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-4">
          Each update should be used independently without mixing it with the
          previous versions. This is because new updates may not have a 100%
          backward compatibility for elements within previous versions.
          Therefore, each update is considered as a new and better version of
          the theme.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="pt-4">
          Here's our recommended steps needed to update
          {{ appName }} effectively:

          <ol class="py-3">
            <li class="py-2">
              Backup all your files (your project folder and {{ appName }}'s's
              older version).
            </li>
            <li class="py-2">
              Review {{ appName }}s changelog to have a quick preview on what
              has been updated.
            </li>
            <li class="py-2">
              Unzip the latest version of {{ appName }} and then browse through
              the elements (layout, widgets, plugins, etc) that you have used on
              your project.
            </li>
            <li class="py-2">
              Get the latest version use it instead the previous version.
            </li>
            <li class="py-2">
              Run <code>yarn install</code> task to update 3rd party plugins
              with their latest versions.
            </li>
            <li class="py-2">
              Compare your project's code vs the latest version of
              {{ appName }}s code for the used elements.
            </li>
            <li class="py-2">
              Manually replace the element HTML, CSS or JS where necessary.
            </li>
            <li class="py-2">
              Save and test your project files with {{ appName }}'s updated
              elements.
            </li>
          </ol>

          If something goes wrong, you can revert your updates by reusing your
          backup files.
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "get-started-updates",
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Updates", ["Getting Started"]);
    });

    return {
      appName,
    };
  },
});
</script>
