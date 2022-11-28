<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="rich-content-example">
      <router-link to="#rich-content-example"></router-link>
      Rich Content Example
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Here's a sample of an alert with rich content elements inside it. Please
      take note of the HTML markup sample below for all required classes.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10 pb-0 d-flex flex-column flex-center">
        <div
          class="alert alert-dismissible bg-light-danger d-flex flex-center flex-column py-10 px-10 px-lg-20 mb-10"
        >
          <!--begin::Close-->
          <button
            type="button"
            class="position-absolute top-0 end-0 m-2 btn btn-icon btn-icon-danger"
            data-bs-dismiss="alert"
          >
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
          </button>
          <!--end::Close-->
          <span class="svg-icon svg-icon-5tx svg-icon-danger mb-5">
            <inline-svg src="media/icons/duotune/general/gen044.svg" />
          </span>
          <!--begin::Content-->
          <div class="text-center text-dark">
            <h1 class="fw-bolder mb-5">This is an alert</h1>
            <div
              class="separator separator-dashed border-danger opacity-25 mb-5"
            ></div>
            <div class="mb-9">
              The alert component can be used to highlight certain parts of your
              page for <strong>higher content visibility</strong>. <br />Please
              read our
              <a href="#" class="fw-bolder me-1">Terms and Conditions</a>for
              more info.
            </div>
            <!--begin::Buttons-->
            <div class="d-flex flex-center flex-wrap">
              <a
                href="#"
                class="btn btn-outline btn-outline-danger btn-active-danger m-2"
                >Cancel</a
              >
              <a href="#" class="btn btn-danger m-2">Ok, I got it</a>
            </div>
            <!--end::Buttons-->
          </div>
          <!--end::Content-->
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html" :field-height="400"
        >{{`<div
          class="
            alert alert-dismissible
            bg-light-danger
            d-flex
            flex-center flex-column
            py-10
            px-10 px-lg-20
            mb-10
          "
      >
        <!--begin::Close-->
        <button
            type="button"
            class="
              position-absolute
              top-0
              end-0
              m-2
              btn btn-icon btn-icon-danger
            "
            data-bs-dismiss="alert"
        >
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/arrows/arr061.svg" />
            </span>
        </button>
        <!--end::Close-->
        <span class="svg-icon svg-icon-5tx svg-icon-danger mb-5">
            <inline-svg src="media/icons/duotune/general/gen044.svg" />
          </span>
        <!--begin::Content-->
        <div class="text-center text-dark">
          <h1 class="fw-bolder mb-5">This is an alert</h1>
          <div
              class="separator separator-dashed border-danger opacity-25 mb-5"
          ></div>
          <div class="mb-9">
            The alert component can be used to highlight certain parts of your
            page for <strong>higher content visibility</strong>. <br />Please
            read our
            <a href="#" class="fw-bolder me-1">Terms and Conditions</a>for
            more info.
          </div>
          <!--begin::Buttons-->
          <div class="d-flex flex-center flex-wrap">
            <a
                href="#"
                class="btn btn-outline btn-outline-danger btn-active-danger m-2"
            >Cancel</a
            >
            <a href="#" class="btn btn-danger m-2">Ok, I got it</a>
          </div>
          <!--end::Buttons-->
        </div>
        <!--end::Content-->
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
