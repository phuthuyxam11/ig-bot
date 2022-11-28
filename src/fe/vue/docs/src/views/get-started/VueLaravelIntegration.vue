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
          Overview
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          <a href="https://laravel.com/">Laravel</a> is one of the most loved
          web frameworks by the developers. It makes development so easier and
          fun. <a href="https://vuejs.org/">VueJs</a> is one of the most loved
          JavaScript and rapidly growing Web Framework. This doc will cover
          steps on implement <a href="https://v3.vuejs.org/">Vue 3</a> in
          <a href="https://www.typescriptlang.org/">TypeScript</a> with
          <a href="https://laravel.com/docs/8.x/mix">Laravel</a>.
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="create_laravel_new_project">
          <router-link to="#create_laravel_new_project"></router-link>
          Create Laravel new project
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li>
            In the beginning we need to create Laravel+Vue project (in this
            example, we wll creat a project named <b>vue_laravel</b>) by running
            following command:

            <CodeHighlighter lang="bash">{{
              `composer create-project --prefer-dist laravel/laravel vue_laravel`
            }}</CodeHighlighter>
          </li>
          <li>
            Navigate to the newly created folder with command below.

            <CodeHighlighter lang="bash">{{
              `cd vue_laravel`
            }}</CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="install_dependencies">
          <router-link to="#install_dependencies"></router-link>
          Install dependencies
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li class="py-4">
            {{ appName }}'s 8 vue theme already contains almost all required
            dependencies for this integration, to get started we can copy list
            of dependencies from package.json in Vue and place them into
            <code>vue_laravel/package.json</code>. Also include devDependencies
            listed below:
            <CodeHighlighter lang="javascript">
              {{
                `"array-sort": "^1.0.0",
"sass": "^1.44.0",
"sass-loader": "^12.4.0",
"vuex-module-decorators": "^1.2.0"`
              }}
            </CodeHighlighter>
          </li>
          <li class="py-4">
            Now we are ready to install some additional dev dependencies for
            this integration:
            <ul class="py-4">
              <li>
                Let’s install Vue loader for webpack:
                <CodeHighlighter lang="bash">
                  {{
                    `npm install --save-dev vue-loader@next @vue/compiler-sfc`
                  }}
                </CodeHighlighter>
              </li>
              <li class="pb-4">
                Also we need to install ts-loader and typescript:

                <CodeHighlighter lang="bash">
                  {{ `npm install --save-dev ts-loader typescript` }}
                </CodeHighlighter>
              </li>
            </ul>
          </li>
          <li class="pb-4">
            Install all dependencies by running following command:

            <CodeHighlighter lang="bash">
              {{ `npm install` }}
            </CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="prepare_resource_files">
          <router-link to="#prepare_resource_files"></router-link>
          Prepare vue files and theme resources.
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li class="py-4">
            To use theme images we need to copy media folder from
            <code>public</code> and paste it into
            <code>vue_laravel/public</code>.
          </li>
          <li class="py-4">
            Create ts folder inside <code>vue_laravel/resources/ts</code> and
            inside this folder create file <code>app.ts</code>.
          </li>
          <li class="py-4">
            All other required files for vue project are placed inside
            <code>src</code> folder we can fully copy this folder and place it
            inside <code>vue_laravel/resources/ts</code>.
          </li>
          <li class="py-4">
            The main entry point of the vue project is <code>main.ts</code>, we
            need to import it in <code>app.ts</code>.
            <CodeHighlighter lang="javascript">
              {{ `import "./src/main";` }}
            </CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="setup_typescript_config">
          <router-link to="#setup_typescript_config"></router-link>
          Setup typescript config.
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li>
            TypeScript requires a configuration file to work properly. We can
            create one using <code>node_modules/.bin/tsc --init</code>. (no need
            for the full path if you have a global TypeScript installation) that
            appears in the newly created <code>tsconfig.json</code> file.
          </li>
          <li class="pt-10">
            Copy below typescript configuration and paste it into
            <code>tsconfig.json</code>
            in laravel app.

            <CodeHighlighter :field-height="400" lang="javascript">{{
              `{
  "compilerOptions": {
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,
    "noImplicitAny": false,
    "target": "es5",
    "module": "esnext",
    "strict": true,
    "jsx": "preserve",
    "importHelpers": true,
    "moduleResolution": "node",
    "skipLibCheck": true,
    "esModuleInterop": true,
    "allowSyntheticDefaultImports": true,
    "sourceMap": true,
    "baseUrl": ".",
    "paths": {
      "@/*": [
        "resources/ts/src/*"
      ]
    },
    "lib": [
      "esnext",
      "dom",
      "dom.iterable",
      "scripthost"
    ]
  },
  "include": [
    "resources/ts/src/**/*.ts",
    "resources/ts/src/**/*.tsx",
    "resources/ts/src/**/*.vue"
  ],
  "exclude": [
    "node_modules"
  ]
}`
            }}</CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="configure_webpack_mix">
          <router-link to="#configure_webpack_mix"></router-link>
          Configure webpack.mix.js
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li>
            To use Vue codebase in Laravel we need to bundle app.ts, copy the
            following webpack.mix.js configuration example and paste into
            <code>webpack.mix.js</code> in vue_laravel project.
            <CodeHighlighter lang="javascript">
              {{
                `const mix = require("laravel-mix");
const path = require("path");

mix.ts("resources/ts/app.ts", "public/js")
    .vue({ version: 3 })
    .webpackConfig({
        module: {
            rules: [
                {
                    test: /\.mjs$/i,
                    resolve: {
                        byDependency: { esm: { fullySpecified: false } },
                    },
                },
            ],
        },
        resolve: {
            alias: {
                "@": path.resolve(__dirname, "resources/ts/src/"),
            },
        },
    });`
              }}
            </CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="build_client">
          <router-link to="#build_client"></router-link>
          Build client
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li>
            We are ready to build client assets to do it run the following
            command:

            <CodeHighlighter lang="bash">{{ `npm run dev` }}</CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="prepare_app_blade_file">
          <router-link to="#prepare_app_blade_file"></router-link>
          Prepare app.blade.php
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <ol class="py-4">
          <li class="pb-5">
            Now we can add a new file inside <code>resources/views</code>. Let's
            call it <code>app.blade.php</code>
          </li>
          <li>
            And add inside <code>app.blade.php</code> file this code.
            <CodeHighlighter lang="html">{{ htmlCode }}</CodeHighlighter>
          </li>
          <li>
            Now we can add new route into file <code>routes/web.php</code>
            <CodeHighlighter lang="javascript">{{
              `Route::get('/', function () {
    return view('app');
});`
            }}</CodeHighlighter>
          </li>
        </ol>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="run_application">
          <router-link to="#run_application"></router-link>
          Run Application
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          <ol>
            <li>
              To run vue+laravel application in your terminal, run the following
              command:
              <CodeHighlighter lang="bash">{{
                `php artisan serve`
              }}</CodeHighlighter>
            </li>
            <li>
              Wait for the app to display that it's listening on
              <code
                ><a href="http://127.0.0.1:8000">http://127.0.0.1:8000</a></code
              >
              and open this url in your browser.
            </li>
          </ol>
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
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "vue-laravel-integration",
  components: {
    CodeHighlighter,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Vue + Laravel integration", [
        "Getting Started",
      ]);
    });

    const htmlCode =
      `<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
   <meta charset="utf-8">
   <meta name="viewport" content="width=device-width, initial-scale=1">
   <meta name="csrf-token" value="{{ csrf_token() }}"/>
   <title>Laravel Vue Example</title>
   <link href="https://fonts.googleapis.com/css?family=Poppins:300,400,500,600,700|Roboto:300,400,500,600,700|Material+Icons" rel="stylesheet">
</head>
<body style="--kt-toolbar-height:55px;--kt-toolbar-height-tablet-and-mobile:55px">
   <div id="app" class="d-flex flex-column flex-root">
   </div>
   <script src="{{ mix('js/app.js') }}" type="text/javascript"></scr` +
      `ipt>
</body>
</html>`;

    return {
      htmlCode,
      appName,
    };
  },
});
</script>
