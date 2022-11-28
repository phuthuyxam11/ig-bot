<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Notice>
        <a href="https://docs.npmjs.com/about-npm">NPM</a> is the recommended
        installation method when building large scale applications with Vue. It
        pairs nicely with module bundlers such as Webpack or Browserify. Vue
        also provides accompanying tools for authoring Single File Components.
      </Notice>
      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="quick_start">
          <router-link to="#quick_start"></router-link>
          Quick Start
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-5">
          <li class="pt-0">
            Start a command prompt window or terminal and change directory to
            <code>[unpacked path]/vue/[demo]</code>.
          </li>
          <div class="pb-5">
            <CodeHighlighter lang="bash">{{
              `cd [unpacked path]/vue/[demo]`
            }}</CodeHighlighter>
          </div>
          <li class="pt-0">
            Install the Vue local project dependencies by executing
            <code>npm install</code> or <code>yarn</code>
            command. This command will install the dependencies in the local
            <code>node_modules</code> folder these dependencies are listed in
            <code>package.json</code>.
          </li>
          <div class="pb-5">
            <CodeHighlighter lang="bash">{{
              `npm install
# OR
yarn`
            }}</CodeHighlighter>
          </div>
          <div class="pb-5">
            Execute <code>npm run serve</code> or
            <code>yarn serve</code> command to launch the server, watches your
            files, and rebuilds the app as you make changes to those files. Open
            your browser on
            <a href="http://localhost:8080/">http://localhost:8080/</a>.
          </div>
          <div class="pb-5">
            <CodeHighlighter lang="bash">{{
              `npm run serve
# OR
yarn serve`
            }}</CodeHighlighter>
          </div>
          <div class="pb-5">
            This command will take some time to finish. Please wait for it to
            complete.
          </div>
        </ol>
        <div class="pb-5">
          For more detailed information to install the Vue application, visit
          this official Vue
          <a href="https://v3.vuejs.org/guide/installation.html#installation"
            >documentation website</a
          >.
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
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import Notice from "@/components/sections/Notice.vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "get-started-build",
  components: {
    CodeHighlighter,
    Notice,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Build", ["Getting Started"]);
    });

    return {
      appName,
    };
  },
});
</script>
