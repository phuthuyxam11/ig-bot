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
        <div class="py-4">
          <b>KTDatatables</b> is datatable component which allows you to display
          your data and manipulate it.
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="basic">
          <router-link to="#basic"></router-link>
          Basic Usage
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          To use the component you need to provide two required props
          <code>table-header</code> and <code>table-data</code> as well as add
          cell configuration slots for each column.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="pt-4">
          <Datatable :table-header="tableHeader" :table-data="tableData">
            <template v-slot:name="{ row: data }">
              {{ data.name }}
            </template>
            <template v-slot:position="{ row: data }">
              {{ data.position }}
            </template>
            <template v-slot:office="{ row: data }">
              {{ data.office }}
            </template>
            <template v-slot:age="{ row: data }">
              {{ data.age }}
            </template>
            <template v-slot:startDate="{ row: data }">
              {{ data.startDate }}
            </template>
            <template v-slot:salary="{ row: data }">
              {{ data.salary }}
            </template>
          </Datatable>
        </div>
        <!--end::Block-->
        <CodeHighlighter2>
          <template v-slot:html>
            {{ basicUsageHTML }}
          </template>
          <template v-slot:js>{{ basicUsageJS }}</template>
        </CodeHighlighter2>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="loading">
          <router-link to="#loading"></router-link>
          Loading
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          Use <code>loading</code> prop to enable and disable table loading. You
          can set loading to false while fetching data from api.
          <div class="mt-2 mb-1">
            <input
              class="me-3"
              type="radio"
              id="enabled"
              name="enabled"
              :value="true"
              v-model="loading"
            />
            <label for="enabled">Enable</label>
          </div>
          <div class="mb-1">
            <input
              class="me-3"
              type="radio"
              id="disabled"
              name="disabled"
              :value="false"
              v-model="loading"
            />
            <label for="disabled">Disable</label>
          </div>
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="pt-4">
          <Datatable
            :loading="loading"
            :table-header="tableHeader"
            :table-data="tableData1"
            :enable-items-per-page-dropdown="false"
          >
            <template v-slot:cell-name="{ row: data }">
              {{ data.name }}
            </template>
            <template v-slot:cell-position="{ row: data }">
              {{ data.position }}
            </template>
            <template v-slot:cell-office="{ row: data }">
              {{ data.office }}
            </template>
            <template v-slot:cell-age="{ row: data }">
              {{ data.age }}
            </template>
            <template v-slot:cell-startDate="{ row: data }">
              {{ data.startDate }}
            </template>
            <template v-slot:cell-salary="{ row: data }">
              {{ data.salary }}
            </template>
          </Datatable>
        </div>
        <!--end::Block-->
        <CodeHighlighter2>
          <template v-slot:html>
            {{ loadingHTML }}
          </template>
          <template v-slot:js>{{ loadingJS }}</template>
        </CodeHighlighter2>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="cell_customization">
          <router-link to="#cell_customization"></router-link>
          Cell customization.
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="pt-4">
          You are able to customize each cell by passing html code into named
          slots. Using slots you can display any html code, this can be used to
          display status, highlight values, display form elements e.t.c.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-4">
          One cell customization example:
          <CodeHighlighter lang="html">{{ slotsCode }}</CodeHighlighter>

          You need to specify such template for each column.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="pt-4">
          <Datatable
            :table-header="tableHeader1"
            :table-data="tableData2"
            :enable-items-per-page-dropdown="false"
          >
            <template v-slot:cell-order="{ row: invoice }">
              {{ invoice.order }}
            </template>
            <template v-slot:cell-amount="{ row: invoice }">
              <span :class="`text-${invoice.color}`">
                {{ invoice.amount }}
              </span>
            </template>
            <template v-slot:cell-status="{ row: invoice }">
              <span :class="`badge badge-light-${invoice.status.state}`">{{
                invoice.status.label
              }}</span>
            </template>
            <template v-slot:cell-date="{ row: invoice }">
              {{ invoice.date }}
            </template>
            <template v-slot:cell-invoice>
              <button class="btn btn-sm btn-light btn-active-light-primary">
                Download
              </button>
            </template>
          </Datatable>
        </div>
        <!--end::Block-->
        <div class="pt-4">
          <CodeHighlighter2>
            <template v-slot:html>
              {{ slotsHTML }}
            </template>
            <template v-slot:js>{{ slotsJS }}</template>
          </CodeHighlighter2>
        </div>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="header_config">
          <router-link to="#header_config"></router-link>
          Header configuration.
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-4">
          Header configuration is used to setup header column names, sorting
          information and column key (which is used in slots names). Header
          configuration is passed into component using
          <code>table-header</code> prop.
        </div>
        <!--end::Block-->

        <div class="table-responsive">
          <table class="table table-striped table-flush align-middle mb-0">
            <thead>
              <tr class="fw-bolder text-dark p-4">
                <th>Name</th>
                <th>Description</th>
                <th>Type</th>
                <th>Required</th>
              </tr>
            </thead>
            <tbody>
              <tr class="p-4">
                <th>name</th>
                <th>
                  Name of the filed, it will be displayed as a name of column in
                  table.
                </th>
                <th>String</th>
                <th class="text-danger">-</th>
              </tr>
              <tr class="p-4">
                <th>key</th>
                <th>Column key, used in named slots.</th>
                <th>String</th>
                <th class="text-success">+</th>
              </tr>
              <tr class="p-4">
                <th>sortingField</th>
                <th>
                  Specify specific filed which will be used in sorting, mostly
                  for columns which use few levels objects.<br />
                  Example: <code>data.status.label</code>
                </th>
                <th>String</th>
                <th class="text-danger">-</th>
              </tr>
              <tr class="p-4">
                <th>sortable</th>
                <th>Enable sorting for specific column.</th>
                <th>Boolean</th>
                <th class="text-danger">-</th>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="events">
          <router-link to="#events"></router-link>
          Events
        </h1>
        <!--end::Heading-->

        <div class="table-responsive">
          <table class="table table-striped table-flush align-middle mb-0">
            <thead>
              <tr class="fw-bolder text-dark p-4">
                <th>Name</th>
                <th>Description</th>
                <th>Parameters</th>
              </tr>
            </thead>
            <tbody>
              <tr class="p-4">
                <th>current-change</th>
                <th>triggers when page is changed</th>
                <th>page number</th>
              </tr>
              <tr class="p-4">
                <th>sort</th>
                <th>triggers when column is clicked</th>
                <th>order and column label</th>
              </tr>
              <tr class="p-4">
                <th>items-per-page-change</th>
                <th>triggers when changing items per page</th>
                <th>rows per page</th>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!--end::Section-->

      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="props">
          <router-link to="#props"></router-link>
          Props
        </h1>
        <!--end::Heading-->
        <div class="table-responsive">
          <table class="table table-striped table-flush align-middle mb-0">
            <thead>
              <tr class="fw-bolder text-dark p-4">
                <th>Name</th>
                <th>Description</th>
                <th>Type</th>
                <th>Required</th>
                <th>Default</th>
              </tr>
            </thead>
            <tbody>
              <tr class="p-4">
                <th>table-header</th>
                <th>table header configuration</th>
                <th>Array</th>
                <th class="text-success">+</th>
                <th>-</th>
              </tr>
              <tr class="p-4">
                <th>table-data</th>
                <th>data to display in table</th>
                <th>Array</th>
                <th class="text-success">+</th>
                <th>-</th>
              </tr>
              <tr class="p-4">
                <th>empty-table-text</th>
                <th>
                  text which will be displayed then <code>table-data</code> prop
                  is empty
                </th>
                <th>String</th>
                <th class="text-danger">-</th>
                <th>No data found</th>
              </tr>
              <tr class="p-4">
                <th>loading</th>
                <th>enable loading</th>
                <th>Boolean</th>
                <th class="text-danger">-</th>
                <th>false</th>
              </tr>
              <tr class="p-4">
                <th>current-page</th>
                <th>set pagination page</th>
                <th>Number</th>
                <th class="text-danger">-</th>
                <th>1</th>
              </tr>
              <tr class="p-4">
                <th>enable-items-per-page-dropdown</th>
                <th>enable items per page dropdown</th>
                <th>Boolean</th>
                <th class="text-danger">-</th>
                <th>true</th>
              </tr>
              <tr class="p-4">
                <th>total</th>
                <th>total items in table</th>
                <th>Number</th>
                <th class="text-danger">-</th>
                <th>0</th>
              </tr>
              <tr class="p-4">
                <th>rows-per-page</th>
                <th>
                  set how many rows display per page, accepts 10/25/50/100
                </th>
                <th>Number</th>
                <th class="text-danger">-</th>
                <th>10</th>
              </tr>
              <tr class="p-4">
                <th>order</th>
                <th>set sorting order, accepts asc/desc</th>
                <th>String</th>
                <th class="text-danger">-</th>
                <th>asc</th>
              </tr>
              <tr class="p-4">
                <th>sort-label</th>
                <th>sets currently sorted column</th>
                <th>String</th>
                <th class="text-danger">-</th>
                <th>-</th>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!--end::Section-->
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";
import CodeHighlighter2 from "@/components/highlighters/CodeHighlighter2.vue";
import Datatable from "@/components/kt-datatable/KTDatatable.vue";

