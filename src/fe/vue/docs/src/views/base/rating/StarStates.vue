<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="star-state">
      <router-link to="#star-state"></router-link>
      Star States
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      By default each star's state is unchecked and to mark them as checked just
      use <code>.checked</code> class for <code>.rating-label</code> element.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <div class="rating">
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
        </div>
        <div class="separator my-10"></div>
        <div class="rating">
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2">
            <i class="bi bi-star fs-2"></i>
          </div>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html" :field-height="400"
        >{{`<div class="rating">
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
      </div>
        <div class="separator my-10"></div>
        <div class="rating">
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2 checked">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2">
            <i class="bi bi-star fs-2"></i>
          </div>
          <div class="rating-label me-2">
            <i class="bi bi-star fs-2"></i>
          </div>
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
  name: "dot-separator",
  components: {
    CodeHighlighter,
  },
});
</script>
