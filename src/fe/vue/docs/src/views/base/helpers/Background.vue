<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="bg-colors">
          <router-link to="#bg-colors"></router-link>
          Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Easily set the background of an element to any contextual class as
          shown below:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template v-for="(item, i) in colors.backgroundColors" :key="i">
              <p
                :class="`bg-${item} text-inverse-${item}`"
                class="p-3 fw-bold fw-6"
              >
                .bg-{{ item }}
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-lighten"></div>
            <div class="bg-body"></div>
            <div class="bg-primary"></div>
            <div class="bg-light"></div>
            <div class="bg-secondary"></div>
            <div class="bg-success"></div>
            <div class="bg-info"></div>
            <div class="bg-warning"></div>
            <div class="bg-danger"></div>
            <div class="bg-dark"></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="light-bg-colors">
          <router-link to="#light-bg-colors"></router-link>
          Light Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.bg-light-{color}</code> class to set an element's light
          background color defined with <code>$theme-light-colors</code> in
          <code>src/assets/sass/core/components/_variables.scss</code>:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template
              v-for="(item, i) in colors.lightBackgroundColors"
              :key="i"
            >
              <p
                :class="`bg-light-${item} text-${item}`"
                class="p-3 fw-bold fw-6"
              >
                .bg-light-{{ item }}
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-light-primary text-primary"></div>
            <div class="bg-light-success text-success"></div>
            <div class="bg-light-info text-info"></div>
            <div class="bg-light-warning text-warning"></div>
            <div class="bg-light-danger text-danger"></div>
            <div class="bg-light-dark text-dark"></div>
            `}}
          </CodeHighlighter>
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="active-bg-color">
          <router-link to="#active-bg-color"></router-link>
          Active Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.bg-active-{color}</code>class to toggle an element's active
          background color state with <code>.active</code>class:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template v-for="(item, i) in colors.backgroundColors" :key="i">
              <p
                :class="`bg-active-${item} text-active-inverse-${item} active`"
                class="p-3 fw-bold fw-6"
              >
                .bg-light-{{ item }}
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!-- prettier-ignore -->
          <CodeHighlighter lang="html">{{
`<div class="bg-active-lighten text-active-inverse-lighten active"></div>
<div class="bg-active-white text-active-inverse-white active"></div>
<div class="bg-active-primary text-active-inverse-primary active"></div>
<div class="bg-active-light text-active-inverse-light active"></div>
<div class="bg-active-secondary text-active-inverse-secondary active"></div>
<div class="bg-active-success text-active-inverse-success active"></div>
<div class="bg-active-info text-active-inverse-info active"></div>
<div class="bg-active-warning text-active-inverse-warning active"></div>
<div class="bg-active-danger text-active-inverse-danger active"></div>
<div class="bg-active-dark text-active-inverse-dark active"></div>`}}
          </CodeHighlighter>
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="hover-bg-colors">
          <router-link to="#hover-bg-colors"></router-link>
          Hover Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.bg-hover-{color}</code> class to set an element's hover
          background color:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template v-for="(item, i) in colors.backgroundColors" :key="i">
              <p
                :class="`bg-hover-${item}`"
                class="text-gray-600 p-3 fw-bold fw-6"
              >
                .bg-hover-{{ item }}
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!-- prettier-ignore -->
          <CodeHighlighter lang="html"
            >{{`<div class="bg-hover-lighten text-hover-inverse-lighten"></div>