export default defineComponent({
  name: "kt-datatables",
  components: {
    Datatable,
    CodeHighlighter,
    CodeHighlighter2,
  },
  setup() {
    const tableHeader = ref([
      {
        name: "Name",
        key: "name",
      },
      {
        name: "Position",
        key: "position",
      },
      {
        name: "Office",
        key: "office",
      },
      {
        name: "Age",
        key: "age",
      },
      {
        name: "Start date",
        key: "startDate",
      },
      {
        name: "Salary",
        key: "salary",
      },
    ]);
    const tableHeader1 = ref([
      {
        name: "Order id",
        key: "order",
        sortable: true,
      },
      {
        name: "Amount",
        key: "amount",
        sortable: true,
      },
      {
        name: "Status",
        key: "status",
        sortingField: "status.label",
        sortable: true,
      },
      {
        name: "Date",
        key: "date",
        sortable: true,
      },
      {
        name: "Invoice",
        key: "invoice",
        sortable: false,
      },
    ]);
    const tableData = ref([
      {
        name: "Tiger Nixon",
        position: "System Architect",
        office: "Edinburgh",
        age: "61",
        startDate: "2011/04/25",
        salary: "$320,800",
      },
      {
        name: "Garrett Winters",
        position: "Accountant",
        office: "Tokyo",
        age: "63",
        startDate: "2011/07/25",
        salary: "$170,750",
      },
      {
        name: "Ashton Cox",
        position: "Junior Technical Author",
        office: "San Francisco",
        age: "66",
        startDate: "2009/01/12",
        salary: "$86,000",
      },
      {
        name: "Cedric Kelly",
        position: "Senior Javascript Developer",
        office: "Edinburgh",
        age: "22",
        startDate: "2012/03/29",
        salary: "$433,060",
      },
      {
        name: "Airi Satou",
        position: "Accountant",
        office: "Tokyo",
        age: "33",
        startDate: "2008/11/28",
        salary: "$162,700",
      },
      {
        name: "Brielle Williamson",
        position: "Integration Specialist",
        office: "New York",
        age: "61",
        startDate: "2012/12/02",
        salary: "$372,000",
      },
      {
        name: "Herrod Chandler",
        position: "Sales Assistant",
        office: "San Francisco",
        age: "59",
        startDate: "2012/08/06",
        salary: "$137,500",
      },
      {
        name: "Rhona Davidson",
        position: "Integration Specialist",
        office: "Tokyo",
        age: "55",
        startDate: "2010/10/14",
        salary: "$327,900",
      },
      {
        name: "Colleen Hurst",
        position: "Javascript Developer",
        office: "San Francisco",
        age: "39",
        startDate: "2009/09/15",
        salary: "$205,500",
      },
      {
        name: "Sonya Frost",
        position: "Software Engineer",
        office: "Edinburgh",
        age: "23",
        startDate: "2008/12/13",
        salary: "$103,600",
      },
      {
        name: "Jena Gaines",
        position: "office Manager",
        office: "London",
        age: "30",
        startDate: "2008/12/19",
        salary: "$90,560",
      },
      {
        name: "Quinn Flynn",
        position: "Support Lead",
        office: "Edinburgh",
        age: "22",
        startDate: "2013/03/03",
        salary: "$342,000",
      },
      {
        name: "Charde Marshall",
        position: "Regional Director",
        office: "San Francisco",
        age: "36",
        startDate: "2008/10/16",
        salary: "$470,600",
      },
      {
        name: "Haley Kennedy",
        position: "Senior Marketing Designer",
        office: "London",
        age: "43",
        startDate: "2012/12/18",
        salary: "$313,500",
      },
      {
        name: "Tatyana Fitzpatrick",
        position: "Regional Director",
        office: "London",
        age: "19",
        startDate: "2010/03/17",
        salary: "$385,750",
      },
      {
        name: "Michael Silva",
        position: "Marketing Designer",
        office: "London",
        age: "66",
        startDate: "2012/11/27",
        salary: "$198,500",
      },
      {
        name: "Paul Byrd",
        position: "Chief Financial officer (CFO)",
        office: "New York",
        age: "64",
        startDate: "2010/06/09",
        salary: "$725,000",
      },
      {
        name: "Gloria Little",
        position: "Systems Administrator",
        office: "New York",
        age: "59",
        startDate: "2009/04/10",
        salary: "$237,500",
      },
      {
        name: "Bradley Greer",
        position: "Software Engineer",
        office: "London",
        age: "41",
        startDate: "2012/10/13",
        salary: "$132,000",
      },
      {
        name: "Dai Rios",
        position: "Personnel Lead",
        office: "Edinburgh",
        age: "35",
        startDate: "2012/09/26",
        salary: "$217,500",
      },
      {
        name: "Jenette Caldwell",
        position: "Development Lead",
        office: "New York",
        age: "30",
        startDate: "2011/09/03",
        salary: "$345,000",
      },
      {
        name: "Yuri Berry",
        position: "Chief Marketing officer (CMO)",
        office: "New York",
        age: "40",
        startDate: "2009/06/25",
        salary: "$675,000",
      },
      {
        name: "Caesar Vance",
        position: "Pre-Sales Support",
        office: "New York",
        age: "21",
        startDate: "2011/12/12",
        salary: "$106,450",
      },
      {
        name: "Doris Wilder",
        position: "Sales Assistant",
        office: "Sydney",
        age: "23",
        startDate: "2010/09/20",
        salary: "$85,600",
      },
      {
        name: "Angelica Ramos",
        position: "Chief Executive officer (CEO)",
        office: "London",
        age: "47",
        startDate: "2009/10/09",
        salary: "$1,200,000",
      },
      {
        name: "Gavin Joyce",
        position: "Developer",
        office: "Edinburgh",
        age: "42",
        startDate: "2010/12/22",
        salary: "$92,575",
      },
      {
        name: "Jennifer Chang",
        position: "Regional Director",
        office: "Singapore",
        age: "28",
        startDate: "2010/11/14",
        salary: "$357,650",
      },
      {
        name: "Brenden Wagner",
        position: "Software Engineer",
        office: "San Francisco",
        age: "28",
        startDate: "2011/06/07",
        salary: "$206,850",
      },
      {
        name: "Fiona Green",
        position: "Chief Operating officer (COO)",
        office: "San Francisco",
        age: "48",
        startDate: "2010/03/11",
        salary: "$850,000",
      },
      {
        name: "Shou Itou",
        position: "Regional Marketing",
        office: "Tokyo",
        age: "20",
        startDate: "2011/08/14",
        salary: "$163,000",
      },
      {
        name: "Michelle House",
        position: "Integration Specialist",
        office: "Sydney",
        age: "37",
        startDate: "2011/06/02",
        salary: "$95,400",
      },
      {
        name: "Suki Burks",
        position: "Developer",
        office: "London",
        age: "53",
        startDate: "2009/10/22",
        salary: "$114,500",
      },
      {
        name: "Prescott Bartlett",
        position: "Technical Author",
        office: "London",
        age: "27",
        startDate: "2011/05/07",
        salary: "$145,000",
      },
      {
        name: "Gavin Cortez",
        position: "Team Leader",
        office: "San Francisco",
        age: "22",
        startDate: "2008/10/26",
        salary: "$235,500",
      },
      {
        name: "Martena Mccray",
        position: "Post-Sales support",
        office: "Edinburgh",
        age: "46",
        startDate: "2011/03/09",
        salary: "$324,050",
      },
      {
        name: "Unity Butler",
        position: "Marketing Designer",
        office: "San Francisco",
        age: "47",
        startDate: "2009/12/09",
        salary: "$85,675",
      },
      {
        name: "Howard Hatfield",
        position: "office Manager",
        office: "San Francisco",
        age: "51",
        startDate: "2008/12/16",
        salary: "$164,500",
      },
      {
        name: "Hope Fuentes",
        position: "Secretary",
        office: "San Francisco",
        age: "41",
        startDate: "2010/02/12",
        salary: "$109,850",
      },
      {
        name: "Vivian Harrell",
        position: "Financial Controller",
        office: "San Francisco",
        age: "62",
        startDate: "2009/02/14",
        salary: "$452,500",
      },
      {
        name: "Timothy Mooney",
        position: "office Manager",
        office: "London",
        age: "37",
        startDate: "2008/12/11",
        salary: "$136,200",
      },
      {
        name: "Jackson Bradshaw",
        position: "Director",
        office: "New York",
        age: "65",
        startDate: "2008/09/26",
        salary: "$645,750",
      },
      {
        name: "Olivia Liang",
        position: "Support Engineer",
        office: "Singapore",
        age: "64",
        startDate: "2011/02/03",
        salary: "$234,500",
      },
      {
        name: "Bruno Nash",
        position: "Software Engineer",
        office: "London",
        age: "38",
        startDate: "2011/05/03",
        salary: "$163,500",
      },
      {
        name: "Sakura Yamamoto",
        position: "Support Engineer",
        office: "Tokyo",
        age: "37",
        startDate: "2009/08/19",
        salary: "$139,575",
      },
      {
        name: "Thor Walton",
        position: "Developer",
        office: "New York",
        age: "61",
        startDate: "2013/08/11",
        salary: "$98,540",
      },
      {
        name: "Finn Camacho",
        position: "Support Engineer",
        office: "San Francisco",
        age: "47",
        startDate: "2009/07/07",
        salary: "$87,500",
      },
      {
        name: "Serge Baldwin",
        position: "Data Coordinator",
        office: "Singapore",
        age: "64",
        startDate: "2012/04/09",
        salary: "$138,575",
      },
      {
        name: "Zenaida Frank",
        position: "Software Engineer",
        office: "New York",
        age: "63",
        startDate: "2010/01/04",
        salary: "$125,250",
      },
      {
        name: "Zorita Serrano",
        position: "Software Engineer",
        office: "San Francisco",
        age: "56",
        startDate: "2012/06/01",
        salary: "$115,000",
      },
      {
        name: "Jennifer Acosta",
        position: "Junior Javascript Developer",
        office: "Edinburgh",
        age: "43",
        startDate: "2013/02/01",
        salary: "$75,650",
      },
      {
        name: "Cara Stevens",
        position: "Sales Assistant",
        office: "New York",
        age: "46",
        startDate: "2011/12/06",
        salary: "$145,600",
      },
      {
        name: "Hermione Butler",
        position: "Regional Director",
        office: "London",
        age: "47",
        startDate: "2011/03/21",
        salary: "$356,250",
      },
      {
        name: "Lael Greer",
        position: "Systems Administrator",
        office: "London",
        age: "21",
        startDate: "2009/02/27",
        salary: "$103,500",
      },
      {
        name: "Jonas Alexander",
        position: "Developer",
        office: "San Francisco",
        age: "30",
        startDate: "2010/07/14",
        salary: "$86,500",
      },
      {
        name: "Shad Decker",
        position: "Regional Director",
        office: "Edinburgh",
        age: "51",
        startDate: "2008/11/13",
        salary: "$183,000",
      },
      {
        name: "Michael Bruce",
        position: "Javascript Developer",
        office: "Singapore",
        age: "29",
        startDate: "2011/06/27",
        salary: "$183,000",
      },
      {
        name: "Donna Snider",
        position: "Customer Support",
        office: "New York",
        age: "27",
        startDate: "2011/01/25",
        salary: "$112,000",
      },
    ]);
    const tableData1 = ref([
      {
        name: "Tiger Nixon",
        position: "System Architect",
        office: "Edinburgh",
        age: "61",
        startDate: "2011/04/25",
        salary: "$320,800",
      },
      {
        name: "Garrett Winters",
        position: "Accountant",
        office: "Tokyo",
        age: "63",
        startDate: "2011/07/25",
        salary: "$170,750",
      },
      {
        name: "Ashton Cox",
        position: "Junior Technical Author",
        office: "San Francisco",
        age: "66",
        startDate: "2009/01/12",
        salary: "$86,000",
      },
      {
        name: "Cedric Kelly",
        position: "Senior Javascript Developer",
        office: "Edinburgh",
        age: "22",
        startDate: "2012/03/29",
        salary: "$433,060",
      },
      {
        name: "Airi Satou",
        position: "Accountant",
        office: "Tokyo",
        age: "33",
        startDate: "2008/11/28",
        salary: "$162,700",
      },
      {
        name: "Brielle Williamson",
        position: "Integration Specialist",
        office: "New York",
        age: "61",
        startDate: "2012/12/02",
        salary: "$372,000",
      },
      {
        name: "Herrod Chandler",
        position: "Sales Assistant",
        office: "San Francisco",
        age: "59",
        startDate: "2012/08/06",
        salary: "$137,500",
      },
      {
        name: "Rhona Davidson",
        position: "Integration Specialist",
        office: "Tokyo",
        age: "55",
        startDate: "2010/10/14",
        salary: "$327,900",
      },
      {
        name: "Colleen Hurst",
        position: "Javascript Developer",
        office: "San Francisco",
        age: "39",
        startDate: "2009/09/15",
        salary: "$205,500",
      },
      {
        name: "Sonya Frost",
        position: "Software Engineer",
        office: "Edinburgh",
        age: "23",
        startDate: "2008/12/13",
        salary: "$103,600",
      },
    ]);
    const tableData2 = ref([
      {
        date: "Nov 01, 2020",
        order: "102445788",
        details: "Darknight transparency  36 Icons Pack",
        color: "success",
        amount: "$38.00",
        status: {
          label: "Pending",
          state: "warning",
        },
      },
      {
        date: "Oct 24, 2020",
        order: "423445721",
        details: "Seller Fee",
        color: "danger",
        amount: "$-2.60",
        status: {
          label: "Failed",
          state: "danger",
        },
      },
      {
        date: "Oct 08, 2020",
        order: "312445984",
        details: "Cartoon Mobile Emoji Phone Pack",
        color: "success",
        amount: "$76.00",
        status: {
          label: "Pending",
          state: "warning",
        },
      },
      {
        date: "Sep 15, 2020",
        order: "312445984",
        details: "Iphone 12 Pro Mockup  Mega Bundle",
        color: "success",
        amount: "$5.00",
        status: {
          label: "Success",
          state: "success",
        },
      },
      {
        date: "May 30, 2020",
        order: "523445943",
        details: "Seller Fee",
        color: "danger",
        amount: "$-1.30",
        status: {
          label: "Failed",
          state: "danger",
        },
      },
      {
        date: "Apr 22, 2020",
        order: "231445943",
        details: "Parcel Shipping / Delivery Service App",
        color: "success",
        amount: "$204.00",
        status: {
          label: "Pending",
          state: "warning",
        },
      },
      {
        date: "Feb 09, 2020",
        order: "426445943",
        details: "Visual Design Illustration",
        color: "success",
        amount: "$31.00",
        status: {
          label: "Pending",
          state: "warning",
        },
      },
      {
        date: "Nov 01, 2020",
        order: "984445943",
        details: "Abstract Vusial Pack",
        color: "success",
        amount: "$52.00",
        status: {
          label: "Success",
          state: "success",
        },
      },
      {
        date: "Jan 04, 2020",
        order: "324442313",
        details: "Seller Fee",
        color: "danger",
        amount: "$-0.80",
        status: {
          label: "Failed",
          state: "danger",
        },
      },
    ]);
    const loading = ref(true);

    const basicUsageHTML = `<Datatable :table-header="tableHeader" :table-data="tableData">
  <template v-slot:cell-name="{ row: data }">
    {{ data.name }}
  </template>
  <template v-slot:cell-position="{ row: data }">
    {{ data.position }}
  </template>
  <template v-slot:cell-office="{ row: data }">
    {{ data.office }}
  </template>
  <template v-slot:cell-age="{ row: data }">
    {{ data.age }}
  </template>
  <template v-slot:cell-startDate="{ row: data }">
    {{ data.startDate }}
  </template>
  <template v-slot:cell-salary="{ row: data }">
    {{ data.salary }}
  </template>
</Datatable>`;
    const basicUsageJS = `setup() {
    const tableHeader = ref([
      {
        name: "Name",
        key: "name",
      },
      {
        name: "Position",
        key: "position",
      },
      {
        name: "Office",
        key: "office",
      },
      {
        name: "Age",
        key: "age",
      },
      {
        name: "Start date",
        key: "startDate",
      },
      {
        name: "Salary",
        key: "salary",
      },
    ]);
    const tableData = ref([
      {
        name: "Tiger Nixon",
        position: "System Architect",
        office: "Edinburgh",
        age: "61",
        startDate: "2011/04/25",
        salary: "$320,800",
      },
      {
        name: "Garrett Winters",
        position: "Accountant",
        office: "Tokyo",
        age: "63",
        startDate: "2011/07/25",
        salary: "$170,750",
      }
      ...
    ]);

    return {
      tableHeader,
      tableData
    }
}`;

    const loadingHTML = `<Datatable
  :loading="loading"
  :table-header="tableHeader"
  :table-data="tableData1"
  :enable-items-per-page-dropdown="false"
>
  <template v-slot:cell-name="{ row: data }">
    {{ data.name }}
  </template>
  <template v-slot:cell-position="{ row: data }">
    {{ data.position }}
  </template>
  <template v-slot:cell-office="{ row: data }">
    {{ data.office }}
  </template>
  <template v-slot:cell-age="{ row: data }">
    {{ data.age }}
  </template>
  <template v-slot:cell-startDate="{ row: data }">
    {{ data.startDate }}
  </template>
  <template v-slot:cell-salary="{ row: data }">
    {{ data.salary }}
  </template>
</Datatable>`;
    const loadingJS = `setup() {
    const tableHeader = ref([
      {
        name: "Name",
        key: "name",
      },
      {
        name: "Position",
        key: "position",
      },
      {
        name: "Office",
        key: "office",
      },
      {
        name: "Age",
        key: "age",
      },
      {
        name: "Start date",
        key: "startDate",
      },
      {
        name: "Salary",
        key: "salary",
      },
    ]);
    const tableData1 = ref([
      {
        name: "Tiger Nixon",
        position: "System Architect",
        office: "Edinburgh",
        age: "61",
        startDate: "2011/04/25",
        salary: "$320,800",
      },
      {
        name: "Garrett Winters",
        position: "Accountant",
        office: "Tokyo",
        age: "63",
        startDate: "2011/07/25",
        salary: "$170,750",
      },
      ...
    ]);
    const loading = ref(true);

    return {
      tableHeader,
      tableData1,
      loading
    };
},`;

    const slotsCode = `<!-- Instead of key put the actual key which you have setup in header config —>
<!-- data represents current cell object -->
<template v-slot:cell-{key}="{ row: data }">
  <span :class="\`text-\${data.color}\`">
    {{ data.amount }}
  </span>
</template>`;
    const slotsHTML = `<Datatable
  :table-header="tableHeader1"
  :table-data="tableData2"
  :enable-items-per-page-dropdown="false"
>
  <template v-slot:cell-order="{ row: invoice }">
    {{ invoice.order }}
  </template>
  <template v-slot:cell-amount="{ row: invoice }">
    <span :class="\`text-\${invoice.color}\`">
      {{ invoice.amount }}
    </span>
  </template>
  <template v-slot:cell-status="{ row: invoice }">
    <span :class="\`badge badge-light-\${invoice.status.state}\`">{{
      invoice.status.label
    }}</span>
  </template>
  <template v-slot:cell-date="{ row: invoice }">
    {{ invoice.date }}
  </template>
  <template v-slot:cell-invoice>
    <button class="btn btn-sm btn-light btn-active-light-primary">
      Download
    </button>
  </template>
</Datatable>`;
    const slotsJS = `setup() {
    const tableHeader1 = ref([
      {
        name: "Order id",
        key: "order",
        sortable: true,
      },
      {
        name: "Amount",
        key: "amount",
        sortable: true,
      },
      {
        name: "Status",
        key: "status",
        sortingField: "status.label",
        sortable: true,
      },
      {
        name: "Date",
        key: "date",
        sortable: true,
      },
      {
        name: "Invoice",
        key: "invoice",
        sortable: false,
      },
    ]);
    const tableData2 = ref([
      {
        date: "Nov 01, 2020",
        order: "102445788",
        details: "Darknight transparency  36 Icons Pack",
        color: "success",
        amount: "$38.00",
        status: {
          label: "Pending",
          state: "warning",
        },
      },
      {
        date: "Oct 24, 2020",
        order: "423445721",
        details: "Seller Fee",
        color: "danger",
        amount: "$-2.60",
        status: {
          label: "Failed",
          state: "danger",
        },
      },
      ...
    ]);

    return {
      tableHeader1,
      tableData2
    };
},`;

    onMounted(() => {
      setCurrentPageBreadcrumbs("Datatables", ["General"]);
    });

    return {
      tableHeader,
      tableHeader1,
      tableData,
      tableData1,
      tableData2,
      loading,
      basicUsageHTML,
      basicUsageJS,
      loadingHTML,
      loadingJS,
      slotsCode,
      slotsHTML,
      slotsJS,
    };
  },
});
</script>
