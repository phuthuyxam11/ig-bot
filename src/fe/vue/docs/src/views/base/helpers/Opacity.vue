<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="pb-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="overview">
          <router-link to="#overview"></router-link>
          Overview
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          {{ appName }} sets custom opacity classes in
          <code>src/assets/sass/core/components/helpers/_opacity.scss</code>and
          through a SASS variable <code>$opacity-values</code>defined in
          <code>src/assets/sass/core/components//_variables.scss</code>to
          provide additional opacity utility classes to cover its own design
          system.
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="opacity">
          <router-link to="#opacity"></router-link>
          Opacity
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use the
          <code>.opacity-{ 0, 5, 10, 15 ,20, 25, 50, 75, 100 }</code>class to
          set an element's opacity style.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="opacity-0 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-0
              </div>
              <div
                class="opacity-5 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-5
              </div>
              <div
                class="opacity-10 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-10
              </div>
              <div
                class="opacity-15 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-15
              </div>
              <div
                class="opacity-20 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-20
              </div>
              <div
                class="opacity-25 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-25
              </div>
              <div
                class="opacity-50 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-50
              </div>
              <div
                class="opacity-75 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-75
              </div>
              <div
                class="opacity-100 bg-primary text-white w-100px h-100px d-flex flex-center me-5 mb-5 fw-bold"
              >
                .opacity-100
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
            >{{`<div class="opacity-0"></div>
<div class="opacity-5"></div>
<div class="opacity-10"></div>
<div class="opacity-15"></div>
<div class="opacity-20"></div>
<div class="opacity-25"></div>
<div class="opacity-50"></div>
<div class="opacity-75"></div>
<div class="opacity-100"></div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="opacity-hover-effects">
          <router-link to="#opacity-hover-effects"></router-link>
          Opacity Hover Effects
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use
          <code>.opacity-{ 0, 5, 10, 15 ,20, 25, 50, 75, 100 }-hover</code>to
          display an element's opacity on hover.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center">
              <div
                class="opacity-0-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-0-hover
              </div>
              <div
                class="opacity-5-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-5-hover
              </div>
              <div
                class="opacity-10-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-10-hover
              </div>
              <div
                class="opacity-15-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-15-hover
              </div>
              <div
                class="opacity-20-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-20-hover
              </div>
              <div
                class="opacity-25-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-25-hover
              </div>
              <div
                class="opacity-50-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-50-hover
              </div>
              <div
                class="opacity-75-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-75-hover
              </div>
              <div
                class="opacity-100-hover bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
              >
                .opacity-100-hover
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<div class="opacity-0-hover"></div>
<div class="opacity-5-hover"></div>
<div class="opacity-10-hover"></div>
<div class="opacity-15-hover"></div>
<div class="opacity-20-hover"></div>
<div class="opacity-25-hover"></div>
<div class="opacity-50-hover"></div>
<div class="opacity-75-hover"></div>
<div class="opacity-100-hover"></div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="opacity-active-state">
          <router-link to="#opacity-active-state"></router-link>
          Opacity Active State
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use
          <code>.opacity-active-{ 0, 5, 10, 15 ,20, 25, 50, 75, 100 }</code>to
          set an element's opacity on active states. Click on the demo below to
          trigger the active state.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center nav">
              <div
                class="nav-link opacity-active-0 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-0
              </div>
              <div
                class="nav-link opacity-active-5 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-5
              </div>
              <div
                class="nav-link opacity-active-10 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-10
              </div>
              <div
                class="nav-link opacity-active-15 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-15
              </div>
              <div
                class="nav-link opacity-active-20 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-20
              </div>
              <div
                class="nav-link opacity-active-25 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-25
              </div>
              <div
                class="nav-link opacity-active-50 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-50
              </div>
              <div
                class="nav-link opacity-active-75 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-75
              </div>
              <div
                class="nav-link opacity-active-100 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-active-100
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<div class="opacity-active-0"></div>
<div class="opacity-active-5"></div>
<div class="opacity-active-10"></div>
<div class="opacity-active-15"></div>
<div class="opacity-active-20"></div>
<div class="opacity-active-25"></div>
<div class="opacity-active-50"></div>
<div class="opacity-active-75"></div>
<div class="opacity-active-100"></div>`}}</CodeHighlighter
          >
        </div>
        <!--end::Code-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-10">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-3" id="opacity-combined">
          <router-link to="#opacity-combined"></router-link>
          Opacity Combined Active & Hover State
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          Use
          <code>.opacity-state-{ 0, 5, 10, 15 ,20, 25, 50, 75, 100 }</code>to
          set an element's opacity on both active and hover states. Click on the
          demo below to trigger the active state.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="pt-5">
          <div class="rounded border p-10">
            <div class="d-flex flex-wrap align-items-center nav">
              <div
                class="nav-link opacity-state-0 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-0
              </div>
              <div
                class="nav-link opacity-state-5 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-5
              </div>
              <div
                class="nav-link opacity-state-10 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-10
              </div>
              <div
                class="nav-link opacity-state-15 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-15
              </div>
              <div
                class="nav-link opacity-state-20 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-20
              </div>
              <div
                class="nav-link opacity-state-25 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-25
              </div>
              <div
                class="nav-link opacity-state-50 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-50
              </div>
              <div
                class="nav-link opacity-state-75 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-75
              </div>
              <div
                class="nav-link opacity-state-100 bg-primary text-white w-100px h-100px d-flex flex-center p-3 me-5 mb-5 fw-bold"
                data-bs-toggle="tab"
              >
                .opacity-state-100
              </div>
            </div>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Code-->
        <div class="py-5">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<div class="opacity-state-0"></div>
<div class="opacity-state-5"></div>
<div class="opacity-state-10"></div>
<div class="opacity-state-15"></div>
<div class="opacity-state-20"></div>
<div class="opacity-state-25"></div>
<div class="opacity-state-50"></div>
<div class="opacity-state-75"></div>
<div class="opacity-state-100"></div>`}}</CodeHighlighter
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
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "helpers-opacity",
  components: {
    CodeHighlighter,
  },
  setup() {
    return {
      appName,
    };
  },
});
</script>
