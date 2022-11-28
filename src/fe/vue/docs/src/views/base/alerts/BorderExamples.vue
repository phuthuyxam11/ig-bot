<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="border-example">
      <router-link to="#border-example"></router-link>
      Border Examples
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.alert</code> with <code>.bg-light-{color}</code> and
      <code>.border</code> classes to set the alert's style with borders.
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
          class="alert alert-dismissible bg-light-primary border border-primary d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span
            class="svg-icon svg-icon-2hx svg-icon-primary me-4 mb-5 mb-sm-0"
          >
            <inline-svg src="media/icons/duotune/files/fil024.svg" />
          </span>

          <!--begin::Content-->
          <div class="d-flex flex-column pe-0 pe-sm-10">
            <h5 class="mb-1">This is an alert with a basic border</h5>
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
            <i class="bi bi-x fs-1 text-primary"></i>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-light-danger border border-danger border-dashed d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span class="svg-icon svg-icon-2hx svg-icon-danger me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/communication/com003.svg" />
          </span>

          <!--begin::Content-->
          <div class="d-flex flex-column pe-0 pe-sm-10">
            <h5 class="mb-1">This is an alert with a dashed border</h5>
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
            <i class="bi bi-x fs-1 text-danger"></i>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-light-success border border-success border-3 d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span
            class="svg-icon svg-icon-2hx svg-icon-success me-4 mb-5 mb-sm-0"
          >
            <inline-svg src="media/icons/duotune/art/art005.svg" />
          </span>

          <!--begin::Content-->
          <div class="d-flex flex-column pe-0 pe-sm-10">
            <h5 class="mb-1">This is an alert with a thick border</h5>
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
            <i class="bi bi-x fs-1 text-success"></i>
          </button>
          <!--end::Close-->
        </div>
        <!--end::Alert-->
        <!--begin::Alert-->
        <div
          class="alert alert-dismissible bg-light-info border border-info border-3 border-dashed d-flex flex-column flex-sm-row w-100 p-5 mb-10"
        >
          <span class="svg-icon svg-icon-2hx svg-icon-info me-4 mb-5 mb-sm-0">
            <inline-svg src="media/icons/duotune/general/gen007.svg" />
          </span>

          <!--begin::Content-->
          <div class="d-flex flex-column pe-0 pe-sm-10">
            <h5 class="mb-1">This is an alert with a thick dashed border</h5>
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
            <i class="bi bi-x fs-1 text-info"></i>
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
            bg-light-primary
            border border-primary
            d-flex
            flex-column flex-sm-row
            w-100
            p-5
            mb-10
          "
      >
          <span
              class="svg-icon svg-icon-2hx svg-icon-primary me-4 mb-5 mb-sm-0"
          >
            <inline-svg src="media/icons/duotune/files/fil024.svg" />
          </span>

        <!--begin::Content-->
        <div class="d-flex flex-column pe-0 pe-sm-10">
          <h5 class="mb-1">This is an alert with a basic border</h5>
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
          <i class="bi bi-x fs-1 text-primary"></i>
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