<div class="bg-hover-white text-hover-inverse-white"></div>
<div class="bg-hover-primary text-hover-inverse-primary"></div>
<div class="bg-hover-light text-hover-inverse-light"></div>
<div class="bg-hover-secondary text-hover-inverse-secondary"></div>
<div class="bg-hover-success text-hover-inverse-success"></div>
<div class="bg-hover-info text-hover-inverse-info"></div>
<div class="bg-hover-warning text-hover-inverse-warning"></div>
<div class="bg-hover-danger text-hover-inverse-danger"></div>
<div class="bg-hover-dark text-hover-inverse-dark"></div>`}}
          </CodeHighlighter>
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="hover-light-bg-colors">
          <router-link to="#hover-light-bg-colors"></router-link>
          Hover Light Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.bg-hover-light-{color}</code> class to set an element's
          hover light background color:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template
              v-for="(item, i) in colors.lightBackgroundColors"
              :key="i"
            >
              <p
                :class="`bg-hover-light-${item}`"
                class="text-gray-600 p-3 fw-bold fw-6"
              >
                .bg-hover-light-{{ item }}
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-hover-light-primary text-hover-primary"></div>
            <div class="bg-hover-light-success text-hover-success"></div>
            <div class="bg-hover-light-info text-hover-info"></div>
            <div class="bg-hover-light-warning text-hover-warning"></div>
            <div class="bg-hover-light-danger text-hover-danger"></div>
            <div class="bg-hover-light-dark text-hover-dark"></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="hoverable-bg-colors">
          <router-link to="#hoverable-bg-colors"></router-link>
          Hoverable Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.hoverable</code> with <code>.bg-{color}</code> class to set
          an element's background color with hover state:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template v-for="(item, i) in colors.backgroundColors" :key="i">
              <p
                :class="`bg-${item}`"
                class="text-gray-600 hoverable p-3 fw-bold fw-6"
              >
                .bg-{{ item }} .hoverable
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-body hoverable"></div>
            <div class="bg-primary hoverable"></div>
            <div class="bg-light hoverable"></div>
            <div class="bg-secondary hoverable"></div>
            <div class="bg-success hoverable"></div>
            <div class="bg-info hoverable"></div>
            <div class="bg-warning hoverable"></div>
            <div class="bg-danger hoverable"></div>
            <div class="bg-dark hoverable"></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="hoverable-light-bg-colors">
          <router-link to="#hoverable-light-bg-colors"></router-link>
          Hoverable Light Background Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.hoverable</code> with <code>.bg-light-{color}</code> class
          to set an element's light background color with hover state:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <template
              v-for="(item, i) in colors.lightBackgroundColors"
              :key="i"
            >
              <p
                :class="`bg-light-${item}`"
                class="text-gray-600 hoverable p-3 fw-bold fw-6"
              >
                .bg-light-{{ item }} .hoverable
              </p>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-light-primary hoverable"></div>
            <div class="bg-light-success hoverable"></div>
            <div class="bg-light-info hoverable"></div>
            <div class="bg-light-warning hoverable"></div>
            <div class="bg-light-danger hoverable"></div>
            <div class="bg-light-dark hoverable"></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="bg-opacity">
          <router-link to="#bg-opacity"></router-link>
          Background Opacity
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.bg-opacity-{0, 5, 10, 15, 20, 25, 50, 75, 100}</code> class
          with <code>.bg-{color}</code> and <code>.bg-light-{color}</code> to
          set an element's background color with opacity:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10 d-flex flex-wrap">
            <template v-for="(item, i) in colors.opacity" :key="i">
              <div class="d-flex flex-column me-5 mb-5">
                <div
                  :class="`bg-opacity-${item}`"
                  class="w-90px h-90px bg-primary flex-grow-1"
                ></div>
                <div class="fw-bold py-2">bg-opacity-{{ item }}</div>
              </div>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="bg-opacity-0 bg-primary"></div>
            <div class="bg-opacity-5 bg-primary"></div>
            <div class="bg-opacity-10 bg-primary"></div>
            <div class="bg-opacity-15 bg-primary"></div>
            <div class="bg-opacity-20 bg-primary"></div>
            <div class="bg-opacity-25 bg-primary"></div>
            <div class="bg-opacity-50 bg-primary"></div>
            <div class="bg-opacity-75 bg-primary"></div>
            <div class="bg-opacity-100 bg-primary"></div>
            ` }}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="hover-bg-opacity">
          <router-link to="#hover-bg-opacity"></router-link>
          Hover Background Opacity
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use
          <code>.bg-hover-opacity-{0, 5, 10, 15, 20, 25, 50, 75, 100}</code
          >class with <code>.bg-{color}</code>or
          <code>.bg-light-{color}</code>to set an element's background color
          with opacity on hover:
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10 d-flex flex-wrap">
            <template v-for="(item, i) in colors.opacity" :key="i">
              <div class="d-flex flex-column me-5 mb-5">
                <div
                  :class="`bg-hover-opacity--${item}`"
                  class="w-90px h-90px bg-primary flex-grow-1"
                ></div>
                <div class="fw-bold py-2">bg-hover-opacity-{{ item }}</div>
              </div>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
            >{{`<div class="bg-hover-opacity-0 bg-primary"></div>
<div class="bg-hover-opacity-5 bg-primary"></div>
<div class="bg-hover-opacity-10 bg-primary"></div>
<div class="bg-hover-opacity-15 bg-primary"></div>
<div class="bg-hover-opacity-20 bg-primary"></div>
<div class="bg-hover-opacity-25 bg-primary"></div>
<div class="bg-hover-opacity-50 bg-primary"></div>
<div class="bg-hover-opacity-75 bg-primary"></div>
<div class="bg-hover-opacity-100 bg-primary"></div>` }}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="active-bg-opacity">
          <router-link to="#active-bg-opacity"></router-link>
          Active Background Opacity
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="py-5">
            Use
            <code>.bg-active-opacity-{0, 5, 10, 15, 20, 25, 50, 75, 100}</code
            >class with <code>.bg-{color}</code>or
            <code>.bg-light-{color}</code>to set an element's background color
            with opacity on active state:
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10 d-flex flex-wrap">
            <template v-for="(item, i) in colors.opacity" :key="i">
              <div class="d-flex flex-column me-5 mb-5">
                <div
                  :class="`bg-active-opacity-${item}`"
                  class="w-90px h-90px bg-primary active flex-grow-1"
                ></div>
                <div class="fw-bold py-2">bg-active-opacity-{{ item }}</div>
              </div>
            </template>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
            >{{`<div class="bg-active-opacity-0 bg-primary active"></div>
<div class="bg-active-opacity-5 bg-primary active"></div>
<div class="bg-active-opacity-10 bg-primary active"></div>
<div class="bg-active-opacity-15 bg-primary active"></div>
<div class="bg-active-opacity-20 bg-primary active"></div>
<div class="bg-active-opacity-25 bg-primary active"></div>
<div class="bg-active-opacity-50 bg-primary active"></div>
<div class="bg-active-opacity-75 bg-primary active"></div>
<div class="bg-active-opacity-100 bg-primary active"></div>` }}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pt-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="bg-image-classes">
          <router-link to="#bg-image-classes"></router-link>
          Background Image Classes
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Assign background image style values to an element with shorthand
          classes with the format <code>bgi-{property}</code>.
        </div>
        <!--end::Block-->

        <div class="py-5">
          Where <code>property</code>is one of:
          <ul class="pt-3">
            <li class="py-2">
              <code>no-repeat</code>to set
              <code>background-repeat: no-repeat</code>
            </li>
            <li class="py-2">
              <code>position-y-top</code>to set
              <code>background-position-y: top</code>
            </li>
            <li class="py-2">
              <code>position-y-bottom</code>to set
              <code>background-position-y: bottom</code>
            </li>
            <li class="py-2">
              <code>position-y-center</code>to set
              <code>background-position-y: center</code>
            </li>
            <li class="py-2">
              <code>position-x-start</code>to set
              <code>background-position-x: left</code>
            </li>
            <li class="py-2">
              <code>position-x-end</code>to set
              <code>background-position-x: right</code>
            </li>
            <li class="py-2">
              <code>position-x-center</code>to set
              <code>background-position-x: center</code>
            </li>
            <li class="py-2">
              <code>position-bottom</code>to set
              <code>background-position: 0 bottom</code>
            </li>
            <li class="py-2">
              <code>position-center</code>to set
              <code>background-position: center</code>
            </li>
          </ul>
        </div>

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div
              style="background-image: url(image.png)"
              class="bgi-size-contain bgi-no-repeat bgi-position-center"
            ></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pt-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="bg-image-size-classes">
          <router-link to="#bg-image-size-classes"></router-link>
          Background Image Size Classes
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use responsive-friendly background image size style values for an
          element with shorthand classes with the format
          <code>.bgi-size-{breakpoint}-{property}</code>. The classes are named
          using the format <code>.bgi-size-{property}</code>for
          <code>xs</code>and <code>.bgi-size-{breakpoint}-{property}</code>for
          <code>sm</code>, <code>md</code>, <code>lg</code>, <code>xl</code>,
          and <code>xxl</code>.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          Where <code>property</code>is one of:
          <ul class="pt-3">
            <li class="py-2">
              <code>auto</code>to set <code>background-size: auto</code>
            </li>
            <li class="py-2">
              <code>cover</code>to set <code>background-size: cover</code>
            </li>
            <li class="py-2">
              <code>contain</code>to set <code>background-size: contain</code>
            </li>
          </ul>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <CodeHighlighter lang="html"
            >{{`
            <div
              style="background-image: url(image.png)"
              class="bgi-size-contain bgi-size-lg-auto bgi-no-repeat bgi-position-center"
            ></div>
            `}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "helpers-background",
  components: {
    CodeHighlighter,
  },
  setup() {
    const colors = ref({
      backgroundColors: [
        "lighten",
        "white",
        "primary",
        "light",
        "secondary",
        "success",
        "info",
        "warning",
        "danger",
        "dark",
      ],
      lightBackgroundColors: [
        "primary",
        "success",
        "info",
        "warning",
        "danger",
        "dark",
      ],
      opacity: [0, 5, 10, 15, 20, 25, 50, 75, 100],
    });

    onMounted(() => {
      setCurrentPageBreadcrumbs("Background", ["Base", "Helpers"]);
    });

    return {
      colors,
    };
  },
});
</script>
