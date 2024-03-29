<script setup>
import { FilterMatchMode } from 'primevue/api'
import { ref, onMounted, onBeforeMount } from 'vue'
import { useToast } from 'primevue/usetoast'
import { buildPlanList, saveBuildPlan, delBuildPlan } from '@/api/build-plan.js'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const toast = useToast()

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
const types = ref([
  { label: 'GIT', value: 'GIT' },
  { label: 'DIR', value: 'DIR' },
  { label: 'SVN', value: 'SVN' },
])

onBeforeMount(() => {
  initFilters()
})

onMounted(() => {
  loadList()
})

const loadList = () => {
  buildPlanList(query.value).then(res => {
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
  if (!tempData.value.name || !tempData.value.type) {
    toast.add({ severity: 'warning', summary: 'Tips', detail: 'Please input data', life: 3000 })
    return
  }
  saveBuildPlan(tempData.value).then(res => {
    toast.add({ severity: 'success', summary: 'Info', detail: res.msg, life: 3000 })
    editDialog.value = false
    tempData.value = {}
    loadList()
  })
}

const toEdit = (data) => {
  tempData.value = { ...data }
  editDialog.value = true
}

const confirmDelete = (data) => {
  tempData.value = data
  deleteDialog.value = true
}

const deleteAction = () => {
  const data = {
    ids: [ tempData.value.id ]
  }
  delBuildPlan(data).then(res => {
    toast.add({ severity: 'success', summary: 'Info', detail: res.msg, life: 3000 })
    deleteDialog.value = false
    tempData.value = {}
    loadList()
  })
}

const exportCSV = () => {
  dt.value.exportCSV()
}

const confirmDeleteSelected = () => {
  deleteDialogs.value = true
}

const deleteSelected = () => {
  const data = {
    ids: []
  }
  selectedData.value.filter(d => data.ids.push(d.id))
  delProject(data).then(res => {
    toast.add({ severity: 'success', summary: 'Info', detail: res.msg, life: 3000 })
    deleteDialogs.value = false
    selectedData.value = null
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
              <Button :label="t('new')" icon="pi pi-plus" class="p-button-success mr-2" @click="openNew" />
              <Button :label="t('delete')" icon="pi pi-trash" class="p-button-danger" @click="confirmDeleteSelected" :disabled="!selectedData || !selectedData.length" />
            </div>
          </template>
          <template v-slot:end>
            <Button :label="t('export')" icon="pi pi-upload" class="p-button-help" @click="exportCSV($event)" />
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
                  <h5 class="m-0">{{ t('build_plan_manage') }}</h5>
                  <span class="block mt-2 md:mt-0 p-input-icon-left">
                      <i class="pi pi-search" />
                      <InputText v-model="filters['global'].value" placeholder="Search..." />
                  </span>
              </div>
          </template>

          <Column selectionMode="multiple" frozen headerStyle="width: 3rem"></Column>
          <Column field="name" :header="t('build_plan.name')" :sortable="true" frozen></Column>
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

        <Dialog v-model:visible="editDialog" w-450px header="Project Details" :modal="true" class="p-fluid">
          <div class="field">
            <label for="name">Name</label>
            <InputText id="name" v-model.trim="tempData.name" required="true" autofocus :class="{ 'p-invalid': submitted && !tempData.name }" />
            <small color-red class="p-invalid" v-if="submitted && !tempData.name">Name is required.</small>
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

        <Dialog v-model:visible="deleteDialog" w-450px header="Confirm" :modal="true">
          <div class="flex align-items-center justify-content-center">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="tempData">
              Are you sure you want to delete 
              <b>{{ tempData.name }}</b>
              ?
            </span>
          </div>
          <template #footer>
            <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteDialog = false" />
            <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteAction" />
          </template>
        </Dialog>

        <Dialog v-model:visible="deleteDialogs" :style="{ width: '450px' }" header="Confirm" :modal="true">
          <div class="flex align-items-center justify-content-center">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="tempData">Are you sure you want to delete the selected tools?</span>
          </div>
          <template #footer>
            <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteDialogs = false" />
            <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteSelected" />
          </template>
        </Dialog>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>