<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="solid-colors">
      <router-link to="#solid-colors"></router-link>
      Basic Example
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.alert</code> with <code>.bg-{color}</code> classes to set the
      alert's style with a solid color.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      Add the class <code>.alert-dismissible</code> to the parent
      <code>.alert</code> wrapper and <code>data-bs-dismiss="alert"</code> to
      the close button to allow dismissing alerts. Please take note of the HTML
      markup sample below for all required responsive classes.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10 pb-0 d-flex flex-column">
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-primary d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <!--begin::Icon-->
          <!--begin::Svg Icon | path: -->
          <span class="svg-icon svg-icon-2hx svg-icon-light me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/files/fil024.svg" />
          </span>
          <!--end::Svg Icon-->
          <!--end::Icon-->
          <!--begin::Content-->
          <div class="d-flex flex-column text-light pe-0 pe-sm-10">
            <h4 class="mb-2 text-light">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
          <!--end::Content-->
          <!--begin::Close-->
          <button
            type="button"
            class="position-absolute position-sm-relative m-2 m-sm-0 top-0 end-0 btn btn-icon ms-sm-auto"
            data-bs-dismiss="alert"
          >
            <span class="svg-icon svg-icon-2x svg-icon-light">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-danger d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span class="svg-icon svg-icon-2hx svg-icon-light me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/communication/com003.svg" />
          </span>
          <!--begin::Content-->
          <div class="d-flex flex-column text-light pe-0 pe-sm-10">
            <h4 class="mb-2 text-light">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
          <!--end::Content-->
          <!--begin::Close-->
          <button
            type="button"
            class="position-absolute position-sm-relative m-2 m-sm-0 top-0 end-0 btn btn-icon ms-sm-auto"
            data-bs-dismiss="alert"
          >
            <span class="svg-icon svg-icon-2x svg-icon-light">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-success d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span class="svg-icon svg-icon-2hx svg-icon-light me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/art/art005.svg" />
          </span>
          <!--begin::Content-->
          <div class="d-flex flex-column text-light pe-0 pe-sm-10">
            <h4 class="mb-2 text-light">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
          <!--end::Content-->
          <!--begin::Close-->
          <button
            type="button"
            class="position-absolute position-sm-relative m-2 m-sm-0 top-0 end-0 btn btn-icon ms-sm-auto"
            data-bs-dismiss="alert"
          >
            <span class="svg-icon svg-icon-2x svg-icon-light">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-info d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span class="svg-icon svg-icon-2hx svg-icon-light me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/general/gen007.svg" />
          </span>
          <!--begin::Content-->
          <div class="d-flex flex-column text-light pe-0 pe-sm-10">
            <h4 class="mb-2 text-light">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
          <!--end::Content-->
          <!--begin::Close-->
          <button
            type="button"
            class="position-absolute position-sm-relative m-2 m-sm-0 top-0 end-0 btn btn-icon ms-sm-auto"
            data-bs-dismiss="alert"
          >
            <span class="svg-icon svg-icon-2x svg-icon-light">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<div
          class="
            alert alert-dismissible
            bg-info
            d-flex
            flex-column flex-sm-row
            w-100
            p-5
            mb-10
          "
      >
          <span class="svg-icon svg-icon-2hx svg-icon-light me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/general/gen007.svg" />
          </span>
        <!--begin::Content-->
        <div class="d-flex flex-column text-light pe-0 pe-sm-10">
          <h4 class="mb-2 text-light">This is an alert</h4>
          <span
          >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
          >
        </div>
        <!--end::Content-->
        <!--begin::Close-->
        <button
            type="button"
            class="
              position-absolute position-sm-relative
              m-2 m-sm-0
              top-0
              end-0
              btn btn-icon
              ms-sm-auto
            "
            data-bs-dismiss="alert"
        >
            <span class="svg-icon svg-icon-2x svg-icon-light">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
        </button>
        <!--end::Close-->
      </div>`}}</CodeHighlighter
      >
    </div>
    <!--end::Code-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "basic-example",
  components: {
    CodeHighlighter,
  },
});
</script>
