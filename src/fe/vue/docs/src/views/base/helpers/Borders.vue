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
        <div class="py-5">
          <b>{{ appName }}</b> extends
          <a
            href="https://getbootstrap.com/docs/5.0/utilities/borders/"
            class="fw-bold"
            >Bootstrap Borders</a
          >
          in
          <code>src/assets/sass/core/components/helpers/_borders.scss</code> to
          provide additional border utility classes to cover its own design
          system.
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="border-styles">
          <router-link to="#border-styles"></router-link>
          Border Styles
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.border-dashed</code>or <code>.border-dotted</code>classes
          to set an element's border style. For border directions use
          <code>.border-{direction}-dashed</code>or
          <code>.border-{direction}-dotted</code>classes format.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          Where <code>direction</code>is one of:
          <ul class="pt-3">
            <li class="py-2">
              <code>top</code>- for classes that set border top style
            </li>
            <li class="py-2">
              <code>bottom</code>- for classes that set border bottom style
            </li>
            <li class="py-2">
              <code>end</code>- for classes that set border right style
            </li>
            <li class="py-2">
              <code>start</code>- for classes that set border start style
            </li>
          </ul>
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="border w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                default style
              </div>
              <div
                class="border-gray-300 border-dotted w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .border-dotted
              </div>
              <div
                class="border-gray-300 border-dashed w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .border-dashed
              </div>
              <div
                class="border-gray-300 border-bottom-dashed w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .border-bottom-dashed
              </div>
              <div
                class="border-gray-300 border-end-dashed w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .border-end-dashed
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
            >{{`<div class="border">
            default style
          </div>
            <div class="border-gray-300 border-dotted">
              .border-dotted
            </div>
            <div class="border-gray-300 border-dashed">
              .border-dashed
            </div>
            <div class="border-gray-300 border-bottom-dashed">
              .border-bottom-dashed
            </div>
            <div class="border-gray-300 border-end-dashed">
              .border-end-dashed
            </div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="hover-state">
          <router-link to="#hover-state"></router-link>
          Hover State
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.border-hover</code>to display an element's border on hover.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="border border-hover w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                default
              </div>
              <div
                class="border border-gray-500 border-hover w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                gray-500
              </div>
              <div
                class="border border-primary border-hover w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                primary
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
            >{{`<div class="border border-hover">
  default
</div>
<div class="border border-gray-500 border-hover">
  gray-500
</div>
<div class="border border-primary border-hover">
  primary
</div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="active-state">
          <router-link to="#active-state"></router-link>
          Active State
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.border-active</code>to display an element's border when
          <code>.active</code>class applied.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="border border-active active w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                default
              </div>
              <div
                class="border border-gray-500 border-active active w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                gray-500
              </div>
              <div
                class="border border-primary border-active active w-125px h-125px d-flex flex-center me-5 mb-5 fw-bold"
              >
                primary
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<div class="border border-active active">
  default
</div>
<div class="border border-gray-500 border-active active">
  gray-500
</div>
<div class="border border-primary border-active active">
  primary
</div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="transparent-state">
          <router-link to="#transparent-state"></router-link>
          Transparent Border
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use <code>.border-transparent</code>to set an element's border color
          to tranparent.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="border border-transparent bg-light w-125px h-125px d-flex flex-center me-5 fw-bold"
              >
                tranparent <br />border
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<div class="border border-transparent">
  default
</div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="pt-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="border-gray-colors">
          <router-link to="#border-gray-colors"></router-link>
          Border Gray Colors
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use
          <code
            >.border-gray-{100, 200, 300, 400, 500, 600, 700, 800, 900}</code
          >
          to set an element's border color.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <template v-for="(item, i) in gray" :key="i">
                <div
                  :class="`border-${item}`"
                  class="w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold border"
                >
                  .border-{{ item }}
                </div>
              </template>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="pt-5">
          <CodeHighlighter lang="html"
            >{{`
            <div class="border border-gray-100"></div>
            <div class="border border-gray-200"></div>
            <div class="border border-gray-300"></div>
            <div class="border border-gray-400"></div>
            <div class="border border-gray-500"></div>
            <div class="border border-gray-600"></div>
            <div class="border border-gray-700"></div>
            <div class="border border-gray-800"></div>
            <div class="border border-gray-900"></div>
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
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "helpers-borders",
  components: {
    CodeHighlighter,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Borders", ["Base", "Helpers"]);
    });

    return {
      appName,
    };
  },
});
</script>
