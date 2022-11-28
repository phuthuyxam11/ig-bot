<template>
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="overview">
          <router-link to="#overview"></router-link>
          Overview
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="pt-4">
          <b>Element Plus</b>, a <b>Vue 3</b> based component library for
          developers, designers and product managers. Is used as a main Vue
          component library in {{ appName }}.
        </div>
        <!--end::Block-->

        <div class="pt-4">
          Default Element UI styles are connected in <code>src/App.vue</code>

          <CodeHighlighter lang="scss">
            {{ `@import "~element-plus/lib/theme-chalk/index.css";` }}
          </CodeHighlighter>
        </div>

        <div class="pt-4">
          All the components are available globally and you can use them without
          component import, just find suitable component in
          <a href="https://element-plus.org/en-US/component/border.html"
            >Element UI official doc</a
          >
          and use it on your page.
        </div>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="basic-usage">
          <router-link to="#basic-usage"></router-link>
          Basic Usage
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="pt-4">
          Here is a basic example of
          <a href="https://element-plus.org/en-US/component/message-box.html"
            >Message box</a
          >
          usage, you can find all available components in
          <a href="https://element-plus.org/en-US/component/border.html"
            >Element UI official doc</a
          >.
        </div>
        <!--end::Block-->

        <div class="pt-4">
          <el-button type="text" @click="open"
            >Click to open the Message Box</el-button
          >
        </div>

        <CodeHighlighter2>
          <template v-slot:html>{{ htmlCode }}</template>
          <template v-slot:js>{{ jsCode }}</template>
        </CodeHighlighter2>
      </div>
      <!--end::Section-->
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { appName } from "@/core/helpers/documentation";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import CodeHighlighter2 from "@/components/highlighters/CodeHighlighter2.vue";
import { ElMessageBox, ElMessage } from "element-plus";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";

export default defineComponent({
  name: "element-ui",
  components: {
    CodeHighlighter,
    CodeHighlighter2,
  },
  setup() {
    const open = () => {
      ElMessageBox.alert("This is a message", "Title", {
        confirmButtonText: "OK",
        callback: (action) => {
          ElMessage({
            type: "info",
            message: `action: ${action}`,
          });
        },
      });
    };

    const htmlCode = `<el-button type="text" @click="open">Click to open the Message Box</el-button>`;

    const jsCode = `import { defineComponent } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'

export default defineComponent({
  setup() {
    const open = () => {
      ElMessageBox.alert('This is a message', 'Title', {
        confirmButtonText: 'OK',
        callback: (action) => {
          ElMessage({
            type: 'info',
            message: \`action: \${action}\`,
          })
        },
      })
    }

    return {
      open,
    }
  },
})`;

    onMounted(() => {
      setCurrentPageBreadcrumbs("Element UI", ["General"]);
    });

    return {
      open,
      appName,
      htmlCode,
      jsCode,
    };
  },
});
</script>
