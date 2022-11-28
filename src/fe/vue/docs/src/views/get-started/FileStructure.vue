<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="file-structure">
          <router-link to="#file-structure"></router-link>
          File structure
        </h1>
        <!--end::Heading-->

        <div class="py-5">
          <table
            class="table table-hover table-rounded table-striped border gy-7 gs-7"
          >
            <thead>
              <tr>
                <th>Path</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>/<strong>dist</strong></td>
                <td>
                  The built output, run command <code>yarn build</code> to
                  create it.
                </td>
              </tr>
              <tr>
                <td>/<strong>node_modules</strong></td>
                <td>
                  The <code>package.json</code> file in the app root defines
                  what libraries will be installed into
                  <code>node_modules/</code> when you run
                  <code>yarn install</code>.
                </td>
              </tr>
              <tr>
                <td>/<strong>public</strong></td>
                <td>
                  Folder contains the <code>index.html</code> file so you can
                  tweak it, change<code>&lt;title&gt;</code>, insert
                  additional<code>&lt;link&gt;</code> and
                  <code>&lt;script&gt;</code> tags. For more details check
                  <a
                    target="_blank"
                    href="https://cli.vuejs.org/guide/html-and-static-assets.html#the-index-file"
                    >official documentation</a
                  >.
                </td>
              </tr>
              <tr>
                <td>/public/<strong>media</strong></td>
                <td>Folder contains all image, icons and SVG files.</td>
              </tr>
              <tr>
                <td>/src/assets/sass/<strong>style.scss</strong></td>
                <td>
                  Main style file, imports all scss file used for Metronic
                  layout and components. This style file is imported inside
                  <code>/src/App.vue</code>.
                </td>
              </tr>
              <tr>
                <td>/src/assets/<strong>ts</strong></td>
                <td>
                  Contains Metronic common ts components. These ts component
                  files are used to initialize theme interactive parts, such as
                  drawer, scroll, menu etc., you can check initialization
                  example in file <code>/src/core/keenthemes.ts</code>.
                </td>
              </tr>
              <tr>
                <td>/src/<strong>components</strong></td>
                <td>
                  Folder contains all components used in the demo pages (folder
                  <code>/src/views</code>).
                </td>
              </tr>
              <tr>
                <td>
                  /src/core/config/<strong>DefaultLayoutConfig.ts</strong>
                </td>
                <td>
                  Layout configuration, you can update layout default values or
                  update and preview them dynamically in theme
                  <a
                    href="https://preview.keenthemes.com/metronic8/vue/demo1/#/builder"
                    >Layout Builder</a
                  >.
                </td>
              </tr>
              <tr>
                <td>/src/core/config/<strong>MainMenuConfig.ts</strong></td>
                <td>
                  Menu configuration, to get started with clean menu you can use
                  <code>/src/core/config/CleanMainMenuConfig.ts</code>.
                </td>
              </tr>
              <tr>
                <td>/src/core/<strong>helpers</strong></td>
                <td>Contains files with theme helper functions.</td>
              </tr>
              <tr>
                <td>/src/core/plugins/<strong>keenthemes.ts</strong></td>
                <td>File contains ts component initialization functions.</td>
              </tr>
              <tr>
                <td>/src/core/services/<strong>ApiService.ts</strong></td>
                <td>
                  Class which can be used to make a http requests to your api,
                  you can find usage examples in
                  <code>/src/store/AuthModule.ts</code> for more info check
                  <router-link to="/api-setup">Api setup doc</router-link>.
                </td>
              </tr>
              <tr>
                <td>/src/core/services/<strong>JwtService.ts</strong></td>
                <td>
                  Class used to get, save and destroy jwt token form
                  localStorage.
                </td>
              </tr>
              <tr>
                <td>/src/core/services/<strong>LayoutService.ts</strong></td>
                <td>Class used to initialize theme layout.</td>
              </tr>
              <tr>
                <td>/src/<strong>layout</strong></td>
                <td>Folder contains all layout vue components.</td>
              </tr>
              <tr>
                <td>/src/<strong>router</strong></td>
                <td>
                  Folder contains <code>index.ts</code> which is default
                  vue-router config used by theme, to get started with clean
                  router use clean.ts config. For more info about the
                  vue-router, check
                  <a href="https://router.vuejs.org/">vue-router official doc</a
                  >.
                </td>
              </tr>
              <tr>
                <td>/src/store/<strong>modules</strong></td>
                <td>Folder contains all vuex modules.</td>
              </tr>
              <tr>
                <td>/src/<strong>views</strong></td>
                <td>Vue components for theme demo pages</td>
              </tr>
              <tr>
                <td>/src/<strong>App.vue</strong></td>
                <td>
                  Core vue component, contains vue root router-view and all
                  style imports
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!--end::Section-->
    </div>
    <!--end::Card Body-->
  </div>
  <!--begin::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "setup-theme-skeleton",
  components: {},
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("File structure", ["Getting Started"]);
    });

    return {
      appName,
    };
  },
});
</script>
