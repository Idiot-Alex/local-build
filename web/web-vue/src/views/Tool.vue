<script setup>
import { FilterMatchMode } from 'primevue/api'
import { ref, onMounted, onBeforeMount } from 'vue'
import { useToast } from 'primevue/usetoast'
import { toolList, saveTool, delTool } from '@/api/tool.js'

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
  { label: 'JDK', value: 'JDK' },
  { label: 'MAVEN', value: 'MAVEN' },
  { label: 'NODE', value: 'NODE' },
])

onBeforeMount(() => {
  initFilters()
})

onMounted(() => {
  loadList()
})

const loadList = () => {
  toolList(query.value).then(res => {
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
  if (!tempData.value.name || !tempData.value.path || !tempData.value.type) {
    toast.add({ severity: 'warning', summary: 'Tips', detail: 'Please input data', life: 3000 })
    return
  }
  saveTool(tempData.value).then(res => {
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
  delTool(data).then(res => {
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
  delTool(data).then(res => {
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
                  <h5 class="m-0">{{ t('tool_manage') }}</h5>
                  <span class="block mt-2 md:mt-0 p-input-icon-left">
                      <i class="pi pi-search" />
                      <InputText v-model="filters['global'].value" placeholder="Search..." />
                  </span>
              </div>
          </template>

          <Column selectionMode="multiple" frozen headerStyle="width: 3rem"></Column>
          <Column field="name" :header="t('tool.name')" :sortable="true" frozen></Column>
          <Column field="type" :header="t('tool.type')" :sortable="true"></Column>
          <Column field="version" :header="t('tool.version')" :sortable="true"></Column>
          <Column field="vendor" :header="t('tool.vender')" :sortable="true"></Column>
          <Column field="path" :header="t('tool.path')" :sortable="true">
            <template #body="slotProps">
              <Chip :label="slotProps.data.path" icon="pi pi-copy" cursor-pointer />
            </template>
          </Column>
          <Column :header="t('operation')" alignFrozen="right" frozen headerStyle="min-width:10rem;">
            <template #body="slotProps">
              <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2" @click="toEdit(slotProps.data)" />
              <Button icon="pi pi-trash" class="p-button-rounded p-button-warning mt-2" @click="confirmDelete(slotProps.data)" />
            </template>
          </Column>
        </DataTable>

        <Dialog v-model:visible="editDialog" w-450px header="编辑构建工具" :modal="true" class="p-fluid">
          <div class="field">
            <label for="name">名称</label>
            <InputText id="name" v-model.trim="tempData.name" required="true" autofocus :class="{ 'p-invalid': submitted && !tempData.name }" />
            <small color-red class="p-invalid" v-if="submitted && !tempData.name">名称不能为空</small>
          </div>
          <div class="field">
            <label for="path">目录</label>
            <Textarea id="path" v-model="tempData.path" rows="2" cols="20" />
          </div>
          <div class="field">
            <label for="type" class="mb-3">类型</label>
            <Dropdown id="type" v-model="tempData.type" :options="types" optionLabel="label" optionValue="value" placeholder="选择一个类型" required="true" :class="{ 'p-invalid': submitted && !tempData.type }">
            </Dropdown>
            <small color-red class="p-invalid" v-if="submitted && !tempData.type">类型不能为空</small>
          </div>
          <div class="field">
            <label for="desc">描述</label>
            <Textarea id="desc" v-model="tempData.desc" required="true" rows="3" cols="20" />
          </div>
          <template #footer>
            <Button label="取消" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
            <Button label="保存" icon="pi pi-check" class="p-button-text" @click="saveData" />
          </template>
        </Dialog>

        <Dialog v-model:visible="deleteDialog" w-450px header="Confirm" :modal="true">
          <div class="flex align-items-center justify-content-center">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="tempData">
              是否确定删除 
              <b>{{ tempData.name }}</b>
              ?
            </span>
          </div>
          <template #footer>
            <Button label="否" icon="pi pi-times" class="p-button-text" @click="deleteDialog = false" />
            <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteAction" />
          </template>
        </Dialog>

        <Dialog v-model:visible="deleteDialogs" :style="{ width: '450px' }" header="Confirm" :modal="true">
          <div class="flex align-items-center justify-content-center">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="tempData">是否确定删除选中的记录?</span>
          </div>
          <template #footer>
            <Button label="否" icon="pi pi-times" class="p-button-text" @click="deleteDialogs = false" />
            <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteSelected" />
          </template>
        </Dialog>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
