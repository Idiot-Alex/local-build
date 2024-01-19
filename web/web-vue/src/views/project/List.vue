<script setup>
import { FilterMatchMode } from 'primevue/api'
import { ref, onMounted, onBeforeMount } from 'vue'
import { projectList, saveProject } from '@/api/project.js'

const query = ref({
  pageNo: 1,
  pageSize: 10
})
const total = ref(null)
const dataList = ref(null)
const editDialog = ref(false)
const deleteDialog = ref(false)
const deleteDialogs = ref(false)
const tempData = ref({})
const selectedData = ref(null)
const dt = ref(null)
const filters = ref({})
const submitted = ref(false)

onBeforeMount(() => {
  initFilters()
})

onMounted(() => {
  loadList()
})

const loadList = () => {
  projectList(query.value).then(res => {
    dataList.value = res.data.dataList
    total.value = res.data.total
  })
}

const handlePage = (event) => {
  console.log(event.page, event.rows)
  query.value.pageNo = event.page + 1
  query.value.pageSize = event.rows
  loadList()
}

const openNew = () => {
  tempData.value = {}
  submitted.value = false
  editDialog.value = true
}

const hideDialog = () => {
  editDialog.value = false
  submitted.value = false
}

const saveData = () => {
  submitted.value = true
  saveProject(tempData.value).then(res => {
    toast.add({ severity: 'success', summary: 'Successful', detail: res.msg, life: 3000 })
    editDialog.value = false
    tempData.value = {}
    loadList()
  })
}

const initFilters = () => {
  filters.value = {
    global: { value: null, matchMode: FilterMatchMode.CONTAINS }
  }
}
</script>

<template>
  <div class="grid">
    <div class="col-12">
      <div class="card">
        <Toast />
        <Toolbar class="mb-4">
          <template v-slot:start>
            <div class="my-2">
              <Button label="New" icon="pi pi-plus" class="p-button-success mr-2" @click="openNew" />
              <Button label="Delete" icon="pi pi-trash" class="p-button-danger" @click="confirmDeleteSelected" :disabled="!selectedData || !selectedData.length" />
            </div>
          </template>
          <template v-slot:end>
            <Button label="Export" icon="pi pi-upload" class="p-button-help" @click="exportCSV($event)" />
          </template>
        </Toolbar>
        
        <DataTable
          ref="dt"
          :value="dataList"
          v-model:selection="selectedData"
          dataKey="id"
          scrollable
          :filters="filters"
          paginator
          :totalRecords="total"
          :page="query.pageNo"
          :rows="query.pageSize"
          @page="handlePage"
          paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
          :rowsPerPageOptions="[10, 25, 50, 100]"
          currentPageReportTemplate="Total {totalRecords}"
          responsiveLayout="scroll"
          lazy
          @lazyload="loadList"
        >
          <template #header>
              <div class="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
                  <h5 class="m-0">Tools Manage</h5>
                  <span class="block mt-2 md:mt-0 p-input-icon-left">
                      <i class="pi pi-search" />
                      <InputText v-model="filters['global'].value" placeholder="Search..." />
                  </span>
              </div>
          </template>

          <Column selectionMode="multiple" frozen headerStyle="width: 3rem"></Column>
          <Column field="name" header="name" :sortable="true" frozen></Column>
          <Column field="type" header="type" :sortable="true"></Column>
          <Column field="path" header="path" :sortable="true">
            <template #body="slotProps">
              <Chip :label="slotProps.data.path" icon="pi pi-copy" cursor-pointer />
            </template>
          </Column>
          <Column header="Operation" alignFrozen="right" frozen headerStyle="min-width:10rem;">
            <template #body="slotProps">
              <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2" @click="toEdit(slotProps.data)" />
              <Button icon="pi pi-trash" class="p-button-rounded p-button-warning mt-2" @click="confirmDelete(slotProps.data)" />
            </template>
          </Column>
        </DataTable>

        <Dialog v-model:visible="editDialog" w-450px header="Tool Details" :modal="true" class="p-fluid">
          <div class="field">
            <label for="name">Name</label>
            <InputText id="name" v-model.trim="tempData.name" required="true" autofocus :class="{ 'p-invalid': submitted && !tempData.name }" />
            <small color-red class="p-invalid" v-if="submitted && !tempData.name">Name is required.</small>
          </div>
          <div class="field">
            <label for="path">Path</label>
            <Textarea id="path" v-model="tempData.path" rows="2" cols="20" />
          </div>
          <div class="field">
            <label for="type" class="mb-3">Type</label>
            <Dropdown id="type" v-model="tempData.type" :options="types" optionLabel="label" optionValue="value" placeholder="Select a Type" required="true" :class="{ 'p-invalid': submitted && !tempData.type }">
            </Dropdown>
            <small color-red class="p-invalid" v-if="submitted && !tempData.type">Type is required.</small>
          </div>
          <div class="field">
            <label for="desc">Description</label>
            <Textarea id="desc" v-model="tempData.desc" required="true" rows="3" cols="20" />
          </div>
          <template #footer>
            <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
            <Button label="Save" icon="pi pi-check" class="p-button-text" @click="saveData" />
          </template>
        </Dialog>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>