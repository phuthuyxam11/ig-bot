<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="vertical-tabs">
      <router-link to="#vertical-tabs"></router-link>
      Vertical Tabs
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      To use a custom Vertical Tabs style just add
      <code>.flex-column</code> class as shown below:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="d-flex flex-column flex-md-row rounded border p-10">
        <ul
          class="nav nav-tabs nav-pills border-0 flex-row flex-md-column me-5 mb-3 mb-md-0 fs-6"
        >
          <li class="nav-item w-md-150px me-0">
            <a
              class="nav-link active"
              data-bs-toggle="tab"
              href="#kt_vtab_pane_1"
              >Link 1</a
            >
          </li>
          <li class="nav-item w-md-150px me-0">
            <a class="nav-link" data-bs-toggle="tab" href="#kt_vtab_pane_2"
              >Link 2</a
            >
          </li>
          <li class="nav-item w-md-150px">
            <a class="nav-link" data-bs-toggle="tab" href="#kt_vtab_pane_3"
              >Link 3</a
            >
          </li>
        </ul>
        <div class="tab-content" id="myTabContent">
          <div
            class="tab-pane fade show active"
            id="kt_vtab_pane_1"
            role="tabpanel"
          >
            Et et consectetur ipsum labore excepteur est proident excepteur ad
            velit occaecat qui minim occaecat veniam. Fugiat veniam incididunt
            anim aliqua enim pariatur veniam sunt est aute sit dolor anim. Velit
            non irure adipisicing aliqua ullamco irure incididunt irure non esse
            consectetur nostrud minim non minim occaecat. Amet duis do nisi duis
            veniam non est eiusmod tempor incididunt tempor dolor ipsum in qui
            sit.
          </div>
          <div class="tab-pane fade" id="kt_vtab_pane_2" role="tabpanel">
            Nulla est ullamco ut irure incididunt nulla Lorem Lorem minim irure
            officia enim reprehenderit. Magna duis labore cillum sint
            adipisicing exercitation ipsum. Nostrud ut anim non exercitation
            velit laboris fugiat cupidatat. Commodo esse dolore fugiat sint
            velit ullamco magna consequat voluptate minim amet aliquip ipsum
            aute laboris nisi. Labore labore veniam irure irure ipsum pariatur
            mollit magna in cupidatat dolore magna irure esse tempor ad mollit.
            Dolore commodo nulla minim amet ipsum officia consectetur amet
            ullamco voluptate nisi commodo ea sit eu.
          </div>
          <div class="tab-pane fade" id="kt_vtab_pane_3" role="tabpanel">
            Sint sit mollit irure quis est nostrud cillum consequat Lorem esse
            do quis dolor esse fugiat sunt do. Eu ex commodo veniam Lorem
            aliquip laborum occaecat qui Lorem esse mollit dolore anim
            cupidatat. eserunt officia id Lorem nostrud aute id commodo elit
            eiusmod enim irure amet eiusmod qui reprehenderit nostrud tempor.
            Fugiat ipsum excepteur in aliqua non et quis aliquip ad irure in
            labore cillum elit enim. Consequat aliquip incididunt ipsum et minim
            laborum laborum laborum et cillum labore. Deserunt adipisicing
            cillum id nulla minim nostrud labore eiusmod et amet.
          </div>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<ul class="nav nav-tabs nav-pills border-0 flex-row flex-md-column me-5 mb-3 mb-md-0 fs-6">
        <li class="nav-item w-md-150px me-0">
          <a class="nav-link active" data-bs-toggle="tab" href="#kt_vtab_pane_1">Link 1</a>
        </li>
        <li class="nav-item w-md-150px me-0">
          <a class="nav-link" data-bs-toggle="tab" href="#kt_vtab_pane_2">Link 2</a>
        </li>
        <li class="nav-item w-md-150px">
          <a class="nav-link" data-bs-toggle="tab" href="#kt_vtab_pane_3">Link 3</a>
        </li>
      </ul>`}}</CodeHighlighter
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
  name: "vertical-tabs",
  components: {
    CodeHighlighter,
  },
});
</script>
